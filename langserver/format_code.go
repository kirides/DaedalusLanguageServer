package langserver

import (
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/kirides/DaedalusLanguageServer/daedalus/parser"
)

func FormatCode(code string) {
	inputStream := antlr.NewInputStream(code)
	lexer := parser.NewDaedalusLexer(inputStream)

	sb := strings.Builder{}
	state := formatState{}
	for token := lexer.NextToken(); token.GetTokenType() != antlr.TokenEOF; token = lexer.NextToken() {
		txt := token.GetText()

		writeToken(txt, &sb, &state)
		updateState(txt, &sb, &state)

		if strings.EqualFold(txt, "func") {
			formatFunc(lexer, &sb)
		} else if strings.EqualFold(txt, "class") {
			formatClass(lexer, &sb)
		} else if strings.EqualFold(txt, "instance") {
			formatClass(lexer, &sb)
		} else if strings.EqualFold(txt, "prototype") {
			formatClass(lexer, &sb)
		}

	}
	final := strings.TrimSpace(sb.String())

	_ = final
}

type formatState struct {
	IsNewLine  bool
	Indent     int
	PrintSpace bool
}

func writeToken(txt string, sb *strings.Builder, state *formatState) {
	if state.IsNewLine {
		state.IsNewLine = false
		sb.WriteString("\r\n")

		if txt != "}" {
			for i := 0; i < state.Indent; i++ {
				sb.WriteString("    ")
			}
		}
	} else if strings.ContainsAny(txt, "();") {

	} else if state.PrintSpace && txt != "," {
		sb.WriteString(" ")
	}

	sb.WriteString(txt)
}

func updateState(txt string, sb *strings.Builder, state *formatState) {
	if strings.TrimSpace(sb.String()) == "};" {
		state.IsNewLine = true
	}

	state.PrintSpace = !strings.ContainsAny(txt, "(")
	if txt == "{" {
		state.IsNewLine = true
	}

	switch {
	case txt == "{":
		state.Indent++
	case txt == "}":
		state.Indent--
		if state.Indent < 0 {
			state.Indent = 0
		}
	}
}

func formatFunc(lexer antlr.Lexer, sb *strings.Builder) {
	state := formatState{PrintSpace: true}
	init := true
	for token := lexer.NextToken(); token.GetTokenType() != antlr.TokenEOF; token = lexer.NextToken() {
		txt := token.GetText()

		writeToken(txt, sb, &state)
		updateState(txt, sb, &state)

		if state.Indent != 0 {
			init = false
		}

		if state.Indent == 0 && !init {
			break
		}
	}
}

func formatClass(lexer antlr.Lexer, sb *strings.Builder) {
	state := formatState{PrintSpace: true}
	init := true
	for token := lexer.NextToken(); token.GetTokenType() != antlr.TokenEOF; token = lexer.NextToken() {
		txt := token.GetText()

		writeToken(txt, sb, &state)
		updateState(txt, sb, &state)

		if txt == ";" {
			state.IsNewLine = true
		}

		if state.Indent != 0 {
			init = false
		}

		if state.Indent == 0 && !init {
			break
		}
	}
}
