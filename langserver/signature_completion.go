package langserver

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	lsp "go.lsp.dev/protocol"
)

func checkInheritance(docs *parseResultsManager, symbol Symbol, symToImplement string) bool {
	sym, ok := symbol.(ProtoTypeOrInstanceSymbol)
	if ok {
		if strings.EqualFold(sym.Parent, symToImplement) {
			return true
		}
		sym2, ok := docs.LookupGlobalSymbol(strings.ToUpper(sym.Parent), SymbolInstance|SymbolPrototype|SymbolClass)
		if ok {
			return checkInheritance(docs, sym2, symToImplement)
		}
	}
	return false
}

func getTypedCompletionItems(h *LspHandler, docs *parseResultsManager, symbol FunctionSymbol, paramIndex int, params *lsp.CompletionParams) ([]lsp.CompletionItem, error) {
	result := make([]lsp.CompletionItem, 0, 200)

	varType := symbol.Parameters[paramIndex].Type
	docu := findJavadocParam(symbol.Documentation(), symbol.Parameters[paramIndex].Name())

	if strings.HasPrefix(docu, "{") || strings.HasPrefix(docu, "[") {
		if strings.HasPrefix(docu, "{") { // if instance list directive
			instances, _ := parseJavadocWithinTokens(docu, "{", "}")

			done := map[string]struct{}{}
			for _, in := range instances {
				if _, ok := done[in]; ok {
					// only process each type once
					continue
				}
				done[in] = struct{}{}

				ci := getLocalsAndParams(h, params.TextDocument.URI, params.Position, in)
				result = append(result, ci...)
				if strings.EqualFold(in, "C_NPC") {
					sortIdx := getHighestSortIndex(result)
					result = append(result, getDefaultC_NPCCompletions(docs, sortIdx)...)
				} else if strings.EqualFold(in, "C_ITEM") {
					sortIdx := getHighestSortIndex(result)
					result = append(result, getDefaultC_ITEMCompletions(docs, sortIdx)...)
				}

				docs.WalkGlobalSymbols(func(s Symbol) error {
					if checkInheritance(docs, s, in) {
						ci, err := completionItemFromSymbol(s)
						if err != nil {
							return nil
						}
						ci.Detail += " (" + strings.ToUpper(in) + " instance)"
						result = append(result, ci)
					}
					return nil
				}, SymbolInstance)
			}
		} else if strings.HasPrefix(docu, "[") { // if enum list directive
			enums, _ := parseJavadocWithinTokens(docu, "[", "]")

			localSortIndex := 0
			for _, enum := range enums {
				symb, ok := docs.LookupGlobalSymbol(strings.ToUpper(enum), SymbolConstant)
				if !ok {
					continue
				}
				ci, err := completionItemFromSymbol(symb)
				if err != nil {
					continue
				}
				localSortIndex++
				ci.Detail += " (enum)"
				ci.SortText = fmt.Sprintf("%000d", localSortIndex)
				result = append(result, ci)
			}
		}
		return result, nil
	} else {
		h.logger.Debugf("varType: %s", varType)
		// TODO: Add support for floats (zParserExtender)
		if strings.EqualFold(varType, "int") || strings.EqualFold(varType, "string") /*|| strings.EqualFold(varType, "float")*/ {

			ci := getLocalsAndParams(h, params.TextDocument.URI, params.Position, varType)
			result = append(result, ci...)

			docs.WalkGlobalSymbols(func(s Symbol) error {
				if strings.EqualFold(s.(ConstantSymbol).Type, varType) {
					ci, err := completionItemFromSymbol(s)
					if err != nil {
						return nil
					}
					result = append(result, ci)
				}
				return nil
			}, SymbolConstant)

			return result, nil
		} else { // it is an instance
			ci := getLocalsAndParams(h, params.TextDocument.URI, params.Position, varType)
			result = append(result, ci...)

			if strings.EqualFold(varType, "C_NPC") {
				sortIdx := getHighestSortIndex(result)
				result = append(result, getDefaultC_NPCCompletions(docs, sortIdx)...)
			} else if strings.EqualFold(varType, "C_ITEM") { // Also offer item global instance
				sortIdx := getHighestSortIndex(result)
				result = append(result, getDefaultC_ITEMCompletions(docs, sortIdx)...)
			}

			docs.WalkGlobalSymbols(func(s Symbol) error {
				if checkInheritance(docs, s, varType) {
					ci, err := completionItemFromSymbol(s)
					if err != nil {
						return nil
					}
					result = append(result, ci)
				}
				return nil
			}, SymbolInstance)
			return result, nil
		}
	}
}

type callContext struct {
	Function FunctionSymbol
	ParamIdx int
}

