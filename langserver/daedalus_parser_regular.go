package langserver

import (
	"langsrv/langserver/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type regularGrammarParser struct {
	*parser.DaedalusParser
}

func (rp *regularGrammarParser) NewDaedalusFile() antlr.Tree {
	return rp.DaedalusFile()
}

var _ DaedalusGrammarParser = (*regularGrammarParser)(nil)
var _ Parser = (*RegularParser)(nil)

type RegularParser struct {
	pooledParsers *parserPool
}

func newRegularParser() *RegularParser {
	return &RegularParser{
		pooledParsers: newParserPool(func() DaedalusGrammarParser { return &regularGrammarParser{parser.NewDaedalusParser(nil)} }),
	}
}

func (rp *RegularParser) Parse(source, content string, listener antlr.ParseTreeListener, errListener antlr.ErrorListener) {
	inputStream := antlr.NewInputStream(content)
	lexer := parser.NewDaedalusLexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)

	p := rp.pooledParsers.Get()
	defer func() {
		p.SetInputStream(nil)
		rp.pooledParsers.Put(p)
	}()
	p.SetInputStream(tokenStream)

	p.RemoveErrorListeners()
	p.AddErrorListener(errListener)
	// Use SLL prediction
	p.GetInterpreter().SetPredictionMode(antlr.PredictionModeSLL)

	antlr.NewParseTreeWalker().Walk(listener, p.NewDaedalusFile())
}
