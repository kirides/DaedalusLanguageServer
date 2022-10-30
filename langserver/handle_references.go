package langserver

import (
	"bytes"
	"context"
	"regexp"

	dls "github.com/kirides/DaedalusLanguageServer"
	"github.com/kirides/DaedalusLanguageServer/daedalus/symbol"
	lsp "github.com/kirides/DaedalusLanguageServer/protocol"
	"go.lsp.dev/uri"
)

func (h *LspHandler) getAllReferences(req context.Context, params lsp.ReferenceParams) <-chan lsp.Location {
	resultCh := make(chan lsp.Location, 1)

	doc := uriToFilename(params.TextDocument.URI)
	if doc == "" {
		close(resultCh)
		return resultCh
	}

	content, err := h.bufferManager.GetBufferCtx(req, doc)
	if err != nil {
		close(resultCh)
		return resultCh
	}
	word := content.GetWordRangeAtPosition(params.Position)
	if word == "" {
		close(resultCh)
		return resultCh
	}

	parsed, err := h.parsedDocuments.GetCtx(req, doc)
	if err != nil {
		close(resultCh)
		return resultCh
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

	go func() {
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
		defer close(resultCh)

		if inScope {
			buffer.WriteString(string(content))
			indices := rxWord.FindAllIndex(buffer.Bytes(), -1)

			for _, startEnd := range indices {
				line, col := getLineCol(&buffer, startEnd[0], startEnd[1])
				if line > bbox.End.Line {
					return
				}

				if !bbox.InBBox(symbol.DefinitionIndex{Line: line, Column: col}) {
					continue
				}
				resultCh <- lsp.Location{
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
					}}
			}
		} else {
			for k, _ := range h.parsedDocuments.parseResults {
				if req.Err() != nil {
					return
				}
				err := h.parsedDocuments.LoadFileBuffer(req, k, &buffer)
				if err != nil {
					continue
				}

				indices := rxWord.FindAllIndex(buffer.Bytes(), -1)

				for _, startEnd := range indices {
					line, col := getLineCol(&buffer, startEnd[0], startEnd[1])

					resultCh <- lsp.Location{
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
						}}
				}
			}
		}
	}()

	return resultCh
}

func (h *LspHandler) handleTextDocumentReferences(req dls.RpcContext, params lsp.ReferenceParams) error {
	var result []lsp.Location

	refsCh := h.getAllReferences(req.Context(), params)
	for v := range refsCh {
		if result == nil {
			result = make([]lsp.Location, 0, 100)
		}
		result = append(result, v)
	}

	return req.Reply(req.Context(), result, nil)
}