package langserver

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	dls "github.com/kirides/DaedalusLanguageServer"
	"github.com/kirides/DaedalusLanguageServer/daedalus/parser"
	"github.com/kirides/DaedalusLanguageServer/daedalus/symbol"
	lsp "github.com/kirides/DaedalusLanguageServer/protocol"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
)

type parseResultsManager struct {
	parseResults map[string]*ParseResult
	mtx          sync.RWMutex
	logger       dls.Logger

	fileEncoding encoding.Encoding
	srcEncoding  encoding.Encoding
	decoderPool  *sync.Pool

	parser           Parser
	NumParserThreads int
}

func newParseResultsManager(logger dls.Logger) *parseResultsManager {
	return &parseResultsManager{
		parseResults:     make(map[string]*ParseResult),
		logger:           logger,
		fileEncoding:     charmap.Windows1252,
		srcEncoding:      charmap.Windows1252,
		decoderPool:      &sync.Pool{New: func() interface{} { return charmap.Windows1252.NewDecoder() }},
		parser:           newRegularParser(),
		NumParserThreads: 0,
	}
}

var encodings = map[string]encoding.Encoding{
	"WINDOWS-1250": charmap.Windows1250,
	"WINDOWS-1251": charmap.Windows1251,
	"WINDOWS-1252": charmap.Windows1252,
	"WINDOWS-1253": charmap.Windows1253,
	"WINDOWS-1254": charmap.Windows1254,
	"WINDOWS-1255": charmap.Windows1255,
	"WINDOWS-1256": charmap.Windows1256,
	"WINDOWS-1257": charmap.Windows1257,
	"WINDOWS-1258": charmap.Windows1258,
}

func (m *parseResultsManager) SetFileEncoding(enc string) error {
	if e, ok := encodings[strings.ToUpper(enc)]; ok {
		m.fileEncoding = e
		m.decoderPool = &sync.Pool{New: func() interface{} { return e.NewDecoder() }}
		return nil
	}
	return fmt.Errorf("unknown encoding %q", enc)
}

func (m *parseResultsManager) SetSrcEncoding(enc string) error {
	if e, ok := encodings[strings.ToUpper(enc)]; ok {
		m.srcEncoding = e
		return nil
	}
	return fmt.Errorf("unknown encoding %q", enc)
}

func (m *parseResultsManager) CountSymbols() int64 {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	n := int64(0)
	for _, p := range m.parseResults {
		n += p.CountSymbols()
	}

	return n
}

