// Code generated from Daedalus.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Daedalus

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type DaedalusParser struct {
	*antlr.BaseParser
}

var DaedalusParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func daedalusParserInit() {
	staticData := &DaedalusParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "','", "'('", "')'", "'['", "']'", "'{'", "'}'", "'&'",
		"'&&'", "'|'", "'||'", "'+'", "'-'", "'/'", "'*'", "'~'", "'!'", "'='",
		"'<'", "'>'", "'<<'", "'>>'", "'<='", "'>='", "'=='", "'!='", "'%'",
		"'+='", "'-='", "'*='", "'/='", "'&='", "'|='", "'.'", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "IntegerLiteral", "FloatLiteral", "StringLiteral", "Const", "Var",
		"If", "Int", "Else", "Func", "StringKeyword", "Class", "Void", "Return",
		"Float", "Prototype", "Instance", "Namespace", "Null", "Meta", "Comma",
		"LeftParen", "RightParen", "LeftBracket", "RightBracket", "LeftBrace",
		"RightBrace", "BitAnd", "And", "BitOr", "Or", "Plus", "Minus", "Div",
		"Star", "Tilde", "Not", "Assign", "Less", "Greater", "Shl", "Shr", "LessEq",
		"GreaterEq", "Equal", "NotEqual", "Percent", "PlusAssign", "MinusAssign",
		"StarAssign", "DivAssign", "AndAssign", "OrAssign", "Dot", "Semi", "Identifier",
		"Whitespace", "Newline", "BlockComment", "LineComment",
	}
	staticData.RuleNames = []string{
		"daedalusFile", "blockDef", "inlineDef", "functionDef", "constDef",
		"classDef", "prototypeDef", "instanceDef", "namespaceDef", "mainBlock",
		"contentBlock", "varDecl", "metaValue", "zParserExtenderMeta", "zParserExtenderMetaBlock",
		"constArrayDef", "constArrayAssignment", "constValueDef", "constValueAssignment",
		"varArrayDecl", "varValueDecl", "parameterList", "parameterDecl", "statementBlock",
		"elseBlock", "elseIfBlock", "ifBlock", "ifBlockStatement", "returnStatement",
		"expression", "arrayIndex", "arraySize", "typeReference", "nameNode",
		"parentReference",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 59, 400, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 1, 0, 3, 0, 72, 8, 0, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 81, 8, 1, 1, 1, 1, 1, 1, 2,
		1, 2, 3, 2, 87, 8, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		4, 1, 4, 1, 4, 1, 4, 3, 4, 101, 8, 4, 1, 4, 1, 4, 1, 4, 3, 4, 106, 8, 4,
		5, 4, 108, 8, 4, 10, 4, 12, 4, 111, 9, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 5, 5, 119, 8, 5, 10, 5, 12, 5, 122, 9, 5, 1, 5, 1, 5, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 137, 8, 7,
		10, 7, 12, 7, 140, 9, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 146, 8, 7, 1, 8,
		1, 8, 1, 8, 1, 8, 5, 8, 152, 8, 8, 10, 8, 12, 8, 155, 9, 8, 1, 8, 1, 8,
		1, 9, 3, 9, 160, 8, 9, 1, 9, 4, 9, 163, 8, 9, 11, 9, 12, 9, 164, 1, 10,
		1, 10, 3, 10, 169, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 175, 8, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 181, 8, 11, 5, 11, 183, 8, 11, 10, 11,
		12, 11, 186, 9, 11, 1, 12, 4, 12, 189, 8, 12, 11, 12, 12, 12, 190, 1, 13,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 5, 14, 201, 8, 14, 10,
		14, 12, 14, 204, 9, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 220, 8, 16, 10,
		16, 12, 16, 223, 9, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18,
		1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1,
		21, 1, 21, 5, 21, 244, 8, 21, 10, 21, 12, 21, 247, 9, 21, 3, 21, 249, 8,
		21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22,
		260, 8, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 267, 8, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 3, 23, 273, 8, 23, 5, 23, 275, 8, 23, 10, 23, 12,
		23, 278, 9, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25,
		1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 5, 27, 296, 8,
		27, 10, 27, 12, 27, 299, 9, 27, 1, 27, 3, 27, 302, 8, 27, 1, 28, 1, 28,
		3, 28, 306, 8, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 5, 29, 324,
		8, 29, 10, 29, 12, 29, 327, 9, 29, 3, 29, 329, 8, 29, 1, 29, 1, 29, 1,
		29, 1, 29, 1, 29, 1, 29, 3, 29, 337, 8, 29, 1, 29, 1, 29, 1, 29, 1, 29,
		1, 29, 1, 29, 3, 29, 345, 8, 29, 3, 29, 347, 8, 29, 3, 29, 349, 8, 29,
		1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29,
		1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 5, 29, 381,
		8, 29, 10, 29, 12, 29, 384, 9, 29, 1, 30, 1, 30, 3, 30, 388, 8, 30, 1,
		31, 1, 31, 3, 31, 392, 8, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34,
		1, 34, 2, 190, 202, 1, 58, 35, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
		24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58,
		60, 62, 64, 66, 68, 0, 10, 2, 0, 31, 32, 35, 36, 5, 0, 7, 7, 9, 12, 14,
		17, 19, 19, 55, 55, 2, 0, 33, 34, 46, 46, 1, 0, 31, 32, 1, 0, 40, 41, 2,
		0, 38, 39, 42, 43, 1, 0, 44, 45, 2, 0, 37, 37, 47, 52, 6, 0, 7, 7, 9, 10,
		12, 12, 14, 14, 16, 16, 55, 55, 5, 0, 5, 5, 7, 7, 9, 12, 14, 19, 55, 55,
		423, 0, 71, 1, 0, 0, 0, 2, 80, 1, 0, 0, 0, 4, 86, 1, 0, 0, 0, 6, 90, 1,
		0, 0, 0, 8, 96, 1, 0, 0, 0, 10, 112, 1, 0, 0, 0, 12, 125, 1, 0, 0, 0, 14,
		132, 1, 0, 0, 0, 16, 147, 1, 0, 0, 0, 18, 159, 1, 0, 0, 0, 20, 168, 1,
		0, 0, 0, 22, 170, 1, 0, 0, 0, 24, 188, 1, 0, 0, 0, 26, 192, 1, 0, 0, 0,
		28, 197, 1, 0, 0, 0, 30, 208, 1, 0, 0, 0, 32, 214, 1, 0, 0, 0, 34, 226,
		1, 0, 0, 0, 36, 229, 1, 0, 0, 0, 38, 232, 1, 0, 0, 0, 40, 237, 1, 0, 0,
		0, 42, 239, 1, 0, 0, 0, 44, 252, 1, 0, 0, 0, 46, 261, 1, 0, 0, 0, 48, 281,
		1, 0, 0, 0, 50, 284, 1, 0, 0, 0, 52, 289, 1, 0, 0, 0, 54, 293, 1, 0, 0,
		0, 56, 303, 1, 0, 0, 0, 58, 348, 1, 0, 0, 0, 60, 387, 1, 0, 0, 0, 62, 391,
		1, 0, 0, 0, 64, 393, 1, 0, 0, 0, 66, 395, 1, 0, 0, 0, 68, 397, 1, 0, 0,
		0, 70, 72, 3, 18, 9, 0, 71, 70, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 73,
		1, 0, 0, 0, 73, 74, 5, 0, 0, 1, 74, 1, 1, 0, 0, 0, 75, 81, 3, 6, 3, 0,
		76, 81, 3, 10, 5, 0, 77, 81, 3, 12, 6, 0, 78, 81, 3, 14, 7, 0, 79, 81,
		3, 16, 8, 0, 80, 75, 1, 0, 0, 0, 80, 76, 1, 0, 0, 0, 80, 77, 1, 0, 0, 0,
		80, 78, 1, 0, 0, 0, 80, 79, 1, 0, 0, 0, 81, 82, 1, 0, 0, 0, 82, 83, 5,
		54, 0, 0, 83, 3, 1, 0, 0, 0, 84, 87, 3, 8, 4, 0, 85, 87, 3, 22, 11, 0,
		86, 84, 1, 0, 0, 0, 86, 85, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 89, 5,
		54, 0, 0, 89, 5, 1, 0, 0, 0, 90, 91, 5, 9, 0, 0, 91, 92, 3, 64, 32, 0,
		92, 93, 3, 66, 33, 0, 93, 94, 3, 42, 21, 0, 94, 95, 3, 46, 23, 0, 95, 7,
		1, 0, 0, 0, 96, 97, 5, 4, 0, 0, 97, 100, 3, 64, 32, 0, 98, 101, 3, 34,
		17, 0, 99, 101, 3, 30, 15, 0, 100, 98, 1, 0, 0, 0, 100, 99, 1, 0, 0, 0,
		101, 109, 1, 0, 0, 0, 102, 105, 5, 20, 0, 0, 103, 106, 3, 34, 17, 0, 104,
		106, 3, 30, 15, 0, 105, 103, 1, 0, 0, 0, 105, 104, 1, 0, 0, 0, 106, 108,
		1, 0, 0, 0, 107, 102, 1, 0, 0, 0, 108, 111, 1, 0, 0, 0, 109, 107, 1, 0,
		0, 0, 109, 110, 1, 0, 0, 0, 110, 9, 1, 0, 0, 0, 111, 109, 1, 0, 0, 0, 112,
		113, 5, 11, 0, 0, 113, 114, 3, 66, 33, 0, 114, 120, 5, 25, 0, 0, 115, 116,
		3, 22, 11, 0, 116, 117, 5, 54, 0, 0, 117, 119, 1, 0, 0, 0, 118, 115, 1,
		0, 0, 0, 119, 122, 1, 0, 0, 0, 120, 118, 1, 0, 0, 0, 120, 121, 1, 0, 0,
		0, 121, 123, 1, 0, 0, 0, 122, 120, 1, 0, 0, 0, 123, 124, 5, 26, 0, 0, 124,
		11, 1, 0, 0, 0, 125, 126, 5, 15, 0, 0, 126, 127, 3, 66, 33, 0, 127, 128,
		5, 21, 0, 0, 128, 129, 3, 68, 34, 0, 129, 130, 5, 22, 0, 0, 130, 131, 3,
		46, 23, 0, 131, 13, 1, 0, 0, 0, 132, 133, 5, 16, 0, 0, 133, 138, 3, 66,
		33, 0, 134, 135, 5, 20, 0, 0, 135, 137, 3, 66, 33, 0, 136, 134, 1, 0, 0,
		0, 137, 140, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0, 139,
		141, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 141, 142, 5, 21, 0, 0, 142, 143,
		3, 68, 34, 0, 143, 145, 5, 22, 0, 0, 144, 146, 3, 46, 23, 0, 145, 144,
		1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 15, 1, 0, 0, 0, 147, 148, 5, 17,
		0, 0, 148, 149, 3, 66, 33, 0, 149, 153, 5, 25, 0, 0, 150, 152, 3, 20, 10,
		0, 151, 150, 1, 0, 0, 0, 152, 155, 1, 0, 0, 0, 153, 151, 1, 0, 0, 0, 153,
		154, 1, 0, 0, 0, 154, 156, 1, 0, 0, 0, 155, 153, 1, 0, 0, 0, 156, 157,
		5, 26, 0, 0, 157, 17, 1, 0, 0, 0, 158, 160, 3, 28, 14, 0, 159, 158, 1,
		0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 162, 1, 0, 0, 0, 161, 163, 3, 20, 10,
		0, 162, 161, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164, 162, 1, 0, 0, 0, 164,
		165, 1, 0, 0, 0, 165, 19, 1, 0, 0, 0, 166, 169, 3, 2, 1, 0, 167, 169, 3,
		4, 2, 0, 168, 166, 1, 0, 0, 0, 168, 167, 1, 0, 0, 0, 169, 21, 1, 0, 0,
		0, 170, 171, 5, 5, 0, 0, 171, 174, 3, 64, 32, 0, 172, 175, 3, 40, 20, 0,
		173, 175, 3, 38, 19, 0, 174, 172, 1, 0, 0, 0, 174, 173, 1, 0, 0, 0, 175,
		184, 1, 0, 0, 0, 176, 180, 5, 20, 0, 0, 177, 181, 3, 22, 11, 0, 178, 181,
		3, 40, 20, 0, 179, 181, 3, 38, 19, 0, 180, 177, 1, 0, 0, 0, 180, 178, 1,
		0, 0, 0, 180, 179, 1, 0, 0, 0, 181, 183, 1, 0, 0, 0, 182, 176, 1, 0, 0,
		0, 183, 186, 1, 0, 0, 0, 184, 182, 1, 0, 0, 0, 184, 185, 1, 0, 0, 0, 185,
		23, 1, 0, 0, 0, 186, 184, 1, 0, 0, 0, 187, 189, 9, 0, 0, 0, 188, 187, 1,
		0, 0, 0, 189, 190, 1, 0, 0, 0, 190, 191, 1, 0, 0, 0, 190, 188, 1, 0, 0,
		0, 191, 25, 1, 0, 0, 0, 192, 193, 3, 66, 33, 0, 193, 194, 5, 37, 0, 0,
		194, 195, 3, 24, 12, 0, 195, 196, 5, 54, 0, 0, 196, 27, 1, 0, 0, 0, 197,
		198, 5, 19, 0, 0, 198, 202, 5, 25, 0, 0, 199, 201, 3, 26, 13, 0, 200, 199,
		1, 0, 0, 0, 201, 204, 1, 0, 0, 0, 202, 203, 1, 0, 0, 0, 202, 200, 1, 0,
		0, 0, 203, 205, 1, 0, 0, 0, 204, 202, 1, 0, 0, 0, 205, 206, 5, 26, 0, 0,
		206, 207, 5, 54, 0, 0, 207, 29, 1, 0, 0, 0, 208, 209, 3, 66, 33, 0, 209,
		210, 5, 23, 0, 0, 210, 211, 3, 62, 31, 0, 211, 212, 5, 24, 0, 0, 212, 213,
		3, 32, 16, 0, 213, 31, 1, 0, 0, 0, 214, 215, 5, 37, 0, 0, 215, 216, 5,
		25, 0, 0, 216, 221, 3, 58, 29, 0, 217, 218, 5, 20, 0, 0, 218, 220, 3, 58,
		29, 0, 219, 217, 1, 0, 0, 0, 220, 223, 1, 0, 0, 0, 221, 219, 1, 0, 0, 0,
		221, 222, 1, 0, 0, 0, 222, 224, 1, 0, 0, 0, 223, 221, 1, 0, 0, 0, 224,
		225, 5, 26, 0, 0, 225, 33, 1, 0, 0, 0, 226, 227, 3, 66, 33, 0, 227, 228,
		3, 36, 18, 0, 228, 35, 1, 0, 0, 0, 229, 230, 5, 37, 0, 0, 230, 231, 3,
		58, 29, 0, 231, 37, 1, 0, 0, 0, 232, 233, 3, 66, 33, 0, 233, 234, 5, 23,
		0, 0, 234, 235, 3, 62, 31, 0, 235, 236, 5, 24, 0, 0, 236, 39, 1, 0, 0,
		0, 237, 238, 3, 66, 33, 0, 238, 41, 1, 0, 0, 0, 239, 248, 5, 21, 0, 0,
		240, 245, 3, 44, 22, 0, 241, 242, 5, 20, 0, 0, 242, 244, 3, 44, 22, 0,
		243, 241, 1, 0, 0, 0, 244, 247, 1, 0, 0, 0, 245, 243, 1, 0, 0, 0, 245,
		246, 1, 0, 0, 0, 246, 249, 1, 0, 0, 0, 247, 245, 1, 0, 0, 0, 248, 240,
		1, 0, 0, 0, 248, 249, 1, 0, 0, 0, 249, 250, 1, 0, 0, 0, 250, 251, 5, 22,
		0, 0, 251, 43, 1, 0, 0, 0, 252, 253, 5, 5, 0, 0, 253, 254, 3, 64, 32, 0,
		254, 259, 3, 66, 33, 0, 255, 256, 5, 23, 0, 0, 256, 257, 3, 62, 31, 0,
		257, 258, 5, 24, 0, 0, 258, 260, 1, 0, 0, 0, 259, 255, 1, 0, 0, 0, 259,
		260, 1, 0, 0, 0, 260, 45, 1, 0, 0, 0, 261, 276, 5, 25, 0, 0, 262, 267,
		3, 56, 28, 0, 263, 267, 3, 8, 4, 0, 264, 267, 3, 22, 11, 0, 265, 267, 3,
		58, 29, 0, 266, 262, 1, 0, 0, 0, 266, 263, 1, 0, 0, 0, 266, 264, 1, 0,
		0, 0, 266, 265, 1, 0, 0, 0, 267, 268, 1, 0, 0, 0, 268, 269, 5, 54, 0, 0,
		269, 275, 1, 0, 0, 0, 270, 272, 3, 54, 27, 0, 271, 273, 5, 54, 0, 0, 272,
		271, 1, 0, 0, 0, 272, 273, 1, 0, 0, 0, 273, 275, 1, 0, 0, 0, 274, 266,
		1, 0, 0, 0, 274, 270, 1, 0, 0, 0, 275, 278, 1, 0, 0, 0, 276, 274, 1, 0,
		0, 0, 276, 277, 1, 0, 0, 0, 277, 279, 1, 0, 0, 0, 278, 276, 1, 0, 0, 0,
		279, 280, 5, 26, 0, 0, 280, 47, 1, 0, 0, 0, 281, 282, 5, 8, 0, 0, 282,
		283, 3, 46, 23, 0, 283, 49, 1, 0, 0, 0, 284, 285, 5, 8, 0, 0, 285, 286,
		5, 6, 0, 0, 286, 287, 3, 58, 29, 0, 287, 288, 3, 46, 23, 0, 288, 51, 1,
		0, 0, 0, 289, 290, 5, 6, 0, 0, 290, 291, 3, 58, 29, 0, 291, 292, 3, 46,
		23, 0, 292, 53, 1, 0, 0, 0, 293, 297, 3, 52, 26, 0, 294, 296, 3, 50, 25,
		0, 295, 294, 1, 0, 0, 0, 296, 299, 1, 0, 0, 0, 297, 295, 1, 0, 0, 0, 297,
		298, 1, 0, 0, 0, 298, 301, 1, 0, 0, 0, 299, 297, 1, 0, 0, 0, 300, 302,
		3, 48, 24, 0, 301, 300, 1, 0, 0, 0, 301, 302, 1, 0, 0, 0, 302, 55, 1, 0,
		0, 0, 303, 305, 5, 13, 0, 0, 304, 306, 3, 58, 29, 0, 305, 304, 1, 0, 0,
		0, 305, 306, 1, 0, 0, 0, 306, 57, 1, 0, 0, 0, 307, 308, 6, 29, -1, 0, 308,
		309, 5, 21, 0, 0, 309, 310, 3, 58, 29, 0, 310, 311, 5, 22, 0, 0, 311, 349,
		1, 0, 0, 0, 312, 313, 7, 0, 0, 0, 313, 349, 3, 58, 29, 17, 314, 349, 5,
		1, 0, 0, 315, 349, 5, 2, 0, 0, 316, 349, 5, 3, 0, 0, 317, 349, 5, 18, 0,
		0, 318, 319, 7, 1, 0, 0, 319, 328, 5, 21, 0, 0, 320, 325, 3, 58, 29, 0,
		321, 322, 5, 20, 0, 0, 322, 324, 3, 58, 29, 0, 323, 321, 1, 0, 0, 0, 324,
		327, 1, 0, 0, 0, 325, 323, 1, 0, 0, 0, 325, 326, 1, 0, 0, 0, 326, 329,
		1, 0, 0, 0, 327, 325, 1, 0, 0, 0, 328, 320, 1, 0, 0, 0, 328, 329, 1, 0,
		0, 0, 329, 330, 1, 0, 0, 0, 330, 349, 5, 22, 0, 0, 331, 336, 7, 1, 0, 0,
		332, 333, 5, 23, 0, 0, 333, 334, 3, 60, 30, 0, 334, 335, 5, 24, 0, 0, 335,
		337, 1, 0, 0, 0, 336, 332, 1, 0, 0, 0, 336, 337, 1, 0, 0, 0, 337, 346,
		1, 0, 0, 0, 338, 339, 5, 53, 0, 0, 339, 344, 7, 1, 0, 0, 340, 341, 5, 23,
		0, 0, 341, 342, 3, 60, 30, 0, 342, 343, 5, 24, 0, 0, 343, 345, 1, 0, 0,
		0, 344, 340, 1, 0, 0, 0, 344, 345, 1, 0, 0, 0, 345, 347, 1, 0, 0, 0, 346,
		338, 1, 0, 0, 0, 346, 347, 1, 0, 0, 0, 347, 349, 1, 0, 0, 0, 348, 307,
		1, 0, 0, 0, 348, 312, 1, 0, 0, 0, 348, 314, 1, 0, 0, 0, 348, 315, 1, 0,
		0, 0, 348, 316, 1, 0, 0, 0, 348, 317, 1, 0, 0, 0, 348, 318, 1, 0, 0, 0,
		348, 331, 1, 0, 0, 0, 349, 382, 1, 0, 0, 0, 350, 351, 10, 16, 0, 0, 351,
		352, 7, 2, 0, 0, 352, 381, 3, 58, 29, 17, 353, 354, 10, 15, 0, 0, 354,
		355, 7, 3, 0, 0, 355, 381, 3, 58, 29, 16, 356, 357, 10, 14, 0, 0, 357,
		358, 7, 4, 0, 0, 358, 381, 3, 58, 29, 15, 359, 360, 10, 13, 0, 0, 360,
		361, 7, 5, 0, 0, 361, 381, 3, 58, 29, 14, 362, 363, 10, 12, 0, 0, 363,
		364, 7, 6, 0, 0, 364, 381, 3, 58, 29, 13, 365, 366, 10, 11, 0, 0, 366,
		367, 5, 27, 0, 0, 367, 381, 3, 58, 29, 12, 368, 369, 10, 10, 0, 0, 369,
		370, 5, 29, 0, 0, 370, 381, 3, 58, 29, 11, 371, 372, 10, 9, 0, 0, 372,
		373, 5, 28, 0, 0, 373, 381, 3, 58, 29, 10, 374, 375, 10, 8, 0, 0, 375,
		376, 5, 30, 0, 0, 376, 381, 3, 58, 29, 9, 377, 378, 10, 7, 0, 0, 378, 379,
		7, 7, 0, 0, 379, 381, 3, 58, 29, 7, 380, 350, 1, 0, 0, 0, 380, 353, 1,
		0, 0, 0, 380, 356, 1, 0, 0, 0, 380, 359, 1, 0, 0, 0, 380, 362, 1, 0, 0,
		0, 380, 365, 1, 0, 0, 0, 380, 368, 1, 0, 0, 0, 380, 371, 1, 0, 0, 0, 380,
		374, 1, 0, 0, 0, 380, 377, 1, 0, 0, 0, 381, 384, 1, 0, 0, 0, 382, 380,
		1, 0, 0, 0, 382, 383, 1, 0, 0, 0, 383, 59, 1, 0, 0, 0, 384, 382, 1, 0,
		0, 0, 385, 388, 5, 1, 0, 0, 386, 388, 3, 66, 33, 0, 387, 385, 1, 0, 0,
		0, 387, 386, 1, 0, 0, 0, 388, 61, 1, 0, 0, 0, 389, 392, 5, 1, 0, 0, 390,
		392, 3, 66, 33, 0, 391, 389, 1, 0, 0, 0, 391, 390, 1, 0, 0, 0, 392, 63,
		1, 0, 0, 0, 393, 394, 7, 8, 0, 0, 394, 65, 1, 0, 0, 0, 395, 396, 7, 9,
		0, 0, 396, 67, 1, 0, 0, 0, 397, 398, 5, 55, 0, 0, 398, 69, 1, 0, 0, 0,
		39, 71, 80, 86, 100, 105, 109, 120, 138, 145, 153, 159, 164, 168, 174,
		180, 184, 190, 202, 221, 245, 248, 259, 266, 272, 274, 276, 297, 301, 305,
		325, 328, 336, 344, 346, 348, 380, 382, 387, 391,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// DaedalusParserInit initializes any static state used to implement DaedalusParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewDaedalusParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func DaedalusParserInit() {
	staticData := &DaedalusParserStaticData
	staticData.once.Do(daedalusParserInit)
}

