package langserver

import (
	"bytes"
	"regexp"

	dls "github.com/kirides/DaedalusLanguageServer"
	"github.com/kirides/DaedalusLanguageServer/daedalus/symbol"
	lsp "github.com/kirides/DaedalusLanguageServer/protocol"
	"go.lsp.dev/uri"
)

func (h *LspHandler) handleTextDocumentReferences(req dls.RpcContext, params lsp.ReferenceParams) error {
	doc := uriToFilename(params.TextDocument.URI)
	if doc == "" {
		req.Reply(req.Context(), nil, nil)
		return nil
	}

	content, err := h.bufferManager.GetBufferCtx(req.Context(), doc)
	if err != nil {
		req.Reply(req.Context(), nil, nil)
		return nil
	}
	word := content.GetWordRangeAtPosition(params.Position)
	if word == "" {
		req.Reply(req.Context(), nil, nil)
		return nil
	}

	parsed, err := h.parsedDocuments.GetCtx(req.Context(), doc)
	if err != nil {
		req.Reply(req.Context(), nil, nil)
		return nil
	}
	wordDefIndex := symbol.DefinitionIndex{
		Line:   int(params.Position.Line) + 1,
		Column: int(params.Position.Character),
	}

	inScope := false
	bbox := symbol.NewDefinition(0, 0, 999999, 999999)

	if _, ok := parsed.LookupScopedVariable(wordDefIndex, word); ok {
		scopeSymbol, ok := parsed.FindContainingScope(wordDefIndex)
		if ok {
			inScope = true
			switch v := scopeSymbol.(type) {
			case symbol.Function:
				bbox = v.BodyDefinition
			case symbol.ProtoTypeOrInstance:
				bbox = v.BodyDefinition
			case symbol.Class:
				bbox = v.BodyDefinition
			}
		}
	}

	var result []lsp.Location
	buffer := bytes.Buffer{}
	rxWord := regexp.MustCompile(`(?i)\b` + regexp.QuoteMeta(word) + `\b`)

	getLineCol := func(buf *bytes.Buffer, start, end int) (int, int) {
		line := 1
		col := 0
		for i := 0; i < buf.Len() && i < start; i++ {
			if buf.Bytes()[i] == '\n' {
				line++
				col = 0
			} else {
				col++
			}
		}
		return line, col
	}

	if inScope {
		buffer.WriteString(string(content))
		indices := rxWord.FindAllIndex(buffer.Bytes(), -1)

		for _, startEnd := range indices {
			line, col := getLineCol(&buffer, startEnd[0], startEnd[1])
			if line > bbox.End.Line {
				break
			}

			if !bbox.InBBox(symbol.DefinitionIndex{Line: line, Column: col}) {
				continue
			}
			result = append(result, lsp.Location{
				URI: lsp.DocumentURI(uri.File(doc)),
				Range: lsp.Range{
					Start: lsp.Position{
						Character: uint32(col),
						Line:      uint32(line - 1),
					},
					End: lsp.Position{
						Character: uint32(col + len(word)),
						Line:      uint32(line - 1),
					},
				}})
		}

		return req.Reply(req.Context(), result, nil)
	}

	for k, _ := range h.parsedDocuments.parseResults {
		if req.Context().Err() != nil {
			return nil
		}
		err := h.parsedDocuments.LoadFileBuffer(req.Context(), k, &buffer)
		if err != nil {
			continue
		}

		indices := rxWord.FindAllIndex(buffer.Bytes(), -1)

		for _, startEnd := range indices {
			line, col := getLineCol(&buffer, startEnd[0], startEnd[1])

			result = append(result, lsp.Location{
				URI: lsp.DocumentURI(uri.File(k)),
				Range: lsp.Range{
					Start: lsp.Position{
						Character: uint32(col),
						Line:      uint32(line - 1),
					},
					End: lsp.Position{
						Character: uint32(col + len(word)),
						Line:      uint32(line - 1),
					},
				}})
		}
	}

	return req.Reply(req.Context(), result, nil)
}
