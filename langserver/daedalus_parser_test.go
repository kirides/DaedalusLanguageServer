package langserver

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"golang.org/x/text/encoding/charmap"
)

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

	result := ParseScript("C:\\temp", script)
	b, _ := json.MarshalIndent(result, "", "  ")
	t.Logf("%s\n", b)
}

func TestParseSingleScriptFromFile(t *testing.T) {
	fileBody, _ := ioutil.ReadFile(`E:\Dev\Gothic II_Mods\_work\Data\Scripts\Content\LeGo\Int64.d`)
	script, _ := charmap.Windows1252.NewDecoder().Bytes(fileBody)

	result := ParseScript("C:\\temp", string(script))
	b, _ := json.MarshalIndent(result, "", "  ")
	t.Logf("%s\n", b)
}

func TestGothicSrc(t *testing.T) {
	paths := resolveSrcPaths(`E:\Dev\Gothic II_Mods\_work\Data\Scripts\Content\Gothic.src`, `E:\Dev\Gothic II_Mods\_work\Data\Scripts\Content`)

	// result := ParseScript("C:\\temp", string(script))
	b, _ := json.MarshalIndent(paths, "", "  ")
	t.Logf("%s\n", b)
}
