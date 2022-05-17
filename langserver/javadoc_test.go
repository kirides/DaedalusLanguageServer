package langserver

import (
	"strings"
	"testing"
)

func TestJavadocParam(t *testing.T) {

	s := NewFunctionSymbol("Fn", "", `@param p decides things
	`,
		SymbolDefinition{},
		"void",
		SymbolDefinition{},
		[]VariableSymbol{NewVariableSymbol("p", "int", "", "", SymbolDefinition{})},
		[]Symbol{})

	jd := simpleJavadocMD(s)

	if !strings.Contains(jd, "- **p** - *decides things*") {
		t.Fatalf("expected markdown property content, actual: %s", jd)
	}
}

func TestJavadocParamWithParser(t *testing.T) {

	script := `
		/// Something
		///
		/// @param amount the amount
		func void InitDamage(var int amount) {};`

	m := newParseResultsManager(nopLogger{})
	result := m.ParseScript("C:\\temp", script)

	fn := result.Functions["INITDAMAGE"]

	jd := simpleJavadocMD(fn)

	if !strings.Contains(jd, "- **amount** - *the amount*") {
		t.Fatalf("expected markdown property content, actual: %s", jd)
	}
}
