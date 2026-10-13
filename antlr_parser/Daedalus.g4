// Daedalus grammar.
//
// The language rules follow the original engine (zParser_Lexer.cpp: ParseFile, ParseBlock, Declare*,
// Parse_Expression), but are kept more lenient where that helps the language server (e.g. a missing
// ';' after an if-block is reported by the validating listener instead of a syntax error).
//
// The tree is kept small on purpose, because parse trees of whole projects are built and walked:
//  * names inside expressions are plain tokens (no nameNode/anyIdentifier rule contexts)
//  * operators are token sets labelled `op` (no operator rule contexts)
//  * there are no single-child wrapper rules (statement, expressionBlock, value, reference, ...)
//  * loops are greedy (`*`), non-greedy loops need full prediction at every iteration
grammar Daedalus;

@lexer::members {
    const (
        // Channel that contains comments
        COMMENTS int = 2;
    )
}

// lexer
IntegerLiteral: Digit+;
FloatLiteral: PointFloat | ExponentFloat;
StringLiteral: '"' (~["\r\n])* '"';

Const: C O N S T;
Var: V A R;
If: I F;
Int: I N T;
Else: E L S E;
Func: F U N C;
StringKeyword: S T R I N G;
Class: C L A S S;
Void: V O I D;
Return: R E T U R N;
Float: F L O A T;
Prototype: P R O T O T Y P E;
Instance: I N S T A N C E;
Namespace: N A M E S P A C E;
Null: N U L L;
Meta: M E T A;

Comma: ',';
LeftParen: '(';
RightParen: ')';

LeftBracket: '[';
RightBracket: ']';

LeftBrace: '{';
RightBrace: '}';

BitAnd: '&';
And: '&&';
BitOr: '|';
Or: '||';
Plus: '+';
Minus: '-';
Div: '/';
Star: '*';
Tilde: '~';
Not: '!';
Assign: '=';
Less: '<';
Greater: '>';
Shl: '<<';
Shr: '>>';
LessEq: '<=';
GreaterEq: '>=';
Equal: '==';
NotEqual: '!=';
Percent: '%';

PlusAssign: '+=';
MinusAssign: '-=';
StarAssign: '*=';
DivAssign: '/=';
AndAssign: '&=';
OrAssign: '|=';

Dot: '.';
Semi: ';';

Identifier: (NonDigit | Digit) IdContinue*;

Whitespace: [ \t]+ -> skip;
Newline: ('\r' '\n'? | '\n') -> skip;
BlockComment: '/*' .*? '*/' -> skip;
LineComment: '//' ~[\r\n]* -> channel(2);

// fragments
fragment NonDigit: GermanCharacter | [a-zA-Z_];
fragment IdContinue: NonDigit | IdSpecial | Digit;
fragment IdSpecial: [@^];
fragment GermanCharacter:
	//   ß     Ä     ä     Ö     ö     Ü     ü
	[\u00DF\u00C4\u00E4\u00D6\u00F6\u00DC\u00FC];
fragment Digit: [0-9];
fragment PointFloat: Digit* '.' Digit+ | Digit+ '.';
fragment ExponentFloat: (Digit+ | PointFloat) Exponent;
fragment Exponent: [eE] [+-]? Digit+;

fragment A: [Aa];
fragment B: [Bb];
fragment C: [Cc];
fragment D: [Dd];
fragment E: [Ee];
fragment F: [Ff];
fragment G: [Gg];
fragment H: [Hh];
fragment I: [Ii];
fragment J: [Jj];
fragment K: [Kk];
fragment L: [Ll];
fragment M: [Mm];
fragment N: [Nn];
fragment O: [Oo];
fragment P: [Pp];
fragment Q: [Qq];
fragment R: [Rr];
fragment S: [Ss];
fragment T: [Tt];
fragment U: [Uu];
fragment V: [Vv];
fragment W: [Ww];
fragment X: [Xx];
fragment Y: [Yy];
fragment Z: [Zz];


//parser
daedalusFile: mainBlock? EOF;
blockDef:
	(
		functionDef
		| classDef
		| prototypeDef
		| instanceDef
		| namespaceDef
	) Semi;
inlineDef: (constDef | varDecl) Semi;

functionDef:
	Func typeReference nameNode parameterList statementBlock;
constDef:
	Const typeReference (constValueDef | constArrayDef) (
		Comma (constValueDef | constArrayDef)
	)*;
classDef: Class nameNode LeftBrace (varDecl Semi)* RightBrace;
prototypeDef:
	Prototype nameNode LeftParen parentReference RightParen statementBlock;
// like the engine: multiple names and an optional body
instanceDef:
	Instance nameNode (Comma nameNode)* LeftParen parentReference RightParen statementBlock?;
namespaceDef:
	Namespace nameNode LeftBrace contentBlock* RightBrace;
mainBlock: zParserExtenderMetaBlock? contentBlock+;
contentBlock: (blockDef | inlineDef);
varDecl:
	Var typeReference (varValueDecl | varArrayDecl) (
		Comma (varDecl | varValueDecl | varArrayDecl)
	)*;

metaValue: .+?;
zParserExtenderMeta: nameNode Assign metaValue Semi;
zParserExtenderMetaBlock:
	Meta LeftBrace zParserExtenderMeta*? RightBrace Semi;
constArrayDef:
	nameNode LeftBracket arraySize RightBracket constArrayAssignment;
constArrayAssignment:
	Assign LeftBrace expression (Comma expression)* RightBrace;

constValueDef: nameNode constValueAssignment;
constValueAssignment: Assign expression;

varArrayDecl: nameNode LeftBracket arraySize RightBracket;
varValueDecl: nameNode;

parameterList:
	LeftParen (parameterDecl (Comma parameterDecl)*)? RightParen;
parameterDecl:
	Var typeReference nameNode (
		LeftBracket arraySize RightBracket
	)?;
statementBlock:
	LeftBrace (
		(returnStatement | constDef | varDecl | expression) Semi
		| ifBlockStatement Semi?
	)* RightBrace;
elseBlock: Else statementBlock;
elseIfBlock: Else If expression statementBlock;
ifBlock: If expression statementBlock;
ifBlockStatement: ifBlock elseIfBlock* elseBlock?;
returnStatement: Return expression?;

// Precedence (like the engine, highest first):
// unary, * / %, + -, << >>, < > <= >=, == !=, &, |, &&, ||, assignments
expression:
	LeftParen expression RightParen								# bracketExpression
	| op = (Plus | Tilde | Minus | Not) expression				# unaryOperation
	| expression op = (Star | Div | Percent) expression			# multExpression
	| expression op = (Plus | Minus) expression					# addExpression
	| expression op = (Shl | Shr) expression					# bitMoveExpression
	| expression op = (Less | Greater | LessEq | GreaterEq) expression	# compExpression
	| expression op = (Equal | NotEqual) expression				# eqExpression
	| expression op = BitAnd expression							# binAndExpression
	| expression op = BitOr expression							# binOrExpression
	| expression op = And expression							# logAndExpression
	| expression op = Or expression								# logOrExpression
	| <assoc = right> expression op = (
		Assign
		| PlusAssign
		| MinusAssign
		| StarAssign
		| DivAssign
		| AndAssign
		| OrAssign
	) expression												# assignmentExpression
	| IntegerLiteral											# integerLiteralExpression
	| FloatLiteral												# floatLiteralExpression
	| StringLiteral												# stringLiteralExpression
	| Null														# nullLiteralExpression
	| name = (Void | Int | Float | StringKeyword | Func | Instance | Class | Prototype | Meta | Namespace | Identifier) LeftParen (expression (Comma expression)*)? RightParen	# funcCallExpression
	| name = (Void | Int | Float | StringKeyword | Func | Instance | Class | Prototype | Meta | Namespace | Identifier) (LeftBracket arrayIndex RightBracket)? (
		Dot member = (Void | Int | Float | StringKeyword | Func | Instance | Class | Prototype | Meta | Namespace | Identifier) (LeftBracket arrayIndex RightBracket)?
	)?															# referenceExpression;

arrayIndex: IntegerLiteral | nameNode;
arraySize: IntegerLiteral | nameNode;

typeReference: (
		Identifier
		| Void
		| Int
		| Float
		| StringKeyword
		| Func
		| Instance
	);

nameNode: (
		Void
		| Var
		| Int
		| Float
		| StringKeyword
		| Func
		| Instance
		| Class
		| Prototype
		| Null
		| Meta
		| Namespace
		| Identifier
	);

parentReference: Identifier;