func getFunctionCallContext(h *LspHandler, docUri lsp.URI, pos lsp.Position) (callContext, error) {
	doc := h.bufferManager.GetBuffer(h.uriToFilename(docUri))

	methodCallLine := doc.GetMethodCall(pos)
	// The expected method call turned out to be a `func void something( ... )` -> a function definition
	if rxFunctionDef.MatchString(methodCallLine) {
		return callContext{}, errors.New("function definition")
	}
	methodCallLine = rxStringValues.ReplaceAllLiteralString(methodCallLine, "")
	oldLen := -1
	for len(methodCallLine) != oldLen {
		oldLen = len(methodCallLine)
		methodCallLine = rxFuncCall.ReplaceAllLiteralString(methodCallLine, "")
	}
	// If for some reason the parenthesis of the methodcall went missing
	idxParen := strings.LastIndexByte(methodCallLine, '(')
	if idxParen < 0 {
		return callContext{}, errors.New("parenthesis went missing")
	}
	word := ""
	for i := idxParen - 1; i > 0; i-- {
		if !isIdentifier(methodCallLine[i]) {
			start := i + 1
			if start+idxParen > len(methodCallLine) {
				return callContext{}, errors.New("index out of bounds")
			}
			word = methodCallLine[start : start+idxParen]
		}
	}
	if word == "" {
		word = methodCallLine[:idxParen]
	}
	word = strings.ToUpper(strings.TrimSpace(word))
	funcSymbol, found := h.parsedDocuments.LookupGlobalSymbol(word, SymbolFunction)
	if !found {
		return callContext{}, errors.New("no function symbol found")
	}
	sigCtx := methodCallLine[idxParen+1:]
	fn := funcSymbol.(FunctionSymbol)
	paramIdx := uint32(strings.Count(sigCtx, ","))
	if int(paramIdx) >= len(fn.Parameters) {
		return callContext{}, errors.New("index bigger than number of elements")
	}

	return callContext{
		Function: fn,
		ParamIdx: int(paramIdx),
	}, nil
}

// // TODO: refactor this - duplicate code
func getSignatureCompletions(params *lsp.CompletionParams, h *LspHandler) ([]lsp.CompletionItem, error) {

	ctx, err := getFunctionCallContext(h, params.TextDocument.URI, params.Position)
	if err != nil {
		return []lsp.CompletionItem{}, err
	}

	// TODO: Lookup parameter from current method context
	return getTypedCompletionItems(h, h.parsedDocuments, ctx.Function, ctx.ParamIdx, params)
}

func globalSignatureCompletionItem(docs *parseResultsManager, symName string, sortIdx int, clsName string) (lsp.CompletionItem, bool) {
	sym, ok := docs.LookupGlobalSymbol(symName, SymbolVariable)
	if !ok {
		return lsp.CompletionItem{}, false
	}

	ci, err := completionItemFromSymbol(sym)
	if err != nil {
		return lsp.CompletionItem{}, false
	}

	ci.Detail += " (global " + clsName + " instance)"
	ci.SortText = fmt.Sprintf("%000d", sortIdx)
	return ci, true
}

func getCompletions(docs *parseResultsManager, sortIdx int, typ string, values []string) (completions []lsp.CompletionItem) {
	result := make([]lsp.CompletionItem, 0, len(values))

	for _, v := range values {
		if ci, ok := globalSignatureCompletionItem(docs, v, sortIdx+1, typ); ok {
			sortIdx++
			result = append(result, ci)
		}
	}

	return result
}

func getDefaultC_NPCCompletions(docs *parseResultsManager, sortIdx int) (completions []lsp.CompletionItem) {
	return getCompletions(docs, sortIdx, "C_NPC", []string{"HERO", "SELF", "OTHER", "VICTIM"})
}

func getDefaultC_ITEMCompletions(docs *parseResultsManager, sortIdx int) (completions []lsp.CompletionItem) {
	return getCompletions(docs, sortIdx, "C_ITEM", []string{"ITEM"})
}

func getHighestSortIndex(items []lsp.CompletionItem) int {
	if len(items) == 0 {
		return 0
	}

	max := 0
	for _, v := range items {
		if v.SortText == "" {
			continue
		}
		n, err := strconv.Atoi(v.SortText)
		if err != nil {
			continue
		}
		if n > max {
			max = n
		}
	}
	return max
}

func getLocalsAndParams(h *LspHandler, docURI lsp.URI, pos lsp.Position, varType string) []lsp.CompletionItem {
	parsedDoc, err := h.parsedDocuments.Get(h.uriToFilename(docURI))
	if err != nil {
		return []lsp.CompletionItem{}
	}

	validType := func(in string) bool {
		return strings.EqualFold(in, varType)
	}

	result := make([]lsp.CompletionItem, 0, 20)
	di := DefinitionIndex{Line: int(pos.Line), Column: int(pos.Character)}
	// locally scoped variables ordered at the top
	localSortIdx := 0
	for _, fn := range parsedDoc.Functions {
		if fn.BodyDefinition.InBBox(di) {
			for _, p := range fn.Parameters {
				if !validType(p.Type) {
					continue
				}
				ci, err := completionItemFromSymbol(p)
				if err != nil {
					continue
				}
				localSortIdx++
				ci.Detail += " (" + p.Type + " parameter)"
				ci.SortText = fmt.Sprintf("%000d", localSortIdx)

				result = append(result, ci)
			}
			for _, p := range fn.LocalVariables {
				typer, ok := p.(interface{ GetType() string })
				if !ok {
					continue
				}
				if !validType(typer.GetType()) {
					continue
				}
				ci, err := completionItemFromSymbol(p)
				if err != nil {
					continue
				}
				localSortIdx++
				ci.Detail += " (" + typer.GetType() + " local)"
				ci.SortText = fmt.Sprintf("%000d", localSortIdx)
				result = append(result, ci)
			}
			break
		}
	}
	return result
}
