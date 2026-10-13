package langserver

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/kirides/DaedalusLanguageServer/daedalus/symbol"
	protocol "github.com/kirides/DaedalusLanguageServer/protocol"
	"golang.org/x/text/encoding/charmap"
)

// TestCorpusDump writes a fingerprint of everything extracted from a script corpus.
// It is meant to compare two implementations (e.g. before/after a grammar change):
//
//	DLS_CORPUS  directories (separated by ';') containing .d files
//	DLS_DUMP    output file, one line per script: name, errors, symbols, identifiers
//	DLS_DUMP_FILE  (optional) print the full dump of scripts whose path contains this text
func TestCorpusDump(t *testing.T) {
	out := os.Getenv("DLS_DUMP")
	roots := os.Getenv("DLS_CORPUS")
	if out == "" || roots == "" {
		t.Skip("DLS_CORPUS / DLS_DUMP not set")
	}

	var files []string
	for _, r := range strings.Split(roots, ";") {
		filepath.WalkDir(r, func(p string, d fs.DirEntry, err error) error {
			if err == nil && !d.IsDir() && strings.HasSuffix(strings.ToLower(p), ".d") {
				files = append(files, p)
			}
			return nil
		})
	}
	sort.Strings(files)

	m := newParseResultsManager(nopLogger{})
	contents := make([]string, len(files))
	for i, f := range files {
		raw, err := os.ReadFile(f)
		if err != nil {
			t.Fatal(err)
		}
		body, _ := charmap.Windows1252.NewDecoder().Bytes(raw)
		contents[i] = string(body)
		// register global symbols for the validation of function calls
		parsed := m.ParseScript(f, contents[i], fileModTime(f))
		parsed.Ast = nil
		m.mtx.Lock()
		m.parseResults[f] = parsed
		m.mtx.Unlock()
	}

	provider := &orderedProvider{parseResultsManager: m, files: files}

	hash := func(v interface{}) (string, []byte) {
		b, err := json.Marshal(v)
		if err != nil {
			t.Fatal(err)
		}
		sum := sha256.Sum256(b)
		return hex.EncodeToString(sum[:6]), b
	}

	type ident struct {
		Name       string
		Start, End [2]int
	}

	var sb strings.Builder
	whole := protocol.Range{Start: protocol.Position{Line: 0, Character: 0}, End: protocol.Position{Line: 1 << 20, Character: 1 << 20}}
	filter := os.Getenv("DLS_DUMP_FILE")
	for i, f := range files {
		// same as parseResultsManager.ParseAndValidateScript, but with deterministic symbol lookups
		stateful := NewDaedalusStatefulListener(f, m)
		validating := NewDaedalusValidatingListener(f, provider)
		errListener := &SyntaxErrorListener{}
		m.ParseScriptListener(f, contents[i], combineListeners(stateful, validating), errListener)
		res := &ParseResult{
			SyntaxErrors: errListener.SyntaxErrors,
			Instances:    stateful.Globals.Instances, GlobalVariables: stateful.Globals.Variables, GlobalConstants: stateful.Globals.Constants,
			Functions: stateful.Globals.Functions, Classes: stateful.Globals.Classes, Prototypes: stateful.Globals.Prototypes,
			Namespaces: stateful.Namespaces,
		}
		errs := make([]string, 0, len(res.SyntaxErrors))
		for _, e := range res.SyntaxErrors {
			errs = append(errs, fmt.Sprintf("%s@%d:%d", e.ErrorCode.Code, e.Line, e.Column))
		}
		sort.Strings(errs)
		symbols := map[string]interface{}{
			"Instances": res.Instances, "Variables": res.GlobalVariables, "Constants": res.GlobalConstants,
			"Functions": res.Functions, "Classes": res.Classes, "Prototypes": res.Prototypes, "Namespaces": res.Namespaces,
		}

		sem, err := m.ParseSemanticsContentDataTypesRange(context.Background(), f, contents[i], whole, DataAll)
		if err != nil {
			t.Fatal(err)
		}
		idents := make([]ident, 0, len(sem.GlobalIdentifiers))
		for _, id := range sem.GlobalIdentifiers {
			d := id.Definition()
			idents = append(idents, ident{id.Name(), [2]int{d.Start.Line, d.Start.Column}, [2]int{d.End.Line, d.End.Column}})
		}
		semSymbols := map[string]interface{}{
			"Instances": sem.Instances, "Variables": sem.GlobalVariables, "Constants": sem.GlobalConstants,
			"Functions": sem.Functions, "Classes": sem.Classes, "Prototypes": sem.Prototypes,
		}

		hErr, bErr := hash(errs)
		hSym, bSym := hash(symbols)
		hSem, bSem := hash(semSymbols)
		hId, bId := hash(idents)
		fmt.Fprintf(&sb, "%s err=%s sym=%s semsym=%s ids=%s\n", f, hErr, hSym, hSem, hId)
		if filter != "" && strings.Contains(f, filter) {
			fmt.Printf("== %s\nERRORS %s\nSYMBOLS %s\nSEMSYMBOLS %s\nIDENTS %s\n", f, bErr, bSym, bSem, bId)
		}
	}
	if err := os.WriteFile(out, []byte(sb.String()), 0o644); err != nil {
		t.Fatal(err)
	}
	t.Logf("fingerprinted %d scripts", len(files))
}

func fileModTime(p string) (t time.Time) {
	if st, err := os.Stat(p); err == nil {
		return st.ModTime()
	}
	return
}

// orderedProvider looks up global symbols in a stable order (the manager iterates a map).
type orderedProvider struct {
	*parseResultsManager
	files []string
}

func (o *orderedProvider) LookupGlobalSymbol(name string, types SymbolType) (symbol.Symbol, bool) {
	o.mtx.RLock()
	defer o.mtx.RUnlock()
	for _, f := range o.files {
		if p, ok := o.parseResults[f]; ok {
			if s, ok := p.LookupGlobalSymbol(name, types); ok {
				return s, true
			}
		}
	}
	return nil, false
}
