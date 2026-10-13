package langserver

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/kirides/DaedalusLanguageServer/daedalus/parser"
)

func dumpToken(t antlr.Token) string {
	return fmt.Sprintf("type=%d ch=%d [%d..%d] %d:%d %q", t.GetTokenType(), t.GetChannel(), t.GetStart(), t.GetStop(), t.GetLine(), t.GetColumn(), t.GetText())
}

func referenceTokens(src string) []string {
	lexer := parser.NewDaedalusLexer(antlr.NewInputStream(src))
	lexer.RemoveErrorListeners()
	var res []string
	for {
		t := lexer.NextToken()
		res = append(res, dumpToken(t))
		if t.GetTokenType() == antlr.TokenEOF {
			return res
		}
	}
}

// compareLexers returns whether the fast path was used.
func compareLexers(t *testing.T, fl *fastLexer, src string) bool {
	t.Helper()
	fl.RemoveErrorListeners()
	fl.Reset(src, antlr.NewInputStream(src))
	defer fl.Release()
	used := fl.fast
	if !used {
		return false
	}
	want := referenceTokens(src)
	var got []string
	for {
		tk := fl.NextToken()
		got = append(got, dumpToken(tk))
		if tk.GetTokenType() == antlr.TokenEOF {
			break
		}
	}
	if len(got) != len(want) {
		t.Fatalf("token count mismatch for %q: got %d, want %d\ngot:  %q\nwant: %q", src, len(got), len(want), got, want)
	}
	for i := range got {
		if got[i] != want[i] {
			t.Fatalf("token %d mismatch for %q:\ngot:  %s\nwant: %s", i, src, got[i], want[i])
		}
	}
	return true
}

func TestFastLexerCorpus(t *testing.T) {
	scripts, _ := loadBenchScripts(t)
	fl := newFastLexer()
	for _, s := range scripts {
		if !compareLexers(t, fl, s.content) {
			t.Errorf("%s: fast lexer fell back to the generated lexer", s.name)
		}
	}
}

func TestFastLexerRandom(t *testing.T) {
	frags := []string{
		"1", "23", "007", ".", "e", "E", "+", "-", "5", "1.5", "1.", ".5", "1e5", "1E+5", "1e", "2.e-3", "abc", "_x", "x1", "1x", "@", "^", "a@b", "ß", "ä", "Ö", "ü", "Ä", "é", "€",
		"\"str\"", "\"", "\"a//b\"", "\"ä\"", "//c\n", "// ö\r\n", "///doc", "/*c*/", "/* a\nb */", "/*", "*/", "/", "*", "=", "==", "!", "!=", "<", "<<", "<=", ">", ">>", ">=",
		"&", "&&", "&=", "|", "||", "|=", "+=", "-=", "*=", "/=", ";", ",", "%", "~", "(", ")", "{", "}", "[", "]",
		" ", "  ", "\t", "\r\n", "\n", "\r", "\x00", "\x0c", "\ufeff",
		"if", "IF", "Const", "var", "VAR", "string", "func", "instance", "prototype", "namespace", "null", "meta", "float", "void", "return", "class", "else", "int", "iff", "constx", "Int1",
	}
	rng := rand.New(rand.NewSource(1))
	fl := newFastLexer()
	fast := 0
	const iterations = 20000
	for n := 0; n < iterations; n++ {
		var sb strings.Builder
		for k := rng.Intn(14) + 1; k > 0; k-- {
			sb.WriteString(frags[rng.Intn(len(frags))])
			if rng.Intn(3) == 0 {
				sb.WriteByte(' ')
			}
		}
		if compareLexers(t, fl, sb.String()) {
			fast++
		}
	}
	t.Logf("fast path used for %d of %d inputs", fast, iterations)
	if fast < iterations/4 {
		t.Errorf("fast path used too rarely (%d of %d)", fast, iterations)
	}
}