// NewDaedalusParser produces a new parser instance for the optional input antlr.TokenStream.
func NewDaedalusParser(input antlr.TokenStream) *DaedalusParser {
	DaedalusParserInit()
	this := new(DaedalusParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &DaedalusParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Daedalus.g4"

	return this
}

// DaedalusParser tokens.
const (
	DaedalusParserEOF            = antlr.TokenEOF
	DaedalusParserIntegerLiteral = 1
	DaedalusParserFloatLiteral   = 2
	DaedalusParserStringLiteral  = 3
	DaedalusParserConst          = 4
	DaedalusParserVar            = 5
	DaedalusParserIf             = 6
	DaedalusParserInt            = 7
	DaedalusParserElse           = 8
	DaedalusParserFunc           = 9
	DaedalusParserStringKeyword  = 10
	DaedalusParserClass          = 11
	DaedalusParserVoid           = 12
	DaedalusParserReturn         = 13
	DaedalusParserFloat          = 14
	DaedalusParserPrototype      = 15
	DaedalusParserInstance       = 16
	DaedalusParserNamespace      = 17
	DaedalusParserNull           = 18
	DaedalusParserMeta           = 19
	DaedalusParserComma          = 20
	DaedalusParserLeftParen      = 21
	DaedalusParserRightParen     = 22
	DaedalusParserLeftBracket    = 23
	DaedalusParserRightBracket   = 24
	DaedalusParserLeftBrace      = 25
	DaedalusParserRightBrace     = 26
	DaedalusParserBitAnd         = 27
	DaedalusParserAnd            = 28
	DaedalusParserBitOr          = 29
	DaedalusParserOr             = 30
	DaedalusParserPlus           = 31
	DaedalusParserMinus          = 32
	DaedalusParserDiv            = 33
	DaedalusParserStar           = 34
	DaedalusParserTilde          = 35
	DaedalusParserNot            = 36
	DaedalusParserAssign         = 37
	DaedalusParserLess           = 38
	DaedalusParserGreater        = 39
	DaedalusParserShl            = 40
	DaedalusParserShr            = 41
	DaedalusParserLessEq         = 42
	DaedalusParserGreaterEq      = 43
	DaedalusParserEqual          = 44
	DaedalusParserNotEqual       = 45
	DaedalusParserPercent        = 46
	DaedalusParserPlusAssign     = 47
	DaedalusParserMinusAssign    = 48
	DaedalusParserStarAssign     = 49
	DaedalusParserDivAssign      = 50
	DaedalusParserAndAssign      = 51
	DaedalusParserOrAssign       = 52
	DaedalusParserDot            = 53
	DaedalusParserSemi           = 54
	DaedalusParserIdentifier     = 55
	DaedalusParserWhitespace     = 56
	DaedalusParserNewline        = 57
	DaedalusParserBlockComment   = 58
	DaedalusParserLineComment    = 59
)

// DaedalusParser rules.
const (
	DaedalusParserRULE_daedalusFile             = 0
	DaedalusParserRULE_blockDef                 = 1
	DaedalusParserRULE_inlineDef                = 2
	DaedalusParserRULE_functionDef              = 3
	DaedalusParserRULE_constDef                 = 4
	DaedalusParserRULE_classDef                 = 5
	DaedalusParserRULE_prototypeDef             = 6
	DaedalusParserRULE_instanceDef              = 7
	DaedalusParserRULE_namespaceDef             = 8
	DaedalusParserRULE_mainBlock                = 9
	DaedalusParserRULE_contentBlock             = 10
	DaedalusParserRULE_varDecl                  = 11
	DaedalusParserRULE_metaValue                = 12
	DaedalusParserRULE_zParserExtenderMeta      = 13
	DaedalusParserRULE_zParserExtenderMetaBlock = 14
	DaedalusParserRULE_constArrayDef            = 15
	DaedalusParserRULE_constArrayAssignment     = 16
	DaedalusParserRULE_constValueDef            = 17
	DaedalusParserRULE_constValueAssignment     = 18
	DaedalusParserRULE_varArrayDecl             = 19
	DaedalusParserRULE_varValueDecl             = 20
	DaedalusParserRULE_parameterList            = 21
	DaedalusParserRULE_parameterDecl            = 22
	DaedalusParserRULE_statementBlock           = 23
	DaedalusParserRULE_elseBlock                = 24
	DaedalusParserRULE_elseIfBlock              = 25
	DaedalusParserRULE_ifBlock                  = 26
	DaedalusParserRULE_ifBlockStatement         = 27
	DaedalusParserRULE_returnStatement          = 28
	DaedalusParserRULE_expression               = 29
	DaedalusParserRULE_arrayIndex               = 30
	DaedalusParserRULE_arraySize                = 31
	DaedalusParserRULE_typeReference            = 32
	DaedalusParserRULE_nameNode                 = 33
	DaedalusParserRULE_parentReference          = 34
)

// IDaedalusFileContext is an interface to support dynamic dispatch.
type IDaedalusFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	MainBlock() IMainBlockContext

	// IsDaedalusFileContext differentiates from other interfaces.
	IsDaedalusFileContext()
}

type DaedalusFileContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDaedalusFileContext() *DaedalusFileContext {
	var p = new(DaedalusFileContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_daedalusFile
	return p
}

func InitEmptyDaedalusFileContext(p *DaedalusFileContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_daedalusFile
}

func (*DaedalusFileContext) IsDaedalusFileContext() {}

func NewDaedalusFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DaedalusFileContext {
	var p = new(DaedalusFileContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_daedalusFile

	return p
}

func (s *DaedalusFileContext) GetParser() antlr.Parser { return s.parser }

func (s *DaedalusFileContext) EOF() antlr.TerminalNode {
	return s.GetToken(DaedalusParserEOF, 0)
}

func (s *DaedalusFileContext) MainBlock() IMainBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMainBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMainBlockContext)
}

func (s *DaedalusFileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DaedalusFileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DaedalusFileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterDaedalusFile(s)
	}
}

func (s *DaedalusFileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitDaedalusFile(s)
	}
}

