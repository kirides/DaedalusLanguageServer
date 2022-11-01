package langserver

import (
	"sort"
	"strings"

	dls "github.com/kirides/DaedalusLanguageServer"
	"github.com/kirides/DaedalusLanguageServer/daedalus/symbol"
	lsp "github.com/kirides/DaedalusLanguageServer/protocol"
)

type SemanticTokensLegend struct {
	TokenTypes     []lsp.SemanticTokenTypes     `json:"tokenTypes"`
	TokenModifiers []lsp.SemanticTokenModifiers `json:"tokenModifiers"`
}
type SemanticTokensOptions struct {
	Legend SemanticTokensLegend `json:"legend"`
	Range  bool                 `json:"range,omitempty"`
	Full   bool                 `json:"full,omitempty"`
}

var (
	semanticTypes = [...]lsp.SemanticTokenTypes{
		"namespace", "type", "class", "enum", "interface",
		"struct", "typeParameter", "parameter", "variable", "property", "enumMember",
		"event", "function", "method", "macro", "keyword", "modifier", "comment",
		"string", "number", "regexp", "operator",
	}
	semanticModifiers = [...]lsp.SemanticTokenModifiers{
		"declaration", "definition", "readonly", "static",
		"deprecated", "abstract", "async", "modification", "documentation", "defaultLibrary",
	}
)

func SemanticTypes() []lsp.SemanticTokenTypes {
	return semanticTypes[:]
}

func SemanticTypeFor(t lsp.SemanticTokenTypes) int {
	for i, v := range SemanticTypes() {
		if v == t {
			return i
		}
	}
	return -1
}

func SemanticModifiers() []lsp.SemanticTokenModifiers {
	return semanticModifiers[:]
}
func SemanticModifiersFor(mods ...lsp.SemanticTokenModifiers) uint32 {
	result := uint32(0)
	for _, v := range mods {
		for j, m := range SemanticModifiers() {
			if v == m {
				result |= 1 << j
			}
		}
	}
	return result
}

type token interface {
	Name() string
	Definition() symbol.Definition
}

type semanticTokensHandler struct {
	parsedDocuments *parseResultsManager
}

func (s semanticTokensHandler) semanticsForToken(t token) (int, uint32) {
	switch s := t.(type) {
	case symbol.Class:
		return SemanticTypeFor(lsp.SemanticTokenClass), 0
	case symbol.ProtoTypeOrInstance:
		mod := uint32(0)
		if !s.IsInstance {
			mod |= SemanticModifiersFor(lsp.SemanticTokenModifierAbstract)
		}
		return SemanticTypeFor(lsp.SemanticTokenClass), mod
	case FunctionWithIdentifiers, symbol.Function:
		return SemanticTypeFor(lsp.SemanticTokenFunction), 0
	case symbol.Constant, symbol.ConstantArray:
		return SemanticTypeFor(lsp.SemanticTokenVariable), SemanticModifiersFor(lsp.SemanticTokenModifierReadonly)
	case symbol.Variable, symbol.ArrayVariable:
		return SemanticTypeFor(lsp.SemanticTokenVariable), 0
	}
	return -1, 0
}

