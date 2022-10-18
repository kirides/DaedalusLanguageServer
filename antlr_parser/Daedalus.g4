grammar Daedalus;
// lexer
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
Null: [Nn][Uu][Ll][Ll];
NoFunc: [Nn][Oo][Ff][Uu][Nn][Cc];

Identifier: (IdStart | IdStartNumeric) IdContinue*;
IntegerLiteral: Digit+;
FloatLiteral: PointFloat | ExponentFloat;
StringLiteral: '"' (~["\\\r\n] | '\\' (. | EOF))* '"';

Whitespace: [ \t]+ -> skip;
TooManyCommentlines: '////' ~[\r\n]* -> skip;
SummaryComment: DocCommentStart ~[\r\n]+;
Newline: [\r\n] -> skip;
BlockComment: '/*' .*? '*/' -> skip;
LineComment: '//' ~[\r\n]* -> skip;

// fragments
fragment IdStart: GermanCharacter | [a-zA-Z_];
fragment IdStartNumeric: Digit+ IdStart;
fragment IdContinue: IdStart | IdSpecial | Digit;
fragment IdSpecial: [@^];
//                             ß     Ä     ä     Ö     ö     Ü     ü
fragment GermanCharacter:
	[\u00DF\u00C4\u00E4\u00D6\u00F6\u00DC\u00FC];
fragment Digit: [0-9];
fragment PointFloat: Digit* '.' Digit+ | Digit+ '.';
fragment ExponentFloat: (Digit+ | PointFloat) Exponent;
fragment Exponent: [eE] [+-]? Digit+;
fragment DocCommentStart: '///';

//parser
symbolSummary: SummaryComment+;

daedalusFile: (blockDef | inlineDef | symbolSummary)*? EOF;
blockDef:
	symbolSummary* (
		functionDef
		| classDef
		| prototypeDef
		| instanceDef
	) ';'?;
inlineDef:
	symbolSummary* (constDef | varDecl | instanceDecl) ';';

functionDef:
	Func typeReference nameNode parameterList statementBlock;
constDef:
	Const typeReference (constValueDef | constArrayDef) (
		',' (constValueDef | constArrayDef)
	)*;
classDef:
	Class nameNode '{' (varDecl ';' | symbolSummary)*? '}';
prototypeDef:
	Prototype nameNode '(' parentReference ')' statementBlock;
instanceDef:
	Instance nameNode '(' parentReference ')' statementBlock;
instanceDecl:
	Instance nameNode (',' nameNode)*? '(' parentReference ')';
varDecl:
	Var typeReference (varValueDecl | varArrayDecl) (
		',' (varDecl | varValueDecl | varArrayDecl)
	)*;

constArrayDef: nameNode '[' arraySize ']' constArrayAssignment;
constArrayAssignment:
	'=' '{' (expressionBlock (',' expressionBlock)*?) '}';

constValueDef: nameNode constValueAssignment;
constValueAssignment: '=' expressionBlock;

varArrayDecl: nameNode '[' arraySize ']';
varValueDecl: nameNode;

parameterList: '(' (parameterDecl (',' parameterDecl)*?)? ')';
parameterDecl: Var typeReference nameNode ('[' arraySize ']')?;
statementBlock:
	symbolSummary* '{' (
		(statement ';')
		| ( ifBlockStatement ( ';')?)
		| symbolSummary
	)*? '}' symbolSummary*;
statement:
	assignment
	| returnStatement
	| constDef
	| varDecl
	| funcCall
	| expression;
funcCall:
	nameNode '(' (funcArgExpression ( ',' funcArgExpression)*?)? ')';
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
	'(' expression ')'						# bracketExpression
	| oneArgOperator expression				# oneArgExpression
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
	| NoFunc		# noFuncLiteralValue
	| funcCall		# funcCallValue
	| reference		# referenceValue
	| nameNode		# anyIdentifierValue;

referenceAtom: nameNode ( '[' arrayIndex ']')?;
reference: referenceAtom ( '.' referenceAtom)?;

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

assignmentOperator: '=' | '+=' | '-=' | '*=' | '/=';
addOperator: '+' | '-';
bitMoveOperator: '<<' | '>>';
compOperator: '<' | '>' | '<=' | '>=';
eqOperator: '==' | '!=';
oneArgOperator: '-' | '!' | '~' | '+';
multOperator: '*' | '/' | '%';
binAndOperator: '&';
binOrOperator: '|';
logAndOperator: '&&';
logOrOperator: '||';