func (m *parseResultsManager) WalkGlobalSymbols(walkFn func(symbol.Symbol) error, types SymbolType) error {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	for _, p := range m.parseResults {
		err := p.WalkGlobalSymbols(func(s symbol.Symbol) error {
			return walkFn(s)
		}, types)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *parseResultsManager) LookupGlobalSymbol(name string, types SymbolType) (symbol.Symbol, bool) {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	for _, p := range m.parseResults {
		if s, ok := p.LookupGlobalSymbol(name, types); ok {
			return s, true
		}
	}

	return nil, false
}

func (m *parseResultsManager) GetGlobalSymbols(types SymbolType) []symbol.Symbol {
	result := make([]symbol.Symbol, 0, 200)

	m.mtx.RLock()
	defer m.mtx.RUnlock()

	for _, p := range m.parseResults {
		p.WalkGlobalSymbols(func(s symbol.Symbol) error {
			result = append(result, s)
			return nil
		}, types)
	}

	return result
}

func (m *parseResultsManager) Get(filePath string) (*ParseResult, error) {
	m.mtx.RLock()
	defer m.mtx.RUnlock()

	if r, ok := m.parseResults[filePath]; ok {
		return r, nil
	}
	for k, v := range m.parseResults {
		if strings.EqualFold(k, filePath) {
			return v, nil
		}
	}
	return nil, fmt.Errorf("document %q not found", filePath)
}

func (m *parseResultsManager) ParseSemantics(ctx context.Context, filePath string) (*SemanticParseResult, error) {
	return m.ParseSemanticsRange(ctx, filePath, lsp.Range{})
}

func (m *parseResultsManager) LoadFile(ctx context.Context, filePath string) (string, error) {
	source := filePath

	buf := bufferPool.Get().(*bytes.Buffer)
	defer bufferPool.Put(buf)

	decoder := m.decoderPool.Get().(*encoding.Decoder)
	defer m.decoderPool.Put(decoder)

	f, err := os.Open(source)
	if err != nil {
		return "", err
	}
	defer f.Close()
	decoder.Reset()
	translated := decoder.Reader(f)
	buf.Reset()
	_, err = buf.ReadFrom(translated)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (m *parseResultsManager) LoadFileBuffer(ctx context.Context, filePath string, buffer *bytes.Buffer) error {
	source := filePath

	decoder := m.decoderPool.Get().(*encoding.Decoder)
	defer m.decoderPool.Put(decoder)

	f, err := os.Open(source)
	if err != nil {
		return err
	}
	defer f.Close()
	decoder.Reset()
	translated := decoder.Reader(f)
	buffer.Reset()
	_, err = buffer.ReadFrom(translated)
	if err != nil {
		return err
	}
	return nil
}

func (m *parseResultsManager) ParseSemanticsRange(ctx context.Context, documentURI string, rang lsp.Range) (*SemanticParseResult, error) {
	content, err := m.LoadFile(ctx, documentURI)
	if err != nil {
		return nil, err
	}

	return m.ParseSemanticsContentRange(ctx, documentURI, content, rang)
}

func (m *parseResultsManager) ParseSemanticsContentRange(ctx context.Context, source, content string, rang lsp.Range) (*SemanticParseResult, error) {
	return m.ParseSemanticsContentDataTypesRange(ctx, source, content, rang, DataAll)
}

type DataType uint32

const (
	DataIdentifiers DataType = 1 << iota
	DataInstance
	DataGlobalVariables
	DataGlobalConstants
	DataFunctions
	DataClasses
	DataPrototypes

	DataAll DataType = 0xFFFFFFFF
)

func (m *parseResultsManager) ParseSemanticsContentDataTypesRange(ctx context.Context, source, content string, rang lsp.Range, dataTypes DataType) (*SemanticParseResult, error) {

	listener := NewDaedalusIdentifierListener(source)
	if rang != (lsp.Range{}) {
		listener.Bbox = symbol.NewDefinition(
			int(rang.Start.Line+1), int(rang.Start.Character),
			int(rang.End.Line+1), int(rang.End.Character),
		)
	}
	listener2 := NewDaedalusStatefulListener(source, m)
	var combined parser.DaedalusListener
	if dataTypes&DataIdentifiers != 0 {
		combined = combineListeners(combined, listener)
	}
	if dataTypes > 1 {
		combined = combineListeners(combined, listener2)
	}
	if combined == nil {
		return nil, fmt.Errorf("datatypes invalid")
	}
	m.ParseScriptListener(source, content, combined, &NoOpErrorListener{})

	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	result := &SemanticParseResult{
		ParseResult: ParseResult{
			Source:          source,
			Instances:       listener2.Globals.Instances,
			GlobalVariables: listener2.Globals.Variables,
			GlobalConstants: listener2.Globals.Constants,
			Functions:       listener2.Globals.Functions,
			Classes:         listener2.Globals.Classes,
			Prototypes:      listener2.Globals.Prototypes,
		},
		GlobalIdentifiers: listener.GlobalIdentifiers,
	}
	return result, nil
}

func (m *parseResultsManager) GetCtx(ctx context.Context, documentURI string) (*ParseResult, error) {
	ticker := time.NewTicker(60 * time.Second) // 60s delay until request is aborted
	defer ticker.Stop()
	for {
		doc, err := m.Get(documentURI)
		if err == nil {
			return doc, err
		} else {
			select {
			case <-ticker.C:
				return nil, fmt.Errorf("timeout: document %q not found", documentURI)
			case <-time.After(100 * time.Millisecond):
			case <-ctx.Done():
				return nil, fmt.Errorf("document %q not found", documentURI)
			}
		}
	}
}

func (m *parseResultsManager) Update(documentURI, content string) (*ParseResult, error) {
	r := m.ParseAndValidateScript(documentURI, content)

	m.mtx.Lock()
	defer m.mtx.Unlock()

	for k := range m.parseResults {
		if strings.EqualFold(k, documentURI) {
			delete(m.parseResults, k)
			break
		}
	}
	m.parseResults[documentURI] = r
	return r, nil
}

func (m *parseResultsManager) resolveSrcPaths(srcFile, prefixDir string) ([]string, error) {
	fileBytes, err := os.ReadFile(srcFile)
	if err != nil {
		return nil, err
	}
	decodedContent, err := m.srcEncoding.NewDecoder().Bytes(fileBytes)
	if err != nil {
		return nil, err
	}
	srcContent := string(decodedContent)

	// Auronen: On Linux and macOS replace backslashes with forwardslashes
	if runtime.GOOS != "windows" {
		srcContent = strings.ReplaceAll(srcContent, "\\", string(os.PathSeparator))
	}

	scanner := bufio.NewScanner(strings.NewReader(srcContent))

	resolvedPaths := []string{}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "//") {
			// likely a comment, like "// LeGo" or "// Ikarus"
			m.logger.Debugf("resolveSrcPaths(%q): skipping line %q", srcFile, line)
			continue
		}
		dir := filepath.Dir(line)
		fname := filepath.Base(line)
		ext := strings.ToLower(filepath.Ext(fname))

		if ext == ".d" {
			if strings.Contains(fname, "*") {
				files, err := ResolvePathsCaseInsensitive(filepath.Join(prefixDir, dir), fname)
				if err != nil {
					m.logger.Warnf("resolveSrcPaths(%q): could not resolve filename %q", srcFile, filepath.Join(prefixDir, dir, fname))
					continue
				}
				for _, e := range files {
					if stat, err := os.Stat(e); err != nil || stat.IsDir() {
						m.logger.Debugf("resolveSrcPaths(%q): ignoring entry %q (error or director)", srcFile, e)
						continue
					}

					resolvedPaths = append(resolvedPaths, e)
				}
			} else {
				found, err := findPath(filepath.Join(prefixDir, dir, fname))
				if err != nil {
					m.logger.Warnf("resolveSrcPaths(%q): file not found on disk %q", srcFile, filepath.Join(prefixDir, dir, fname))
					continue
				}
				if stat, err := os.Stat(found); err != nil || stat.IsDir() {
					m.logger.Debugf("resolveSrcPaths(%q): ignoring entry %q (error or director)", srcFile, found)
					continue
				}
				resolvedPaths = append(resolvedPaths, found)
			}
		} else if ext == ".src" {
			found, err := findPath(filepath.Join(prefixDir, dir, fname))
			if err != nil {
				m.logger.Warnf("resolveSrcPaths(%q): file not found on disk %q", srcFile, filepath.Join(prefixDir, dir, fname))
				continue
			}
			m.logger.Infof("Collecting scripts from %q", found)

			inner, err := m.resolveSrcPaths(found, filepath.Dir(found))
			if err == nil {
				resolvedPaths = append(resolvedPaths, inner...)
			}
		}
	}

	return resolvedPaths, nil
}

