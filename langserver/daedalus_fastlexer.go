package langserver

import (
	"strings"
	"unicode/utf8"

	"github.com/antlr4-go/antlr/v4"
	"github.com/kirides/DaedalusLanguageServer/daedalus/parser"
)

// fastLexer is a hand-written replacement for the generated DaedalusLexer.
//
// It tokenizes the whole input eagerly and produces exactly the tokens the
// generated lexer would (type, channel, start/stop, line, column and text).
// As soon as it encounters anything it can not reproduce with certainty
// (unknown characters, unterminated strings or block comments, invalid
// UTF-8), it discards its result and delegates to the generated lexer for the
// complete input, so error reporting and recovery behave exactly as before.
//
// The generated lexer is embedded to satisfy antlr.TokenSource (which has an
// unexported method) and to serve as the fallback.
type fastLexer struct {
	*parser.DaedalusLexer

	tokens []*fastToken
	slab   tokenSlab
	pos    int
	fast   bool
}

var _ antlr.TokenSource = (*fastLexer)(nil)

// maxPooledTokenCap bounds the token buffer retained between parses.
const maxPooledTokenCap = 1 << 18

func newFastLexer() *fastLexer {
	return &fastLexer{DaedalusLexer: parser.NewDaedalusLexer(nil)}
}

// ResetString prepares the lexer for src, without converting it to runes unless necessary.
func (l *fastLexer) ResetString(src string) {
	l.Reset(src, newLazyCharStream(src))
}

// Reset prepares the lexer for src. input must be a CharStream over src.
func (l *fastLexer) Reset(src string, input antlr.CharStream) {
	l.DaedalusLexer.SetInputStream(input)
	l.pos = 0
	l.tokens = l.tokens[:0]
	l.slab.reset(len(src))
	l.fast = l.tokenize(src, input)
	if !l.fast {
		clear(l.tokens)
		l.tokens = l.tokens[:0]
	}
}

// TokenStream returns the token stream for the last input.
func (l *fastLexer) TokenStream() antlr.TokenStream {
	if l.fast {
		return newFastTokenStream(l.tokens, l)
	}
	return antlr.NewCommonTokenStream(l, antlr.TokenDefaultChannel)
}

// Release drops all references to the last input.
func (l *fastLexer) Release() {
	l.DaedalusLexer.SetInputStream(nil)
	clear(l.tokens)
	l.slab = tokenSlab{}
	if cap(l.tokens) > maxPooledTokenCap {
		l.tokens = nil
	} else {
		l.tokens = l.tokens[:0]
	}
	l.pos = 0
	l.fast = false
}

func (l *fastLexer) NextToken() antlr.Token {
	if !l.fast {
		return l.DaedalusLexer.NextToken()
	}
	if l.pos < len(l.tokens) {
		t := l.tokens[l.pos]
		l.pos++
		return t
	}
	// stay on EOF
	return l.tokens[len(l.tokens)-1]
}

func (l *fastLexer) GetLine() int {
	if !l.fast {
		return l.DaedalusLexer.GetLine()
	}
	if l.pos > 0 {
		return l.tokens[l.pos-1].GetLine()
	}
	return 1
}

func (l *fastLexer) GetCharPositionInLine() int {
	if !l.fast {
		return l.DaedalusLexer.GetCharPositionInLine()
	}
	if l.pos > 0 {
		return l.tokens[l.pos-1].GetColumn()
	}
	return 0
}

func isDigit(c byte) bool { return c >= '0' && c <= '9' }

func isIdentStartASCII(c byte) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c == '_'
}

func isIdentContinueASCII(c byte) bool {
	return isIdentStartASCII(c) || isDigit(c) || c == '@' || c == '^'
}

// isGermanChar reports whether r is one of ß Ä ä Ö ö Ü ü.
func isGermanChar(r rune) bool {
	switch r {
	case 0xDF, 0xC4, 0xE4, 0xD6, 0xF6, 0xDC, 0xFC:
		return true
	}
	return false
}

// scanIdent returns the byte length and rune count of the identifier
// characters (including digits) starting at src[i].
func scanIdent(src string, i int) (bytes, runes int) {
	j := i
	for j < len(src) {
		c := src[j]
		if c < utf8.RuneSelf {
			if !isIdentContinueASCII(c) {
				break
			}
			j++
			runes++
			continue
		}
		r, w := utf8.DecodeRuneInString(src[j:])
		if !isGermanChar(r) {
			break
		}
		j += w
		runes++
	}
	return j - i, runes
}

