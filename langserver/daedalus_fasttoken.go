package langserver

import (
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

// tokenInfo is shared by all tokens of one lexed input.
type tokenInfo struct {
	pair  *antlr.TokenSourceCharStreamPair
	src   antlr.TokenSource
	input antlr.CharStream
}

// fastToken is a compact antlr.Token (48 bytes instead of 96 for antlr.CommonToken).
// Tokens are allocated in chunks, see tokenSlab.
type fastToken struct {
	text   string
	info   *tokenInfo
	start  int32
	stop   int32
	line   int32
	column int32
	index  int32
	typ    int16
	ch     int16
}

var _ antlr.Token = (*fastToken)(nil)

func (t *fastToken) GetSource() *antlr.TokenSourceCharStreamPair { return t.info.pair }
func (t *fastToken) GetTokenType() int                           { return int(t.typ) }
func (t *fastToken) GetChannel() int                             { return int(t.ch) }
func (t *fastToken) GetStart() int                               { return int(t.start) }
func (t *fastToken) GetStop() int                                { return int(t.stop) }
func (t *fastToken) GetLine() int                                { return int(t.line) }
func (t *fastToken) GetColumn() int                              { return int(t.column) }
func (t *fastToken) GetTokenIndex() int                          { return int(t.index) }
func (t *fastToken) SetTokenIndex(v int)                         { t.index = int32(v) }
func (t *fastToken) SetText(s string)                            { t.text = s }
func (t *fastToken) GetTokenSource() antlr.TokenSource           { return t.info.src }
func (t *fastToken) GetInputStream() antlr.CharStream            { return t.info.input }

func (t *fastToken) GetText() string {
	if t.text != "" {
		return t.text
	}
	input := t.info.input
	if input == nil {
		return ""
	}
	n := input.Size()
	if int(t.start) < n && int(t.stop) < n {
		return input.GetTextFromInterval(antlr.NewInterval(int(t.start), int(t.stop)))
	}
	return "<EOF>"
}

func (t *fastToken) String() string {
	txt := t.GetText()
	if txt != "" {
		txt = strings.ReplaceAll(txt, "\n", "\\n")
		txt = strings.ReplaceAll(txt, "\r", "\\r")
		txt = strings.ReplaceAll(txt, "\t", "\\t")
	} else {
		txt = "<no text>"
	}
	ch := ""
	if t.ch > 0 {
		ch = ",channel=" + strconv.Itoa(int(t.ch))
	}
	return "[@" + strconv.Itoa(int(t.index)) + "," + strconv.Itoa(int(t.start)) + ":" + strconv.Itoa(int(t.stop)) + "='" +
		txt + "',<" + strconv.Itoa(int(t.typ)) + ">" + ch + "," + strconv.Itoa(int(t.line)) + ":" + strconv.Itoa(int(t.column)) + "]"
}

const (
	tokenSlabSize    = 1024
	maxFirstSlabSize = 4096
)

// tokenSlab hands out fastTokens from chunks, to avoid one allocation per token.
// The tokens of a parse result stay alive together, so the first chunk is sized
// from the input to not waste memory on small files.
type tokenSlab struct {
	chunk     []fastToken
	firstSize int
}

// reset prepares the slab for an input of n bytes.
func (s *tokenSlab) reset(n int) {
	s.chunk = nil
	s.firstSize = min(n/5+4, maxFirstSlabSize)
}

func (s *tokenSlab) next() *fastToken {
	if len(s.chunk) == 0 {
		size := tokenSlabSize
		if s.firstSize > 0 {
			size, s.firstSize = s.firstSize, 0
		}
		s.chunk = make([]fastToken, size)
	}
	t := &s.chunk[0]
	s.chunk = s.chunk[1:]
	return t
}

// fastTokenStream is an antlr.TokenStream over fully lexed tokens.
// It mirrors the semantics of antlr.CommonTokenStream for the default channel,
// without the on-demand fetching.
type fastTokenStream struct {
	tokens []*fastToken // all tokens, the last one is EOF
	index  int          // index of the current (on-channel) token
	src    antlr.TokenSource
}

var _ antlr.TokenStream = (*fastTokenStream)(nil)

func newFastTokenStream(tokens []*fastToken, src antlr.TokenSource) *fastTokenStream {
	s := &fastTokenStream{tokens: tokens, src: src}
	s.index = s.nextOnChannel(0)
	return s
}

// nextOnChannel returns i if tokens[i] is on the default channel, otherwise the next one that is.
// The last token (EOF) is always on the default channel.
func (s *fastTokenStream) nextOnChannel(i int) int {
	last := len(s.tokens) - 1
	for i < last && s.tokens[i].ch != antlr.TokenDefaultChannel {
		i++
	}
	if i > last {
		return last
	}
	return i
}

func (s *fastTokenStream) previousOnChannel(i, channel int) int {
	for i >= 0 && int(s.tokens[i].ch) != channel {
		i--
	}
	return i
}

func (s *fastTokenStream) Consume() {
	if s.tokens[s.index].typ == antlr.TokenEOF {
		panic("cannot consume EOF")
	}
	s.index = s.nextOnChannel(s.index + 1)
}

func (s *fastTokenStream) LA(i int) int { return s.LT(i).GetTokenType() }

func (s *fastTokenStream) Mark() int { return 0 }

func (s *fastTokenStream) Release(_ int) {}

func (s *fastTokenStream) Index() int { return s.index }

func (s *fastTokenStream) Seek(index int) { s.index = s.nextOnChannel(index) }

func (s *fastTokenStream) Size() int { return len(s.tokens) }

func (s *fastTokenStream) GetSourceName() string { return s.src.GetSourceName() }

func (s *fastTokenStream) LT(k int) antlr.Token {
	switch {
	case k == 1:
		return s.tokens[s.index]
	case k == 0:
		return nil
	case k < 0:
		return s.lb(-k)
	}
	i := s.index
	for n := 1; n < k; n++ {
		if i+1 < len(s.tokens) {
			i = s.nextOnChannel(i + 1)
		}
	}
	return s.tokens[i]
}

func (s *fastTokenStream) lb(k int) antlr.Token {
	if k == 0 || s.index-k < 0 {
		return nil
	}
	i := s.index
	for n := 1; n <= k; n++ {
		i = s.previousOnChannel(i-1, antlr.TokenDefaultChannel)
	}
	if i < 0 {
		return nil
	}
	return s.tokens[i]
}

func (s *fastTokenStream) Reset() { s.index = s.nextOnChannel(0) }

func (s *fastTokenStream) Get(index int) antlr.Token { return s.tokens[index] }

func (s *fastTokenStream) GetTokenSource() antlr.TokenSource { return s.src }

func (s *fastTokenStream) SetTokenSource(src antlr.TokenSource) { s.src = src }

func (s *fastTokenStream) GetAllText() string {
	return s.GetTextFromInterval(antlr.NewInterval(0, len(s.tokens)-1))
}

func (s *fastTokenStream) GetTextFromTokens(start, end antlr.Token) string {
	if start == nil || end == nil {
		return ""
	}
	return s.GetTextFromInterval(antlr.NewInterval(start.GetTokenIndex(), end.GetTokenIndex()))
}

func (s *fastTokenStream) GetTextFromRuleContext(interval antlr.RuleContext) string {
	return s.GetTextFromInterval(interval.GetSourceInterval())
}

func (s *fastTokenStream) GetTextFromInterval(interval antlr.Interval) string {
	start, stop := interval.Start, interval.Stop
	if start < 0 || stop < 0 {
		return ""
	}
	if stop >= len(s.tokens) {
		stop = len(s.tokens) - 1
	}
	var sb strings.Builder
	for i := start; i <= stop; i++ {
		t := s.tokens[i]
		if t.typ == antlr.TokenEOF {
			break
		}
		sb.WriteString(t.GetText())
	}
	return sb.String()
}

// GetHiddenTokensToLeft collects all tokens on channel to the left of the
// token at tokenIndex until a token on the default channel is seen.
// If channel is -1, it finds any non default channel token.
func (s *fastTokenStream) GetHiddenTokensToLeft(tokenIndex, channel int) []antlr.Token {
	if tokenIndex < 0 || tokenIndex >= len(s.tokens) {
		panic(strconv.Itoa(tokenIndex) + " not in 0.." + strconv.Itoa(len(s.tokens)-1))
	}
	prev := s.previousOnChannel(tokenIndex-1, antlr.TokenDefaultChannel)
	if prev == tokenIndex-1 {
		return nil
	}
	return s.filterForChannel(prev+1, tokenIndex-1, channel)
}

// GetHiddenTokensToRight collects all tokens on channel to the right of the
// token at tokenIndex until a token on the default channel or EOF is seen.
// If channel is -1, it finds any non default channel token.
func (s *fastTokenStream) GetHiddenTokensToRight(tokenIndex, channel int) []antlr.Token {
	if tokenIndex < 0 || tokenIndex >= len(s.tokens) {
		panic(strconv.Itoa(tokenIndex) + " not in 0.." + strconv.Itoa(len(s.tokens)-1))
	}
	to := s.nextOnChannel(tokenIndex + 1)
	return s.filterForChannel(tokenIndex+1, to, channel)
}

func (s *fastTokenStream) filterForChannel(left, right, channel int) []antlr.Token {
	var hidden []antlr.Token
	for i := left; i < right+1 && i < len(s.tokens); i++ {
		t := s.tokens[i]
		if channel == -1 {
			if t.ch != antlr.TokenDefaultChannel {
				hidden = append(hidden, t)
			}
		} else if int(t.ch) == channel {
			hidden = append(hidden, t)
		}
	}
	return hidden
}