func (p *DaedalusParser) DaedalusFile() (localctx IDaedalusFileContext) {
	localctx = NewDaedalusFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DaedalusParserRULE_daedalusFile)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&756272) != 0 {
		{
			p.SetState(70)
			p.MainBlock()
		}

	}
	{
		p.SetState(73)
		p.Match(DaedalusParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockDefContext is an interface to support dynamic dispatch.
type IBlockDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Semi() antlr.TerminalNode
	FunctionDef() IFunctionDefContext
	ClassDef() IClassDefContext
	PrototypeDef() IPrototypeDefContext
	InstanceDef() IInstanceDefContext
	NamespaceDef() INamespaceDefContext

	// IsBlockDefContext differentiates from other interfaces.
	IsBlockDefContext()
}

type BlockDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockDefContext() *BlockDefContext {
	var p = new(BlockDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_blockDef
	return p
}

func InitEmptyBlockDefContext(p *BlockDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_blockDef
}

func (*BlockDefContext) IsBlockDefContext() {}

func NewBlockDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockDefContext {
	var p = new(BlockDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_blockDef

	return p
}

func (s *BlockDefContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockDefContext) Semi() antlr.TerminalNode {
	return s.GetToken(DaedalusParserSemi, 0)
}

func (s *BlockDefContext) FunctionDef() IFunctionDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionDefContext)
}

func (s *BlockDefContext) ClassDef() IClassDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClassDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClassDefContext)
}

func (s *BlockDefContext) PrototypeDef() IPrototypeDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrototypeDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrototypeDefContext)
}

func (s *BlockDefContext) InstanceDef() IInstanceDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstanceDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstanceDefContext)
}

func (s *BlockDefContext) NamespaceDef() INamespaceDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamespaceDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamespaceDefContext)
}

func (s *BlockDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterBlockDef(s)
	}
}

func (s *BlockDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitBlockDef(s)
	}
}

func (p *DaedalusParser) BlockDef() (localctx IBlockDefContext) {
	localctx = NewBlockDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DaedalusParserRULE_blockDef)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case DaedalusParserFunc:
		{
			p.SetState(75)
			p.FunctionDef()
		}

	case DaedalusParserClass:
		{
			p.SetState(76)
			p.ClassDef()
		}

	case DaedalusParserPrototype:
		{
			p.SetState(77)
			p.PrototypeDef()
		}

	case DaedalusParserInstance:
		{
			p.SetState(78)
			p.InstanceDef()
		}

	case DaedalusParserNamespace:
		{
			p.SetState(79)
			p.NamespaceDef()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(82)
		p.Match(DaedalusParserSemi)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInlineDefContext is an interface to support dynamic dispatch.
type IInlineDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Semi() antlr.TerminalNode
	ConstDef() IConstDefContext
	VarDecl() IVarDeclContext

	// IsInlineDefContext differentiates from other interfaces.
	IsInlineDefContext()
}

type InlineDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInlineDefContext() *InlineDefContext {
	var p = new(InlineDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_inlineDef
	return p
}

func InitEmptyInlineDefContext(p *InlineDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_inlineDef
}

func (*InlineDefContext) IsInlineDefContext() {}

func NewInlineDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InlineDefContext {
	var p = new(InlineDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_inlineDef

	return p
}

func (s *InlineDefContext) GetParser() antlr.Parser { return s.parser }

func (s *InlineDefContext) Semi() antlr.TerminalNode {
	return s.GetToken(DaedalusParserSemi, 0)
}

func (s *InlineDefContext) ConstDef() IConstDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstDefContext)
}

func (s *InlineDefContext) VarDecl() IVarDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDeclContext)
}

func (s *InlineDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InlineDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InlineDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterInlineDef(s)
	}
}

func (s *InlineDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitInlineDef(s)
	}
}

func (p *DaedalusParser) InlineDef() (localctx IInlineDefContext) {
	localctx = NewInlineDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, DaedalusParserRULE_inlineDef)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case DaedalusParserConst:
		{
			p.SetState(84)
			p.ConstDef()
		}

	case DaedalusParserVar:
		{
			p.SetState(85)
			p.VarDecl()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(88)
		p.Match(DaedalusParserSemi)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionDefContext is an interface to support dynamic dispatch.
type IFunctionDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Func() antlr.TerminalNode
	TypeReference() ITypeReferenceContext
	NameNode() INameNodeContext
	ParameterList() IParameterListContext
	StatementBlock() IStatementBlockContext

	// IsFunctionDefContext differentiates from other interfaces.
	IsFunctionDefContext()
}

type FunctionDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDefContext() *FunctionDefContext {
	var p = new(FunctionDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_functionDef
	return p
}

func InitEmptyFunctionDefContext(p *FunctionDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_functionDef
}

func (*FunctionDefContext) IsFunctionDefContext() {}

func NewFunctionDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDefContext {
	var p = new(FunctionDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_functionDef

	return p
}

func (s *FunctionDefContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDefContext) Func() antlr.TerminalNode {
	return s.GetToken(DaedalusParserFunc, 0)
}

func (s *FunctionDefContext) TypeReference() ITypeReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeReferenceContext)
}

func (s *FunctionDefContext) NameNode() INameNodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameNodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameNodeContext)
}

func (s *FunctionDefContext) ParameterList() IParameterListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterListContext)
}

func (s *FunctionDefContext) StatementBlock() IStatementBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementBlockContext)
}

func (s *FunctionDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterFunctionDef(s)
	}
}

func (s *FunctionDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitFunctionDef(s)
	}
}

func (p *DaedalusParser) FunctionDef() (localctx IFunctionDefContext) {
	localctx = NewFunctionDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, DaedalusParserRULE_functionDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(90)
		p.Match(DaedalusParserFunc)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(91)
		p.TypeReference()
	}
	{
		p.SetState(92)
		p.NameNode()
	}
	{
		p.SetState(93)
		p.ParameterList()
	}
	{
		p.SetState(94)
		p.StatementBlock()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstDefContext is an interface to support dynamic dispatch.
type IConstDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Const() antlr.TerminalNode
	TypeReference() ITypeReferenceContext
	AllConstValueDef() []IConstValueDefContext
	ConstValueDef(i int) IConstValueDefContext
	AllConstArrayDef() []IConstArrayDefContext
	ConstArrayDef(i int) IConstArrayDefContext
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsConstDefContext differentiates from other interfaces.
	IsConstDefContext()
}

type ConstDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstDefContext() *ConstDefContext {
	var p = new(ConstDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_constDef
	return p
}

func InitEmptyConstDefContext(p *ConstDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_constDef
}

func (*ConstDefContext) IsConstDefContext() {}

func NewConstDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstDefContext {
	var p = new(ConstDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_constDef

	return p
}

func (s *ConstDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstDefContext) Const() antlr.TerminalNode {
	return s.GetToken(DaedalusParserConst, 0)
}

func (s *ConstDefContext) TypeReference() ITypeReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeReferenceContext)
}

func (s *ConstDefContext) AllConstValueDef() []IConstValueDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConstValueDefContext); ok {
			len++
		}
	}

	tst := make([]IConstValueDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConstValueDefContext); ok {
			tst[i] = t.(IConstValueDefContext)
			i++
		}
	}

	return tst
}

func (s *ConstDefContext) ConstValueDef(i int) IConstValueDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstValueDefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstValueDefContext)
}

func (s *ConstDefContext) AllConstArrayDef() []IConstArrayDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConstArrayDefContext); ok {
			len++
		}
	}

	tst := make([]IConstArrayDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConstArrayDefContext); ok {
			tst[i] = t.(IConstArrayDefContext)
			i++
		}
	}

	return tst
}

func (s *ConstDefContext) ConstArrayDef(i int) IConstArrayDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstArrayDefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstArrayDefContext)
}

func (s *ConstDefContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(DaedalusParserComma)
}

func (s *ConstDefContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(DaedalusParserComma, i)
}

func (s *ConstDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterConstDef(s)
	}
}

func (s *ConstDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitConstDef(s)
	}
}

func (p *DaedalusParser) ConstDef() (localctx IConstDefContext) {
	localctx = NewConstDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, DaedalusParserRULE_constDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(96)
		p.Match(DaedalusParserConst)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(97)
		p.TypeReference()
	}
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(98)
			p.ConstValueDef()
		}

	case 2:
		{
			p.SetState(99)
			p.ConstArrayDef()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == DaedalusParserComma {
		{
			p.SetState(102)
			p.Match(DaedalusParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(103)
				p.ConstValueDef()
			}

		case 2:
			{
				p.SetState(104)
				p.ConstArrayDef()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IClassDefContext is an interface to support dynamic dispatch.
type IClassDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Class() antlr.TerminalNode
	NameNode() INameNodeContext
	LeftBrace() antlr.TerminalNode
	RightBrace() antlr.TerminalNode
	AllVarDecl() []IVarDeclContext
	VarDecl(i int) IVarDeclContext
	AllSemi() []antlr.TerminalNode
	Semi(i int) antlr.TerminalNode

	// IsClassDefContext differentiates from other interfaces.
	IsClassDefContext()
}

type ClassDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassDefContext() *ClassDefContext {
	var p = new(ClassDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_classDef
	return p
}

func InitEmptyClassDefContext(p *ClassDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_classDef
}

func (*ClassDefContext) IsClassDefContext() {}

func NewClassDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassDefContext {
	var p = new(ClassDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_classDef

	return p
}

func (s *ClassDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassDefContext) Class() antlr.TerminalNode {
	return s.GetToken(DaedalusParserClass, 0)
}

func (s *ClassDefContext) NameNode() INameNodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameNodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameNodeContext)
}

func (s *ClassDefContext) LeftBrace() antlr.TerminalNode {
	return s.GetToken(DaedalusParserLeftBrace, 0)
}

func (s *ClassDefContext) RightBrace() antlr.TerminalNode {
	return s.GetToken(DaedalusParserRightBrace, 0)
}

func (s *ClassDefContext) AllVarDecl() []IVarDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarDeclContext); ok {
			len++
		}
	}

	tst := make([]IVarDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarDeclContext); ok {
			tst[i] = t.(IVarDeclContext)
			i++
		}
	}

	return tst
}

func (s *ClassDefContext) VarDecl(i int) IVarDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDeclContext)
}

func (s *ClassDefContext) AllSemi() []antlr.TerminalNode {
	return s.GetTokens(DaedalusParserSemi)
}

func (s *ClassDefContext) Semi(i int) antlr.TerminalNode {
	return s.GetToken(DaedalusParserSemi, i)
}

func (s *ClassDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterClassDef(s)
	}
}

func (s *ClassDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitClassDef(s)
	}
}

func (p *DaedalusParser) ClassDef() (localctx IClassDefContext) {
	localctx = NewClassDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, DaedalusParserRULE_classDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)
		p.Match(DaedalusParserClass)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(113)
		p.NameNode()
	}
	{
		p.SetState(114)
		p.Match(DaedalusParserLeftBrace)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == DaedalusParserVar {
		{
			p.SetState(115)
			p.VarDecl()
		}
		{
			p.SetState(116)
			p.Match(DaedalusParserSemi)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(123)
		p.Match(DaedalusParserRightBrace)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrototypeDefContext is an interface to support dynamic dispatch.
type IPrototypeDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Prototype() antlr.TerminalNode
	NameNode() INameNodeContext
	LeftParen() antlr.TerminalNode
	ParentReference() IParentReferenceContext
	RightParen() antlr.TerminalNode
	StatementBlock() IStatementBlockContext

	// IsPrototypeDefContext differentiates from other interfaces.
	IsPrototypeDefContext()
}

type PrototypeDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrototypeDefContext() *PrototypeDefContext {
	var p = new(PrototypeDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_prototypeDef
	return p
}

func InitEmptyPrototypeDefContext(p *PrototypeDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_prototypeDef
}

func (*PrototypeDefContext) IsPrototypeDefContext() {}

func NewPrototypeDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrototypeDefContext {
	var p = new(PrototypeDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_prototypeDef

	return p
}

func (s *PrototypeDefContext) GetParser() antlr.Parser { return s.parser }

func (s *PrototypeDefContext) Prototype() antlr.TerminalNode {
	return s.GetToken(DaedalusParserPrototype, 0)
}

func (s *PrototypeDefContext) NameNode() INameNodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameNodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameNodeContext)
}

func (s *PrototypeDefContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(DaedalusParserLeftParen, 0)
}

func (s *PrototypeDefContext) ParentReference() IParentReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParentReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParentReferenceContext)
}

func (s *PrototypeDefContext) RightParen() antlr.TerminalNode {
	return s.GetToken(DaedalusParserRightParen, 0)
}

func (s *PrototypeDefContext) StatementBlock() IStatementBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementBlockContext)
}

func (s *PrototypeDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrototypeDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrototypeDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterPrototypeDef(s)
	}
}

func (s *PrototypeDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitPrototypeDef(s)
	}
}

func (p *DaedalusParser) PrototypeDef() (localctx IPrototypeDefContext) {
	localctx = NewPrototypeDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, DaedalusParserRULE_prototypeDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(125)
		p.Match(DaedalusParserPrototype)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(126)
		p.NameNode()
	}
	{
		p.SetState(127)
		p.Match(DaedalusParserLeftParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(128)
		p.ParentReference()
	}
	{
		p.SetState(129)
		p.Match(DaedalusParserRightParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(130)
		p.StatementBlock()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IInstanceDefContext is an interface to support dynamic dispatch.
type IInstanceDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Instance() antlr.TerminalNode
	AllNameNode() []INameNodeContext
	NameNode(i int) INameNodeContext
	LeftParen() antlr.TerminalNode
	ParentReference() IParentReferenceContext
	RightParen() antlr.TerminalNode
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode
	StatementBlock() IStatementBlockContext

	// IsInstanceDefContext differentiates from other interfaces.
	IsInstanceDefContext()
}

type InstanceDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstanceDefContext() *InstanceDefContext {
	var p = new(InstanceDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_instanceDef
	return p
}

func InitEmptyInstanceDefContext(p *InstanceDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_instanceDef
}

func (*InstanceDefContext) IsInstanceDefContext() {}

func NewInstanceDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstanceDefContext {
	var p = new(InstanceDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_instanceDef

	return p
}

func (s *InstanceDefContext) GetParser() antlr.Parser { return s.parser }

func (s *InstanceDefContext) Instance() antlr.TerminalNode {
	return s.GetToken(DaedalusParserInstance, 0)
}

func (s *InstanceDefContext) AllNameNode() []INameNodeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INameNodeContext); ok {
			len++
		}
	}

	tst := make([]INameNodeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INameNodeContext); ok {
			tst[i] = t.(INameNodeContext)
			i++
		}
	}

	return tst
}

func (s *InstanceDefContext) NameNode(i int) INameNodeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameNodeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameNodeContext)
}

func (s *InstanceDefContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(DaedalusParserLeftParen, 0)
}

func (s *InstanceDefContext) ParentReference() IParentReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParentReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParentReferenceContext)
}

func (s *InstanceDefContext) RightParen() antlr.TerminalNode {
	return s.GetToken(DaedalusParserRightParen, 0)
}

func (s *InstanceDefContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(DaedalusParserComma)
}

func (s *InstanceDefContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(DaedalusParserComma, i)
}

func (s *InstanceDefContext) StatementBlock() IStatementBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementBlockContext)
}

func (s *InstanceDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstanceDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstanceDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterInstanceDef(s)
	}
}

func (s *InstanceDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitInstanceDef(s)
	}
}

