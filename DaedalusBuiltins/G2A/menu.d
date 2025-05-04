// Vanilla Menu externals with docu comments for DLS implementation
//
// Source: https://github.com/muczc1wek/daedalus-externals

var instance instance_help;

/// Updates `C_MENU_ITEM` choice box with the current setting
/// This specific boxes are hardcoded in the function itself: 
/// `MENUITEM_AUDIO_PROVIDER_CHOICE`,
/// `MENUITEM_VID_DEVICE_CHOICE`,
/// `MENUITEM_VID_RESOLUTION_CHOICE`
///
/// @param choicebox name of the one of the hardcoded choice boxes
func void Update_ChoiceBox(var string choicebox) {};

/// Applies the options for the performance, analyzes the system and changes the settings if necessary
func void Apply_Options_Performance() {};

/// Applies the options for the video, changes the resolution and graphics device if necessary
func void Apply_Options_Video() {};

/// [deprecated] Meant to apply the options for the audio, does nothing apart playing an apply sound
func void Apply_Options_Audio() {};

/// [deprecated] Meant to apply the options for the game, does nothing apart playing an apply sound
func void Apply_Options_Game() {};

/// [deprecated] Meant to apply the options for the controls, does nothing apart playing an apply sound
func void Apply_Options_Controls() {};

/// Plays a video
///
/// @param filename name of the video file
/// @return TRUE if the video was played successfully, FALSE otherwise
func int PlayVideo(var string filename) {};
