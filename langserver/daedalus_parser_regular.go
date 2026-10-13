package langserver

import (
	"github.com/kirides/DaedalusLanguageServer/daedalus/parser"

	"github.com/antlr4-go/antlr/v4"
)

type regularGrammarParser struct {
	*parser.DaedalusParser
}

func (rp *regularGrammarParser) NewDaedalusFile() parser.IDaedalusFileContext {
	return rp.DaedalusFile()
}

var _ DaedalusGrammarParser = (*regularGrammarParser)(nil)
var _ Parser = (*RegularParser)(nil)

type RegularParser struct {
	pooledParsers *parserPool
	pooledLexers  *lexerPool
}

func newRegularParser() *RegularParser {
	return &RegularParser{
		pooledParsers: newParserPool(func() DaedalusGrammarParser { return &regularGrammarParser{parser.NewDaedalusParser(nil)} }),
		pooledLexers:  newLexerPool(newFastLexer),
	}
}

func (rp *RegularParser) Parse(source, content string, listener antlr.ParseTreeListener, errListener antlr.ErrorListener) parser.IDaedalusFileContext {
	lexer := rp.pooledLexers.Get()
	p := rp.pooledParsers.Get()
	defer func() {
		lexer.Release()
		rp.pooledLexers.Put(lexer)
		p.SetInputStream(nil)
		rp.pooledParsers.Put(p)
	}()
	lexer.ResetString(content)
	tokenStream := lexer.TokenStream()

	p.SetInputStream(tokenStream)

	p.RemoveErrorListeners()
	p.AddErrorListener(errListener)
	// Use SLL prediction
	p.GetInterpreter().SetPredictionMode(antlr.PredictionModeSLL)

	parsed := p.NewDaedalusFile()
	antlr.NewParseTreeWalker().Walk(listener, parsed)
	return parsed
}