func (h *semanticTokensHandler) getSemanticTokens(r *SemanticParseResult, area *lsp.Range) []uint32 {
	type item struct {
		token token
		t     uint32
		m     uint32
	}
	syms := make([]item, 0, r.CountSymbols())

	onIdentifier := func(id *Identifier, isLocal bool) (t, m uint32, rok bool) {
		var identified symbol.Symbol
		isParam := false
		if isLocal {
			r.WalkScopedVariables(id.Definition().Start, func(t symbol.Symbol, isp bool) bool {
				if strings.EqualFold(t.Name(), id.Name()) {
					identified = t
					isParam = isp
					return false
				}
				return true
			})
		}

		if identified != nil {
			if !isParam {
				switch identified.(type) {
				case symbol.Constant, symbol.ConstantArray:
				default:
					// do not highlight local VARIABLEs
					return
				}
			}
			if typ, mod := h.semanticsForToken(identified); typ != -1 {
				t = uint32(typ)
				if isParam {
					typ = SemanticTypeFor(lsp.SemanticTokenParameter)
					if typ != -1 {
						t = uint32(typ)
					}
				}
				m = mod
				rok = true
			}
		} else {
			found, ok := h.parsedDocuments.LookupGlobalSymbol(strings.ToUpper(id.Name()), SymbolAll)
			if !ok {
				return
			}
			if typ, mod := h.semanticsForToken(found); typ != -1 {
				t = uint32(typ)
				// global symbol, mark static
				m = mod | SemanticModifiersFor(lsp.SemanticTokenModifierStatic)
				rok = true
			}
		}
		return
	}

	r.WalkSymbols(func(s token) error {
		itm := item{
			token: s,
			m:     SemanticModifiersFor(lsp.SemanticTokenModifierDeclaration),
		}

		switch s := s.(type) {
		case Identifier:
			t, m, ok := onIdentifier(&s, false)
			if ok {
				itm.t, itm.m = t, m
				syms = append(syms, itm)
			}
		case symbol.Class:
			if t := SemanticTypeFor(lsp.SemanticTokenClass); t != -1 {
				itm.t = uint32(t)
				syms = append(syms, itm)
			}
		case symbol.ProtoTypeOrInstance:
			if t := SemanticTypeFor(lsp.SemanticTokenClass); t != -1 {
				mod := uint32(0)
				if !s.IsInstance {
					mod |= SemanticModifiersFor(lsp.SemanticTokenModifierAbstract)
				}
				itm.t = uint32(t)
				itm.m = mod
				syms = append(syms, itm)
			}
		case FunctionWithIdentifiers:
			if t := SemanticTypeFor(lsp.SemanticTokenFunction); t != -1 {
				itm.t = uint32(t)
				// ignore functions, they're handled without semantics
				// syms = append(syms, itm)

				if t := SemanticTypeFor(lsp.SemanticTokenParameter); t != -1 {
					for _, v := range s.Parameters {
						syms = append(syms, item{
							token: v,
							t:     uint32(t),
							m:     SemanticModifiersFor(lsp.SemanticTokenModifierDeclaration),
						})
					}
				}
				if t := SemanticTypeFor(lsp.SemanticTokenVariable); t != -1 {
					for _, v := range s.LocalVariables {
						switch v.(type) {
						// We don't highlight local variables
						// case symbol.Variable, symbol.ArrayVariable:
						// 	syms = append(syms, item{
						// 		token: v,
						// 		t:     uint32(t),
						// 		m:     SemanticModifiersFor(lsp.SemanticTokenModifierDeclaration),
						// 	})
						case symbol.Constant, symbol.ConstantArray:
							syms = append(syms, item{
								token: v,
								t:     uint32(t),
								m:     SemanticModifiersFor(lsp.SemanticTokenModifierDeclaration, lsp.SemanticTokenModifierReadonly),
							})
						}
					}
				}
				for _, id := range s.Identifiers {
					t, m, ok := onIdentifier(&id, true)
					if ok {
						syms = append(syms, item{
							token: id,
							t:     t,
							m:     m,
						})
					}
				}
			}
		case symbol.Constant, symbol.ConstantArray:
			if t := SemanticTypeFor(lsp.SemanticTokenVariable); t != -1 {
				itm.t = uint32(t)
				itm.m |= SemanticModifiersFor(lsp.SemanticTokenModifierReadonly, lsp.SemanticTokenModifierStatic)
				syms = append(syms, itm)
			}
		case symbol.Variable, symbol.ArrayVariable:
			if t := SemanticTypeFor(lsp.SemanticTokenVariable); t != -1 {
				itm.t = uint32(t)
				itm.m |= SemanticModifiersFor(lsp.SemanticTokenModifierStatic)
				syms = append(syms, itm)
			}
		}

		return nil
	})

	sort.Slice(syms, func(i, j int) bool {
		if syms[i].token.Definition().Start.Line != syms[j].token.Definition().Start.Line {
			return syms[i].token.Definition().Start.Line < syms[j].token.Definition().Start.Line
		}
		return syms[i].token.Definition().Start.Column < syms[j].token.Definition().Start.Column
	})

	x := make([]uint32, 5*len(syms))
	var j int
	var last token
	for i := 0; i < len(syms); i++ {
		item := syms[i]

		start := item.token.Definition().Start.Line - 1
		if j == 0 {
			x[0] = uint32(syms[0].token.Definition().Start.Line - 1)
		} else {
			x[j] = uint32(start - (last.Definition().Start.Line - 1))
		}
		x[j+1] = uint32(item.token.Definition().Start.Column)
		if j > 0 && x[j] == 0 {
			x[j+1] = uint32(item.token.Definition().Start.Column - last.Definition().Start.Column)
		}
		x[j+2] = uint32(len(item.token.Name()))
		x[j+3] = item.t
		x[j+4] = item.m
		j += 5
		last = item.token
	}

	return x[:j]
}

func (h *LspHandler) handleSemanticTokensFull(req dls.RpcContext, params lsp.SemanticTokensParams) error {
	ws := h.GetWorkspace(params.TextDocument.URI)
	if ws == nil {
		return req.Reply(req.Context(), nil, nil)
	}
	handler := &semanticTokensHandler{parsedDocuments: ws.parsedDocuments}

	source := uriToFilename(params.TextDocument.URI)
	if source == "" {
		req.Reply(req.Context(), nil, nil)
		return nil
	}

	buf, err := ws.bufferManager.GetBufferCtx(req.Context(), source)
	if err != nil {
		req.Reply(req.Context(), nil, err)
		return err
	}

	parsed, err := handler.parsedDocuments.ParseSemanticsContentRange(req.Context(), source, string(buf), lsp.Range{})
	if err != nil {
		req.Reply(req.Context(), nil, err)
		return err
	}

	result := &lsp.SemanticTokens{
		Data: handler.getSemanticTokens(parsed, nil),
	}

	// h.LogInfo("MethodSemanticTokensFull:\n%v", result)
	req.Reply(req.Context(), result, nil)
	return nil
}

func (h *LspHandler) handleSemanticTokensRange(req dls.RpcContext, params lsp.SemanticTokensRangeParams) error {
	ws := h.GetWorkspace(params.TextDocument.URI)
	if ws == nil {
		return req.Reply(req.Context(), nil, nil)
	}

	handler := &semanticTokensHandler{parsedDocuments: ws.parsedDocuments}

	source := uriToFilename(params.TextDocument.URI)
	if source == "" {
		req.Reply(req.Context(), nil, nil)
		return nil
	}

	buf, err := ws.bufferManager.GetBufferCtx(req.Context(), source)
	if err != nil {
		req.Reply(req.Context(), nil, err)
		return err
	}

	parsed, err := handler.parsedDocuments.ParseSemanticsContentRange(req.Context(), source, string(buf), params.Range)
	if err != nil {
		req.Reply(req.Context(), nil, err)
		return err
	}

	result := &lsp.SemanticTokens{
		Data: handler.getSemanticTokens(parsed, nil),
	}

	// h.LogInfo("MethodSemanticTokensFull:\n%v", result)
	req.Reply(req.Context(), result, nil)
	return nil
}
