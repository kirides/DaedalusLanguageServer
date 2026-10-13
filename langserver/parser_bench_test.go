package langserver

import (
	"context"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/antlr4-go/antlr/v4"
	"github.com/kirides/DaedalusLanguageServer/daedalus/parser"
	"golang.org/x/text/encoding/charmap"
)

type benchScript struct{ name, content string }

func loadBenchScripts(tb testing.TB) (scripts []benchScript, total int) {
	tb.Helper()
	err := filepath.WalkDir("testdata", func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(strings.ToLower(path), ".d") {
			return err
		}
		raw, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		body, _ := charmap.Windows1252.NewDecoder().Bytes(raw)
		scripts = append(scripts, benchScript{path, string(body)})
		total += len(body)
		return nil
	})
	if err != nil || len(scripts) == 0 {
		tb.Fatalf("no bench scripts: %v", err)
	}
	return scripts, total
}

// BenchmarkLex measures tokenization only.
func BenchmarkLex(b *testing.B) {
	scripts, total := loadBenchScripts(b)
	lexer := parser.NewDaedalusLexer(nil)
	b.SetBytes(int64(total))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, s := range scripts {
			lexer.SetInputStream(antlr.NewInputStream(s.content))
			for t := lexer.NextToken(); t.GetTokenType() != antlr.TokenEOF; t = lexer.NextToken() {
			}
		}
	}
}

// BenchmarkParseTree measures lex + parse (no listeners).
func BenchmarkParseTree(b *testing.B) {
	scripts, total := loadBenchScripts(b)
	p := newRegularParser()
	b.SetBytes(int64(total))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, s := range scripts {
			p.Parse(s.name, s.content, &parser.BaseDaedalusListener{}, &SyntaxErrorListener{})
		}
	}
}

// BenchmarkParseScripts measures the full ParseScript path (parser + stateful listener).
func BenchmarkParseScripts(b *testing.B) {
	scripts, total := loadBenchScripts(b)
	m := newParseResultsManager(nopLogger{})
	b.SetBytes(int64(total))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, s := range scripts {
			m.ParseScript(s.name, s.content, time.Now())
		}
	}
}

// BenchmarkParseScriptsParallel measures how well parsing scales across goroutines.
func BenchmarkParseScriptsParallel(b *testing.B) {
	scripts, total := loadBenchScripts(b)
	m := newParseResultsManager(nopLogger{})
	b.SetBytes(int64(total))
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for _, s := range scripts {
				m.ParseScript(s.name, s.content, time.Now())
			}
		}
	})
}

// loadCorpusScripts loads the scripts of the directories in $DLS_CORPUS (separated by ';'),
// taking every n-th file so that at most maxBytes are used.
func loadCorpusScripts(tb testing.TB, maxBytes int) (scripts []benchScript, total int) {
	tb.Helper()
	roots := os.Getenv("DLS_CORPUS")
	if roots == "" {
		tb.Skip("DLS_CORPUS not set")
	}
	var files []string
	var all int
	for _, r := range strings.Split(roots, ";") {
		filepath.WalkDir(r, func(p string, d fs.DirEntry, err error) error {
			if err == nil && !d.IsDir() && strings.HasSuffix(strings.ToLower(p), ".d") {
				files = append(files, p)
				if st, err := d.Info(); err == nil {
					all += int(st.Size())
				}
			}
			return nil
		})
	}
	sort.Strings(files)
	step := 1
	if maxBytes > 0 && all > maxBytes {
		step = (all + maxBytes - 1) / maxBytes
	}
	for i := 0; i < len(files); i += step {
		raw, err := os.ReadFile(files[i])
		if err != nil {
			continue
		}
		body, _ := charmap.Windows1252.NewDecoder().Bytes(raw)
		scripts = append(scripts, benchScript{files[i], string(body)})
		total += len(body)
	}
	return scripts, total
}

// BenchmarkParseCorpus measures ParseScript on a sample of real scripts ($DLS_CORPUS).
func BenchmarkParseCorpus(b *testing.B) {
	scripts, total := loadCorpusScripts(b, 4<<20)
	m := newParseResultsManager(nopLogger{})
	for _, s := range scripts { // warm up the prediction caches
		m.ParseScript(s.name, s.content, time.Now())
	}
	b.SetBytes(int64(total))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, s := range scripts {
			m.ParseScript(s.name, s.content, time.Now())
		}
	}
}

// TestParseSourceMemory reports peak and retained heap of parsing a whole project:
//
//	DLS_SRC='path\to\Gothic.src' go test ./langserver -run ParseSourceMemory -v
func TestParseSourceMemory(t *testing.T) {
	src := os.Getenv("DLS_SRC")
	if src == "" {
		t.Skip("DLS_SRC not set")
	}
	var peak atomic.Uint64
	stop := make(chan struct{})
	done := make(chan struct{})
	go func() {
		defer close(done)
		var ms runtime.MemStats
		for {
			select {
			case <-stop:
				return
			case <-time.After(5 * time.Millisecond):
			}
			runtime.ReadMemStats(&ms)
			if ms.HeapInuse > peak.Load() {
				peak.Store(ms.HeapInuse)
			}
		}
	}()
	m := newParseResultsManager(nopLogger{})
	start := time.Now()
	res, err := m.ParseSource(context.Background(), src)
	elapsed := time.Since(start)
	close(stop)
	<-done
	if err != nil {
		t.Fatal(err)
	}
	runtime.GC()
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	t.Logf("scripts=%d time=%v peakHeapInuse=%dMB retainedAfterGC=%dMB totalAlloc=%dMB numGC=%d",
		len(res), elapsed.Round(time.Millisecond), peak.Load()>>20, ms.HeapAlloc>>20, ms.TotalAlloc>>20, ms.NumGC)
}
