package langserver

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"langsrv/langserver/parser"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
)

// ParseResult ...
type ParseResult struct {
	SyntaxErrors    []SyntaxError
	GlobalVariables map[string]VariableSymbol
	GlobalConstants map[string]ConstantSymbol
	Functions       map[string]FunctionSymbol
	Classes         map[string]ClassSymbol
	Prototypes      map[string]ProtoTypeOrInstanceSymbol
	Instances       map[string]ProtoTypeOrInstanceSymbol
	Source          string
}

type parserPool struct {
	inner sync.Pool
}

func newParserPool() *parserPool {
	return &parserPool{
		inner: sync.Pool{
			New: func() interface{} { return parser.NewDaedalusParser(nil) },
		},
	}
}
func (p *parserPool) Get() *parser.DaedalusParser {
	return p.inner.Get().(*parser.DaedalusParser)
}
func (p *parserPool) Put(v *parser.DaedalusParser) {
	p.inner.Put(v)
}

var pooledParsers = newParserPool()

// ParseAndValidateScript ...
func (m *parseResultsManager) ParseAndValidateScript(source, content string) *ParseResult {
	r := m.ParseScript(source, content)
	e := m.ValidateScript(source, content)
	if len(e) > 0 {
		r.SyntaxErrors = append(r.SyntaxErrors, e...)
	}
	return r
}

// ParseScript ...
func (m *parseResultsManager) ParseScript(source, content string) *ParseResult {
	inputStream := antlr.NewInputStream(content)
	lexer := parser.NewDaedalusLexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	p := pooledParsers.Get()
	defer func() {
		p.SetInputStream(nil)
		pooledParsers.Put(p)
	}()
	p.SetInputStream(tokenStream)

	errListener := &SyntaxErrorListener{}
	p.RemoveErrorListeners()
	p.AddErrorListener(errListener)
	// Use SLL prediction
	p.Interpreter.SetPredictionMode(antlr.PredictionModeSLL)
	listener := NewDaedalusStatefulListener(source, m)

	antlr.NewParseTreeWalker().Walk(listener, p.DaedalusFile())

	result := &ParseResult{
		SyntaxErrors:    errListener.SyntaxErrors,
		GlobalVariables: listener.GlobalVariables,
		GlobalConstants: listener.GlobalConstants,
		Functions:       listener.Functions,
		Classes:         listener.Classes,
		Prototypes:      listener.Prototypes,
		Instances:       listener.Instances,
		Source:          source,
	}
	return result
}

