package langserver

import (
	"context"
	"slices"
	"testing"

	"github.com/kirides/DaedalusLanguageServer/protocol"
)

func TestSemanticForParameterVariable(t *testing.T) {
	script := `func void FuncName(var int CurrentLevel) {
	CurrentLevel = ADDONWORLD_ZEN;
};`

	m := newParseResultsManager(nopLogger{})
	fullArea := protocol.Range{Start: protocol.Position{Line: 0, Character: 0}, End: protocol.Position{Line: 999, Character: 999}}

	result, _ := m.ParseSemanticsContentDataTypesRange(context.Background(), "C:\\temp", script, fullArea, DataAll)
	h := semanticTokensHandler{parsedDocuments: m}

	tokens := h.getSemanticTokens(result, &fullArea)

	if !slices.Equal(tokens, []uint32{0x1, 0x1, 0xc, 0x7, 0x0}) {
		t.Fatalf("Expected exactly one set of semantic tokens")
	}
}