// exponentLen returns the length of an Exponent ([eE] [+-]? Digit+) at src[i], or 0.
func exponentLen(src string, i int) int {
	if i >= len(src) || (src[i] != 'e' && src[i] != 'E') {
		return 0
	}
	j := i + 1
	if j < len(src) && (src[j] == '+' || src[j] == '-') {
		j++
	}
	k := j
	for k < len(src) && isDigit(src[k]) {
		k++
	}
	if k == j {
		return 0
	}
	return k - i
}

func countDigits(src string, i int) int {
	j := i
	for j < len(src) && isDigit(src[j]) {
		j++
	}
	return j - i
}

// floatLen returns the length of the longest FloatLiteral at src[i], or 0.
func floatLen(src string, i int) int {
	d1 := countDigits(src, i)
	j := i + d1
	best := 0
	if j < len(src) && src[j] == '.' {
		d2 := countDigits(src, j+1)
		pf := 0
		switch {
		case d2 > 0:
			pf = d1 + 1 + d2
		case d1 > 0:
			pf = d1 + 1
		}
		if pf > 0 {
			best = pf
			if e := exponentLen(src, i+pf); e > 0 {
				best = pf + e
			}
		}
	}
	if d1 > 0 {
		if e := exponentLen(src, j); e > 0 && d1+e > best {
			best = d1 + e
		}
	}
	return best
}

func keywordType(s string) int {
	switch len(s) {
	case 2:
		if asciiEqualFold(s, "if") {
			return parser.DaedalusLexerIf
		}
	case 3:
		if asciiEqualFold(s, "var") {
			return parser.DaedalusLexerVar
		}
		if asciiEqualFold(s, "int") {
			return parser.DaedalusLexerInt
		}
	case 4:
		switch {
		case asciiEqualFold(s, "else"):
			return parser.DaedalusLexerElse
		case asciiEqualFold(s, "func"):
			return parser.DaedalusLexerFunc
		case asciiEqualFold(s, "void"):
			return parser.DaedalusLexerVoid
		case asciiEqualFold(s, "null"):
			return parser.DaedalusLexerNull
		case asciiEqualFold(s, "meta"):
			return parser.DaedalusLexerMeta
		}
	case 5:
		switch {
		case asciiEqualFold(s, "const"):
			return parser.DaedalusLexerConst
		case asciiEqualFold(s, "class"):
			return parser.DaedalusLexerClass
		case asciiEqualFold(s, "float"):
			return parser.DaedalusLexerFloat
		}
	case 6:
		switch {
		case asciiEqualFold(s, "string"):
			return parser.DaedalusLexerStringKeyword
		case asciiEqualFold(s, "return"):
			return parser.DaedalusLexerReturn
		}
	case 8:
		if asciiEqualFold(s, "instance") {
			return parser.DaedalusLexerInstance
		}
	case 9:
		switch {
		case asciiEqualFold(s, "prototype"):
			return parser.DaedalusLexerPrototype
		case asciiEqualFold(s, "namespace"):
			return parser.DaedalusLexerNamespace
		}
	}
	return parser.DaedalusLexerIdentifier
}

// asciiEqualFold compares s to the lowercase ASCII word lower.
func asciiEqualFold(s, lower string) bool {
	if len(s) != len(lower) {
		return false
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= 'A' && c <= 'Z' {
			c += 'a' - 'A'
		}
		if c != lower[i] {
			return false
		}
	}
	return true
}

// lexState tracks the position while tokenizing.
type lexState struct {
	src       string
	i         int // byte offset
	ri        int // rune index
	line, col int
}

// skipNoNewline advances over the bytes src[st.i:end] that contain no newlines,
// counting runes. It reports false on invalid UTF-8.
func (st *lexState) skipNoNewline(end int) bool {
	for st.i < end {
		c := st.src[st.i]
		if c < utf8.RuneSelf {
			st.i++
		} else {
			r, w := utf8.DecodeRuneInString(st.src[st.i:end])
			if r == utf8.RuneError && w <= 1 {
				return false
			}
			st.i += w
		}
		st.ri++
		st.col++
	}
	return true
}