// ValidateScript ...
func (m *parseResultsManager) ValidateScript(source, content string) []SyntaxError {
	inputStream := antlr.NewInputStream(content)
	lexer := parser.NewDaedalusLexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	p := pooledParsers.Get()
	defer func() {
		p.SetInputStream(nil)
		pooledParsers.Put(p)
	}()
	p.SetInputStream(tokenStream)

	errListener := &SyntaxErrorListener{}
	p.RemoveErrorListeners()
	p.AddErrorListener(errListener)
	// Use SLL prediction
	p.Interpreter.SetPredictionMode(antlr.PredictionModeSLL)
	listener := NewDaedalusValidatingListener(source, m)

	antlr.NewParseTreeWalker().Walk(listener, p.DaedalusFile())

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
func (p *ParseResult) LookupGlobalSymbol(name string, types SymbolType) (Symbol, bool) {
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
func (p *ParseResult) WalkGlobalSymbols(walkFn func(Symbol) error, types SymbolType) error {
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

type parseResultsManager struct {
	mtx          sync.RWMutex
	parseResults map[string]*ParseResult
}

func newParseResultsManager() *parseResultsManager {
	return &parseResultsManager{
		parseResults: make(map[string]*ParseResult),
	}
}

func (m *parseResultsManager) WalkGlobalSymbols(walkFn func(Symbol) error, types SymbolType) error {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	for _, p := range m.parseResults {
		err := p.WalkGlobalSymbols(func(s Symbol) error {
			return walkFn(s)
		}, types)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *parseResultsManager) LookupGlobalSymbol(name string, types SymbolType) (Symbol, bool) {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	for _, p := range m.parseResults {
		if s, ok := p.LookupGlobalSymbol(name, types); ok {
			return s, true
		}
	}

	return nil, false
}

func (m *parseResultsManager) GetGlobalSymbols(types SymbolType) ([]Symbol, error) {
	result := make([]Symbol, 0, 200)

	m.mtx.RLock()
	defer m.mtx.RUnlock()

	for _, p := range m.parseResults {
		p.WalkGlobalSymbols(func(s Symbol) error {
			result = append(result, s)
			return nil
		}, types)
	}

	return result, nil
}

func (m *parseResultsManager) Get(documentURI string) (*ParseResult, error) {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	if r, ok := m.parseResults[documentURI]; ok {
		return r, nil
	}
	return nil, fmt.Errorf("Document %q not found", documentURI)
}

func (m *parseResultsManager) Update(documentURI, content string) (*ParseResult, error) {
	r := m.ParseAndValidateScript(documentURI, content)

	m.mtx.Lock()
	defer m.mtx.Unlock()

	m.parseResults[documentURI] = r
	return r, nil
}

func resolveSrcPaths(srcFile, prefixDir string) []string {
	fileBytes, err := ioutil.ReadFile(srcFile)
	if err != nil {
		return []string{}
	}
	decodedContent, err := charmap.Windows1252.NewDecoder().Bytes(fileBytes)
	if err != nil {
		return []string{}
	}
	srcContent := string(decodedContent)
	scanner := bufio.NewScanner(strings.NewReader(srcContent))

	resolvedPaths := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}
		dir := filepath.Dir(line)
		fname := filepath.Base(line)
		ext := strings.ToLower(filepath.Ext(fname))

		if ext == ".d" {
			if strings.Contains(fname, "*") {
				rxFilePath := regexp.MustCompile(strings.ReplaceAll(regexp.QuoteMeta(strings.ToLower(fname)), "\\*", ".*"))

				entries, err := ioutil.ReadDir(filepath.Join(prefixDir, dir))
				if err != nil {
					continue
				}
				for _, e := range entries {
					if e.IsDir() {
						continue
					}
					if rxFilePath.MatchString(strings.ToLower(e.Name())) {
						absPath, err := filepath.Abs(filepath.Join(prefixDir, dir, e.Name()))
						if err != nil {
							continue
						}
						resolvedPaths = append(resolvedPaths, absPath)
					}
				}
			} else {
				absPath, err := filepath.Abs(filepath.Join(prefixDir, dir, fname))
				if err != nil {
					continue
				}
				resolvedPaths = append(resolvedPaths, absPath)
			}
		} else if ext == ".src" {
			fmt.Fprintf(os.Stderr, "Collecting scripts from %q\n", filepath.Join(prefixDir, dir, fname))
			resolvedPaths = append(resolvedPaths, resolveSrcPaths(filepath.Join(prefixDir, line), filepath.Join(prefixDir, dir))...)
		}
	}

	return resolvedPaths
}

var (
	bufferPool  = sync.Pool{New: func() interface{} { b := new(bytes.Buffer); b.Grow(2048); return b }}
	decoderPool = sync.Pool{New: func() interface{} { return charmap.Windows1252.NewDecoder() }}
)

func (m *parseResultsManager) validateFiles(resolvedPaths []string) map[string][]SyntaxError {
	results := make(map[string][]SyntaxError)

	chanPaths := make(chan string, len(resolvedPaths))
	for _, r := range resolvedPaths {
		chanPaths <- r
	}
	close(chanPaths)

	var wg sync.WaitGroup
	numWorkers := runtime.NumCPU()
	if numWorkers > 2 {
		numWorkers = numWorkers / 2
	}
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()

			buf := bufferPool.Get().(*bytes.Buffer)
			defer bufferPool.Put(buf)

			decoder := decoderPool.Get().(*encoding.Decoder)
			defer decoderPool.Put(decoder)

			for r := range chanPaths {
				f, err := os.Open(r)
				if err != nil {
					continue
				}
				decoder.Reset()
				translated := decoder.Reader(f)
				buf.Reset()
				_, err = buf.ReadFrom(translated)
				f.Close()
				if err != nil {
					continue
				}

				parsed := m.ValidateScript(r, string(buf.Bytes()))
				if len(parsed) > 0 {
					m.mtx.Lock()
					results[r] = parsed
					m.mtx.Unlock()
				}
			}
		}(&wg)
	}

	wg.Wait()
	return results
}

func (m *parseResultsManager) ParseSource(srcFile string) ([]*ParseResult, error) {
	resolvedPaths := resolveSrcPaths(srcFile, filepath.Dir(srcFile))
	fmt.Fprintf(os.Stderr, "Parsing %q. This might take a while.\n", srcFile)

	results := make([]*ParseResult, 0, len(resolvedPaths))

	chanPaths := make(chan string, len(resolvedPaths))
	for _, r := range resolvedPaths {
		chanPaths <- r
	}
	close(chanPaths)

	var wg sync.WaitGroup
	numWorkers := runtime.NumCPU()
	if numWorkers > 2 {
		numWorkers = numWorkers / 2
	}
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()

			buf := bufferPool.Get().(*bytes.Buffer)
			defer bufferPool.Put(buf)

			decoder := decoderPool.Get().(*encoding.Decoder)
			defer decoderPool.Put(decoder)

			for r := range chanPaths {
				f, err := os.Open(r)
				if err != nil {
					continue
				}
				decoder.Reset()
				translated := decoder.Reader(f)
				buf.Reset()
				_, err = buf.ReadFrom(translated)
				f.Close()
				if err != nil {
					continue
				}

				parsed := m.ParseScript(r, string(buf.Bytes()))

				m.mtx.Lock()
				m.parseResults[parsed.Source] = parsed
				results = append(results, parsed)
				m.mtx.Unlock()
			}
		}(&wg)
	}

	wg.Wait()

	validations := m.validateFiles(resolvedPaths)

	for _, v := range results {
		if e, ok := validations[v.Source]; ok {
			v.SyntaxErrors = append(v.SyntaxErrors, e...)
		}
	}

	fmt.Fprintf(os.Stderr, "Done parsing %q: %d scripts.\n", srcFile, len(results))
	return results, nil
}
