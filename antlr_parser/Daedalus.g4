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

Const: [Cc][Oo][Nn][Ss][Tt];
Var: [Vv][Aa][Rr];
If: [Ii][Ff];
Int: [Ii][Nn][Tt];
Else: [Ee][Ll][Ss][Ee];
Func: [Ff][Uu][Nn][Cc];
StringKeyword: [Ss][Tt][Rr][Ii][Nn][Gg];
Class: [Cc][Ll][Aa][Ss][Ss];
Void: [Vv][Oo][Ii][Dd];
Return: [Rr][Ee][Tt][Uu][Rr][Nn];
Float: [Ff][Ll][Oo][Aa][Tt];
Prototype: [Pp][Rr][Oo][Tt][Oo][Tt][Yy][Pp][Ee];
Instance: [Ii][Nn][Ss][Tt][Aa][Nn][Cc][Ee];
Namespace: [Nn][Aa][Mm][Ee][Ss][Pp][Aa][Cc][Ee];
Null: [Nn][Uu][Ll][Ll];

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
//                             ß     Ä     ä     Ö     ö     Ü     ü
fragment GermanCharacter:
	[\u00DF\u00C4\u00E4\u00D6\u00F6\u00DC\u00FC];
fragment Digit: [0-9];
fragment PointFloat: Digit* '.' Digit+ | Digit+ '.';
fragment ExponentFloat: (Digit+ | PointFloat) Exponent;
fragment Exponent: [eE] [+-]? Digit+;

//parser
daedalusFile: mainBlock EOF;
blockDef:
	(functionDef | classDef | prototypeDef | instanceDef | namespaceDef) Semi?;
inlineDef: (constDef | varDecl | instanceDecl) Semi;

functionDef:
	Func typeReference nameNode parameterList statementBlock;
constDef:
	Const typeReference (constValueDef | constArrayDef) (
		',' (constValueDef | constArrayDef)
	)*;
classDef: Class nameNode LeftBrace (varDecl Semi)*? RightBrace;
prototypeDef:
	Prototype nameNode LeftParen parentReference RightParen statementBlock;
instanceDef:
	Instance nameNode LeftParen parentReference RightParen statementBlock;
instanceDecl:
	Instance nameNode (',' nameNode)*? LeftParen parentReference RightParen;
namespaceDef:
	Namespace nameNode LeftBrace mainBlock RightBrace Semi;
mainBlock:
	(blockDef | inlineDef)*?;
varDecl:
	Var typeReference (varValueDecl | varArrayDecl) (
		',' (varDecl | varValueDecl | varArrayDecl)
	)*;

constArrayDef:
	nameNode LeftBracket arraySize RightBracket constArrayAssignment;
constArrayAssignment:
	Assign LeftBrace (expressionBlock (',' expressionBlock)*?) RightBrace;

constValueDef: nameNode constValueAssignment;
constValueAssignment: Assign expressionBlock;

varArrayDecl: nameNode LeftBracket arraySize RightBracket;
varValueDecl: nameNode;

parameterList:
	LeftParen (parameterDecl (',' parameterDecl)*?)? RightParen;
parameterDecl:
	Var typeReference nameNode (
		LeftBracket arraySize RightBracket
	)?;
statementBlock:
	LeftBrace ((statement Semi) | ( ifBlockStatement Semi?))*? RightBrace;
statement:
	assignment
	| returnStatement
	| constDef
	| varDecl
	| expression;
funcCall:
	nameNode LeftParen (
		funcArgExpression (',' funcArgExpression)*?
	)? RightParen;
assignment: reference assignmentOperator expressionBlock;
ifCondition: expressionBlock;
elseBlock: Else statementBlock;
elseIfBlock: Else If ifCondition statementBlock;
ifBlock: If ifCondition statementBlock;
ifBlockStatement: ifBlock ( elseIfBlock)*? ( elseBlock)?;
returnStatement: Return ( expressionBlock)?;

funcArgExpression:
	expressionBlock; // we use that to detect func call args
expressionBlock:
	expression; // we use that expression to force parser threat expression as a block

expression:
	LeftParen expression RightParen			# bracketExpression
	| unaryOperator expression				# unaryOperation
	| expression multOperator expression	# multExpression
	| expression addOperator expression		# addExpression
	| expression bitMoveOperator expression	# bitMoveExpression
	| expression compOperator expression	# compExpression
	| expression eqOperator expression		# eqExpression
	| expression binAndOperator expression	# binAndExpression
	| expression binOrOperator expression	# binOrExpression
	| expression logAndOperator expression	# logAndExpression
	| expression logOrOperator expression	# logOrExpression
	| value									# valExpression;

arrayIndex: IntegerLiteral | referenceAtom;
arraySize: IntegerLiteral | referenceAtom;

value:
	IntegerLiteral	# integerLiteralValue
	| FloatLiteral	# floatLiteralValue
	| StringLiteral	# stringLiteralValue
	| Null			# nullLiteralValue
	| funcCall		# funcCallValue
	| reference		# referenceValue;

referenceAtom: nameNode ( LeftBracket arrayIndex RightBracket)?;
reference: referenceAtom ( Dot referenceAtom)?;

typeReference: (
		Identifier
		| Void
		| Int
		| Float
		| StringKeyword
		| Func
		| Instance
	);
anyIdentifier: (
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
		| Identifier
	);

nameNode: anyIdentifier;

parentReference: Identifier;

assignmentOperator:
	Assign
	| StarAssign
	| DivAssign
	| PlusAssign
	| MinusAssign
	| AndAssign
	| OrAssign;

unaryOperator: Plus | Tilde | Minus | Not;

addOperator: '+' | '-';
bitMoveOperator: '<<' | '>>';
compOperator: '<' | '>' | '<=' | '>=';
eqOperator: '==' | '!=';
multOperator: '*' | '/' | '%';
binAndOperator: '&';
binOrOperator: '|';
logAndOperator: '&&';
logOrOperator: '||';