func (p *DaedalusParser) InstanceDef() (localctx IInstanceDefContext) {
	localctx = NewInstanceDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, DaedalusParserRULE_instanceDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(132)
		p.Match(DaedalusParserInstance)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(133)
		p.NameNode()
	}
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == DaedalusParserComma {
		{
			p.SetState(134)
			p.Match(DaedalusParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(135)
			p.NameNode()
		}

		p.SetState(140)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(141)
		p.Match(DaedalusParserLeftParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(142)
		p.ParentReference()
	}
	{
		p.SetState(143)
		p.Match(DaedalusParserRightParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DaedalusParserLeftBrace {
		{
			p.SetState(144)
			p.StatementBlock()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INamespaceDefContext is an interface to support dynamic dispatch.
type INamespaceDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Namespace() antlr.TerminalNode
	NameNode() INameNodeContext
	LeftBrace() antlr.TerminalNode
	RightBrace() antlr.TerminalNode
	AllContentBlock() []IContentBlockContext
	ContentBlock(i int) IContentBlockContext

	// IsNamespaceDefContext differentiates from other interfaces.
	IsNamespaceDefContext()
}

type NamespaceDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamespaceDefContext() *NamespaceDefContext {
	var p = new(NamespaceDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_namespaceDef
	return p
}

func InitEmptyNamespaceDefContext(p *NamespaceDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_namespaceDef
}

func (*NamespaceDefContext) IsNamespaceDefContext() {}

func NewNamespaceDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamespaceDefContext {
	var p = new(NamespaceDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_namespaceDef

	return p
}

func (s *NamespaceDefContext) GetParser() antlr.Parser { return s.parser }

func (s *NamespaceDefContext) Namespace() antlr.TerminalNode {
	return s.GetToken(DaedalusParserNamespace, 0)
}

func (s *NamespaceDefContext) NameNode() INameNodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameNodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameNodeContext)
}

func (s *NamespaceDefContext) LeftBrace() antlr.TerminalNode {
	return s.GetToken(DaedalusParserLeftBrace, 0)
}

func (s *NamespaceDefContext) RightBrace() antlr.TerminalNode {
	return s.GetToken(DaedalusParserRightBrace, 0)
}

func (s *NamespaceDefContext) AllContentBlock() []IContentBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IContentBlockContext); ok {
			len++
		}
	}

	tst := make([]IContentBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IContentBlockContext); ok {
			tst[i] = t.(IContentBlockContext)
			i++
		}
	}

	return tst
}

func (s *NamespaceDefContext) ContentBlock(i int) IContentBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContentBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContentBlockContext)
}

func (s *NamespaceDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamespaceDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamespaceDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterNamespaceDef(s)
	}
}

func (s *NamespaceDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitNamespaceDef(s)
	}
}

func (p *DaedalusParser) NamespaceDef() (localctx INamespaceDefContext) {
	localctx = NewNamespaceDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, DaedalusParserRULE_namespaceDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(147)
		p.Match(DaedalusParserNamespace)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(148)
		p.NameNode()
	}
	{
		p.SetState(149)
		p.Match(DaedalusParserLeftBrace)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&231984) != 0 {
		{
			p.SetState(150)
			p.ContentBlock()
		}

		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(156)
		p.Match(DaedalusParserRightBrace)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMainBlockContext is an interface to support dynamic dispatch.
type IMainBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ZParserExtenderMetaBlock() IZParserExtenderMetaBlockContext
	AllContentBlock() []IContentBlockContext
	ContentBlock(i int) IContentBlockContext

	// IsMainBlockContext differentiates from other interfaces.
	IsMainBlockContext()
}

type MainBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMainBlockContext() *MainBlockContext {
	var p = new(MainBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_mainBlock
	return p
}

func InitEmptyMainBlockContext(p *MainBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_mainBlock
}

func (*MainBlockContext) IsMainBlockContext() {}

func NewMainBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MainBlockContext {
	var p = new(MainBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_mainBlock

	return p
}

func (s *MainBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *MainBlockContext) ZParserExtenderMetaBlock() IZParserExtenderMetaBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IZParserExtenderMetaBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IZParserExtenderMetaBlockContext)
}

func (s *MainBlockContext) AllContentBlock() []IContentBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IContentBlockContext); ok {
			len++
		}
	}

	tst := make([]IContentBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IContentBlockContext); ok {
			tst[i] = t.(IContentBlockContext)
			i++
		}
	}

	return tst
}

func (s *MainBlockContext) ContentBlock(i int) IContentBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContentBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContentBlockContext)
}

func (s *MainBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MainBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MainBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterMainBlock(s)
	}
}

func (s *MainBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitMainBlock(s)
	}
}

func (p *DaedalusParser) MainBlock() (localctx IMainBlockContext) {
	localctx = NewMainBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, DaedalusParserRULE_mainBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DaedalusParserMeta {
		{
			p.SetState(158)
			p.ZParserExtenderMetaBlock()
		}

	}
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&231984) != 0) {
		{
			p.SetState(161)
			p.ContentBlock()
		}

		p.SetState(164)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IContentBlockContext is an interface to support dynamic dispatch.
type IContentBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BlockDef() IBlockDefContext
	InlineDef() IInlineDefContext

	// IsContentBlockContext differentiates from other interfaces.
	IsContentBlockContext()
}

type ContentBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContentBlockContext() *ContentBlockContext {
	var p = new(ContentBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_contentBlock
	return p
}

func InitEmptyContentBlockContext(p *ContentBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_contentBlock
}

func (*ContentBlockContext) IsContentBlockContext() {}

func NewContentBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContentBlockContext {
	var p = new(ContentBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_contentBlock

	return p
}

func (s *ContentBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ContentBlockContext) BlockDef() IBlockDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockDefContext)
}

func (s *ContentBlockContext) InlineDef() IInlineDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInlineDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInlineDefContext)
}

func (s *ContentBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContentBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContentBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterContentBlock(s)
	}
}

func (s *ContentBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitContentBlock(s)
	}
}

func (p *DaedalusParser) ContentBlock() (localctx IContentBlockContext) {
	localctx = NewContentBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, DaedalusParserRULE_contentBlock)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case DaedalusParserFunc, DaedalusParserClass, DaedalusParserPrototype, DaedalusParserInstance, DaedalusParserNamespace:
		{
			p.SetState(166)
			p.BlockDef()
		}

	case DaedalusParserConst, DaedalusParserVar:
		{
			p.SetState(167)
			p.InlineDef()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarDeclContext is an interface to support dynamic dispatch.
type IVarDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Var() antlr.TerminalNode
	TypeReference() ITypeReferenceContext
	AllVarValueDecl() []IVarValueDeclContext
	VarValueDecl(i int) IVarValueDeclContext
	AllVarArrayDecl() []IVarArrayDeclContext
	VarArrayDecl(i int) IVarArrayDeclContext
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode
	AllVarDecl() []IVarDeclContext
	VarDecl(i int) IVarDeclContext

	// IsVarDeclContext differentiates from other interfaces.
	IsVarDeclContext()
}

type VarDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDeclContext() *VarDeclContext {
	var p = new(VarDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_varDecl
	return p
}

func InitEmptyVarDeclContext(p *VarDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_varDecl
}

func (*VarDeclContext) IsVarDeclContext() {}

func NewVarDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDeclContext {
	var p = new(VarDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_varDecl

	return p
}

func (s *VarDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDeclContext) Var() antlr.TerminalNode {
	return s.GetToken(DaedalusParserVar, 0)
}

func (s *VarDeclContext) TypeReference() ITypeReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeReferenceContext)
}

func (s *VarDeclContext) AllVarValueDecl() []IVarValueDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarValueDeclContext); ok {
			len++
		}
	}

	tst := make([]IVarValueDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarValueDeclContext); ok {
			tst[i] = t.(IVarValueDeclContext)
			i++
		}
	}

	return tst
}

func (s *VarDeclContext) VarValueDecl(i int) IVarValueDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarValueDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarValueDeclContext)
}

func (s *VarDeclContext) AllVarArrayDecl() []IVarArrayDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarArrayDeclContext); ok {
			len++
		}
	}

	tst := make([]IVarArrayDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarArrayDeclContext); ok {
			tst[i] = t.(IVarArrayDeclContext)
			i++
		}
	}

	return tst
}

func (s *VarDeclContext) VarArrayDecl(i int) IVarArrayDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarArrayDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarArrayDeclContext)
}

func (s *VarDeclContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(DaedalusParserComma)
}

func (s *VarDeclContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(DaedalusParserComma, i)
}

func (s *VarDeclContext) AllVarDecl() []IVarDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarDeclContext); ok {
			len++
		}
	}

	tst := make([]IVarDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarDeclContext); ok {
			tst[i] = t.(IVarDeclContext)
			i++
		}
	}

	return tst
}

func (s *VarDeclContext) VarDecl(i int) IVarDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDeclContext)
}

func (s *VarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterVarDecl(s)
	}
}

func (s *VarDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitVarDecl(s)
	}
}

func (p *DaedalusParser) VarDecl() (localctx IVarDeclContext) {
	localctx = NewVarDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, DaedalusParserRULE_varDecl)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(170)
		p.Match(DaedalusParserVar)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(171)
		p.TypeReference()
	}
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(172)
			p.VarValueDecl()
		}

	case 2:
		{
			p.SetState(173)
			p.VarArrayDecl()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(176)
				p.Match(DaedalusParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(180)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(177)
					p.VarDecl()
				}

			case 2:
				{
					p.SetState(178)
					p.VarValueDecl()
				}

			case 3:
				{
					p.SetState(179)
					p.VarArrayDecl()
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(186)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 15, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMetaValueContext is an interface to support dynamic dispatch.
type IMetaValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsMetaValueContext differentiates from other interfaces.
	IsMetaValueContext()
}

type MetaValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMetaValueContext() *MetaValueContext {
	var p = new(MetaValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_metaValue
	return p
}

func InitEmptyMetaValueContext(p *MetaValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_metaValue
}

func (*MetaValueContext) IsMetaValueContext() {}

func NewMetaValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MetaValueContext {
	var p = new(MetaValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_metaValue

	return p
}

func (s *MetaValueContext) GetParser() antlr.Parser { return s.parser }
func (s *MetaValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MetaValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MetaValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterMetaValue(s)
	}
}

func (s *MetaValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitMetaValue(s)
	}
}

func (p *DaedalusParser) MetaValue() (localctx IMetaValueContext) {
	localctx = NewMetaValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, DaedalusParserRULE_metaValue)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1 + 1
	for ok := true; ok; ok = _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1 + 1:
			p.SetState(187)
			p.MatchWildcard()

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(190)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IZParserExtenderMetaContext is an interface to support dynamic dispatch.
type IZParserExtenderMetaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NameNode() INameNodeContext
	Assign() antlr.TerminalNode
	MetaValue() IMetaValueContext
	Semi() antlr.TerminalNode

	// IsZParserExtenderMetaContext differentiates from other interfaces.
	IsZParserExtenderMetaContext()
}

type ZParserExtenderMetaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyZParserExtenderMetaContext() *ZParserExtenderMetaContext {
	var p = new(ZParserExtenderMetaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_zParserExtenderMeta
	return p
}

func InitEmptyZParserExtenderMetaContext(p *ZParserExtenderMetaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_zParserExtenderMeta
}

func (*ZParserExtenderMetaContext) IsZParserExtenderMetaContext() {}

func NewZParserExtenderMetaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ZParserExtenderMetaContext {
	var p = new(ZParserExtenderMetaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_zParserExtenderMeta

	return p
}

func (s *ZParserExtenderMetaContext) GetParser() antlr.Parser { return s.parser }

func (s *ZParserExtenderMetaContext) NameNode() INameNodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameNodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameNodeContext)
}

func (s *ZParserExtenderMetaContext) Assign() antlr.TerminalNode {
	return s.GetToken(DaedalusParserAssign, 0)
}

func (s *ZParserExtenderMetaContext) MetaValue() IMetaValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMetaValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMetaValueContext)
}

func (s *ZParserExtenderMetaContext) Semi() antlr.TerminalNode {
	return s.GetToken(DaedalusParserSemi, 0)
}

func (s *ZParserExtenderMetaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ZParserExtenderMetaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ZParserExtenderMetaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterZParserExtenderMeta(s)
	}
}

func (s *ZParserExtenderMetaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitZParserExtenderMeta(s)
	}
}

func (p *DaedalusParser) ZParserExtenderMeta() (localctx IZParserExtenderMetaContext) {
	localctx = NewZParserExtenderMetaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, DaedalusParserRULE_zParserExtenderMeta)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		p.NameNode()
	}
	{
		p.SetState(193)
		p.Match(DaedalusParserAssign)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(194)
		p.MetaValue()
	}
	{
		p.SetState(195)
		p.Match(DaedalusParserSemi)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IZParserExtenderMetaBlockContext is an interface to support dynamic dispatch.
type IZParserExtenderMetaBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Meta() antlr.TerminalNode
	LeftBrace() antlr.TerminalNode
	RightBrace() antlr.TerminalNode
	Semi() antlr.TerminalNode
	AllZParserExtenderMeta() []IZParserExtenderMetaContext
	ZParserExtenderMeta(i int) IZParserExtenderMetaContext

	// IsZParserExtenderMetaBlockContext differentiates from other interfaces.
	IsZParserExtenderMetaBlockContext()
}

type ZParserExtenderMetaBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyZParserExtenderMetaBlockContext() *ZParserExtenderMetaBlockContext {
	var p = new(ZParserExtenderMetaBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_zParserExtenderMetaBlock
	return p
}

func InitEmptyZParserExtenderMetaBlockContext(p *ZParserExtenderMetaBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_zParserExtenderMetaBlock
}

func (*ZParserExtenderMetaBlockContext) IsZParserExtenderMetaBlockContext() {}

func NewZParserExtenderMetaBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ZParserExtenderMetaBlockContext {
	var p = new(ZParserExtenderMetaBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_zParserExtenderMetaBlock

	return p
}

func (s *ZParserExtenderMetaBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ZParserExtenderMetaBlockContext) Meta() antlr.TerminalNode {
	return s.GetToken(DaedalusParserMeta, 0)
}

func (s *ZParserExtenderMetaBlockContext) LeftBrace() antlr.TerminalNode {
	return s.GetToken(DaedalusParserLeftBrace, 0)
}

func (s *ZParserExtenderMetaBlockContext) RightBrace() antlr.TerminalNode {
	return s.GetToken(DaedalusParserRightBrace, 0)
}

func (s *ZParserExtenderMetaBlockContext) Semi() antlr.TerminalNode {
	return s.GetToken(DaedalusParserSemi, 0)
}

func (s *ZParserExtenderMetaBlockContext) AllZParserExtenderMeta() []IZParserExtenderMetaContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IZParserExtenderMetaContext); ok {
			len++
		}
	}

	tst := make([]IZParserExtenderMetaContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IZParserExtenderMetaContext); ok {
			tst[i] = t.(IZParserExtenderMetaContext)
			i++
		}
	}

	return tst
}

func (s *ZParserExtenderMetaBlockContext) ZParserExtenderMeta(i int) IZParserExtenderMetaContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IZParserExtenderMetaContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IZParserExtenderMetaContext)
}

func (s *ZParserExtenderMetaBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ZParserExtenderMetaBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ZParserExtenderMetaBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterZParserExtenderMetaBlock(s)
	}
}

func (s *ZParserExtenderMetaBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitZParserExtenderMetaBlock(s)
	}
}

