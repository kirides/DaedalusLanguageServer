package javadoc_test

import (
	"strings"
	"testing"

	"github.com/kirides/DaedalusLanguageServer/daedalus/symbol"
	"github.com/kirides/DaedalusLanguageServer/langserver"
	"github.com/kirides/DaedalusLanguageServer/langserver/javadoc"
)

func TestJavadocParam(t *testing.T) {

	s := symbol.NewFunction("Fn", "", `@param p decides things
	`,
		symbol.Definition{},
		"void",
		symbol.Definition{},
		[]symbol.Variable{symbol.NewVariable("p", "int", "", "", symbol.Definition{})},
		[]symbol.Symbol{})

	jd := javadoc.MarkdownSimple(s)

	if !strings.Contains(jd, "- **p** - *decides things*") {
		t.Fatalf("expected markdown property content, actual: %s", jd)
	}
}

func TestJavadocParamWithParser(t *testing.T) {
	result := &langserver.ParseResult{
		Functions: map[string]symbol.Function{
			"INITDAMAGE": symbol.NewFunction(
				"INITDAMAGE",
				".",
				"Something\n\n@param amount the amount",
				symbol.Definition{},
				"void",
				symbol.Definition{},
				[]symbol.Variable{},
				[]symbol.Symbol{},
			),
		},
	}

	fn := result.Functions["INITDAMAGE"]

	jd := javadoc.MarkdownSimple(fn)

	if !strings.Contains(jd, "- **amount** - *the amount*") {
		t.Fatalf("expected markdown property content, actual: %s", jd)
	}
}
