package langserver

import (
	"sync"
	"time"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/kirides/DaedalusLanguageServer/daedalus/parser"
	"github.com/kirides/DaedalusLanguageServer/daedalus/symbol"
)

type DaedalusGrammarParser interface {
	RemoveErrorListeners()
	AddErrorListener(antlr.ErrorListener)
	SetInputStream(antlr.TokenStream)
	GetInterpreter() *antlr.ParserATNSimulator

	NewDaedalusFile() antlr.Tree
}

type Parser interface {
	Parse(source, content string, listener antlr.ParseTreeListener, errListener antlr.ErrorListener)
}

type parserPool struct {
	inner sync.Pool
}

func newParserPool(newFn func() DaedalusGrammarParser) *parserPool {
	return &parserPool{
		inner: sync.Pool{New: func() interface{} { return newFn() }},
	}
}
func (p *parserPool) Get() DaedalusGrammarParser {
	return p.inner.Get().(DaedalusGrammarParser)
}
func (p *parserPool) Put(v DaedalusGrammarParser) {
	p.inner.Put(v)
}

// ParseAndValidateScript ...
func (m *parseResultsManager) ParseAndValidateScript(source, content string) *ParseResult {
	stateful := NewDaedalusStatefulListener(source, m)
	validating := NewDaedalusValidatingListener(source, m)
	errListener := &SyntaxErrorListener{}
	m.ParseScriptListener(source, content, combineListeners(stateful, validating), errListener)

	result := &ParseResult{
		SyntaxErrors:    errListener.SyntaxErrors,
		GlobalVariables: stateful.Globals.Variables,
		GlobalConstants: stateful.Globals.Constants,
		Functions:       stateful.Globals.Functions,
		Classes:         stateful.Globals.Classes,
		Prototypes:      stateful.Globals.Prototypes,
		Instances:       stateful.Globals.Instances,
		Namespaces:      stateful.Namespaces,
		Source:          source,
		lastModifiedAt:  time.Now(),
	}
	return result
}

// ParseScriptListener ...
func (m *parseResultsManager) ParseScriptListener(source, content string, listener parser.DaedalusListener, errListener antlr.ErrorListener) {
	m.parser.Parse(source, content, listener, errListener)
}

// ParseScript ...
func (m *parseResultsManager) ParseScript(source, content string, lastModifiedAt time.Time) *ParseResult {
	m.mtx.Lock()
	if existing, ok := m.parseResults[source]; ok && existing.lastModifiedAt == lastModifiedAt {
		m.mtx.Unlock()
		return existing
	}
	m.mtx.Unlock()

	listener := NewDaedalusStatefulListener(source, m)
	errListener := &SyntaxErrorListener{}

	m.ParseScriptListener(source, content, listener, errListener)

	result := &ParseResult{
		SyntaxErrors:    errListener.SyntaxErrors,
		GlobalVariables: listener.Globals.Variables,
		GlobalConstants: listener.Globals.Constants,
		Functions:       listener.Globals.Functions,
		Classes:         listener.Globals.Classes,
		Prototypes:      listener.Globals.Prototypes,
		Instances:       listener.Globals.Instances,
		Namespaces:      listener.Namespaces,
		Source:          source,
		lastModifiedAt:  lastModifiedAt,
	}
	return result
}

// ValidateScript ...
func (m *parseResultsManager) ValidateScript(source, content string) []SyntaxError {
	listener := NewDaedalusValidatingListener(source, m)
	errListener := &SyntaxErrorListener{}
	m.ParseScriptListener(source, content, listener, errListener)

	return errListener.SyntaxErrors
}

// SymbolType ...
type SymbolType uint

const (
	// SymbolNone ...
	SymbolNone SymbolType = 0
	// SymbolClass ...
	SymbolClass SymbolType = 1 << 0
	// SymbolConstant ...
	SymbolConstant SymbolType = 1 << 1
	// SymbolFunction ...
	SymbolFunction SymbolType = 1 << 2
	// SymbolInstance ...
	SymbolInstance SymbolType = 1 << 3
	// SymbolPrototype ...
	SymbolPrototype SymbolType = 1 << 4
	// SymbolVariable ...
	SymbolVariable SymbolType = 1 << 5
	// SymbolAll ...
	SymbolAll SymbolType = 0xFFFFFFFF
)

// LookupGlobalSymbol ...
func (p *ParseResult) LookupGlobalSymbol(name string, types SymbolType) (symbol.Symbol, bool) {
	if (types & SymbolClass) != 0 {
		if s, ok := p.Classes[name]; ok {
			return s, true
		}
	}
	if (types & SymbolConstant) != 0 {
		if s, ok := p.GlobalConstants[name]; ok {
			return s, true
		}
	}
	if (types & SymbolFunction) != 0 {
		if s, ok := p.Functions[name]; ok {
			return s, true
		}
	}
	if (types & SymbolInstance) != 0 {
		if s, ok := p.Instances[name]; ok {
			return s, true
		}
	}
	if (types & SymbolPrototype) != 0 {
		if s, ok := p.Prototypes[name]; ok {
			return s, true
		}
	}
	if (types & SymbolVariable) != 0 {
		if s, ok := p.GlobalVariables[name]; ok {
			return s, true
		}
	}
	return nil, false
}

// WalkGlobalSymbols ...
func (p *ParseResult) WalkGlobalSymbols(walkFn func(symbol.Symbol) error, types SymbolType) error {
	if (types & SymbolClass) != 0 {
		for _, s := range p.Classes {
			if err := walkFn(s); err != nil {
				return err
			}
		}
	}
	if (types & SymbolConstant) != 0 {
		for _, s := range p.GlobalConstants {
			if err := walkFn(s); err != nil {
				return err
			}
		}
	}
	if (types & SymbolFunction) != 0 {
		for _, s := range p.Functions {
			if err := walkFn(s); err != nil {
				return err
			}
		}
	}
	if (types & SymbolInstance) != 0 {
		for _, s := range p.Instances {
			if err := walkFn(s); err != nil {
				return err
			}
		}
	}
	if (types & SymbolPrototype) != 0 {
		for _, s := range p.Prototypes {
			if err := walkFn(s); err != nil {
				return err
			}
		}
	}
	if (types & SymbolVariable) != 0 {
		for _, s := range p.GlobalVariables {
			if err := walkFn(s); err != nil {
				return err
			}
		}
	}
	return nil
}