func (p *DaedalusParser) ZParserExtenderMetaBlock() (localctx IZParserExtenderMetaBlockContext) {
	localctx = NewZParserExtenderMetaBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, DaedalusParserRULE_zParserExtenderMetaBlock)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(197)
		p.Match(DaedalusParserMeta)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(198)
		p.Match(DaedalusParserLeftBrace)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1+1 {
			{
				p.SetState(199)
				p.ZParserExtenderMeta()
			}

		}
		p.SetState(204)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(205)
		p.Match(DaedalusParserRightBrace)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(206)
		p.Match(DaedalusParserSemi)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstArrayDefContext is an interface to support dynamic dispatch.
type IConstArrayDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NameNode() INameNodeContext
	LeftBracket() antlr.TerminalNode
	ArraySize() IArraySizeContext
	RightBracket() antlr.TerminalNode
	ConstArrayAssignment() IConstArrayAssignmentContext

	// IsConstArrayDefContext differentiates from other interfaces.
	IsConstArrayDefContext()
}

type ConstArrayDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstArrayDefContext() *ConstArrayDefContext {
	var p = new(ConstArrayDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_constArrayDef
	return p
}

func InitEmptyConstArrayDefContext(p *ConstArrayDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_constArrayDef
}

func (*ConstArrayDefContext) IsConstArrayDefContext() {}

func NewConstArrayDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstArrayDefContext {
	var p = new(ConstArrayDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_constArrayDef

	return p
}

func (s *ConstArrayDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstArrayDefContext) NameNode() INameNodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameNodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameNodeContext)
}

func (s *ConstArrayDefContext) LeftBracket() antlr.TerminalNode {
	return s.GetToken(DaedalusParserLeftBracket, 0)
}

func (s *ConstArrayDefContext) ArraySize() IArraySizeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArraySizeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArraySizeContext)
}

func (s *ConstArrayDefContext) RightBracket() antlr.TerminalNode {
	return s.GetToken(DaedalusParserRightBracket, 0)
}

func (s *ConstArrayDefContext) ConstArrayAssignment() IConstArrayAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstArrayAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstArrayAssignmentContext)
}

func (s *ConstArrayDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstArrayDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstArrayDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterConstArrayDef(s)
	}
}

func (s *ConstArrayDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitConstArrayDef(s)
	}
}

func (p *DaedalusParser) ConstArrayDef() (localctx IConstArrayDefContext) {
	localctx = NewConstArrayDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, DaedalusParserRULE_constArrayDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(208)
		p.NameNode()
	}
	{
		p.SetState(209)
		p.Match(DaedalusParserLeftBracket)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(210)
		p.ArraySize()
	}
	{
		p.SetState(211)
		p.Match(DaedalusParserRightBracket)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(212)
		p.ConstArrayAssignment()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstArrayAssignmentContext is an interface to support dynamic dispatch.
type IConstArrayAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Assign() antlr.TerminalNode
	LeftBrace() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	RightBrace() antlr.TerminalNode
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsConstArrayAssignmentContext differentiates from other interfaces.
	IsConstArrayAssignmentContext()
}

type ConstArrayAssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstArrayAssignmentContext() *ConstArrayAssignmentContext {
	var p = new(ConstArrayAssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_constArrayAssignment
	return p
}

func InitEmptyConstArrayAssignmentContext(p *ConstArrayAssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_constArrayAssignment
}

func (*ConstArrayAssignmentContext) IsConstArrayAssignmentContext() {}

func NewConstArrayAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstArrayAssignmentContext {
	var p = new(ConstArrayAssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_constArrayAssignment

	return p
}

func (s *ConstArrayAssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstArrayAssignmentContext) Assign() antlr.TerminalNode {
	return s.GetToken(DaedalusParserAssign, 0)
}

func (s *ConstArrayAssignmentContext) LeftBrace() antlr.TerminalNode {
	return s.GetToken(DaedalusParserLeftBrace, 0)
}

func (s *ConstArrayAssignmentContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ConstArrayAssignmentContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConstArrayAssignmentContext) RightBrace() antlr.TerminalNode {
	return s.GetToken(DaedalusParserRightBrace, 0)
}

func (s *ConstArrayAssignmentContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(DaedalusParserComma)
}

func (s *ConstArrayAssignmentContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(DaedalusParserComma, i)
}

func (s *ConstArrayAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstArrayAssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstArrayAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterConstArrayAssignment(s)
	}
}

func (s *ConstArrayAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitConstArrayAssignment(s)
	}
}

func (p *DaedalusParser) ConstArrayAssignment() (localctx IConstArrayAssignmentContext) {
	localctx = NewConstArrayAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, DaedalusParserRULE_constArrayAssignment)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(214)
		p.Match(DaedalusParserAssign)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(215)
		p.Match(DaedalusParserLeftBrace)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(216)
		p.expression(0)
	}
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == DaedalusParserComma {
		{
			p.SetState(217)
			p.Match(DaedalusParserComma)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(218)
			p.expression(0)
		}

		p.SetState(223)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(224)
		p.Match(DaedalusParserRightBrace)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstValueDefContext is an interface to support dynamic dispatch.
type IConstValueDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NameNode() INameNodeContext
	ConstValueAssignment() IConstValueAssignmentContext

	// IsConstValueDefContext differentiates from other interfaces.
	IsConstValueDefContext()
}

type ConstValueDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstValueDefContext() *ConstValueDefContext {
	var p = new(ConstValueDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_constValueDef
	return p
}

func InitEmptyConstValueDefContext(p *ConstValueDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_constValueDef
}

func (*ConstValueDefContext) IsConstValueDefContext() {}

func NewConstValueDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstValueDefContext {
	var p = new(ConstValueDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_constValueDef

	return p
}

func (s *ConstValueDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstValueDefContext) NameNode() INameNodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameNodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameNodeContext)
}

func (s *ConstValueDefContext) ConstValueAssignment() IConstValueAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstValueAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstValueAssignmentContext)
}

func (s *ConstValueDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstValueDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstValueDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterConstValueDef(s)
	}
}

func (s *ConstValueDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitConstValueDef(s)
	}
}

func (p *DaedalusParser) ConstValueDef() (localctx IConstValueDefContext) {
	localctx = NewConstValueDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, DaedalusParserRULE_constValueDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(226)
		p.NameNode()
	}
	{
		p.SetState(227)
		p.ConstValueAssignment()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstValueAssignmentContext is an interface to support dynamic dispatch.
type IConstValueAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Assign() antlr.TerminalNode
	Expression() IExpressionContext

	// IsConstValueAssignmentContext differentiates from other interfaces.
	IsConstValueAssignmentContext()
}

type ConstValueAssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstValueAssignmentContext() *ConstValueAssignmentContext {
	var p = new(ConstValueAssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_constValueAssignment
	return p
}

func InitEmptyConstValueAssignmentContext(p *ConstValueAssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_constValueAssignment
}

func (*ConstValueAssignmentContext) IsConstValueAssignmentContext() {}

func NewConstValueAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstValueAssignmentContext {
	var p = new(ConstValueAssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_constValueAssignment

	return p
}

func (s *ConstValueAssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstValueAssignmentContext) Assign() antlr.TerminalNode {
	return s.GetToken(DaedalusParserAssign, 0)
}

func (s *ConstValueAssignmentContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConstValueAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstValueAssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstValueAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterConstValueAssignment(s)
	}
}

func (s *ConstValueAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitConstValueAssignment(s)
	}
}

func (p *DaedalusParser) ConstValueAssignment() (localctx IConstValueAssignmentContext) {
	localctx = NewConstValueAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, DaedalusParserRULE_constValueAssignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(229)
		p.Match(DaedalusParserAssign)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(230)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarArrayDeclContext is an interface to support dynamic dispatch.
type IVarArrayDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NameNode() INameNodeContext
	LeftBracket() antlr.TerminalNode
	ArraySize() IArraySizeContext
	RightBracket() antlr.TerminalNode

	// IsVarArrayDeclContext differentiates from other interfaces.
	IsVarArrayDeclContext()
}

type VarArrayDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarArrayDeclContext() *VarArrayDeclContext {
	var p = new(VarArrayDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_varArrayDecl
	return p
}

func InitEmptyVarArrayDeclContext(p *VarArrayDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_varArrayDecl
}

func (*VarArrayDeclContext) IsVarArrayDeclContext() {}

func NewVarArrayDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarArrayDeclContext {
	var p = new(VarArrayDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_varArrayDecl

	return p
}

func (s *VarArrayDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *VarArrayDeclContext) NameNode() INameNodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameNodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameNodeContext)
}

func (s *VarArrayDeclContext) LeftBracket() antlr.TerminalNode {
	return s.GetToken(DaedalusParserLeftBracket, 0)
}

func (s *VarArrayDeclContext) ArraySize() IArraySizeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArraySizeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArraySizeContext)
}

func (s *VarArrayDeclContext) RightBracket() antlr.TerminalNode {
	return s.GetToken(DaedalusParserRightBracket, 0)
}

func (s *VarArrayDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarArrayDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarArrayDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterVarArrayDecl(s)
	}
}

func (s *VarArrayDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitVarArrayDecl(s)
	}
}

func (p *DaedalusParser) VarArrayDecl() (localctx IVarArrayDeclContext) {
	localctx = NewVarArrayDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, DaedalusParserRULE_varArrayDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(232)
		p.NameNode()
	}
	{
		p.SetState(233)
		p.Match(DaedalusParserLeftBracket)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(234)
		p.ArraySize()
	}
	{
		p.SetState(235)
		p.Match(DaedalusParserRightBracket)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarValueDeclContext is an interface to support dynamic dispatch.
type IVarValueDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NameNode() INameNodeContext

	// IsVarValueDeclContext differentiates from other interfaces.
	IsVarValueDeclContext()
}

type VarValueDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarValueDeclContext() *VarValueDeclContext {
	var p = new(VarValueDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_varValueDecl
	return p
}

func InitEmptyVarValueDeclContext(p *VarValueDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_varValueDecl
}

func (*VarValueDeclContext) IsVarValueDeclContext() {}

func NewVarValueDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarValueDeclContext {
	var p = new(VarValueDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_varValueDecl

	return p
}

func (s *VarValueDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *VarValueDeclContext) NameNode() INameNodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameNodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameNodeContext)
}

func (s *VarValueDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarValueDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarValueDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterVarValueDecl(s)
	}
}

func (s *VarValueDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitVarValueDecl(s)
	}
}

func (p *DaedalusParser) VarValueDecl() (localctx IVarValueDeclContext) {
	localctx = NewVarValueDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, DaedalusParserRULE_varValueDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(237)
		p.NameNode()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParameterListContext is an interface to support dynamic dispatch.
type IParameterListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LeftParen() antlr.TerminalNode
	RightParen() antlr.TerminalNode
	AllParameterDecl() []IParameterDeclContext
	ParameterDecl(i int) IParameterDeclContext
	AllComma() []antlr.TerminalNode
	Comma(i int) antlr.TerminalNode

	// IsParameterListContext differentiates from other interfaces.
	IsParameterListContext()
}

type ParameterListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterListContext() *ParameterListContext {
	var p = new(ParameterListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_parameterList
	return p
}

func InitEmptyParameterListContext(p *ParameterListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_parameterList
}

func (*ParameterListContext) IsParameterListContext() {}

func NewParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterListContext {
	var p = new(ParameterListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_parameterList

	return p
}

func (s *ParameterListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterListContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(DaedalusParserLeftParen, 0)
}

func (s *ParameterListContext) RightParen() antlr.TerminalNode {
	return s.GetToken(DaedalusParserRightParen, 0)
}

func (s *ParameterListContext) AllParameterDecl() []IParameterDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParameterDeclContext); ok {
			len++
		}
	}

	tst := make([]IParameterDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParameterDeclContext); ok {
			tst[i] = t.(IParameterDeclContext)
			i++
		}
	}

	return tst
}

func (s *ParameterListContext) ParameterDecl(i int) IParameterDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterDeclContext)
}

func (s *ParameterListContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(DaedalusParserComma)
}

func (s *ParameterListContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(DaedalusParserComma, i)
}

func (s *ParameterListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterParameterList(s)
	}
}

func (s *ParameterListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitParameterList(s)
	}
}

func (p *DaedalusParser) ParameterList() (localctx IParameterListContext) {
	localctx = NewParameterListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, DaedalusParserRULE_parameterList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(239)
		p.Match(DaedalusParserLeftParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DaedalusParserVar {
		{
			p.SetState(240)
			p.ParameterDecl()
		}
		p.SetState(245)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == DaedalusParserComma {
			{
				p.SetState(241)
				p.Match(DaedalusParserComma)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(242)
				p.ParameterDecl()
			}

			p.SetState(247)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(250)
		p.Match(DaedalusParserRightParen)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParameterDeclContext is an interface to support dynamic dispatch.
type IParameterDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Var() antlr.TerminalNode
	TypeReference() ITypeReferenceContext
	NameNode() INameNodeContext
	LeftBracket() antlr.TerminalNode
	ArraySize() IArraySizeContext
	RightBracket() antlr.TerminalNode

	// IsParameterDeclContext differentiates from other interfaces.
	IsParameterDeclContext()
}

type ParameterDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterDeclContext() *ParameterDeclContext {
	var p = new(ParameterDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_parameterDecl
	return p
}

func InitEmptyParameterDeclContext(p *ParameterDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_parameterDecl
}

func (*ParameterDeclContext) IsParameterDeclContext() {}

func NewParameterDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterDeclContext {
	var p = new(ParameterDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_parameterDecl

	return p
}

func (s *ParameterDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterDeclContext) Var() antlr.TerminalNode {
	return s.GetToken(DaedalusParserVar, 0)
}

func (s *ParameterDeclContext) TypeReference() ITypeReferenceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeReferenceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeReferenceContext)
}

func (s *ParameterDeclContext) NameNode() INameNodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameNodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameNodeContext)
}

func (s *ParameterDeclContext) LeftBracket() antlr.TerminalNode {
	return s.GetToken(DaedalusParserLeftBracket, 0)
}

func (s *ParameterDeclContext) ArraySize() IArraySizeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArraySizeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArraySizeContext)
}

func (s *ParameterDeclContext) RightBracket() antlr.TerminalNode {
	return s.GetToken(DaedalusParserRightBracket, 0)
}

func (s *ParameterDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterParameterDecl(s)
	}
}

func (s *ParameterDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitParameterDecl(s)
	}
}

