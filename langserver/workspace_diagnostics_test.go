package langserver

import (
	"context"
	"os"
	"path/filepath"
	"sort"
	"testing"
)

// Diagnostics of a workspace must not depend on keeping the parse trees around,
// and calls are checked against functions of other scripts.
func TestParseSourceReportsDiagnostics(t *testing.T) {
	dir := t.TempDir()
	write := func(name, content string) {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	write("a.d", "func void Target(var int a, var int b) {};\n")
	write("b.d", "func void Caller() {\n\tTarget(1);\n\tTarget(1, 2);\n\tTarget(1, 2, 3);\n\tvar int 1abc;\n\tx = ;\n};\n")
	write("c.d", "func void Fine() { Target(1, 2); };\n")
	write("t.src", "a.d\nb.d\nc.d\n")

	m := newParseResultsManager(nopLogger{})
	results, err := m.ParseSource(context.Background(), filepath.Join(dir, "t.src"))
	if err != nil {
		t.Fatal(err)
	}

	codes := map[string][]string{}
	for _, r := range results {
		for _, e := range r.SyntaxErrors {
			codes[filepath.Base(r.Source)] = append(codes[filepath.Base(r.Source)], e.ErrorCode.Code)
		}
		if len(r.pendingCalls) != 0 {
			t.Errorf("%s: pending calls were not resolved", r.Source)
		}
	}
	for k := range codes {
		sort.Strings(codes[k])
	}

	if len(codes["a.d"]) != 0 || len(codes["c.d"]) != 0 {
		t.Errorf("unexpected diagnostics: %v", codes)
	}
	want := map[string]int{"D0000": 1, "D0001": 1, "D0004": 1, "D0005": 1}
	got := map[string]int{}
	for _, c := range codes["b.d"] {
		got[c]++
	}
	for code, n := range want {
		if got[code] != n {
			t.Errorf("b.d: want %d x %s, got %v", n, code, codes["b.d"])
		}
	}

	// the same results as for a single opened document
	single := m.ParseAndValidateScript(filepath.Join(dir, "b.d"), "func void Caller() {\n\tTarget(1);\n\tTarget(1, 2);\n\tTarget(1, 2, 3);\n\tvar int 1abc;\n\tx = ;\n};\n")
	if len(single.SyntaxErrors) != len(codes["b.d"]) {
		t.Errorf("single document: %d diagnostics, workspace: %d", len(single.SyntaxErrors), len(codes["b.d"]))
	}
}
