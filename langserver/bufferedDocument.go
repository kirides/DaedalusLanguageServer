package langserver

import (
	"strings"
	"unicode"

	lsp "go.lsp.dev/protocol"
)

// BufferedDocument ...
type BufferedDocument string

// GetWordRangeAtPosition ...
func (m BufferedDocument) GetWordRangeAtPosition(position lsp.Position) string {
	if position.Character < 2 && position.Line < 1 {
		return ""
	}
	line := 0
	offset := 0
	wordLine := int(position.Line)
	doc := string(m)
	for line < wordLine {
		line++
		lineEnd := strings.IndexRune(doc, '\n')
		if lineEnd != -1 {
			offset += lineEnd
			if len(doc) < lineEnd+1 {
				break
			}
			offset++
			doc = doc[lineEnd+1:]
		}
	}
	center := int(position.Character) - 1
	start := center
	end := center

	for isIdentifier(doc[start]) {
		start--
	}
	for isIdentifier(doc[end]) {
		end++
	}
	if start < center {
		start++ // skip the first bad character
	}
	return doc[start : start+(end-start)]
}

func isIdentifier(b byte) bool {
	if unicode.IsDigit(rune(b)) || unicode.IsLetter(rune(b)) {
		return true
	}
	return b == '_' || b == '@' || b == '^'
}
