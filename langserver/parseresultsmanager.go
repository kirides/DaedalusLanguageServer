package langserver

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/kirides/DaedalusLanguageServer/daedalus/symbol"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
)

type parseResultsManager struct {
	parseResults map[string]*ParseResult
	mtx          sync.RWMutex
	logger       Logger

	fileEncoding encoding.Encoding
	srcEncoding  encoding.Encoding
	decoderPool  *sync.Pool

	parser           Parser
	NumParserThreads int
}

func newParseResultsManager(logger Logger) *parseResultsManager {
	return &parseResultsManager{
		parseResults:     make(map[string]*ParseResult),
		logger:           logger,
		fileEncoding:     charmap.Windows1252,
		srcEncoding:      charmap.Windows1252,
		decoderPool:      &sync.Pool{New: func() interface{} { return charmap.Windows1252.NewDecoder() }},
		parser:           newRegularParser(),
		NumParserThreads: 0,
	}
}

var encodings = map[string]encoding.Encoding{
	"WINDOWS-1250": charmap.Windows1250,
	"WINDOWS-1251": charmap.Windows1251,
	"WINDOWS-1252": charmap.Windows1252,
	"WINDOWS-1253": charmap.Windows1253,
	"WINDOWS-1254": charmap.Windows1254,
	"WINDOWS-1255": charmap.Windows1255,
	"WINDOWS-1256": charmap.Windows1256,
	"WINDOWS-1257": charmap.Windows1257,
	"WINDOWS-1258": charmap.Windows1258,
}

func (m *parseResultsManager) SetFileEncoding(enc string) error {
	if e, ok := encodings[strings.ToUpper(enc)]; ok {
		m.fileEncoding = e
		m.decoderPool = &sync.Pool{New: func() interface{} { return e.NewDecoder() }}
		return nil
	}
	return fmt.Errorf("unknown encoding %q", enc)
}

func (m *parseResultsManager) SetSrcEncoding(enc string) error {
	if e, ok := encodings[strings.ToUpper(enc)]; ok {
		m.srcEncoding = e
		return nil
	}
	return fmt.Errorf("unknown encoding %q", enc)
}

func (m *parseResultsManager) CountSymbols() int64 {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	n := int64(0)
	for _, p := range m.parseResults {
		n += p.CountSymbols()
	}

	return n
}

