package langserver

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/kirides/DaedalusLanguageServer/daedalus/symbol"
	"github.com/kirides/DaedalusLanguageServer/javadoc"
	lsp "go.lsp.dev/protocol"
)

func checkInheritance(docs SymbolProvider, sym symbol.Symbol, symToImplement string) bool {
	if strings.EqualFold(sym.Name(), symToImplement) {
		return true
	}
	inst, ok := sym.(symbol.ProtoTypeOrInstance)
	if ok {
		if strings.EqualFold(inst.Parent, symToImplement) {
			return true
		}
		sym2, ok := docs.LookupGlobalSymbol(strings.ToUpper(inst.Parent), SymbolInstance|SymbolPrototype|SymbolClass)
		if ok {
			return checkInheritance(docs, sym2, symToImplement)
		}
	}
	return false
}

func noFilter(symbol.Symbol) bool { return true }

func getCompletionItemsByJavadoc(result []lsp.CompletionItem, h *LspHandler, parentDocu, paramDocu, varType string, docs SymbolProvider, params *lsp.CompletionParams) ([]lsp.CompletionItem, bool, error) {
	found := true
	if strings.HasPrefix(paramDocu, "{") { // if instance list directive
		instances, _ := javadoc.ParseWithin(paramDocu, "{", "}")

		done := map[string]struct{}{}
		for _, in := range instances {
			if _, ok := done[in]; ok {
				// only process each type once
				continue
			}
			done[in] = struct{}{}

			ci := getLocalsAndParams(h, params.TextDocument.URI, params.Position, in, noFilter)
			result = append(result, ci...)
			if strings.EqualFold(in, "C_NPC") {
				sortIdx := getHighestSortIndex(result)
				result = append(result, getDefaultC_NPCCompletions(docs, sortIdx)...)
			} else if strings.EqualFold(in, "C_ITEM") {
				sortIdx := getHighestSortIndex(result)
				result = append(result, getDefaultC_ITEMCompletions(docs, sortIdx)...)
			}

			docs.WalkGlobalSymbols(func(s symbol.Symbol) error {
				if checkInheritance(docs, s, in) {
					ci, err := completionItemFromSymbol(docs, s)
					if err != nil {
						return nil
					}
					ci.Detail += " (" + strings.ToUpper(in) + " instance)"
					result = append(result, ci)
				}
				return nil
			}, SymbolInstance)
		}
	} else if strings.HasPrefix(paramDocu, "[") { // if enum list directive
		enums, _ := javadoc.ParseWithin(paramDocu, "[", "]")

		ci := getLocalsAndParams(h, params.TextDocument.URI, params.Position, varType, noFilter)
		result = append(result, ci...)

		localSortIndex := getHighestSortIndex(result)
		for _, enum := range enums {
			symb, ok := docs.LookupGlobalSymbol(strings.ToUpper(enum), SymbolConstant)
			if !ok {
				continue
			}
			ci, err := completionItemFromSymbol(docs, symb)
			if err != nil {
				continue
			}
			localSortIndex++
			ci.Detail += " (enum)"
			ci.SortText = fmt.Sprintf("%000d", localSortIndex)
			result = append(result, ci)
		}
	} else if strings.HasPrefix(paramDocu, "<") { // if enum list directive
		signature, _ := javadoc.ParseWithin(paramDocu, "<", ">")
		if len(signature) == 0 {
			return result, false, nil
		}
		fn := javadoc.NewFunc(signature[0], signature[1:]...)

		ci := getLocalsAndParams(h, params.TextDocument.URI, params.Position, varType, func(s symbol.Symbol) bool {
			// Only show functions that match the signature or do not have a signature
			pdoc := javadoc.FindParam(parentDocu, s.Name())
			if pdoc == "" {
				return true
			}

			signature, _ := javadoc.ParseWithin(paramDocu, "<", ">")
			if len(signature) == 0 {
				return true
			}
			localFn := javadoc.NewFunc(signature[0], signature[1:]...)
			return fn.Equal(localFn)
		})
		result = append(result, ci...)

		docs.WalkGlobalSymbols(func(s symbol.Symbol) error {
			if fnSymb, ok := s.(symbol.Function); ok {
				if !fn.EqualSym(fnSymb) {
					return nil
				}
				ci, err := completionItemFromSymbol(docs, s)
				if err != nil {
					return nil
				}
				result = append(result, ci)
			}
			return nil
		}, SymbolFunction)
	} else {
		found = false
	}
	return result, found, nil
}

