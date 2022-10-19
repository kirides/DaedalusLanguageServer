package langserver

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

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
	result := m.ParseScript("C:\\temp", script)
	if _, ok := result.Functions[strings.ToUpper("InitDamage")]; !ok {
		t.Fail()
	}
}

func TestParseSingleScriptFromFile(t *testing.T) {
	src := filepath.Join("testdata", "demo.d")
	fileBody, _ := os.ReadFile(src)
	script, _ := charmap.Windows1252.NewDecoder().Bytes(fileBody)

	m := newParseResultsManager(nopLogger{})
	result := m.ParseScript(src, string(script))
	if _, ok := result.Functions[strings.ToUpper("Do")]; !ok {
		t.Fail()
	}
}

func TestGothicSrc(t *testing.T) {
	m := newParseResultsManager(nopLogger{})
	m.NumParserThreads = 8
	result, err := m.ParseSource(filepath.Join("testdata", "Gothic.src"))
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
		_, err := m.ParseSource(filepath.Join("testdata", "Gothic.src"))
		if err != nil {
			b.Error(err)
		}
	}
}
