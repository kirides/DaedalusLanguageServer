package javadoc

import (
	"bufio"
	"bytes"
	"strings"

	"github.com/kirides/DaedalusLanguageServer/daedalus/symbol"
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

	const (
		PREFIX_INST = "{"
		PREFIX_ENUM = "["
		PREFIX_FUNC = "<"

		INST_OPEN  = PREFIX_INST
		INST_CLOSE = "}"

		ENUM_OPEN  = PREFIX_ENUM
		ENUM_CLOSE = "]"

		FUNC_OPEN  = PREFIX_FUNC
		FUNC_CLOSE = ">"
	)

	if strings.HasPrefix(desc, PREFIX_INST) || strings.HasPrefix(desc, PREFIX_ENUM) || strings.HasPrefix(desc, PREFIX_FUNC) {
		insts, desc := ParseWithinDedup(desc, INST_OPEN, INST_CLOSE)
		enums, desc := ParseWithinDedup(desc, ENUM_OPEN, ENUM_CLOSE)
		fnSigdata, desc := ParseWithinDedup(desc, FUNC_OPEN, FUNC_CLOSE)

		appendMarkdownEscaped(sb, desc)
		sb.WriteString("*  \n")

		if len(insts) != 0 {
			sb.WriteString("Types: ")
			for i := 0; i < len(insts); i++ {
				insts[i] = "`" + insts[i] + "`"
			}
			sb.WriteString(strings.Join(insts, ", "))
			sb.WriteString("\n")
		}
		if len(enums) != 0 {
			sb.WriteString("Valid Values: ")
			for i := 0; i < len(enums); i++ {
				enums[i] = "`" + enums[i] + "`"
			}
			sb.WriteString(strings.Join(enums, ", "))
			sb.WriteString("\n")
		}
		if len(fnSigdata) != 0 {
			sb.WriteString("Required function signature: ")
			sig, err := getFuncSignatureString(fnSigdata)
			if err == nil {
				// I wanted daedalus syntax highlightin here, but it does not work for me
				sb.WriteString("```daedalus\n" + sig + "\n```")
			} else {
				sb.WriteString("<invalid signature>")
			}
			sb.WriteString("\n")
		}
	} else {
		appendMarkdownEscaped(sb, desc)
		sb.WriteString("*\n")
	}
}

// RemoveTokens removes all javadoc tokens from a text and makes it plain text
func RemoveTokens(desc string) string {
	if strings.HasPrefix(desc, "{") || strings.HasPrefix(desc, "[") || strings.HasPrefix(desc, "<") {
		_, desc = ParseWithinDedup(desc, "{", "}")
		_, desc = ParseWithinDedup(desc, "[", "]")
		_, desc = ParseWithinDedup(desc, "<", ">")
	}
	return desc
}

func parseJavadocMdEscaped(sym symbol.Symbol) javadoc {
	r := javadoc{
		Summary: sym.Documentation(),
	}
	fn, ok := sym.(symbol.Function)
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

func MarkdownSimple(sym symbol.Symbol) string {
	doc := parseJavadocMdEscaped(sym)
	return Markdown(doc)
}

func Markdown(doc javadoc) string {
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

func FindParam(doc, key string) string {
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

// dedupI returns a slice without any duplicates (insensitive)
func dedupI(in []string) []string {
	var result []string
	for _, v := range in {
		add := true
		for _, existing := range result {
			if strings.EqualFold(existing, v) {
				add = false
				break
			}
		}
		if add {
			result = append(result, v)
		}
	}
	return result
}

func ParseWithinDedup(doc, open, close string) (instances []string, remaining string) {

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

	return dedupI(instances), strings.TrimSpace(rem)
}
func ParseWithin(doc, open, close string) (instances []string, remaining string) {

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