func getCompletionItemsSimple(result []lsp.CompletionItem, h *LspHandler, varType string, docs SymbolProvider, params *lsp.CompletionParams) ([]lsp.CompletionItem, bool, error) {
	ci := getLocalsAndParams(h, params.TextDocument.URI, params.Position, varType, noFilter)
	result = append(result, ci...)

	types := SymbolConstant | SymbolVariable | SymbolFunction
	if strings.EqualFold(varType, "int") {
		// instance, prototype and class all qualify as "int"
		types |= SymbolInstance | SymbolClass
	} else if strings.EqualFold(varType, "func") {
		// instance, prototype and class all qualify as "int"
		types = SymbolFunction
	}

	docs.WalkGlobalSymbols(func(s symbol.Symbol) error {
		useIt := false
		if typer, ok := s.(interface{ GetType() string }); ok {
			if strings.EqualFold(typer.GetType(), varType) {
				useIt = true
			} else if _, ok := s.(symbol.Function); ok && types == SymbolFunction {
				useIt = true
			}
		} else if _, ok := s.(symbol.ProtoTypeOrInstance); ok {
			useIt = true
		} else if _, ok := s.(symbol.Class); ok {
			useIt = true
		}
		if useIt {
			ci, err := completionItemFromSymbol(docs, s)
			if err != nil {
				return nil
			}
			result = append(result, ci)
		}
		return nil
	}, types)

	return result, true, nil
}

func getCompletionItemsComplex(result []lsp.CompletionItem, h *LspHandler, varType string, docs SymbolProvider, params *lsp.CompletionParams) ([]lsp.CompletionItem, bool, error) {
	ci := getLocalsAndParams(h, params.TextDocument.URI, params.Position, varType, noFilter)
	result = append(result, ci...)

	if strings.EqualFold(varType, "C_NPC") {
		sortIdx := getHighestSortIndex(result)
		result = append(result, getDefaultC_NPCCompletions(docs, sortIdx)...)
	} else if strings.EqualFold(varType, "C_ITEM") { // Also offer item global instance
		sortIdx := getHighestSortIndex(result)
		result = append(result, getDefaultC_ITEMCompletions(docs, sortIdx)...)
	}

	docs.WalkGlobalSymbols(func(s symbol.Symbol) error {
		if checkInheritance(docs, s, varType) {
			ci, err := completionItemFromSymbol(docs, s)
			if err != nil {
				return nil
			}
			result = append(result, ci)
		}
		return nil
	}, SymbolInstance)
	return result, true, nil
}

func getTypedCompletionItems(h *LspHandler, docs *parseResultsManager, sym symbol.Function, paramIndex int, params *lsp.CompletionParams) ([]lsp.CompletionItem, bool, error) {
	// Pre-allocate buffer
	result := make([]lsp.CompletionItem, 0, 200)

	varType := sym.Parameters[paramIndex].Type
	paramName := sym.Parameters[paramIndex].Name()
	docu := javadoc.FindParam(sym.Documentation(), paramName)

	if strings.HasPrefix(docu, "{") || strings.HasPrefix(docu, "[") || strings.HasPrefix(docu, "<") {
		return getCompletionItemsByJavadoc(result, h, sym.Documentation(), docu, varType, docs, params)
	} else {
		h.logger.Debugf("varType: %s", varType)
		if strings.EqualFold(varType, "int") || strings.EqualFold(varType, "string") || strings.EqualFold(varType, "float") || strings.EqualFold(varType, "func") {
			return getCompletionItemsSimple(result, h, varType, docs, params)
		} else { // it is an instance
			return getCompletionItemsComplex(result, h, varType, docs, params)
		}
	}
}

