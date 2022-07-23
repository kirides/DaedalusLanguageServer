package langserver

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"

	lsp "go.lsp.dev/protocol"
)

var (
	rxFunctionDef  = regexp.MustCompile(`^\s*func\s+`)
	rxStringValues = regexp.MustCompile(`(".*? "|'.*?')`)
	rxFuncCall     = regexp.MustCompile(`\([\w@^_,:\/"'=\s\[\]]*\)`)
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
	if center < 0 {
		center = 0
	}
	start := center
	end := center
	if start >= len(doc) {
		return ""
	}
	if end >= len(doc) {
		end = len(doc)
	}
	for start >= 0 && isIdentifier(doc[start]) {
		start--
	}
	for end < len(doc) && isIdentifier(doc[end]) {
		end++
	}
	if start < center {
		start++ // skip the first bad character
	}
	return doc[start : start+(end-start)]
}

func (m BufferedDocument) GetIdentifier(position lsp.Position) (partial string, err error) {
	doc := string(m)
	c := doc
	currentLine := 0
	offset := 0
	line := int(position.Line)
	startOfLine := 0
	for currentLine < line && offset < len(doc) {
		currentLine++
		lineEnd := strings.IndexRune(c, '\n')
		if lineEnd == -1 {
			break
		}
		offset += lineEnd
		if len(c) < lineEnd+1 {
			break
		}
		offset++
		c = c[lineEnd+1:]
	}

	startOfLine = offset
	if position.Character > 0 {
		offset += int(position.Character)
	}

	lineContent := doc[startOfLine:offset]

	end := len(lineContent)
	o := end - 1

	for o >= 0 {
		token := lineContent[o]
		r := rune(token)
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' {
			o--
			continue
		}
		break
	}

	return lineContent[o+1 : end], nil
}

func (m BufferedDocument) GetParentSymbolReference(position lsp.Position) (parent string, partial string, err error) {
	doc := string(m)
	c := doc
	currentLine := 0
	offset := 0
	line := int(position.Line)
	startOfLine := 0
	for currentLine < line && offset < len(doc) {
		currentLine++
		lineEnd := strings.IndexRune(c, '\n')
		if lineEnd == -1 {
			break
		}
		offset += lineEnd
		if len(c) < lineEnd+1 {
			break
		}
		offset++
		c = c[lineEnd+1:]
	}

	startOfLine = offset
	if position.Character > 0 {
		offset += int(position.Character)
	}

	lineContent := doc[startOfLine:offset]

	o := len(lineContent) - 1
	start := o

	foundDot := false
	for o >= 0 {
		token := lineContent[o]
		r := rune(token)
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' {
			o--
			continue
		}
		if r == '.' {
			foundDot = true
		}
		break
	}
	if !foundDot {
		return "", "", fmt.Errorf("not found")
	}
	idxDot := o
	o--
	foundIdentifier := false
	for o >= 0 {
		token := lineContent[o]
		r := rune(token)
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' {
			o--
			foundIdentifier = true
			continue
		}
		break
	}

	if foundIdentifier {
		o++
		return lineContent[o:idxDot], lineContent[idxDot:start], nil
	}
	return "", "", fmt.Errorf("not found")
}

// GetMethodCall ...
func (m BufferedDocument) GetMethodCall(position lsp.Position) string {
	doc := string(m)
	c := doc
	currentLine := 0
	offset := 0
	line := int(position.Line)

	for currentLine < line && offset < len(doc) {
		currentLine++
		lineEnd := strings.IndexRune(c, '\n')
		if lineEnd == -1 {
			break
		}
		offset += lineEnd
		if len(c) < lineEnd+1 {
			break
		}
		offset++
		c = c[lineEnd+1:]
	}

	if position.Character > 0 {
		offset += int(position.Character)
	}
	o := offset
	skipOpen := 1

	for o >= 0 {
		token := doc[o]
		o--
		if token == ';' || token == '}' || token == '{' {
			break
		} else if token == '(' && skipOpen <= 0 {
			break
		} else if token == ')' {
			skipOpen++
		} else if token == '(' {
			skipOpen--
		} else if token == '\n' || token == '\r' {
			continue
		}
	}
	if o+1 > len(doc) {
		return doc[o:]
	}
	st := o + 2
	if st >= len(doc) {
		st = len(doc) - 1
	}
	ed := st + (offset - o - 2)
	if ed < st {
		ed = st + 1
	}
	methodCallLine := doc[st:ed]
	for i := 0; i < len(methodCallLine); i++ {
		if !unicode.IsSpace(rune(methodCallLine[i])) {
			methodCallLine = methodCallLine[i:]
			break
		}
	}

	return methodCallLine
}

func isIdentifier(b byte) bool {
	if unicode.IsDigit(rune(b)) || unicode.IsLetter(rune(b)) {
		return true
	}
	return b == '_' || b == '@' || b == '^'
}
