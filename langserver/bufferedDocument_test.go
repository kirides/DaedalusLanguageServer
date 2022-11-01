package langserver

import (
	"testing"

	lsp "github.com/kirides/DaedalusLanguageServer/protocol"
)

func TestBufferedDocumentWordRange(t *testing.T) {
	doc := BufferedDocument(`
    const int X_PRINTSCREEN_USE_SMOOTH_PRINTS = TRUE;
    
    func void X_AI_PrintScreen(var string txt, var int x, var int y, var string fontName, var int timeout) {
        AI_PrintScreen(txt, x, y, fontName, timeout);  
    };
    
    func void X_PrintScreen(var string txt, var int x, var int y, var string fontName, var int timeout) {
        if (X_PRINTSCREEN_USE_SMOOTH_PRINTS) {
            if (y <= 100 && y > 0) {
                PrintScreenS_Ext(txt, x, (PS_VMax * y / 100), COL_White, fontName);
            } else {
                PrintScreenS_Ext(txt, x, (PS_VMax * 30 / 100), COL_White, fontName);
                // PrintScreenCXS_Ext(txt, COL_White, fontName);
            };
        } else {
            PrintScreen(txt, x, y, fontName, timeout);
        };
    };`)

	word, pos := doc.GetWordRangeAtPosition(lsp.Position{
		Character: 29,
		Line:      8,
	})
	t.Logf("Found word: %q at %v", word, pos)
}
func TestBufferedDocumentMethodCall(t *testing.T) {
	doc := BufferedDocument(`
    const int X_PRINTSCREEN_USE_SMOOTH_PRINTS = TRUE;
    
    func void X_AI_PrintScreen(var string txt, var int x, var int y, var string fontName, var int timeout) {
        AI_PrintScreen(txt, x, y, fontName, timeout);  
    };
    
    func void X_PrintScreen(var string txt, var int x, var int y, var string fontName, var int timeout) {
        if (X_PRINTSCREEN_USE_SMOOTH_PRINTS) {
            if (y <= 100 && y > 0) {
                PrintScreenS_Ext(txt, x, (PS_VMax * y / 100), COL_White, fontName);
            } else {
                PrintScreenS_Ext(txt, x, (PS_VMax * 30 / 100), COL_White, fontName);
                // PrintScreenCXS_Ext(txt, COL_White, fontName);
            };
        } else {
            PrintScreen(txt, x, y, fontName, timeout);
        };
    };`)

	word := doc.GetMethodCall(lsp.Position{
		Character: 40,
		Line:      4,
	})
	t.Logf("Found Call: %q", word)
}
