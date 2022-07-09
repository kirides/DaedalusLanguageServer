package langserver

import (
	"bufio"
	"bytes"
	"strings"
)

var (
	javadocParam  = []byte("@param")
	javadocReturn = []byte("@return")
)

type javadoc struct {
	Summary    string
	Parameters string
	Return     string
}

func appendMarkdownEscaped(sb *strings.Builder, text string) {
	for _, v := range text {
		if v == '*' {
			sb.WriteRune('\\')
		}
		sb.WriteRune(v)
	}
}
func formatParams(sb *strings.Builder, param, desc string) {
	sb.WriteString("- **")
	appendMarkdownEscaped(sb, param)
	sb.WriteString("** - *")
	appendMarkdownEscaped(sb, desc)
	sb.WriteString("*\n")
}

func parseJavadocMdEscaped(sym Symbol) javadoc {
	r := javadoc{
		Summary: sym.Documentation(),
	}
	fn, ok := sym.(FunctionSymbol)
	if !ok {
		return r
	}

	summary := strings.Builder{}
	params := strings.Builder{}
	returns := strings.Builder{}

	scn := bufio.NewScanner(strings.NewReader(sym.Documentation()))

	for scn.Scan() {
		line := bytes.TrimSpace(scn.Bytes())

		if bytes.HasPrefix(line, javadocReturn) {
			line = bytes.TrimSpace(bytes.TrimPrefix(line, javadocReturn))
			appendMarkdownEscaped(&returns, string(line))
		} else if bytes.HasPrefix(line, javadocParam) {
			param, desc, ok := parseJavadocParam(string(line))
			if !ok {
				continue
			}
			for _, p := range fn.Parameters {
				if strings.EqualFold(p.Name(), param) {
					formatParams(&params, param, desc)
					break
				}
			}
		} else {
			appendMarkdownEscaped(&summary, string(line))
			summary.WriteString("  \n")
		}
	}
	r.Summary = strings.TrimSpace(summary.String())
	r.Parameters = strings.TrimSpace(params.String())
	r.Return = strings.TrimSpace(returns.String())
	return r
}

func simpleJavadocMD(sym Symbol) string {
	doc := parseJavadocMdEscaped(sym)
	return javadocMD(doc)
}

func javadocMD(doc javadoc) string {
	sb := strings.Builder{}

	sb.WriteString(doc.Summary)
	sb.WriteString("\n")
	sb.WriteString("\n")
	sb.WriteString(doc.Parameters)
	if doc.Return != "" {
		sb.WriteString("\n\n*returns ")
		sb.WriteString(doc.Return)
		sb.WriteString("*")
	}

	return strings.TrimSpace(sb.String())
}

func parseJavadocParam(line string) (param, desc string, ok bool) {
	line = strings.TrimSpace(line[7:])

	if len(line) < 8 {
		// minimum "@param a"
		return "", "", false
	}

	param, desc, ok = strings.Cut(line, " ")

	param = strings.TrimSpace(param)
	desc = strings.TrimSpace(desc)
	return
}

func findJavadocParam(doc, key string) string {
	scn := bufio.NewScanner(strings.NewReader(doc))

	for scn.Scan() {
		line := bytes.TrimSpace(scn.Bytes())
		if !bytes.HasPrefix(line, javadocParam) {
			continue
		}
		line = bytes.TrimSpace(bytes.TrimPrefix(line, javadocParam))
		param, doc, ok := bytes.Cut(line, []byte(" "))

		if ok && strings.EqualFold(key, string(param)) {
			return string(bytes.TrimSpace(doc))
		}
	}

	return ""
}

func parseJavadocWithinTokens(doc, open, close string) (instances []string, remaining string) {

	idxOpen := strings.Index(doc, open)
	if idxOpen == -1 {
		return nil, doc
	}

	idxClose := strings.Index(doc, close)
	if idxClose == -1 || idxClose < idxOpen {
		return nil, doc
	}

	rem := doc[idxClose+1:]

	instances = strings.Split(doc[idxOpen+1:idxClose], ",")
	for i := 0; i < len(instances); i++ {
		instances[i] = strings.TrimSpace(instances[i])
	}

	return instances, strings.TrimSpace(rem)
}
