package langserver

import (
	"encoding/json"
	"os"
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
	b, _ := json.MarshalIndent(result, "", "  ")
	t.Logf("%s\n", b)
}

func TestParseSingleScriptFromFile(t *testing.T) {
	fileBody, _ := os.ReadFile(`E:\Dev\Gothic II_Mods\_work\Data\Scripts\Content\LeGo\Int64.d`)
	script, _ := charmap.Windows1252.NewDecoder().Bytes(fileBody)

	m := newParseResultsManager(nopLogger{})
	result := m.ParseScript("C:\\temp", string(script))
	b, _ := json.MarshalIndent(result, "", "  ")
	t.Logf("%s\n", b)
}

func TestGothicSrc(t *testing.T) {
	m := newParseResultsManager(nopLogger{})
	paths := m.resolveSrcPaths(`E:\Dev\Gothic II_Mods\_work\Data\Scripts\Content\Gothic.src`, `E:\Dev\Gothic II_Mods\_work\Data\Scripts\Content`)

	// result := ParseScript("C:\\temp", string(script))
	b, _ := json.MarshalIndent(paths, "", "  ")
	t.Logf("%s\n", b)
}