func (m *parseResultsManager) WalkGlobalSymbols(walkFn func(symbol.Symbol) error, types SymbolType) error {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	for _, p := range m.parseResults {
		err := p.WalkGlobalSymbols(func(s symbol.Symbol) error {
			return walkFn(s)
		}, types)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *parseResultsManager) LookupGlobalSymbol(name string, types SymbolType) (symbol.Symbol, bool) {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	for _, p := range m.parseResults {
		if s, ok := p.LookupGlobalSymbol(name, types); ok {
			return s, true
		}
	}

	return nil, false
}

func (m *parseResultsManager) GetGlobalSymbols(types SymbolType) ([]symbol.Symbol, error) {
	result := make([]symbol.Symbol, 0, 200)

	m.mtx.RLock()
	defer m.mtx.RUnlock()

	for _, p := range m.parseResults {
		p.WalkGlobalSymbols(func(s symbol.Symbol) error {
			result = append(result, s)
			return nil
		}, types)
	}

	return result, nil
}

func (m *parseResultsManager) Get(documentURI string) (*ParseResult, error) {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	if r, ok := m.parseResults[documentURI]; ok {
		return r, nil
	}
	return nil, fmt.Errorf("document %q not found", documentURI)
}

func (m *parseResultsManager) Update(documentURI, content string) (*ParseResult, error) {
	r := m.ParseAndValidateScript(documentURI, content)

	m.mtx.Lock()
	defer m.mtx.Unlock()

	m.parseResults[documentURI] = r
	return r, nil
}

func (m *parseResultsManager) resolveSrcPaths(srcFile, prefixDir string) ([]string, error) {
	fileBytes, err := os.ReadFile(srcFile)
	if err != nil {
		return nil, err
	}
	decodedContent, err := m.srcEncoding.NewDecoder().Bytes(fileBytes)
	if err != nil {
		return nil, err
	}
	srcContent := string(decodedContent)

	// Auronen: On Linux and macOS replace backslashes with forwardslashes
	if runtime.GOOS != "windows" {
		srcContent = strings.ReplaceAll(srcContent, "\\", string(os.PathSeparator))
	}

	scanner := bufio.NewScanner(strings.NewReader(srcContent))

	resolvedPaths := []string{}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "//") {
			// likely a comment, like "// LeGo" or "// Ikarus"
			m.logger.Debugf("resolveSrcPaths(%q): skipping line %q", srcFile, line)
			continue
		}
		dir := filepath.Dir(line)
		fname := filepath.Base(line)
		ext := strings.ToLower(filepath.Ext(fname))

		if ext == ".d" {
			if strings.Contains(fname, "*") {
				files, err := ResolvePathsCaseInsensitive(filepath.Join(prefixDir, dir), fname)
				if err != nil {
					m.logger.Warnf("resolveSrcPaths(%q): could not resolve filename %q", srcFile, filepath.Join(prefixDir, dir, fname))
					continue
				}
				for _, e := range files {
					if stat, err := os.Stat(e); err != nil || stat.IsDir() {
						m.logger.Debugf("resolveSrcPaths(%q): ignoring entry %q (error or director)", srcFile, e)
						continue
					}

					resolvedPaths = append(resolvedPaths, e)
				}
			} else {
				found, err := findPath(filepath.Join(prefixDir, dir, fname))
				if err != nil {
					m.logger.Warnf("resolveSrcPaths(%q): file not found on disk %q", srcFile, filepath.Join(prefixDir, dir, fname))
					continue
				}
				if stat, err := os.Stat(found); err != nil || stat.IsDir() {
					m.logger.Debugf("resolveSrcPaths(%q): ignoring entry %q (error or director)", srcFile, found)
					continue
				}
				resolvedPaths = append(resolvedPaths, found)
			}
		} else if ext == ".src" {
			found, err := findPath(filepath.Join(prefixDir, dir, fname))
			if err != nil {
				m.logger.Warnf("resolveSrcPaths(%q): file not found on disk %q", srcFile, filepath.Join(prefixDir, dir, fname))
				continue
			}
			m.logger.Infof("Collecting scripts from %q", found)

			inner, err := m.resolveSrcPaths(found, filepath.Dir(found))
			if err == nil {
				resolvedPaths = append(resolvedPaths, inner...)
			}
		}
	}

	return resolvedPaths, nil
}

var (
	bufferPool = sync.Pool{New: func() interface{} { b := new(bytes.Buffer); b.Grow(4096); return b }}
)

func (m *parseResultsManager) getConcurrency() int {
	max := runtime.NumCPU()

	if m.NumParserThreads <= max && m.NumParserThreads > 0 {
		return m.NumParserThreads
	}

	numWorkers := max
	if numWorkers > 2 {
		numWorkers = numWorkers / 2
	}

	return numWorkers
}

func (m *parseResultsManager) validateFiles(resolvedPaths []string) map[string][]SyntaxError {
	results := make(map[string][]SyntaxError)

	chanPaths := make(chan string, len(resolvedPaths))
	for _, r := range resolvedPaths {
		chanPaths <- r
	}
	close(chanPaths)

	var wg sync.WaitGroup

	numWorkers := m.getConcurrency()

	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()

			buf := bufferPool.Get().(*bytes.Buffer)
			defer bufferPool.Put(buf)

			decoder := m.decoderPool.Get().(*encoding.Decoder)
			defer m.decoderPool.Put(decoder)

			for r := range chanPaths {
				f, err := os.Open(r)
				if err != nil {
					continue
				}
				decoder.Reset()
				translated := decoder.Reader(f)
				buf.Reset()
				_, err = buf.ReadFrom(translated)
				f.Close()
				if err != nil {
					continue
				}

				parsed := m.ValidateScript(r, buf.String())
				if len(parsed) > 0 {
					m.mtx.Lock()
					results[r] = parsed
					m.mtx.Unlock()
				}
			}
		}(&wg)
	}

	wg.Wait()
	return results
}