func (p *DaedalusParser) ParameterDecl() (localctx IParameterDeclContext) {
	localctx = NewParameterDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, DaedalusParserRULE_parameterDecl)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(252)
		p.Match(DaedalusParserVar)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(253)
		p.TypeReference()
	}
	{
		p.SetState(254)
		p.NameNode()
	}
	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DaedalusParserLeftBracket {
		{
			p.SetState(255)
			p.Match(DaedalusParserLeftBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(256)
			p.ArraySize()
		}
		{
			p.SetState(257)
			p.Match(DaedalusParserRightBracket)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementBlockContext is an interface to support dynamic dispatch.
type IStatementBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LeftBrace() antlr.TerminalNode
	RightBrace() antlr.TerminalNode
	AllSemi() []antlr.TerminalNode
	Semi(i int) antlr.TerminalNode
	AllIfBlockStatement() []IIfBlockStatementContext
	IfBlockStatement(i int) IIfBlockStatementContext
	AllReturnStatement() []IReturnStatementContext
	ReturnStatement(i int) IReturnStatementContext
	AllConstDef() []IConstDefContext
	ConstDef(i int) IConstDefContext
	AllVarDecl() []IVarDeclContext
	VarDecl(i int) IVarDeclContext
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsStatementBlockContext differentiates from other interfaces.
	IsStatementBlockContext()
}

type StatementBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementBlockContext() *StatementBlockContext {
	var p = new(StatementBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_statementBlock
	return p
}

func InitEmptyStatementBlockContext(p *StatementBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_statementBlock
}

func (*StatementBlockContext) IsStatementBlockContext() {}

func NewStatementBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementBlockContext {
	var p = new(StatementBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_statementBlock

	return p
}

func (s *StatementBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementBlockContext) LeftBrace() antlr.TerminalNode {
	return s.GetToken(DaedalusParserLeftBrace, 0)
}

func (s *StatementBlockContext) RightBrace() antlr.TerminalNode {
	return s.GetToken(DaedalusParserRightBrace, 0)
}

func (s *StatementBlockContext) AllSemi() []antlr.TerminalNode {
	return s.GetTokens(DaedalusParserSemi)
}

func (s *StatementBlockContext) Semi(i int) antlr.TerminalNode {
	return s.GetToken(DaedalusParserSemi, i)
}

func (s *StatementBlockContext) AllIfBlockStatement() []IIfBlockStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIfBlockStatementContext); ok {
			len++
		}
	}

	tst := make([]IIfBlockStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIfBlockStatementContext); ok {
			tst[i] = t.(IIfBlockStatementContext)
			i++
		}
	}

	return tst
}

func (s *StatementBlockContext) IfBlockStatement(i int) IIfBlockStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfBlockStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfBlockStatementContext)
}

func (s *StatementBlockContext) AllReturnStatement() []IReturnStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IReturnStatementContext); ok {
			len++
		}
	}

	tst := make([]IReturnStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IReturnStatementContext); ok {
			tst[i] = t.(IReturnStatementContext)
			i++
		}
	}

	return tst
}

func (s *StatementBlockContext) ReturnStatement(i int) IReturnStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnStatementContext)
}

func (s *StatementBlockContext) AllConstDef() []IConstDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConstDefContext); ok {
			len++
		}
	}

	tst := make([]IConstDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConstDefContext); ok {
			tst[i] = t.(IConstDefContext)
			i++
		}
	}

	return tst
}

func (s *StatementBlockContext) ConstDef(i int) IConstDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstDefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstDefContext)
}

func (s *StatementBlockContext) AllVarDecl() []IVarDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarDeclContext); ok {
			len++
		}
	}

	tst := make([]IVarDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarDeclContext); ok {
			tst[i] = t.(IVarDeclContext)
			i++
		}
	}

	return tst
}

func (s *StatementBlockContext) VarDecl(i int) IVarDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDeclContext)
}

func (s *StatementBlockContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *StatementBlockContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StatementBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterStatementBlock(s)
	}
}

func (s *StatementBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitStatementBlock(s)
	}
}

func (p *DaedalusParser) StatementBlock() (localctx IStatementBlockContext) {
	localctx = NewStatementBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, DaedalusParserRULE_statementBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(261)
		p.Match(DaedalusParserLeftBrace)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(276)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&36028906543775486) != 0 {
		p.SetState(274)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case DaedalusParserIntegerLiteral, DaedalusParserFloatLiteral, DaedalusParserStringLiteral, DaedalusParserConst, DaedalusParserVar, DaedalusParserInt, DaedalusParserFunc, DaedalusParserStringKeyword, DaedalusParserClass, DaedalusParserVoid, DaedalusParserReturn, DaedalusParserFloat, DaedalusParserPrototype, DaedalusParserInstance, DaedalusParserNamespace, DaedalusParserNull, DaedalusParserMeta, DaedalusParserLeftParen, DaedalusParserPlus, DaedalusParserMinus, DaedalusParserTilde, DaedalusParserNot, DaedalusParserIdentifier:
			p.SetState(266)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case DaedalusParserReturn:
				{
					p.SetState(262)
					p.ReturnStatement()
				}

			case DaedalusParserConst:
				{
					p.SetState(263)
					p.ConstDef()
				}

			case DaedalusParserVar:
				{
					p.SetState(264)
					p.VarDecl()
				}

			case DaedalusParserIntegerLiteral, DaedalusParserFloatLiteral, DaedalusParserStringLiteral, DaedalusParserInt, DaedalusParserFunc, DaedalusParserStringKeyword, DaedalusParserClass, DaedalusParserVoid, DaedalusParserFloat, DaedalusParserPrototype, DaedalusParserInstance, DaedalusParserNamespace, DaedalusParserNull, DaedalusParserMeta, DaedalusParserLeftParen, DaedalusParserPlus, DaedalusParserMinus, DaedalusParserTilde, DaedalusParserNot, DaedalusParserIdentifier:
				{
					p.SetState(265)
					p.expression(0)
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}
			{
				p.SetState(268)
				p.Match(DaedalusParserSemi)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case DaedalusParserIf:
			{
				p.SetState(270)
				p.IfBlockStatement()
			}
			p.SetState(272)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			if _la == DaedalusParserSemi {
				{
					p.SetState(271)
					p.Match(DaedalusParserSemi)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(278)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(279)
		p.Match(DaedalusParserRightBrace)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElseBlockContext is an interface to support dynamic dispatch.
type IElseBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Else() antlr.TerminalNode
	StatementBlock() IStatementBlockContext

	// IsElseBlockContext differentiates from other interfaces.
	IsElseBlockContext()
}

type ElseBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseBlockContext() *ElseBlockContext {
	var p = new(ElseBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_elseBlock
	return p
}

func InitEmptyElseBlockContext(p *ElseBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_elseBlock
}

func (*ElseBlockContext) IsElseBlockContext() {}

func NewElseBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseBlockContext {
	var p = new(ElseBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_elseBlock

	return p
}

func (s *ElseBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseBlockContext) Else() antlr.TerminalNode {
	return s.GetToken(DaedalusParserElse, 0)
}

func (s *ElseBlockContext) StatementBlock() IStatementBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementBlockContext)
}

func (s *ElseBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterElseBlock(s)
	}
}

func (s *ElseBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitElseBlock(s)
	}
}

func (p *DaedalusParser) ElseBlock() (localctx IElseBlockContext) {
	localctx = NewElseBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, DaedalusParserRULE_elseBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(281)
		p.Match(DaedalusParserElse)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(282)
		p.StatementBlock()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElseIfBlockContext is an interface to support dynamic dispatch.
type IElseIfBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Else() antlr.TerminalNode
	If() antlr.TerminalNode
	Expression() IExpressionContext
	StatementBlock() IStatementBlockContext

	// IsElseIfBlockContext differentiates from other interfaces.
	IsElseIfBlockContext()
}

type ElseIfBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseIfBlockContext() *ElseIfBlockContext {
	var p = new(ElseIfBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_elseIfBlock
	return p
}

func InitEmptyElseIfBlockContext(p *ElseIfBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_elseIfBlock
}

func (*ElseIfBlockContext) IsElseIfBlockContext() {}

func NewElseIfBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseIfBlockContext {
	var p = new(ElseIfBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_elseIfBlock

	return p
}

func (s *ElseIfBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseIfBlockContext) Else() antlr.TerminalNode {
	return s.GetToken(DaedalusParserElse, 0)
}

func (s *ElseIfBlockContext) If() antlr.TerminalNode {
	return s.GetToken(DaedalusParserIf, 0)
}

func (s *ElseIfBlockContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ElseIfBlockContext) StatementBlock() IStatementBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementBlockContext)
}

func (s *ElseIfBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseIfBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseIfBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterElseIfBlock(s)
	}
}

func (s *ElseIfBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitElseIfBlock(s)
	}
}

func (p *DaedalusParser) ElseIfBlock() (localctx IElseIfBlockContext) {
	localctx = NewElseIfBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, DaedalusParserRULE_elseIfBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(284)
		p.Match(DaedalusParserElse)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(285)
		p.Match(DaedalusParserIf)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(286)
		p.expression(0)
	}
	{
		p.SetState(287)
		p.StatementBlock()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfBlockContext is an interface to support dynamic dispatch.
type IIfBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	If() antlr.TerminalNode
	Expression() IExpressionContext
	StatementBlock() IStatementBlockContext

	// IsIfBlockContext differentiates from other interfaces.
	IsIfBlockContext()
}

type IfBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfBlockContext() *IfBlockContext {
	var p = new(IfBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_ifBlock
	return p
}

func InitEmptyIfBlockContext(p *IfBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_ifBlock
}

func (*IfBlockContext) IsIfBlockContext() {}

func NewIfBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfBlockContext {
	var p = new(IfBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_ifBlock

	return p
}

func (s *IfBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *IfBlockContext) If() antlr.TerminalNode {
	return s.GetToken(DaedalusParserIf, 0)
}

func (s *IfBlockContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfBlockContext) StatementBlock() IStatementBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementBlockContext)
}

func (s *IfBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterIfBlock(s)
	}
}

func (s *IfBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitIfBlock(s)
	}
}

func (p *DaedalusParser) IfBlock() (localctx IIfBlockContext) {
	localctx = NewIfBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, DaedalusParserRULE_ifBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(289)
		p.Match(DaedalusParserIf)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(290)
		p.expression(0)
	}
	{
		p.SetState(291)
		p.StatementBlock()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfBlockStatementContext is an interface to support dynamic dispatch.
type IIfBlockStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IfBlock() IIfBlockContext
	AllElseIfBlock() []IElseIfBlockContext
	ElseIfBlock(i int) IElseIfBlockContext
	ElseBlock() IElseBlockContext

	// IsIfBlockStatementContext differentiates from other interfaces.
	IsIfBlockStatementContext()
}

type IfBlockStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfBlockStatementContext() *IfBlockStatementContext {
	var p = new(IfBlockStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_ifBlockStatement
	return p
}

func InitEmptyIfBlockStatementContext(p *IfBlockStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_ifBlockStatement
}

func (*IfBlockStatementContext) IsIfBlockStatementContext() {}

func NewIfBlockStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfBlockStatementContext {
	var p = new(IfBlockStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_ifBlockStatement

	return p
}

func (s *IfBlockStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfBlockStatementContext) IfBlock() IIfBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfBlockContext)
}

func (s *IfBlockStatementContext) AllElseIfBlock() []IElseIfBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElseIfBlockContext); ok {
			len++
		}
	}

	tst := make([]IElseIfBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElseIfBlockContext); ok {
			tst[i] = t.(IElseIfBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfBlockStatementContext) ElseIfBlock(i int) IElseIfBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseIfBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseIfBlockContext)
}

func (s *IfBlockStatementContext) ElseBlock() IElseBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseBlockContext)
}

func (s *IfBlockStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfBlockStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfBlockStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterIfBlockStatement(s)
	}
}

func (s *IfBlockStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitIfBlockStatement(s)
	}
}

func (p *DaedalusParser) IfBlockStatement() (localctx IIfBlockStatementContext) {
	localctx = NewIfBlockStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, DaedalusParserRULE_ifBlockStatement)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(293)
		p.IfBlock()
	}
	p.SetState(297)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(294)
				p.ElseIfBlock()
			}

		}
		p.SetState(299)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(301)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == DaedalusParserElse {
		{
			p.SetState(300)
			p.ElseBlock()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturnStatementContext is an interface to support dynamic dispatch.
type IReturnStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Return() antlr.TerminalNode
	Expression() IExpressionContext

	// IsReturnStatementContext differentiates from other interfaces.
	IsReturnStatementContext()
}

type ReturnStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStatementContext() *ReturnStatementContext {
	var p = new(ReturnStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_returnStatement
	return p
}

func InitEmptyReturnStatementContext(p *ReturnStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_returnStatement
}

func (*ReturnStatementContext) IsReturnStatementContext() {}

func NewReturnStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_returnStatement

	return p
}

func (s *ReturnStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStatementContext) Return() antlr.TerminalNode {
	return s.GetToken(DaedalusParserReturn, 0)
}

func (s *ReturnStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterReturnStatement(s)
	}
}

func (s *ReturnStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitReturnStatement(s)
	}
}

func (p *DaedalusParser) ReturnStatement() (localctx IReturnStatementContext) {
	localctx = NewReturnStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, DaedalusParserRULE_returnStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(303)
		p.Match(DaedalusParserReturn)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(305)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&36028906543767182) != 0 {
		{
			p.SetState(304)
			p.expression(0)
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyAll(ctx *ExpressionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BitMoveExpressionContext struct {
	ExpressionContext
	op antlr.Token
}

func NewBitMoveExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitMoveExpressionContext {
	var p = new(BitMoveExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BitMoveExpressionContext) GetOp() antlr.Token { return s.op }

func (s *BitMoveExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *BitMoveExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitMoveExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BitMoveExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BitMoveExpressionContext) Shl() antlr.TerminalNode {
	return s.GetToken(DaedalusParserShl, 0)
}

func (s *BitMoveExpressionContext) Shr() antlr.TerminalNode {
	return s.GetToken(DaedalusParserShr, 0)
}

func (s *BitMoveExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterBitMoveExpression(s)
	}
}

func (s *BitMoveExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitBitMoveExpression(s)
	}
}

type ReferenceExpressionContext struct {
	ExpressionContext
	name   antlr.Token
	member antlr.Token
}

func NewReferenceExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReferenceExpressionContext {
	var p = new(ReferenceExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *ReferenceExpressionContext) GetName() antlr.Token { return s.name }

func (s *ReferenceExpressionContext) GetMember() antlr.Token { return s.member }

func (s *ReferenceExpressionContext) SetName(v antlr.Token) { s.name = v }

func (s *ReferenceExpressionContext) SetMember(v antlr.Token) { s.member = v }

func (s *ReferenceExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReferenceExpressionContext) AllVoid() []antlr.TerminalNode {
	return s.GetTokens(DaedalusParserVoid)
}

func (s *ReferenceExpressionContext) Void(i int) antlr.TerminalNode {
	return s.GetToken(DaedalusParserVoid, i)
}

func (s *ReferenceExpressionContext) AllInt() []antlr.TerminalNode {
	return s.GetTokens(DaedalusParserInt)
}

func (s *ReferenceExpressionContext) Int(i int) antlr.TerminalNode {
	return s.GetToken(DaedalusParserInt, i)
}

func (s *ReferenceExpressionContext) AllFloat() []antlr.TerminalNode {
	return s.GetTokens(DaedalusParserFloat)
}

func (s *ReferenceExpressionContext) Float(i int) antlr.TerminalNode {
	return s.GetToken(DaedalusParserFloat, i)
}

func (s *ReferenceExpressionContext) AllStringKeyword() []antlr.TerminalNode {
	return s.GetTokens(DaedalusParserStringKeyword)
}

func (s *ReferenceExpressionContext) StringKeyword(i int) antlr.TerminalNode {
	return s.GetToken(DaedalusParserStringKeyword, i)
}

func (s *ReferenceExpressionContext) AllFunc() []antlr.TerminalNode {
	return s.GetTokens(DaedalusParserFunc)
}

func (s *ReferenceExpressionContext) Func(i int) antlr.TerminalNode {
	return s.GetToken(DaedalusParserFunc, i)
}

func (s *ReferenceExpressionContext) AllInstance() []antlr.TerminalNode {
	return s.GetTokens(DaedalusParserInstance)
}

func (s *ReferenceExpressionContext) Instance(i int) antlr.TerminalNode {
	return s.GetToken(DaedalusParserInstance, i)
}

func (s *ReferenceExpressionContext) AllClass() []antlr.TerminalNode {
	return s.GetTokens(DaedalusParserClass)
}

func (s *ReferenceExpressionContext) Class(i int) antlr.TerminalNode {
	return s.GetToken(DaedalusParserClass, i)
}

func (s *ReferenceExpressionContext) AllPrototype() []antlr.TerminalNode {
	return s.GetTokens(DaedalusParserPrototype)
}

func (s *ReferenceExpressionContext) Prototype(i int) antlr.TerminalNode {
	return s.GetToken(DaedalusParserPrototype, i)
}

func (s *ReferenceExpressionContext) AllMeta() []antlr.TerminalNode {
	return s.GetTokens(DaedalusParserMeta)
}

func (s *ReferenceExpressionContext) Meta(i int) antlr.TerminalNode {
	return s.GetToken(DaedalusParserMeta, i)
}

func (s *ReferenceExpressionContext) AllNamespace() []antlr.TerminalNode {
	return s.GetTokens(DaedalusParserNamespace)
}

func (s *ReferenceExpressionContext) Namespace(i int) antlr.TerminalNode {
	return s.GetToken(DaedalusParserNamespace, i)
}

func (s *ReferenceExpressionContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(DaedalusParserIdentifier)
}

func (s *ReferenceExpressionContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(DaedalusParserIdentifier, i)
}

func (s *ReferenceExpressionContext) AllLeftBracket() []antlr.TerminalNode {
	return s.GetTokens(DaedalusParserLeftBracket)
}

func (s *ReferenceExpressionContext) LeftBracket(i int) antlr.TerminalNode {
	return s.GetToken(DaedalusParserLeftBracket, i)
}

func (s *ReferenceExpressionContext) AllArrayIndex() []IArrayIndexContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IArrayIndexContext); ok {
			len++
		}
	}

	tst := make([]IArrayIndexContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IArrayIndexContext); ok {
			tst[i] = t.(IArrayIndexContext)
			i++
		}
	}

	return tst
}

