package langserver

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/kirides/DaedalusLanguageServer/daedalus/parser"
)

type recordingErrorListener struct {
	*antlr.DefaultErrorListener
	errors []string
}

func (l *recordingErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	l.errors = append(l.errors, fmt.Sprintf("%d:%d %s", line, column, msg))
}

type hiddenTokenReader interface {
	GetHiddenTokensToLeft(tokenIndex, channel int) []antlr.Token
	GetHiddenTokensToRight(tokenIndex, channel int) []antlr.Token
}

type pipelineResult struct {
	fast   bool
	tree   string
	errors []string
	hidden []string
}

func describeHidden(ts antlr.TokenStream) []string {
	h := ts.(hiddenTokenReader)
	var res []string
	for i := 0; i < ts.Size(); i++ {
		for _, side := range []func(int, int) []antlr.Token{h.GetHiddenTokensToLeft, h.GetHiddenTokensToRight} {
			var sb strings.Builder
			for _, t := range side(i, parser.COMMENTS) {
				sb.WriteString(t.GetText())
				sb.WriteByte('|')
			}
			res = append(res, sb.String())
		}
	}
	return res
}

func runPipeline(src string, reference bool, fl *fastLexer) pipelineResult {
	in := antlr.NewInputStream(src) // reference only
	errs := &recordingErrorListener{DefaultErrorListener: antlr.NewDefaultErrorListener()}

	var ts antlr.TokenStream
	var lex antlr.Lexer
	if reference {
		l := parser.NewDaedalusLexer(in)
		l.RemoveErrorListeners()
		lex = l
		ts = antlr.NewCommonTokenStream(l, 0)
	} else {
		fl.RemoveErrorListeners()
		fl.ResetString(src)
		defer fl.Release()
		ts = fl.TokenStream()
	}
	_ = lex
	p := parser.NewDaedalusParser(ts)
	p.RemoveErrorListeners()
	p.AddErrorListener(errs)
	p.GetInterpreter().SetPredictionMode(antlr.PredictionModeSLL)
	tree := p.DaedalusFile()
	return pipelineResult{
		fast:   !reference && fl.fast,
		tree:   antlr.TreesStringTree(tree, nil, p),
		errors: errs.errors,
		hidden: describeHidden(ts),
	}
}

var pipelineStats struct{ total, fast, withErrors int }

func comparePipelines(t *testing.T, fl *fastLexer, name, src string) {
	t.Helper()
	want := runPipeline(src, true, nil)
	got := runPipeline(src, false, fl)
	pipelineStats.total++
	if got.fast {
		pipelineStats.fast++
	}
	if len(want.errors) > 0 {
		pipelineStats.withErrors++
	}
	if got.tree != want.tree {
		t.Fatalf("%s: parse tree differs\ngot:  %.400s\nwant: %.400s", name, got.tree, want.tree)
	}
	if strings.Join(got.errors, "\n") != strings.Join(want.errors, "\n") {
		t.Fatalf("%s: syntax errors differ\ngot:  %q\nwant: %q", name, got.errors, want.errors)
	}
	if strings.Join(got.hidden, "\n") != strings.Join(want.hidden, "\n") {
		t.Fatalf("%s: hidden tokens differ", name)
	}
}

func TestFastPipelineMatchesGenerated(t *testing.T) {
	scripts, _ := loadBenchScripts(t)
	fl := newFastLexer()
	rng := rand.New(rand.NewSource(42))

	for _, s := range scripts {
		comparePipelines(t, fl, s.name, s.content)

		// damaged variants to exercise error recovery
		for n := 0; n < 40 && len(s.content) > 20; n++ {
			b := []byte(s.content)
			switch rng.Intn(4) {
			case 0: // delete a range
				a := rng.Intn(len(b))
				e := min(len(b), a+1+rng.Intn(12))
				b = append(b[:a:a], b[e:]...)
			case 1: // duplicate a range
				a := rng.Intn(len(b))
				e := min(len(b), a+1+rng.Intn(12))
				b = append(b[:e:e], append(append([]byte{}, b[a:e]...), b[e:]...)...)
			case 2: // insert a random punctuation
				a := rng.Intn(len(b))
				p := []string{";", "(", ")", "{", "}", ",", "=", "var", "if", "\"", "/*", "."}[rng.Intn(12)]
				b = append(b[:a:a], append([]byte(p), b[a:]...)...)
			case 3: // truncate
				b = b[:rng.Intn(len(b))]
			}
			comparePipelines(t, fl, fmt.Sprintf("%s (variant %d)", s.name, n), string(b))
		}
	}
	t.Logf("compared %d inputs: %d lexed on the fast path, %d with syntax errors", pipelineStats.total, pipelineStats.fast, pipelineStats.withErrors)
}