func (m *parseResultsManager) validateFile(dPath string) ([]SyntaxError, error) {
	buf := bufferPool.Get().(*bytes.Buffer)
	defer bufferPool.Put(buf)

	decoder := m.decoderPool.Get().(*encoding.Decoder)
	defer m.decoderPool.Put(decoder)

	f, err := os.Open(dPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	decoder.Reset()
	translated := decoder.Reader(f)
	buf.Reset()
	_, err = buf.ReadFrom(translated)
	if err != nil {
		return nil, err
	}
	return m.ValidateScript(dPath, buf.String()), nil
}

func (m *parseResultsManager) ParseSource(srcFile string) ([]*ParseResult, error) {
	resolvedPaths, err := m.resolveSrcPaths(srcFile, filepath.Dir(srcFile))
	if err != nil {
		return nil, err
	}
	m.logger.Infof("Parsing %q. This might take a while.", srcFile)

	results := make([]*ParseResult, 0, len(resolvedPaths))

	chanPaths := make(chan string, len(resolvedPaths))
	for _, r := range resolvedPaths {
		chanPaths <- r
	}
	close(chanPaths)

	var wg sync.WaitGroup
	numWorkers := m.getConcurrency()
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()

			buf := bufferPool.Get().(*bytes.Buffer)
			defer bufferPool.Put(buf)

			decoder := m.decoderPool.Get().(*encoding.Decoder)
			defer m.decoderPool.Put(decoder)

			for r := range chanPaths {
				f, err := os.Open(r)
				if err != nil {
					continue
				}
				decoder.Reset()
				translated := decoder.Reader(f)
				buf.Reset()
				_, err = buf.ReadFrom(translated)
				f.Close()
				if err != nil {
					continue
				}

				parsed := m.ParseScript(r, buf.String())

				m.mtx.Lock()
				m.parseResults[parsed.Source] = parsed
				results = append(results, parsed)
				m.mtx.Unlock()
			}
		}(&wg)
	}

	wg.Wait()

	validations := m.validateFiles(resolvedPaths)

	for _, v := range results {
		if e, ok := validations[v.Source]; ok {
			v.SyntaxErrors = append(v.SyntaxErrors, e...)
		}
	}

	m.logger.Infof("Done parsing %q: %d scripts.", srcFile, len(results))
	return results, nil
}

func (m *parseResultsManager) ParseFile(dFile string) (*ParseResult, error) {
	buf := bufferPool.Get().(*bytes.Buffer)
	defer bufferPool.Put(buf)

	decoder := m.decoderPool.Get().(*encoding.Decoder)
	defer m.decoderPool.Put(decoder)

	f, err := os.Open(dFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	decoder.Reset()
	translated := decoder.Reader(f)
	buf.Reset()
	_, err = buf.ReadFrom(translated)
	if err != nil {
		return nil, err
	}

	parsed := m.ParseScript(dFile, buf.String())

	m.mtx.Lock()
	m.parseResults[parsed.Source] = parsed
	m.mtx.Unlock()

	validations, err := m.validateFile(parsed.Source)

	if err == nil && len(validations) > 0 {
		parsed.SyntaxErrors = append(parsed.SyntaxErrors, validations...)
	}
	return parsed, nil
}

func (h *parseResultsManager) resolveIntConstant(c string) int {
	n, err := strconv.Atoi(c)
	if err == nil {
		return n
	}
	// probably symbol...?
	found, ok := h.LookupGlobalSymbol(c, SymbolConstant)
	if !ok {
		return -1
	}
	cs, ok := found.(symbol.Constant)
	if !ok {
		return -1
	}
	if n, err := strconv.Atoi(cs.Value); err == nil {
		return n
	}
	return -1
}

func (h *parseResultsManager) getSymbolCode(s symbol.Symbol) string {
	var codeText string
	if cas, ok := s.(symbol.ConstantArray); ok {
		sb := strings.Builder{}
		resolvedSize := h.resolveIntConstant(cas.ArraySizeText)
		cas.Format(&sb, resolvedSize)
		codeText = sb.String()
	} else if cas, ok := s.(symbol.ArrayVariable); ok {
		sb := strings.Builder{}
		resolvedSize := h.resolveIntConstant(cas.ArraySizeText)
		cas.Format(&sb, resolvedSize)
		codeText = sb.String()
	} else {
		codeText = s.String()
	}
	return codeText
}