// skipWithNewlines is like skipNoNewline, but tracks line breaks ('\n' only).
func (st *lexState) skipWithNewlines(end int) bool {
	for st.i < end {
		c := st.src[st.i]
		if c < utf8.RuneSelf {
			st.i++
			st.ri++
			if c == '\n' {
				st.line++
				st.col = 0
			} else {
				st.col++
			}
			continue
		}
		r, w := utf8.DecodeRuneInString(st.src[st.i:end])
		if r == utf8.RuneError && w <= 1 {
			return false
		}
		st.i += w
		st.ri++
		st.col++
	}
	return true
}

// tokenize fills l.tokens. It reports false if the input needs the generated lexer.
func (l *fastLexer) tokenize(src string, input antlr.CharStream) bool {
	info := &tokenInfo{pair: l.GetTokenSourceCharStreamPair(), src: l.DaedalusLexer, input: input}
	slab := &l.slab

	if cap(l.tokens) == 0 {
		l.tokens = make([]*fastToken, 0, len(src)/4+16)
	}
	toks := l.tokens[:0]
	emit := func(ttype int, text string, channel, start, stop, line, col int) {
		t := slab.next()
		*t = fastToken{
			text: text, info: info, start: int32(start), stop: int32(stop), line: int32(line), column: int32(col),
			index: int32(len(toks)), typ: int16(ttype), ch: int16(channel),
		}
		toks = append(toks, t)
	}

	st := lexState{src: src, line: 1}
	n := len(src)

	for st.i < n {
		c := src[st.i]
		start, rstart, line, col := st.i, st.ri, st.line, st.col

		ttype := 0
		channel := antlr.TokenDefaultChannel
		// Number of bytes / runes of the token for the simple cases.
		blen, rlen := 1, 1

		switch c {
		case ' ', '\t', '\r':
			st.i++
			st.ri++
			st.col++
			continue
		case '\n':
			st.i++
			st.ri++
			st.line++
			st.col = 0
			continue
		case '"':
			end := -1
			for j := st.i + 1; j < n; j++ {
				b := src[j]
				if b == '"' {
					end = j
					break
				}
				if b == '\r' || b == '\n' {
					break
				}
			}
			if end < 0 {
				return false
			}
			if !st.skipNoNewline(end + 1) {
				return false
			}
			emit(parser.DaedalusLexerStringLiteral, src[start:st.i], channel, rstart, st.ri-1, line, col)
			continue
		case '/':
			var next byte
			if st.i+1 < n {
				next = src[st.i+1]
			}
			switch next {
			case '/':
				end := st.i + 2
				for end < n && src[end] != '\r' && src[end] != '\n' {
					end++
				}
				if !st.skipNoNewline(end) {
					return false
				}
				emit(parser.DaedalusLexerLineComment, src[start:st.i], parser.COMMENTS, rstart, st.ri-1, line, col)
				continue
			case '*':
				idx := strings.Index(src[st.i+2:], "*/")
				if idx < 0 {
					return false
				}
				if !st.skipWithNewlines(st.i + 2 + idx + 2) {
					return false
				}
				continue
			case '=':
				ttype, blen, rlen = parser.DaedalusLexerDivAssign, 2, 2
			default:
				ttype = parser.DaedalusLexerDiv
			}
		case '.':
			if st.i+1 < n && isDigit(src[st.i+1]) {
				fl := floatLen(src, st.i)
				ttype, blen, rlen = parser.DaedalusLexerFloatLiteral, fl, fl
			} else {
				ttype = parser.DaedalusLexerDot
			}
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			// IntegerLiteral, FloatLiteral and Identifier compete, the longest wins.
			// On equal length the earlier rule wins (Integer, Float, Identifier).
			il := countDigits(src, st.i)
			fl := floatLen(src, st.i)
			idb, idr := scanIdent(src, st.i)
			switch {
			case idb > il && idb > fl:
				ttype, blen, rlen = parser.DaedalusLexerIdentifier, idb, idr
			case fl > il:
				ttype, blen, rlen = parser.DaedalusLexerFloatLiteral, fl, fl
			default:
				ttype, blen, rlen = parser.DaedalusLexerIntegerLiteral, il, il
			}
		case '(':
			ttype = parser.DaedalusLexerLeftParen
		case ')':
			ttype = parser.DaedalusLexerRightParen
		case '[':
			ttype = parser.DaedalusLexerLeftBracket
		case ']':
			ttype = parser.DaedalusLexerRightBracket
		case '{':
			ttype = parser.DaedalusLexerLeftBrace
		case '}':
			ttype = parser.DaedalusLexerRightBrace
		case ';':
			ttype = parser.DaedalusLexerSemi
		case ',':
			ttype = parser.DaedalusLexerComma
		case '%':
			ttype = parser.DaedalusLexerPercent
		case '~':
			ttype = parser.DaedalusLexerTilde
		case '&':
			switch peek(src, st.i+1) {
			case '&':
				ttype, blen, rlen = parser.DaedalusLexerAnd, 2, 2
			case '=':
				ttype, blen, rlen = parser.DaedalusLexerAndAssign, 2, 2
			default:
				ttype = parser.DaedalusLexerBitAnd
			}
		case '|':
			switch peek(src, st.i+1) {
			case '|':
				ttype, blen, rlen = parser.DaedalusLexerOr, 2, 2
			case '=':
				ttype, blen, rlen = parser.DaedalusLexerOrAssign, 2, 2
			default:
				ttype = parser.DaedalusLexerBitOr
			}
		case '+':
			if peek(src, st.i+1) == '=' {
				ttype, blen, rlen = parser.DaedalusLexerPlusAssign, 2, 2
			} else {
				ttype = parser.DaedalusLexerPlus
			}
		case '-':
			if peek(src, st.i+1) == '=' {
				ttype, blen, rlen = parser.DaedalusLexerMinusAssign, 2, 2
			} else {
				ttype = parser.DaedalusLexerMinus
			}
		case '*':
			if peek(src, st.i+1) == '=' {
				ttype, blen, rlen = parser.DaedalusLexerStarAssign, 2, 2
			} else {
				ttype = parser.DaedalusLexerStar
			}
		case '!':
			if peek(src, st.i+1) == '=' {
				ttype, blen, rlen = parser.DaedalusLexerNotEqual, 2, 2
			} else {
				ttype = parser.DaedalusLexerNot
			}
		case '=':
			if peek(src, st.i+1) == '=' {
				ttype, blen, rlen = parser.DaedalusLexerEqual, 2, 2
			} else {
				ttype = parser.DaedalusLexerAssign
			}
		case '<':
			switch peek(src, st.i+1) {
			case '<':
				ttype, blen, rlen = parser.DaedalusLexerShl, 2, 2
			case '=':
				ttype, blen, rlen = parser.DaedalusLexerLessEq, 2, 2
			default:
				ttype = parser.DaedalusLexerLess
			}
		case '>':
			switch peek(src, st.i+1) {
			case '>':
				ttype, blen, rlen = parser.DaedalusLexerShr, 2, 2
			case '=':
				ttype, blen, rlen = parser.DaedalusLexerGreaterEq, 2, 2
			default:
				ttype = parser.DaedalusLexerGreater
			}
		default:
			// identifiers (and keywords) starting with a letter, '_' or a German letter
			if isIdentStartASCII(c) {
				blen, rlen = scanIdent(src, st.i)
				ttype = keywordType(src[st.i : st.i+blen])
			} else if c >= utf8.RuneSelf {
				r, _ := utf8.DecodeRuneInString(src[st.i:])
				if !isGermanChar(r) {
					return false
				}
				blen, rlen = scanIdent(src, st.i)
				ttype = parser.DaedalusLexerIdentifier
			} else {
				return false
			}
		}

		st.i += blen
		st.ri += rlen
		st.col += rlen
		emit(ttype, src[start:st.i], channel, rstart, st.ri-1, line, col)
	}

	emit(antlr.TokenEOF, "", antlr.TokenDefaultChannel, st.ri, st.ri-1, st.line, st.col)
	if lazy, ok := input.(*lazyCharStream); ok {
		lazy.size = st.ri
	}
	l.tokens = toks
	return true
}

func peek(src string, i int) byte {
	if i < len(src) {
		return src[i]
	}
	return 0
}