func (s *ReferenceExpressionContext) ArrayIndex(i int) IArrayIndexContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayIndexContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayIndexContext)
}

func (s *ReferenceExpressionContext) AllRightBracket() []antlr.TerminalNode {
	return s.GetTokens(DaedalusParserRightBracket)
}

func (s *ReferenceExpressionContext) RightBracket(i int) antlr.TerminalNode {
	return s.GetToken(DaedalusParserRightBracket, i)
}

func (s *ReferenceExpressionContext) Dot() antlr.TerminalNode {
	return s.GetToken(DaedalusParserDot, 0)
}

func (s *ReferenceExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterReferenceExpression(s)
	}
}

func (s *ReferenceExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitReferenceExpression(s)
	}
}

type AssignmentExpressionContext struct {
	ExpressionContext
	op antlr.Token
}

func NewAssignmentExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentExpressionContext {
	var p = new(AssignmentExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AssignmentExpressionContext) GetOp() antlr.Token { return s.op }

func (s *AssignmentExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *AssignmentExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AssignmentExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentExpressionContext) Assign() antlr.TerminalNode {
	return s.GetToken(DaedalusParserAssign, 0)
}

func (s *AssignmentExpressionContext) PlusAssign() antlr.TerminalNode {
	return s.GetToken(DaedalusParserPlusAssign, 0)
}

func (s *AssignmentExpressionContext) MinusAssign() antlr.TerminalNode {
	return s.GetToken(DaedalusParserMinusAssign, 0)
}

func (s *AssignmentExpressionContext) StarAssign() antlr.TerminalNode {
	return s.GetToken(DaedalusParserStarAssign, 0)
}

func (s *AssignmentExpressionContext) DivAssign() antlr.TerminalNode {
	return s.GetToken(DaedalusParserDivAssign, 0)
}

func (s *AssignmentExpressionContext) AndAssign() antlr.TerminalNode {
	return s.GetToken(DaedalusParserAndAssign, 0)
}

func (s *AssignmentExpressionContext) OrAssign() antlr.TerminalNode {
	return s.GetToken(DaedalusParserOrAssign, 0)
}

func (s *AssignmentExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterAssignmentExpression(s)
	}
}

func (s *AssignmentExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitAssignmentExpression(s)
	}
}

type BracketExpressionContext struct {
	ExpressionContext
}

func NewBracketExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BracketExpressionContext {
	var p = new(BracketExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BracketExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BracketExpressionContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(DaedalusParserLeftParen, 0)
}

func (s *BracketExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BracketExpressionContext) RightParen() antlr.TerminalNode {
	return s.GetToken(DaedalusParserRightParen, 0)
}

func (s *BracketExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterBracketExpression(s)
	}
}

func (s *BracketExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitBracketExpression(s)
	}
}

type UnaryOperationContext struct {
	ExpressionContext
	op antlr.Token
}

func NewUnaryOperationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryOperationContext {
	var p = new(UnaryOperationContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *UnaryOperationContext) GetOp() antlr.Token { return s.op }

func (s *UnaryOperationContext) SetOp(v antlr.Token) { s.op = v }

func (s *UnaryOperationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryOperationContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *UnaryOperationContext) Plus() antlr.TerminalNode {
	return s.GetToken(DaedalusParserPlus, 0)
}

func (s *UnaryOperationContext) Tilde() antlr.TerminalNode {
	return s.GetToken(DaedalusParserTilde, 0)
}

func (s *UnaryOperationContext) Minus() antlr.TerminalNode {
	return s.GetToken(DaedalusParserMinus, 0)
}

func (s *UnaryOperationContext) Not() antlr.TerminalNode {
	return s.GetToken(DaedalusParserNot, 0)
}

func (s *UnaryOperationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterUnaryOperation(s)
	}
}

func (s *UnaryOperationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitUnaryOperation(s)
	}
}

type FloatLiteralExpressionContext struct {
	ExpressionContext
}

func NewFloatLiteralExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloatLiteralExpressionContext {
	var p = new(FloatLiteralExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *FloatLiteralExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatLiteralExpressionContext) FloatLiteral() antlr.TerminalNode {
	return s.GetToken(DaedalusParserFloatLiteral, 0)
}

func (s *FloatLiteralExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterFloatLiteralExpression(s)
	}
}

func (s *FloatLiteralExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitFloatLiteralExpression(s)
	}
}

type LogAndExpressionContext struct {
	ExpressionContext
	op antlr.Token
}

func NewLogAndExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogAndExpressionContext {
	var p = new(LogAndExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *LogAndExpressionContext) GetOp() antlr.Token { return s.op }

func (s *LogAndExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *LogAndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogAndExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *LogAndExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LogAndExpressionContext) And() antlr.TerminalNode {
	return s.GetToken(DaedalusParserAnd, 0)
}

func (s *LogAndExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterLogAndExpression(s)
	}
}

func (s *LogAndExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitLogAndExpression(s)
	}
}

type FuncCallExpressionContext struct {
	ExpressionContext
	name antlr.Token
}

func NewFuncCallExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FuncCallExpressionContext {
	var p = new(FuncCallExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *FuncCallExpressionContext) GetName() antlr.Token { return s.name }

func (s *FuncCallExpressionContext) SetName(v antlr.Token) { s.name = v }

func (s *FuncCallExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncCallExpressionContext) LeftParen() antlr.TerminalNode {
	return s.GetToken(DaedalusParserLeftParen, 0)
}

func (s *FuncCallExpressionContext) RightParen() antlr.TerminalNode {
	return s.GetToken(DaedalusParserRightParen, 0)
}

func (s *FuncCallExpressionContext) Void() antlr.TerminalNode {
	return s.GetToken(DaedalusParserVoid, 0)
}

func (s *FuncCallExpressionContext) Int() antlr.TerminalNode {
	return s.GetToken(DaedalusParserInt, 0)
}

func (s *FuncCallExpressionContext) Float() antlr.TerminalNode {
	return s.GetToken(DaedalusParserFloat, 0)
}

func (s *FuncCallExpressionContext) StringKeyword() antlr.TerminalNode {
	return s.GetToken(DaedalusParserStringKeyword, 0)
}

func (s *FuncCallExpressionContext) Func() antlr.TerminalNode {
	return s.GetToken(DaedalusParserFunc, 0)
}

func (s *FuncCallExpressionContext) Instance() antlr.TerminalNode {
	return s.GetToken(DaedalusParserInstance, 0)
}

func (s *FuncCallExpressionContext) Class() antlr.TerminalNode {
	return s.GetToken(DaedalusParserClass, 0)
}

func (s *FuncCallExpressionContext) Prototype() antlr.TerminalNode {
	return s.GetToken(DaedalusParserPrototype, 0)
}

func (s *FuncCallExpressionContext) Meta() antlr.TerminalNode {
	return s.GetToken(DaedalusParserMeta, 0)
}

func (s *FuncCallExpressionContext) Namespace() antlr.TerminalNode {
	return s.GetToken(DaedalusParserNamespace, 0)
}

func (s *FuncCallExpressionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(DaedalusParserIdentifier, 0)
}

func (s *FuncCallExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *FuncCallExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FuncCallExpressionContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(DaedalusParserComma)
}

func (s *FuncCallExpressionContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(DaedalusParserComma, i)
}

func (s *FuncCallExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterFuncCallExpression(s)
	}
}

func (s *FuncCallExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitFuncCallExpression(s)
	}
}

type StringLiteralExpressionContext struct {
	ExpressionContext
}

func NewStringLiteralExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringLiteralExpressionContext {
	var p = new(StringLiteralExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *StringLiteralExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralExpressionContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(DaedalusParserStringLiteral, 0)
}

func (s *StringLiteralExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterStringLiteralExpression(s)
	}
}

func (s *StringLiteralExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitStringLiteralExpression(s)
	}
}

type EqExpressionContext struct {
	ExpressionContext
	op antlr.Token
}

func NewEqExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqExpressionContext {
	var p = new(EqExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *EqExpressionContext) GetOp() antlr.Token { return s.op }

func (s *EqExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *EqExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *EqExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *EqExpressionContext) Equal() antlr.TerminalNode {
	return s.GetToken(DaedalusParserEqual, 0)
}

func (s *EqExpressionContext) NotEqual() antlr.TerminalNode {
	return s.GetToken(DaedalusParserNotEqual, 0)
}

func (s *EqExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterEqExpression(s)
	}
}

func (s *EqExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitEqExpression(s)
	}
}

type AddExpressionContext struct {
	ExpressionContext
	op antlr.Token
}

func NewAddExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddExpressionContext {
	var p = new(AddExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *AddExpressionContext) GetOp() antlr.Token { return s.op }

func (s *AddExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AddExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AddExpressionContext) Plus() antlr.TerminalNode {
	return s.GetToken(DaedalusParserPlus, 0)
}

func (s *AddExpressionContext) Minus() antlr.TerminalNode {
	return s.GetToken(DaedalusParserMinus, 0)
}

func (s *AddExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterAddExpression(s)
	}
}

func (s *AddExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitAddExpression(s)
	}
}

type CompExpressionContext struct {
	ExpressionContext
	op antlr.Token
}

func NewCompExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompExpressionContext {
	var p = new(CompExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *CompExpressionContext) GetOp() antlr.Token { return s.op }

func (s *CompExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *CompExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *CompExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CompExpressionContext) Less() antlr.TerminalNode {
	return s.GetToken(DaedalusParserLess, 0)
}

func (s *CompExpressionContext) Greater() antlr.TerminalNode {
	return s.GetToken(DaedalusParserGreater, 0)
}

func (s *CompExpressionContext) LessEq() antlr.TerminalNode {
	return s.GetToken(DaedalusParserLessEq, 0)
}

func (s *CompExpressionContext) GreaterEq() antlr.TerminalNode {
	return s.GetToken(DaedalusParserGreaterEq, 0)
}

func (s *CompExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterCompExpression(s)
	}
}

func (s *CompExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitCompExpression(s)
	}
}

type LogOrExpressionContext struct {
	ExpressionContext
	op antlr.Token
}

func NewLogOrExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogOrExpressionContext {
	var p = new(LogOrExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *LogOrExpressionContext) GetOp() antlr.Token { return s.op }

func (s *LogOrExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *LogOrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogOrExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *LogOrExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LogOrExpressionContext) Or() antlr.TerminalNode {
	return s.GetToken(DaedalusParserOr, 0)
}

func (s *LogOrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterLogOrExpression(s)
	}
}

func (s *LogOrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitLogOrExpression(s)
	}
}

type BinAndExpressionContext struct {
	ExpressionContext
	op antlr.Token
}

func NewBinAndExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinAndExpressionContext {
	var p = new(BinAndExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BinAndExpressionContext) GetOp() antlr.Token { return s.op }

func (s *BinAndExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *BinAndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinAndExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BinAndExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BinAndExpressionContext) BitAnd() antlr.TerminalNode {
	return s.GetToken(DaedalusParserBitAnd, 0)
}

func (s *BinAndExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterBinAndExpression(s)
	}
}

func (s *BinAndExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitBinAndExpression(s)
	}
}

type IntegerLiteralExpressionContext struct {
	ExpressionContext
}

func NewIntegerLiteralExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntegerLiteralExpressionContext {
	var p = new(IntegerLiteralExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *IntegerLiteralExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerLiteralExpressionContext) IntegerLiteral() antlr.TerminalNode {
	return s.GetToken(DaedalusParserIntegerLiteral, 0)
}

func (s *IntegerLiteralExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterIntegerLiteralExpression(s)
	}
}

func (s *IntegerLiteralExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitIntegerLiteralExpression(s)
	}
}

type BinOrExpressionContext struct {
	ExpressionContext
	op antlr.Token
}

func NewBinOrExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinOrExpressionContext {
	var p = new(BinOrExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *BinOrExpressionContext) GetOp() antlr.Token { return s.op }

func (s *BinOrExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *BinOrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinOrExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BinOrExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BinOrExpressionContext) BitOr() antlr.TerminalNode {
	return s.GetToken(DaedalusParserBitOr, 0)
}

func (s *BinOrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterBinOrExpression(s)
	}
}

func (s *BinOrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitBinOrExpression(s)
	}
}

type MultExpressionContext struct {
	ExpressionContext
	op antlr.Token
}

func NewMultExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultExpressionContext {
	var p = new(MultExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *MultExpressionContext) GetOp() antlr.Token { return s.op }

func (s *MultExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *MultExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *MultExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MultExpressionContext) Star() antlr.TerminalNode {
	return s.GetToken(DaedalusParserStar, 0)
}

func (s *MultExpressionContext) Div() antlr.TerminalNode {
	return s.GetToken(DaedalusParserDiv, 0)
}

func (s *MultExpressionContext) Percent() antlr.TerminalNode {
	return s.GetToken(DaedalusParserPercent, 0)
}

func (s *MultExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterMultExpression(s)
	}
}

func (s *MultExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitMultExpression(s)
	}
}

