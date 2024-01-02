package langserver

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/kirides/DaedalusLanguageServer/protocol"
	"golang.org/x/exp/slices"
	"golang.org/x/text/encoding/charmap"
)

type nopLogger struct{}

func (nopLogger) Debugf(template string, args ...interface{}) {}
func (nopLogger) Infof(template string, args ...interface{})  {}
func (nopLogger) Warnf(template string, args ...interface{})  {}
func (nopLogger) Errorf(template string, args ...interface{}) {}

func TestParseSingleScript(t *testing.T) {
	script := `
	/// Summary line 1
	/// Summary Line 2
	/// ---------------
	/// summary line 4
	func void InitDamage() {
		const int dmg = 0;
		if (dmg) { return; };
		HookEngineF(6736583/*0x66CAC7*/, 5, _DMG_OnDmg_Post);
		const int oCNpc__OnDamage_Hit = 6710800;
		HookEngineF(oCNpc__OnDamage_Hit, 7, _DMG_OnDmg_Pre);
		dmg = 1;
	};`

	m := newParseResultsManager(nopLogger{})
	result := m.ParseScript("C:\\temp", script, time.Now())
	if _, ok := result.Functions[strings.ToUpper("InitDamage")]; !ok {
		t.Fail()
	}
}

func TestZParserExtender(t *testing.T) {
	script := `
	namespace zPE {
		func void InitDamage() {
			const int dmg = 0;
			if (dmg) { return; };
			HookEngineF(6736583/*0x66CAC7*/, 5, _DMG_OnDmg_Post);
			const int oCNpc__OnDamage_Hit = 6710800;
			HookEngineF(oCNpc__OnDamage_Hit, 7, _DMG_OnDmg_Pre);
			dmg = 1;
		};
	};
	`

	m := newParseResultsManager(nopLogger{})
	result := m.ParseScript("C:\\temp", script, time.Now())
	if _, ok := result.Namespaces[strings.ToUpper("zPE")].Functions[strings.ToUpper("InitDamage")]; !ok {
		for _, v := range result.SyntaxErrors {
			t.Logf("%d:%d %s: %s", v.Line, v.Column, v.ErrorCode.Code, v.ErrorCode.Description)
		}
		t.Fail()
	}
}

func TestParseSingleScriptFromFile(t *testing.T) {
	src := filepath.Join("testdata", "demo.d")
	fileBody, _ := os.ReadFile(src)
	script, _ := charmap.Windows1252.NewDecoder().Bytes(fileBody)

	m := newParseResultsManager(nopLogger{})
	result := m.ParseScript(src, string(script), time.Now())
	if _, ok := result.Functions[strings.ToUpper("Do")]; !ok {
		t.Fail()
	}
}

func TestGothicSrc(t *testing.T) {
	m := newParseResultsManager(nopLogger{})
	m.NumParserThreads = 8
	result, err := m.ParseSource(context.TODO(), filepath.Join("testdata", "Gothic.src"))
	if err != nil {
		t.Error(err)
	}
	_ = result
}

func BenchmarkGothicSrc(b *testing.B) {
	b.Logf("Running %d iterations", b.N)
	for i := 0; i < b.N; i++ {
		m := newParseResultsManager(nopLogger{})
		m.NumParserThreads = 1
		_, err := m.ParseSource(context.TODO(), filepath.Join("testdata", "Gothic.src"))
		if err != nil {
			b.Error(err)
		}
	}
}

func TestParseSemanticForGlobals(t *testing.T) {
	script := `var int CurrentLevel;
const int ADDONWORLD_ZEN = 4;

func void FuncName() {
	CurrentLevel = ADDONWORLD_ZEN;
};`

	data := []struct {
		expected string
	}{
		{expected: "ADDONWORLD_ZEN"},
		{expected: "CurrentLevel"},
	}
	m := newParseResultsManager(nopLogger{})
	result, _ := m.ParseSemanticsContentDataTypesRange(context.Background(), "C:\\temp", script, protocol.Range{Start: protocol.Position{Line: 0, Character: 0}, End: protocol.Position{Line: 999, Character: 999}}, DataAll)

	for _, v := range data {
		t.Run(fmt.Sprintf("Expect %q", v.expected), func(t *testing.T) {
			if !slices.ContainsFunc(result.GlobalIdentifiers, func(i Identifier) bool {
				return i.Name() == v.expected
			}) {
				t.Fail()
			}
		})
	}
}

func TestParseSemanticForLocals(t *testing.T) {
	script := `
func void FuncName() {
	var int CurrentLevel;
	const int ADDONWORLD_ZEN = 4;

	CurrentLevel = ADDONWORLD_ZEN;
};`

	data := []struct {
		expected string
	}{
		{expected: "ADDONWORLD_ZEN"},
		{expected: "CurrentLevel"},
	}
	m := newParseResultsManager(nopLogger{})
	result, _ := m.ParseSemanticsContentDataTypesRange(context.Background(), "C:\\temp", script, protocol.Range{Start: protocol.Position{Line: 0, Character: 0}, End: protocol.Position{Line: 999, Character: 999}}, DataAll)

	for _, v := range data {
		t.Run(fmt.Sprintf("Expect %q", v.expected), func(t *testing.T) {
			if !slices.ContainsFunc(result.GlobalIdentifiers, func(i Identifier) bool {
				return i.Name() == v.expected
			}) {
				t.Fail()
			}
		})
	}
}