var (
	bufferPool = sync.Pool{New: func() interface{} { b := new(bytes.Buffer); b.Grow(4096); return b }}
)

func (m *parseResultsManager) getConcurrency() int {
	max := runtime.NumCPU()

	if m.NumParserThreads <= max && m.NumParserThreads > 0 {
		return m.NumParserThreads
	}

	numWorkers := max
	if numWorkers > 2 {
		numWorkers = numWorkers / 2
	}

	return numWorkers
}

func (m *parseResultsManager) validateFiles(ctx context.Context, resolvedPaths []string) map[string][]SyntaxError {
	results := make(map[string][]SyntaxError)

	chanPaths := make(chan string, len(resolvedPaths))
	for _, r := range resolvedPaths {
		chanPaths <- r
	}
	close(chanPaths)

	var wg sync.WaitGroup

	numWorkers := m.getConcurrency()

	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()

			buf := bufferPool.Get().(*bytes.Buffer)
			defer bufferPool.Put(buf)

			decoder := m.decoderPool.Get().(*encoding.Decoder)
			defer m.decoderPool.Put(decoder)

			for r := range chanPaths {
				if ctx.Err() != nil {
					return
				}
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

				parsed := m.ValidateScript(r, buf.String())
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

func (m *parseResultsManager) validateFile(dPath string) ([]SyntaxError, error) {
	buf := bufferPool.Get().(*bytes.Buffer)
	defer bufferPool.Put(buf)

	decoder := m.decoderPool.Get().(*encoding.Decoder)
	defer m.decoderPool.Put(decoder)

	f, err := os.Open(dPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	decoder.Reset()
	translated := decoder.Reader(f)
	buf.Reset()
	_, err = buf.ReadFrom(translated)
	if err != nil {
		return nil, err
	}
	return m.ValidateScript(dPath, buf.String()), nil
}

func (m *parseResultsManager) ParseSource(ctx context.Context, srcFile string) ([]*ParseResult, error) {
	resolvedPaths, err := m.resolveSrcPaths(srcFile, filepath.Dir(srcFile))
	if err != nil {
		return nil, err
	}
	m.logger.Infof("Parsing %q. This might take a while.", srcFile)

	results := make([]*ParseResult, 0, len(resolvedPaths))

	chanPaths := make(chan string, len(resolvedPaths))
	for _, r := range resolvedPaths {
		chanPaths <- r
	}
	close(chanPaths)

	var wg sync.WaitGroup
	numWorkers := m.getConcurrency()
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()

			buf := bufferPool.Get().(*bytes.Buffer)
			defer bufferPool.Put(buf)

			decoder := m.decoderPool.Get().(*encoding.Decoder)
			defer m.decoderPool.Put(decoder)

			for r := range chanPaths {
				if ctx.Err() != nil {
					return
				}
				f, err := os.Open(r)
				if err != nil {
					continue
				}
				var lastMod time.Time
				if stat, err := f.Stat(); err == nil {
					lastMod = stat.ModTime()
				} else {
					lastMod = time.Now()
				}
				decoder.Reset()
				translated := decoder.Reader(f)
				buf.Reset()
				_, err = buf.ReadFrom(translated)
				f.Close()
				if err != nil {
					continue
				}

				parsed := m.ParseScript(r, buf.String(), lastMod)

				m.mtx.Lock()
				m.parseResults[parsed.Source] = parsed
				results = append(results, parsed)
				m.mtx.Unlock()
			}
		}(&wg)
	}

	wg.Wait()

	validations := m.validateFiles(ctx, resolvedPaths)

	for _, v := range results {
		if e, ok := validations[v.Source]; ok {
			v.SyntaxErrors = append(v.SyntaxErrors, e...)
		}
	}

	m.logger.Infof("Done parsing %q: %d scripts.", srcFile, len(results))
	return results, nil
}

func (m *parseResultsManager) ParseFile(dFile string) (*ParseResult, error) {
	buf := bufferPool.Get().(*bytes.Buffer)
	defer bufferPool.Put(buf)

	decoder := m.decoderPool.Get().(*encoding.Decoder)
	defer m.decoderPool.Put(decoder)

	f, err := os.Open(dFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	decoder.Reset()
	translated := decoder.Reader(f)
	buf.Reset()
	_, err = buf.ReadFrom(translated)
	if err != nil {
		return nil, err
	}

	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}

	parsed := m.ParseScript(dFile, buf.String(), stat.ModTime())

	m.mtx.Lock()
	m.parseResults[parsed.Source] = parsed
	m.mtx.Unlock()

	validations, err := m.validateFile(parsed.Source)

	if err == nil && len(validations) > 0 {
		parsed.SyntaxErrors = append(parsed.SyntaxErrors, validations...)
	}
	return parsed, nil
}

func resolveIntConstant(h SymbolProvider, c string) int {
	n, err := strconv.Atoi(c)
	if err == nil {
		return n
	}
	// probably symbol...?
	found, ok := h.LookupGlobalSymbol(c, SymbolConstant)
	if !ok {
		return -1
	}
	cs, ok := found.(symbol.Constant)
	if !ok {
		return -1
	}
	if n, err := strconv.Atoi(cs.Value); err == nil {
		return n
	}
	return -1
}

func SymbolToReadableCode(symbols SymbolProvider, s symbol.Symbol) string {
	var codeText string
	switch s := s.(type) {
	case symbol.ConstantArray:
		sb := strings.Builder{}
		resolvedSize := resolveIntConstant(symbols, s.ArraySizeText)
		s.Format(&sb, resolvedSize)
		codeText = sb.String()
	case symbol.ArrayVariable:
		sb := strings.Builder{}
		resolvedSize := resolveIntConstant(symbols, s.ArraySizeText)
		s.Format(&sb, resolvedSize)
		codeText = sb.String()
	default:
		codeText = s.String()
	}

	return codeText
}
