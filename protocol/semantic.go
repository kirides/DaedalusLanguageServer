package protocol

// SemanticTokenTypes represents a type of semantic token.
//
// @since 3.16.0.
type SemanticTokenTypes string

// list of SemanticTokenTypes.
const (
	SemanticTokenNamespace SemanticTokenTypes = "namespace"

	// Represents a generic type. Acts as a fallback for types which
	// can't be mapped to a specific type like class or enum.
	SemanticTokenType          SemanticTokenTypes = "type"
	SemanticTokenClass         SemanticTokenTypes = "class"
	SemanticTokenEnum          SemanticTokenTypes = "enum"
	SemanticTokenInterface     SemanticTokenTypes = "interface"
	SemanticTokenStruct        SemanticTokenTypes = "struct"
	SemanticTokenTypeParameter SemanticTokenTypes = "typeParameter"
	SemanticTokenParameter     SemanticTokenTypes = "parameter"
	SemanticTokenVariable      SemanticTokenTypes = "variable"
	SemanticTokenProperty      SemanticTokenTypes = "property"
	SemanticTokenEnumMember    SemanticTokenTypes = "enumMember"
	SemanticTokenEvent         SemanticTokenTypes = "event"
	SemanticTokenFunction      SemanticTokenTypes = "function"
	SemanticTokenMethod        SemanticTokenTypes = "method"
	SemanticTokenMacro         SemanticTokenTypes = "macro"
	SemanticTokenKeyword       SemanticTokenTypes = "keyword"
	SemanticTokenModifier      SemanticTokenTypes = "modifier"
	SemanticTokenComment       SemanticTokenTypes = "comment"
	SemanticTokenString        SemanticTokenTypes = "string"
	SemanticTokenNumber        SemanticTokenTypes = "number"
	SemanticTokenRegexp        SemanticTokenTypes = "regexp"
	SemanticTokenOperator      SemanticTokenTypes = "operator"
)

// SemanticTokenModifiers represents a modifiers of semantic token.
//
// @since 3.16.0.
type SemanticTokenModifiers string

// list of SemanticTokenModifiers.
const (
	SemanticTokenModifierDeclaration    SemanticTokenModifiers = "declaration"
	SemanticTokenModifierDefinition     SemanticTokenModifiers = "definition"
	SemanticTokenModifierReadonly       SemanticTokenModifiers = "readonly"
	SemanticTokenModifierStatic         SemanticTokenModifiers = "static"
	SemanticTokenModifierDeprecated     SemanticTokenModifiers = "deprecated"
	SemanticTokenModifierAbstract       SemanticTokenModifiers = "abstract"
	SemanticTokenModifierAsync          SemanticTokenModifiers = "async"
	SemanticTokenModifierModification   SemanticTokenModifiers = "modification"
	SemanticTokenModifierDocumentation  SemanticTokenModifiers = "documentation"
	SemanticTokenModifierDefaultLibrary SemanticTokenModifiers = "defaultLibrary"
)
