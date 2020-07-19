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
	"strings"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"golang.org/x/text/encoding/charmap"
)

// ParseResult ...
type ParseResult struct {
	SyntaxErrors    []SyntaxError
	GlobalVariables []VariableSymbol
	GlobalConstants []ConstantSymbol
	Functions       []FunctionSymbol
	Classes         []ClassSymbol
	Prototypes      []ProtoTypeOrInstanceSymbol
	Instances       []ProtoTypeOrInstanceSymbol
	Source          string
}

// ParseScript ...
func ParseScript(source, content string) *ParseResult {
	inputStream := antlr.NewInputStream(content)
	lexer := parser.NewDaedalusLexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewDaedalusParser(tokenStream)
	errListener := &SyntaxErrorListener{}
	p.AddErrorListener(errListener)
	// Use SLL prediction
	p.Interpreter.SetPredictionMode(antlr.PredictionModeSLL)
	listener := NewDaedalusStatefulListener(p, source)

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

// WalkGlobalSymbols ...
func (p *ParseResult) WalkGlobalSymbols(walkFn func(Symbol) error) error {
	for _, s := range p.Classes {
		if err := walkFn(s); err != nil {
			return err
		}
	}
	for _, s := range p.GlobalConstants {
		if err := walkFn(s); err != nil {
			return err
		}
	}
	for _, s := range p.Functions {
		if err := walkFn(s); err != nil {
			return err
		}
	}
	for _, s := range p.Instances {
		if err := walkFn(s); err != nil {
			return err
		}
	}
	for _, s := range p.Prototypes {
		if err := walkFn(s); err != nil {
			return err
		}
	}
	for _, s := range p.GlobalVariables {
		if err := walkFn(s); err != nil {
			return err
		}
	}
	return nil
}

type parseResultsManager struct {
	mtx          *sync.Mutex
	parseResults map[string]*ParseResult
}

func newParseResultsManager() *parseResultsManager {
	return &parseResultsManager{
		mtx:          &sync.Mutex{},
		parseResults: make(map[string]*ParseResult),
	}
}

func (m *parseResultsManager) WalkGlobalSymbols(walkFn func(Symbol) error) error {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	for _, p := range m.parseResults {
		err := p.WalkGlobalSymbols(func(s Symbol) error {
			return walkFn(s)
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *parseResultsManager) GetGlobalSymbols() ([]Symbol, error) {
	result := make([]Symbol, 0, 200)

	m.mtx.Lock()
	defer m.mtx.Unlock()

	for _, p := range m.parseResults {
		p.WalkGlobalSymbols(func(s Symbol) error {
			result = append(result, s)
			return nil
		})
	}

	return result, nil
}

func (m *parseResultsManager) Get(documentURI string) (*ParseResult, error) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	if r, ok := m.parseResults[documentURI]; ok {
		return r, nil
	}
	return nil, fmt.Errorf("Document %q not found", documentURI)
}

func (m *parseResultsManager) Update(documentURI, content string) (*ParseResult, error) {
	r := ParseScript(documentURI, content)

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
				rxFilePath := regexp.MustCompile(strings.ReplaceAll(regexp.QuoteMeta(fname), "\\*", ".*"))

				entries, err := ioutil.ReadDir(dir)
				if err != nil {
					continue
				}
				for _, e := range entries {
					if e.IsDir() {
						continue
					}
					if rxFilePath.MatchString(e.Name()) {
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
			resolvedPaths = append(resolvedPaths, resolveSrcPaths(line, dir)...)
		}
	}

	return resolvedPaths
}

func (m *parseResultsManager) ParseSource(srcFile string) ([]*ParseResult, error) {

	fmt.Fprintf(os.Stderr, "Parsing %q. This might take a while.\n", srcFile)
	resolvedPaths := resolveSrcPaths(srcFile, filepath.Dir(srcFile))

	results := make([]*ParseResult, 0, len(resolvedPaths))
	decoder := charmap.Windows1252.NewDecoder()
	buf := new(bytes.Buffer)
	buf.Grow(2048)

	for _, r := range resolvedPaths {
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

		parsed := ParseScript(r, string(buf.Bytes()))
		results = append(results, parsed)
	}

	// maxConcurrency := runtime.NumCPU()
	// if maxConcurrency > 4 {
	// 	maxConcurrency = maxConcurrency / 2
	// }
	// filesToParse := make(chan string, maxConcurrency)
	// mtx := &sync.Mutex{}

	// results := make([]*ParseResult, 0, len(resolvedPaths))
	// wg := &sync.WaitGroup{}
	// wg.Add(maxConcurrency)
	// for i := 0; i < maxConcurrency; i++ {
	// 	go func() {
	// 		defer wg.Done()
	// 		for r := range filesToParse {
	// 			fileBody, _ := ioutil.ReadFile(r)
	// 			script, _ := charmap.Windows1252.NewDecoder().Bytes(fileBody)
	// 			parsed := ParseScript(r, string(script))

	// 			mtx.Lock()
	// 			results = append(results, parsed)
	// 			mtx.Unlock()
	// 		}
	// 	}()
	// }

	// for _, v := range resolvedPaths {
	// 	filesToParse <- v
	// }
	// close(filesToParse)
	// wg.Wait()

	m.mtx.Lock()
	defer m.mtx.Unlock()
	for _, r := range results {
		m.parseResults[r.Source] = r
	}
	fmt.Fprintf(os.Stderr, "Done parsing %q: %d scripts.\n", srcFile, len(results))
	return results, nil
}
