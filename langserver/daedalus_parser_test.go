package langserver

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"golang.org/x/text/encoding/charmap"
)

func TestParseSingleScript(t *testing.T) {
	script := `// Look for the function "DMG_OnDmg" to modify

	funcint DMG_OnDmg(var int victimPtr, var int attackerPtr, var int dmg, var int dmgDescriptorPtr) {	
		if (!attackerPtr && !victimPtr) { return dmg; };
	
		var oSDamageDescriptor dmgDesc; dmgDesc = _^(dmgDescriptorPtr);
	
		if (!attackerptr) {
			attackerptr = dmgDesc.attackerNpc;
		};
	
		if (dmg > 0 && attackerptr && victimPtr) {
			var c_npc attackerNpc; attackerNpc = _^(attackerptr);
			var c_npc victimNpc; victimNpc = _^(victimPtr);
	
			if (Npc_IsPlayer(attackerNpc) && (victimNpc.aivar[AIV_VictoryXPGiven] == FALSE)){
				SNC_OnDamage(dmg, dmgDescriptorPtr);
			};
	
			X_Apply_Buff_Burn(attackerNpc, victimNpc, dmgDesc);
		};
	
		// Diese Funktion anpassen, wenn ihr den Schaden verändern wollt! 'dmg' ist der von Gothic berechnete Schaden
		return dmg;
	};
		
	
	var int _DMG_DmgDesc;
	
	func void _DMG_OnDmg_Post() {
		EDI = DMG_OnDmg(EBP, MEM_ReadInt(MEM_ReadInt(ESP+644)+8), EDI, _DMG_DmgDesc);
	};
	
	func void _DMG_OnDmg_Pre() {
		_DMG_DmgDesc = ESI; // I'm preeeeetty sure it won't get moved in the meantime...
	};
	
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

	ParseScript("C:\\temp", string(script))
	// result := ParseScript("C:\\temp", string(script))
	// b, _ := json.MarshalIndent(result, "", "  ")
	// t.Logf("%s\n", b)
}
