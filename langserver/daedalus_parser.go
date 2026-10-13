package langserver

//go:generate go run ../cmd/gen-combined-listener

import (
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/antlr4-go/antlr/v4"
	"github.com/kirides/DaedalusLanguageServer/daedalus/parser"
	"github.com/kirides/DaedalusLanguageServer/daedalus/symbol"
)

type DaedalusGrammarParser interface {
	RemoveErrorListeners()
	AddErrorListener(antlr.ErrorListener)
	SetInputStream(antlr.TokenStream)
	GetInterpreter() *antlr.ParserATNSimulator

	NewDaedalusFile() parser.IDaedalusFileContext
}

type Parser interface {
	Parse(source, content string, listener antlr.ParseTreeListener, errListener antlr.ErrorListener) parser.IDaedalusFileContext
}

type parserPool struct {
	inner sync.Pool
}

func newParserPool(newFn func() DaedalusGrammarParser) *parserPool {
	return &parserPool{
		inner: sync.Pool{New: func() interface{} { return newFn() }},
	}
}
func (p *parserPool) Get() DaedalusGrammarParser  { return p.inner.Get().(DaedalusGrammarParser) }
func (p *parserPool) Put(v DaedalusGrammarParser) { p.inner.Put(v) }

type lexerPool struct {
	inner sync.Pool
}

func newLexerPool(newFn func() *fastLexer) *lexerPool {
	return &lexerPool{
		inner: sync.Pool{New: func() interface{} { return newFn() }},
	}
}
func (p *lexerPool) Get() *fastLexer  { return p.inner.Get().(*fastLexer) }
func (p *lexerPool) Put(v *fastLexer) { p.inner.Put(v) }

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
func (m *parseResultsManager) ParseScriptListener(source, content string, listener parser.DaedalusListener, errListener antlr.ErrorListener) parser.IDaedalusFileContext {
	return m.parser.Parse(source, content, listener, errListener)
}

// ParseScript ...
func (m *parseResultsManager) ParseScript(source, content string, lastModifiedAt time.Time) *ParseResult {
	m.mtx.RLock()
	if existing, ok := m.parseResults[source]; ok && existing.lastModifiedAt.Equal(lastModifiedAt) {
		m.mtx.RUnlock()
		return existing
	}
	m.mtx.RUnlock()

	listener := NewDaedalusStatefulListener(source, m)
	validating := NewDeferringValidatingListener(source)
	errListener := &SyntaxErrorListener{}

	// Validate while walking the tree, so that the (large) tree does not have to be kept around.
	m.ParseScriptListener(source, content, combineListeners(listener, validating), errListener)

	result := &ParseResult{
		SyntaxErrors:    errListener.SyntaxErrors,
		pendingCalls:    validating.pendingCalls,
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

// functionParameterCounts returns the number of parameters of all known global functions,
// by upper case name. If a function is defined multiple times, the first definition
// (by source path) wins.
func (m *parseResultsManager) functionParameterCounts() map[string]int {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	sources := make([]string, 0, len(m.parseResults))
	for k := range m.parseResults {
		sources = append(sources, k)
	}
	sort.Strings(sources)

	counts := make(map[string]int, 4096)
	for _, src := range sources {
		for name, fn := range m.parseResults[src].Functions {
			if _, exists := counts[name]; !exists {
				counts[name] = len(fn.Parameters)
			}
		}
	}
	return counts
}

// ResolvePendingCalls checks the argument counts of the function calls collected
// while parsing against all known functions.
func (m *parseResultsManager) ResolvePendingCalls(results ...*ParseResult) {
	var counts map[string]int
	for _, r := range results {
		if len(r.pendingCalls) == 0 {
			continue
		}
		if counts == nil {
			counts = m.functionParameterCounts()
		}
		for _, c := range r.pendingCalls {
			params, ok := counts[strings.ToUpper(c.name)]
			switch {
			case !ok:
			case int(c.args) < params:
				r.SyntaxErrors = append(r.SyntaxErrors, NewSyntaxError(int(c.line), int(c.col), D0004NotEnoughArgumentsSpecified))
			case int(c.args) > params:
				r.SyntaxErrors = append(r.SyntaxErrors, NewSyntaxError(int(c.line), int(c.col), D0005TooManyArgumentsSpecified))
			}
		}
		r.pendingCalls = nil
	}
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