type NullLiteralExpressionContext struct {
	ExpressionContext
}

func NewNullLiteralExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullLiteralExpressionContext {
	var p = new(NullLiteralExpressionContext)

	InitEmptyExpressionContext(&p.ExpressionContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExpressionContext))

	return p
}

func (s *NullLiteralExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullLiteralExpressionContext) Null() antlr.TerminalNode {
	return s.GetToken(DaedalusParserNull, 0)
}

func (s *NullLiteralExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterNullLiteralExpression(s)
	}
}

func (s *NullLiteralExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitNullLiteralExpression(s)
	}
}

func (p *DaedalusParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *DaedalusParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 58
	p.EnterRecursionRule(localctx, 58, DaedalusParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(348)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		localctx = NewBracketExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(308)
			p.Match(DaedalusParserLeftParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(309)
			p.expression(0)
		}
		{
			p.SetState(310)
			p.Match(DaedalusParserRightParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewUnaryOperationContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(312)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*UnaryOperationContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&109521666048) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*UnaryOperationContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(313)
			p.expression(17)
		}

	case 3:
		localctx = NewIntegerLiteralExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(314)
			p.Match(DaedalusParserIntegerLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewFloatLiteralExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(315)
			p.Match(DaedalusParserFloatLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewStringLiteralExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(316)
			p.Match(DaedalusParserStringLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		localctx = NewNullLiteralExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(317)
			p.Match(DaedalusParserNull)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewFuncCallExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(318)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*FuncCallExpressionContext).name = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&36028797019741824) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*FuncCallExpressionContext).name = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(319)
			p.Match(DaedalusParserLeftParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(328)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&36028906543767182) != 0 {
			{
				p.SetState(320)
				p.expression(0)
			}
			p.SetState(325)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)

			for _la == DaedalusParserComma {
				{
					p.SetState(321)
					p.Match(DaedalusParserComma)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(322)
					p.expression(0)
				}

				p.SetState(327)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(330)
			p.Match(DaedalusParserRightParen)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewReferenceExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(331)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ReferenceExpressionContext).name = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&36028797019741824) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ReferenceExpressionContext).name = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(336)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(332)
				p.Match(DaedalusParserLeftBracket)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(333)
				p.ArrayIndex()
			}
			{
				p.SetState(334)
				p.Match(DaedalusParserRightBracket)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(346)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(338)
				p.Match(DaedalusParserDot)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(339)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*ReferenceExpressionContext).member = _lt

				_la = p.GetTokenStream().LA(1)

				if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&36028797019741824) != 0) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*ReferenceExpressionContext).member = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			p.SetState(344)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext()) == 1 {
				{
					p.SetState(340)
					p.Match(DaedalusParserLeftBracket)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(341)
					p.ArrayIndex()
				}
				{
					p.SetState(342)
					p.Match(DaedalusParserRightBracket)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			} else if p.HasError() { // JIM
				goto errorExit
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(382)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(380)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMultExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DaedalusParserRULE_expression)
				p.SetState(350)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(351)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MultExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&70394513981440) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MultExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(352)
					p.expression(17)
				}

			case 2:
				localctx = NewAddExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DaedalusParserRULE_expression)
				p.SetState(353)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(354)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == DaedalusParserPlus || _la == DaedalusParserMinus) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(355)
					p.expression(16)
				}

			case 3:
				localctx = NewBitMoveExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DaedalusParserRULE_expression)
				p.SetState(356)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(357)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*BitMoveExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == DaedalusParserShl || _la == DaedalusParserShr) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*BitMoveExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(358)
					p.expression(15)
				}

			case 4:
				localctx = NewCompExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DaedalusParserRULE_expression)
				p.SetState(359)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(360)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*CompExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&14018773254144) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*CompExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(361)
					p.expression(14)
				}

			case 5:
				localctx = NewEqExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DaedalusParserRULE_expression)
				p.SetState(362)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(363)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*EqExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == DaedalusParserEqual || _la == DaedalusParserNotEqual) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*EqExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(364)
					p.expression(13)
				}

			case 6:
				localctx = NewBinAndExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DaedalusParserRULE_expression)
				p.SetState(365)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(366)

					var _m = p.Match(DaedalusParserBitAnd)

					localctx.(*BinAndExpressionContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(367)
					p.expression(12)
				}

			case 7:
				localctx = NewBinOrExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DaedalusParserRULE_expression)
				p.SetState(368)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(369)

					var _m = p.Match(DaedalusParserBitOr)

					localctx.(*BinOrExpressionContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(370)
					p.expression(11)
				}

			case 8:
				localctx = NewLogAndExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DaedalusParserRULE_expression)
				p.SetState(371)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(372)

					var _m = p.Match(DaedalusParserAnd)

					localctx.(*LogAndExpressionContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(373)
					p.expression(10)
				}

			case 9:
				localctx = NewLogOrExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DaedalusParserRULE_expression)
				p.SetState(374)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(375)

					var _m = p.Match(DaedalusParserOr)

					localctx.(*LogOrExpressionContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(376)
					p.expression(9)
				}

			case 10:
				localctx = NewAssignmentExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, DaedalusParserRULE_expression)
				p.SetState(377)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(378)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AssignmentExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8866599205339136) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AssignmentExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(379)
					p.expression(7)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(384)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 36, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArrayIndexContext is an interface to support dynamic dispatch.
type IArrayIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IntegerLiteral() antlr.TerminalNode
	NameNode() INameNodeContext

	// IsArrayIndexContext differentiates from other interfaces.
	IsArrayIndexContext()
}

type ArrayIndexContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayIndexContext() *ArrayIndexContext {
	var p = new(ArrayIndexContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_arrayIndex
	return p
}

func InitEmptyArrayIndexContext(p *ArrayIndexContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_arrayIndex
}

func (*ArrayIndexContext) IsArrayIndexContext() {}

func NewArrayIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayIndexContext {
	var p = new(ArrayIndexContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_arrayIndex

	return p
}

func (s *ArrayIndexContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayIndexContext) IntegerLiteral() antlr.TerminalNode {
	return s.GetToken(DaedalusParserIntegerLiteral, 0)
}

func (s *ArrayIndexContext) NameNode() INameNodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameNodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameNodeContext)
}

func (s *ArrayIndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayIndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayIndexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterArrayIndex(s)
	}
}

func (s *ArrayIndexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitArrayIndex(s)
	}
}

func (p *DaedalusParser) ArrayIndex() (localctx IArrayIndexContext) {
	localctx = NewArrayIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, DaedalusParserRULE_arrayIndex)
	p.SetState(387)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case DaedalusParserIntegerLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(385)
			p.Match(DaedalusParserIntegerLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case DaedalusParserVar, DaedalusParserInt, DaedalusParserFunc, DaedalusParserStringKeyword, DaedalusParserClass, DaedalusParserVoid, DaedalusParserFloat, DaedalusParserPrototype, DaedalusParserInstance, DaedalusParserNamespace, DaedalusParserNull, DaedalusParserMeta, DaedalusParserIdentifier:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(386)
			p.NameNode()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArraySizeContext is an interface to support dynamic dispatch.
type IArraySizeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IntegerLiteral() antlr.TerminalNode
	NameNode() INameNodeContext

	// IsArraySizeContext differentiates from other interfaces.
	IsArraySizeContext()
}

type ArraySizeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArraySizeContext() *ArraySizeContext {
	var p = new(ArraySizeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_arraySize
	return p
}

func InitEmptyArraySizeContext(p *ArraySizeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_arraySize
}

func (*ArraySizeContext) IsArraySizeContext() {}

func NewArraySizeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArraySizeContext {
	var p = new(ArraySizeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_arraySize

	return p
}

func (s *ArraySizeContext) GetParser() antlr.Parser { return s.parser }

func (s *ArraySizeContext) IntegerLiteral() antlr.TerminalNode {
	return s.GetToken(DaedalusParserIntegerLiteral, 0)
}

func (s *ArraySizeContext) NameNode() INameNodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INameNodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INameNodeContext)
}

func (s *ArraySizeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArraySizeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArraySizeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterArraySize(s)
	}
}

func (s *ArraySizeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitArraySize(s)
	}
}

func (p *DaedalusParser) ArraySize() (localctx IArraySizeContext) {
	localctx = NewArraySizeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, DaedalusParserRULE_arraySize)
	p.SetState(391)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case DaedalusParserIntegerLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(389)
			p.Match(DaedalusParserIntegerLiteral)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case DaedalusParserVar, DaedalusParserInt, DaedalusParserFunc, DaedalusParserStringKeyword, DaedalusParserClass, DaedalusParserVoid, DaedalusParserFloat, DaedalusParserPrototype, DaedalusParserInstance, DaedalusParserNamespace, DaedalusParserNull, DaedalusParserMeta, DaedalusParserIdentifier:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(390)
			p.NameNode()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypeReferenceContext is an interface to support dynamic dispatch.
type ITypeReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode
	Void() antlr.TerminalNode
	Int() antlr.TerminalNode
	Float() antlr.TerminalNode
	StringKeyword() antlr.TerminalNode
	Func() antlr.TerminalNode
	Instance() antlr.TerminalNode

	// IsTypeReferenceContext differentiates from other interfaces.
	IsTypeReferenceContext()
}

type TypeReferenceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeReferenceContext() *TypeReferenceContext {
	var p = new(TypeReferenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_typeReference
	return p
}

func InitEmptyTypeReferenceContext(p *TypeReferenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_typeReference
}

func (*TypeReferenceContext) IsTypeReferenceContext() {}

func NewTypeReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeReferenceContext {
	var p = new(TypeReferenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_typeReference

	return p
}

func (s *TypeReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeReferenceContext) Identifier() antlr.TerminalNode {
	return s.GetToken(DaedalusParserIdentifier, 0)
}

func (s *TypeReferenceContext) Void() antlr.TerminalNode {
	return s.GetToken(DaedalusParserVoid, 0)
}

func (s *TypeReferenceContext) Int() antlr.TerminalNode {
	return s.GetToken(DaedalusParserInt, 0)
}

func (s *TypeReferenceContext) Float() antlr.TerminalNode {
	return s.GetToken(DaedalusParserFloat, 0)
}

func (s *TypeReferenceContext) StringKeyword() antlr.TerminalNode {
	return s.GetToken(DaedalusParserStringKeyword, 0)
}

func (s *TypeReferenceContext) Func() antlr.TerminalNode {
	return s.GetToken(DaedalusParserFunc, 0)
}

func (s *TypeReferenceContext) Instance() antlr.TerminalNode {
	return s.GetToken(DaedalusParserInstance, 0)
}

func (s *TypeReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterTypeReference(s)
	}
}

func (s *TypeReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitTypeReference(s)
	}
}

func (p *DaedalusParser) TypeReference() (localctx ITypeReferenceContext) {
	localctx = NewTypeReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, DaedalusParserRULE_typeReference)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(393)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&36028797019051648) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INameNodeContext is an interface to support dynamic dispatch.
type INameNodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Void() antlr.TerminalNode
	Var() antlr.TerminalNode
	Int() antlr.TerminalNode
	Float() antlr.TerminalNode
	StringKeyword() antlr.TerminalNode
	Func() antlr.TerminalNode
	Instance() antlr.TerminalNode
	Class() antlr.TerminalNode
	Prototype() antlr.TerminalNode
	Null() antlr.TerminalNode
	Meta() antlr.TerminalNode
	Namespace() antlr.TerminalNode
	Identifier() antlr.TerminalNode

	// IsNameNodeContext differentiates from other interfaces.
	IsNameNodeContext()
}

type NameNodeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameNodeContext() *NameNodeContext {
	var p = new(NameNodeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_nameNode
	return p
}

func InitEmptyNameNodeContext(p *NameNodeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_nameNode
}

func (*NameNodeContext) IsNameNodeContext() {}

func NewNameNodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameNodeContext {
	var p = new(NameNodeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_nameNode

	return p
}

func (s *NameNodeContext) GetParser() antlr.Parser { return s.parser }

func (s *NameNodeContext) Void() antlr.TerminalNode {
	return s.GetToken(DaedalusParserVoid, 0)
}

func (s *NameNodeContext) Var() antlr.TerminalNode {
	return s.GetToken(DaedalusParserVar, 0)
}

func (s *NameNodeContext) Int() antlr.TerminalNode {
	return s.GetToken(DaedalusParserInt, 0)
}

func (s *NameNodeContext) Float() antlr.TerminalNode {
	return s.GetToken(DaedalusParserFloat, 0)
}

func (s *NameNodeContext) StringKeyword() antlr.TerminalNode {
	return s.GetToken(DaedalusParserStringKeyword, 0)
}

func (s *NameNodeContext) Func() antlr.TerminalNode {
	return s.GetToken(DaedalusParserFunc, 0)
}

func (s *NameNodeContext) Instance() antlr.TerminalNode {
	return s.GetToken(DaedalusParserInstance, 0)
}

func (s *NameNodeContext) Class() antlr.TerminalNode {
	return s.GetToken(DaedalusParserClass, 0)
}

func (s *NameNodeContext) Prototype() antlr.TerminalNode {
	return s.GetToken(DaedalusParserPrototype, 0)
}

func (s *NameNodeContext) Null() antlr.TerminalNode {
	return s.GetToken(DaedalusParserNull, 0)
}

func (s *NameNodeContext) Meta() antlr.TerminalNode {
	return s.GetToken(DaedalusParserMeta, 0)
}

func (s *NameNodeContext) Namespace() antlr.TerminalNode {
	return s.GetToken(DaedalusParserNamespace, 0)
}

func (s *NameNodeContext) Identifier() antlr.TerminalNode {
	return s.GetToken(DaedalusParserIdentifier, 0)
}

func (s *NameNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameNodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterNameNode(s)
	}
}

func (s *NameNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitNameNode(s)
	}
}

func (p *DaedalusParser) NameNode() (localctx INameNodeContext) {
	localctx = NewNameNodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, DaedalusParserRULE_nameNode)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(395)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&36028797020004000) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParentReferenceContext is an interface to support dynamic dispatch.
type IParentReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() antlr.TerminalNode

	// IsParentReferenceContext differentiates from other interfaces.
	IsParentReferenceContext()
}

type ParentReferenceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParentReferenceContext() *ParentReferenceContext {
	var p = new(ParentReferenceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_parentReference
	return p
}

func InitEmptyParentReferenceContext(p *ParentReferenceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = DaedalusParserRULE_parentReference
}

func (*ParentReferenceContext) IsParentReferenceContext() {}

func NewParentReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParentReferenceContext {
	var p = new(ParentReferenceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = DaedalusParserRULE_parentReference

	return p
}

func (s *ParentReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *ParentReferenceContext) Identifier() antlr.TerminalNode {
	return s.GetToken(DaedalusParserIdentifier, 0)
}

func (s *ParentReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParentReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParentReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.EnterParentReference(s)
	}
}

func (s *ParentReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DaedalusListener); ok {
		listenerT.ExitParentReference(s)
	}
}

func (p *DaedalusParser) ParentReference() (localctx IParentReferenceContext) {
	localctx = NewParentReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, DaedalusParserRULE_parentReference)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(397)
		p.Match(DaedalusParserIdentifier)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *DaedalusParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 29:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *DaedalusParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 7)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