type callContext struct {
	Function symbol.Function
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

	fnName := strings.TrimSpace(methodCallLine[:idxParen])
	for i := len(fnName) - 1; i > 0; i-- {
		if !isIdentifier(fnName[i]) {
			if i+1 >= len(fnName) {
				return callContext{}, errors.New("index out of bounds")
			}
			fnName = fnName[i+1:]
			break
		}
	}

	fnName = strings.ToUpper(strings.TrimSpace(fnName))
	funcSymbol, found := h.parsedDocuments.LookupGlobalSymbol(fnName, SymbolFunction)
	if !found {
		return callContext{}, errors.New("no function symbol found")
	}
	sigCtx := methodCallLine[idxParen+1:]
	fn := funcSymbol.(symbol.Function)
	paramIdx := uint32(strings.Count(sigCtx, ","))
	if int(paramIdx) >= len(fn.Parameters) {
		return callContext{}, errors.New("index bigger than number of elements")
	}

	return callContext{
		Function: fn,
		ParamIdx: int(paramIdx),
	}, nil
}

// TODO: refactor this - duplicate code
func getSignatureCompletions(params *lsp.CompletionParams, h *LspHandler) ([]lsp.CompletionItem, bool, error) {
	ctx, err := getFunctionCallContext(h, params.TextDocument.URI, params.Position)
	if err != nil {
		return []lsp.CompletionItem{}, false, err
	}
	return getTypedCompletionItems(h, h.parsedDocuments, ctx.Function, ctx.ParamIdx, params)
}

func globalSignatureCompletionItem(docs SymbolProvider, symName string, sortIdx int, clsName string) (lsp.CompletionItem, bool) {
	sym, ok := docs.LookupGlobalSymbol(symName, SymbolVariable)
	if !ok {
		return lsp.CompletionItem{}, false
	}

	ci, err := completionItemFromSymbol(docs, sym)
	if err != nil {
		return lsp.CompletionItem{}, false
	}

	ci.Detail += " (global " + clsName + " instance)"
	ci.SortText = fmt.Sprintf("%000d", sortIdx)
	return ci, true
}

func getCompletions(docs SymbolProvider, sortIdx int, typ string, values []string) (completions []lsp.CompletionItem) {
	result := make([]lsp.CompletionItem, 0, len(values))

	for _, v := range values {
		if ci, ok := globalSignatureCompletionItem(docs, v, sortIdx+1, typ); ok {
			sortIdx++
			result = append(result, ci)
		}
	}

	return result
}

func getDefaultC_NPCCompletions(docs SymbolProvider, sortIdx int) (completions []lsp.CompletionItem) {
	return getCompletions(docs, sortIdx, "C_NPC", []string{"HERO", "SELF", "OTHER", "VICTIM"})
}

func getDefaultC_ITEMCompletions(docs SymbolProvider, sortIdx int) (completions []lsp.CompletionItem) {
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

func getLocalsAndParams(h *LspHandler, docURI lsp.URI, pos lsp.Position, varType string, filter func(symbol.Symbol) bool) []lsp.CompletionItem {
	parsedDoc, err := h.parsedDocuments.Get(h.uriToFilename(docURI))
	if err != nil {
		return []lsp.CompletionItem{}
	}

	validType := func(in string) bool {
		return strings.EqualFold(in, varType)
	}

	result := make([]lsp.CompletionItem, 0, 20)
	di := symbol.DefinitionIndex{Line: int(pos.Line), Column: int(pos.Character)}
	// locally scoped variables ordered at the top
	localSortIdx := 0
	for _, fn := range parsedDoc.Functions {
		if fn.BodyDefinition.InBBox(di) {
			for _, p := range fn.Parameters {
				if !validType(p.Type) {
					continue
				}
				if !filter(p) {
					continue
				}
				ci, err := completionItemFromSymbol(h.parsedDocuments, p)
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
				if !filter(p) {
					continue
				}
				ci, err := completionItemFromSymbol(h.parsedDocuments, p)
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
