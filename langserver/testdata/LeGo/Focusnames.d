/***********************************\
             FOCUSNAMES
\***********************************/

func int Focusnames_Color_Friendly() {
    return RGBA(0,   255, 0,   255); // Grün
};

func int Focusnames_Color_Neutral() {
    return RGBA(255, 255, 255, 255); // Weiß
};

func int Focusnames_Color_Angry() {
    return RGBA(255, 180, 0,   255); // Orange
};

func int Focusnames_Color_Hostile() {
    return RGBA(255, 0,   0,   255); // Rot
};

/// Only used in LeGo Focusnames!!!
func void _FocusNames_X_CallOnNpc(var c_npc slf, var string fnc) {
	const int fnId = -1;
	fnId = MEM_FindParserSymbol(fnc);
	if (fnId != -1) {
		MEM_PushInstParam(slf);
		MEM_CallByID(fnId);
	};
};

//========================================
// [intern] Färben der Namen
//========================================
func void _Focusnames() {
    var int col; col = -1; // Stupid pseudo-locals
	const int lastColor = -1;
    var oCNpc her; her = Hlp_GetNpc(hero);

	col = Focusnames_Color_Neutral();

	if(Hlp_Is_oCNpc(her.focus_vob)) {
		var c_npc oth; oth = MEM_PtrToInst(her.focus_vob);
		
		if (!Npc_IsDead(oth)) {
			var int att; att = Npc_GetPermAttitude(hero, oth);
			if (Npc_IsInState(oth, ZS_Attack)) {
				var c_npc oldTarget; oldTarget = Hlp_GetNpc(other);
				if (Npc_GetTarget(oth)) {
					if (Hlp_GetInstanceID(other) == Hlp_GetInstanceID(hero)) {
						if   (att == ATT_HOSTILE) { col = Focusnames_Color_Hostile(); } 
						else                      { col = Focusnames_Color_Angry();   };
					};
				};
				other = Hlp_GetNpc(oldTarget);
			};

			_FocusNames_X_CallOnNpc(oth, "B_GetPlayerCrime");
			if (MEM_PopIntResult() == CRIME_MURDER) {
				// Player is murderer, color focus read!
				col = Focusnames_Color_Hostile();
			};
			if (col == Focusnames_Color_Neutral()) {
				if (oth.aivar[AIV_PARTYMEMBER] == TRUE) { col = Focusnames_Color_Friendly(); }
				else if(att == ATT_FRIENDLY) { col = Focusnames_Color_Friendly(); }
				else if(att == ATT_NEUTRAL)  { col = Focusnames_Color_Neutral();  }
				else if(att == ATT_ANGRY)    { col = Focusnames_Color_Angry();    }
				else if(att == ATT_HOSTILE)  {
					col = Focusnames_Color_Hostile();
					if (oth.guild == GIL_BDT && C_PlayerIsFakeBandit(oth, hero)) {
						col = Focusnames_Color_Angry();
					} else if (CurrentLevel == ADDONWORLD_ZEN) {
					// ------- Banditenlager nach Bloodwyns Tod -----------
						if (oth.guild == GIL_BDT && SC_IsWellKnownBanit_AddonWorld) {
							col = Focusnames_Color_Angry();
						};
					};
				};
			};
		};
	}
	else if(Hlp_Is_oCItem(her.focus_vob)) {
		// var c_item itm; itm = MEM_PtrToInst(her.focus_vob);
		col = Focusnames_Color_Neutral();
	}
	else if (Hlp_Is_oCMobLockable(her.focus_vob)) {
		if(Hlp_Is_oCMobContainer(her.focus_vob)) {
			var oCMobContainer ocCont; ocCont = MEM_PtrToInst(her.focus_vob);
			// If there is something inside the container.
			if (ocCont.containList_next) {
				col = Focusnames_Color_Friendly();
			};
		};
		var oCMobLockable lockable; lockable = MEM_PtrToInst(her.focus_vob);

		// If it's locked and won't automatically open upon use.
		if (lockable.bitfield & oCMobLockable_bitfield_locked) {
			if (!(lockable.bitfield & oCMobLockable_bitfield_autoOpen)) { 
				col = Focusnames_Color_Hostile();
			};
			if (!Hlp_StrCmp("", lockable.pickLockStr)) {
				col = Focusnames_Color_Angry();
			};
		};
	};

	if (lastColor == col) {
		return;
	};
	lastColor = col;

    var int ptr; ptr = MEM_Alloc(4);
    MEM_WriteInt(ptr, col);
    CALL_IntParam(ptr);
    CALL__thiscall(MEM_ReadInt(screen_offset), zCView__SetFontColor);
    MEM_Free(ptr);
};