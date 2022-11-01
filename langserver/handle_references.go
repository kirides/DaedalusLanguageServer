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

func (h *LspWorkspace) getAllReferences(req context.Context, params lsp.ReferenceParams) <-chan lsp.Location {
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
	word, pos := content.GetWordRangeAtPosition(params.Position)
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
		Line:   int(pos.Start.Line) + 1,
		Column: int(pos.Start.Character),
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

		rxBlockComment := regexp.MustCompile(`(?s)\/\*.*?\*\/`)

		getLineCol := func(buf *bytes.Buffer, start, end int) ([]byte, int, int) {
			line := 1
			col := 0
			lineStart := 0
			for i := 0; i < buf.Len() && i < start; i++ {
				if buf.Bytes()[i] == '\n' {
					line++
					col = 0
					lineStart = i + 1
				} else {
					col++
				}
			}
			return buf.Bytes()[lineStart : lineStart+col], line, col
		}
		defer close(resultCh)

		isComment := func(buf *bytes.Buffer, blockComments [][]int, segment []byte, idxInSegment int, startEnd []int) bool {
			segment = bytes.TrimSpace(segment)
			if bytes.HasPrefix(segment, []byte("//")) {
				// definitely inside comment
				return true
			}

			isString := false
			for i := 0; i < len(segment); i++ {
				if segment[i] == '"' {
					isString = !isString
					continue
				}
				if !isString && bytes.HasPrefix(segment[i:], []byte("//")) {
					// is inline comment at the end somewhere
					return true
				}
			}
			if isString {
				// if it's not a comment and still in a string, then it's not a reference
				return true
			}

			for _, block := range blockComments {
				if startEnd[0] >= block[0] && startEnd[1] <= block[1] {
					// definitely inside block-comment
					return true
				}
			}
			return false
		}

		if inScope {
			buffer.WriteString(string(content))
			indices := rxWord.FindAllIndex(buffer.Bytes(), -1)
			var blockComments [][]int
			if len(indices) > 0 {
				blockComments = rxBlockComment.FindAllIndex(buffer.Bytes(), -1)
			}

			for _, startEnd := range indices {
				segment, line, col := getLineCol(&buffer, startEnd[0], startEnd[1])
				if line > bbox.End.Line {
					return
				}

				if !bbox.InBBox(symbol.DefinitionIndex{Line: line, Column: col}) {
					continue
				}

				if wordDefIndex.Line == line+1 && wordDefIndex.Column == col {
					// ignore itself
					continue
				}

				if isComment(&buffer, blockComments, segment, col, startEnd) {
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
			for k := range h.parsedDocuments.parseResults {
				if req.Err() != nil {
					return
				}
				err := h.parsedDocuments.LoadFileBuffer(req, k, &buffer)
				if err != nil {
					continue
				}

				indices := rxWord.FindAllIndex(buffer.Bytes(), -1)

				var blockComments [][]int
				if len(indices) > 0 {
					blockComments = rxBlockComment.FindAllIndex(buffer.Bytes(), -1)
				}
				for _, startEnd := range indices {
					segment, line, col := getLineCol(&buffer, startEnd[0], startEnd[1])

					if k == doc {
						if wordDefIndex.Line == line && wordDefIndex.Column == col {
							// ignore itself
							continue
						}
					}

					if isComment(&buffer, blockComments, segment, col, startEnd) {
						continue
					}

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
	ws := h.GetWorkspace(params.TextDocument.URI)
	if ws == nil {
		return req.Reply(req.Context(), nil, nil)
	}
	var result []lsp.Location

	refsCh := ws.getAllReferences(req.Context(), params)
	for v := range refsCh {
		if result == nil {
			result = make([]lsp.Location, 0, 100)
		}
		result = append(result, v)
	}

	return req.Reply(req.Context(), result, nil)
}
