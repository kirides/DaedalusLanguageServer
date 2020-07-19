// Code generated from Daedalus.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Daedalus

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseDaedalusListener is a complete listener for a parse tree produced by DaedalusParser.
type BaseDaedalusListener struct{}

var _ DaedalusListener = &BaseDaedalusListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDaedalusListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDaedalusListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDaedalusListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDaedalusListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSymbolSummary is called when production symbolSummary is entered.
func (s *BaseDaedalusListener) EnterSymbolSummary(ctx *SymbolSummaryContext) {}

// ExitSymbolSummary is called when production symbolSummary is exited.
func (s *BaseDaedalusListener) ExitSymbolSummary(ctx *SymbolSummaryContext) {}

// EnterDaedalusFile is called when production daedalusFile is entered.
func (s *BaseDaedalusListener) EnterDaedalusFile(ctx *DaedalusFileContext) {}

// ExitDaedalusFile is called when production daedalusFile is exited.
func (s *BaseDaedalusListener) ExitDaedalusFile(ctx *DaedalusFileContext) {}

// EnterBlockDef is called when production blockDef is entered.
func (s *BaseDaedalusListener) EnterBlockDef(ctx *BlockDefContext) {}

// ExitBlockDef is called when production blockDef is exited.
func (s *BaseDaedalusListener) ExitBlockDef(ctx *BlockDefContext) {}

// EnterInlineDef is called when production inlineDef is entered.
func (s *BaseDaedalusListener) EnterInlineDef(ctx *InlineDefContext) {}

// ExitInlineDef is called when production inlineDef is exited.
func (s *BaseDaedalusListener) ExitInlineDef(ctx *InlineDefContext) {}

// EnterFunctionDef is called when production functionDef is entered.
func (s *BaseDaedalusListener) EnterFunctionDef(ctx *FunctionDefContext) {}

// ExitFunctionDef is called when production functionDef is exited.
func (s *BaseDaedalusListener) ExitFunctionDef(ctx *FunctionDefContext) {}

// EnterConstDef is called when production constDef is entered.
func (s *BaseDaedalusListener) EnterConstDef(ctx *ConstDefContext) {}

// ExitConstDef is called when production constDef is exited.
func (s *BaseDaedalusListener) ExitConstDef(ctx *ConstDefContext) {}

// EnterClassDef is called when production classDef is entered.
func (s *BaseDaedalusListener) EnterClassDef(ctx *ClassDefContext) {}

// ExitClassDef is called when production classDef is exited.
func (s *BaseDaedalusListener) ExitClassDef(ctx *ClassDefContext) {}

// EnterPrototypeDef is called when production prototypeDef is entered.
func (s *BaseDaedalusListener) EnterPrototypeDef(ctx *PrototypeDefContext) {}

// ExitPrototypeDef is called when production prototypeDef is exited.
func (s *BaseDaedalusListener) ExitPrototypeDef(ctx *PrototypeDefContext) {}

// EnterInstanceDef is called when production instanceDef is entered.
func (s *BaseDaedalusListener) EnterInstanceDef(ctx *InstanceDefContext) {}

// ExitInstanceDef is called when production instanceDef is exited.
func (s *BaseDaedalusListener) ExitInstanceDef(ctx *InstanceDefContext) {}

// EnterInstanceDecl is called when production instanceDecl is entered.
func (s *BaseDaedalusListener) EnterInstanceDecl(ctx *InstanceDeclContext) {}

// ExitInstanceDecl is called when production instanceDecl is exited.
func (s *BaseDaedalusListener) ExitInstanceDecl(ctx *InstanceDeclContext) {}

// EnterVarDecl is called when production varDecl is entered.
func (s *BaseDaedalusListener) EnterVarDecl(ctx *VarDeclContext) {}

// ExitVarDecl is called when production varDecl is exited.
func (s *BaseDaedalusListener) ExitVarDecl(ctx *VarDeclContext) {}

// EnterConstArrayDef is called when production constArrayDef is entered.
func (s *BaseDaedalusListener) EnterConstArrayDef(ctx *ConstArrayDefContext) {}

// ExitConstArrayDef is called when production constArrayDef is exited.
func (s *BaseDaedalusListener) ExitConstArrayDef(ctx *ConstArrayDefContext) {}

// EnterConstArrayAssignment is called when production constArrayAssignment is entered.
func (s *BaseDaedalusListener) EnterConstArrayAssignment(ctx *ConstArrayAssignmentContext) {}

// ExitConstArrayAssignment is called when production constArrayAssignment is exited.
func (s *BaseDaedalusListener) ExitConstArrayAssignment(ctx *ConstArrayAssignmentContext) {}

// EnterConstValueDef is called when production constValueDef is entered.
func (s *BaseDaedalusListener) EnterConstValueDef(ctx *ConstValueDefContext) {}

// ExitConstValueDef is called when production constValueDef is exited.
func (s *BaseDaedalusListener) ExitConstValueDef(ctx *ConstValueDefContext) {}

// EnterConstValueAssignment is called when production constValueAssignment is entered.
func (s *BaseDaedalusListener) EnterConstValueAssignment(ctx *ConstValueAssignmentContext) {}

// ExitConstValueAssignment is called when production constValueAssignment is exited.
func (s *BaseDaedalusListener) ExitConstValueAssignment(ctx *ConstValueAssignmentContext) {}

// EnterVarArrayDecl is called when production varArrayDecl is entered.
func (s *BaseDaedalusListener) EnterVarArrayDecl(ctx *VarArrayDeclContext) {}

// ExitVarArrayDecl is called when production varArrayDecl is exited.
func (s *BaseDaedalusListener) ExitVarArrayDecl(ctx *VarArrayDeclContext) {}

// EnterVarValueDecl is called when production varValueDecl is entered.
func (s *BaseDaedalusListener) EnterVarValueDecl(ctx *VarValueDeclContext) {}

// ExitVarValueDecl is called when production varValueDecl is exited.
func (s *BaseDaedalusListener) ExitVarValueDecl(ctx *VarValueDeclContext) {}

// EnterParameterList is called when production parameterList is entered.
func (s *BaseDaedalusListener) EnterParameterList(ctx *ParameterListContext) {}

// ExitParameterList is called when production parameterList is exited.
func (s *BaseDaedalusListener) ExitParameterList(ctx *ParameterListContext) {}

// EnterParameterDecl is called when production parameterDecl is entered.
func (s *BaseDaedalusListener) EnterParameterDecl(ctx *ParameterDeclContext) {}

// ExitParameterDecl is called when production parameterDecl is exited.
func (s *BaseDaedalusListener) ExitParameterDecl(ctx *ParameterDeclContext) {}

// EnterStatementBlock is called when production statementBlock is entered.
func (s *BaseDaedalusListener) EnterStatementBlock(ctx *StatementBlockContext) {}

// ExitStatementBlock is called when production statementBlock is exited.
func (s *BaseDaedalusListener) ExitStatementBlock(ctx *StatementBlockContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseDaedalusListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseDaedalusListener) ExitStatement(ctx *StatementContext) {}

// EnterFuncCall is called when production funcCall is entered.
func (s *BaseDaedalusListener) EnterFuncCall(ctx *FuncCallContext) {}

// ExitFuncCall is called when production funcCall is exited.
func (s *BaseDaedalusListener) ExitFuncCall(ctx *FuncCallContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseDaedalusListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseDaedalusListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterIfCondition is called when production ifCondition is entered.
func (s *BaseDaedalusListener) EnterIfCondition(ctx *IfConditionContext) {}

// ExitIfCondition is called when production ifCondition is exited.
func (s *BaseDaedalusListener) ExitIfCondition(ctx *IfConditionContext) {}

// EnterElseBlock is called when production elseBlock is entered.
func (s *BaseDaedalusListener) EnterElseBlock(ctx *ElseBlockContext) {}

// ExitElseBlock is called when production elseBlock is exited.
func (s *BaseDaedalusListener) ExitElseBlock(ctx *ElseBlockContext) {}

// EnterElseIfBlock is called when production elseIfBlock is entered.
func (s *BaseDaedalusListener) EnterElseIfBlock(ctx *ElseIfBlockContext) {}

// ExitElseIfBlock is called when production elseIfBlock is exited.
func (s *BaseDaedalusListener) ExitElseIfBlock(ctx *ElseIfBlockContext) {}

// EnterIfBlock is called when production ifBlock is entered.
func (s *BaseDaedalusListener) EnterIfBlock(ctx *IfBlockContext) {}

// ExitIfBlock is called when production ifBlock is exited.
func (s *BaseDaedalusListener) ExitIfBlock(ctx *IfBlockContext) {}

// EnterIfBlockStatement is called when production ifBlockStatement is entered.
func (s *BaseDaedalusListener) EnterIfBlockStatement(ctx *IfBlockStatementContext) {}

// ExitIfBlockStatement is called when production ifBlockStatement is exited.
func (s *BaseDaedalusListener) ExitIfBlockStatement(ctx *IfBlockStatementContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseDaedalusListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseDaedalusListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterFuncArgExpression is called when production funcArgExpression is entered.
func (s *BaseDaedalusListener) EnterFuncArgExpression(ctx *FuncArgExpressionContext) {}

// ExitFuncArgExpression is called when production funcArgExpression is exited.
func (s *BaseDaedalusListener) ExitFuncArgExpression(ctx *FuncArgExpressionContext) {}

// EnterExpressionBlock is called when production expressionBlock is entered.
func (s *BaseDaedalusListener) EnterExpressionBlock(ctx *ExpressionBlockContext) {}

// ExitExpressionBlock is called when production expressionBlock is exited.
func (s *BaseDaedalusListener) ExitExpressionBlock(ctx *ExpressionBlockContext) {}

// EnterBitMoveExpression is called when production bitMoveExpression is entered.
func (s *BaseDaedalusListener) EnterBitMoveExpression(ctx *BitMoveExpressionContext) {}

// ExitBitMoveExpression is called when production bitMoveExpression is exited.
func (s *BaseDaedalusListener) ExitBitMoveExpression(ctx *BitMoveExpressionContext) {}

// EnterOneArgExpression is called when production oneArgExpression is entered.
func (s *BaseDaedalusListener) EnterOneArgExpression(ctx *OneArgExpressionContext) {}

// ExitOneArgExpression is called when production oneArgExpression is exited.
func (s *BaseDaedalusListener) ExitOneArgExpression(ctx *OneArgExpressionContext) {}

// EnterEqExpression is called when production eqExpression is entered.
func (s *BaseDaedalusListener) EnterEqExpression(ctx *EqExpressionContext) {}

// ExitEqExpression is called when production eqExpression is exited.
func (s *BaseDaedalusListener) ExitEqExpression(ctx *EqExpressionContext) {}

// EnterValExpression is called when production valExpression is entered.
func (s *BaseDaedalusListener) EnterValExpression(ctx *ValExpressionContext) {}

// ExitValExpression is called when production valExpression is exited.
func (s *BaseDaedalusListener) ExitValExpression(ctx *ValExpressionContext) {}

// EnterAddExpression is called when production addExpression is entered.
func (s *BaseDaedalusListener) EnterAddExpression(ctx *AddExpressionContext) {}

// ExitAddExpression is called when production addExpression is exited.
func (s *BaseDaedalusListener) ExitAddExpression(ctx *AddExpressionContext) {}

// EnterCompExpression is called when production compExpression is entered.
func (s *BaseDaedalusListener) EnterCompExpression(ctx *CompExpressionContext) {}

// ExitCompExpression is called when production compExpression is exited.
func (s *BaseDaedalusListener) ExitCompExpression(ctx *CompExpressionContext) {}

// EnterLogOrExpression is called when production logOrExpression is entered.
func (s *BaseDaedalusListener) EnterLogOrExpression(ctx *LogOrExpressionContext) {}

// ExitLogOrExpression is called when production logOrExpression is exited.
func (s *BaseDaedalusListener) ExitLogOrExpression(ctx *LogOrExpressionContext) {}

// EnterBinAndExpression is called when production binAndExpression is entered.
func (s *BaseDaedalusListener) EnterBinAndExpression(ctx *BinAndExpressionContext) {}

// ExitBinAndExpression is called when production binAndExpression is exited.
func (s *BaseDaedalusListener) ExitBinAndExpression(ctx *BinAndExpressionContext) {}

// EnterBinOrExpression is called when production binOrExpression is entered.
func (s *BaseDaedalusListener) EnterBinOrExpression(ctx *BinOrExpressionContext) {}

// ExitBinOrExpression is called when production binOrExpression is exited.
func (s *BaseDaedalusListener) ExitBinOrExpression(ctx *BinOrExpressionContext) {}

// EnterMultExpression is called when production multExpression is entered.
func (s *BaseDaedalusListener) EnterMultExpression(ctx *MultExpressionContext) {}

// ExitMultExpression is called when production multExpression is exited.
func (s *BaseDaedalusListener) ExitMultExpression(ctx *MultExpressionContext) {}

// EnterBracketExpression is called when production bracketExpression is entered.
func (s *BaseDaedalusListener) EnterBracketExpression(ctx *BracketExpressionContext) {}

// ExitBracketExpression is called when production bracketExpression is exited.
func (s *BaseDaedalusListener) ExitBracketExpression(ctx *BracketExpressionContext) {}

// EnterLogAndExpression is called when production logAndExpression is entered.
func (s *BaseDaedalusListener) EnterLogAndExpression(ctx *LogAndExpressionContext) {}

// ExitLogAndExpression is called when production logAndExpression is exited.
func (s *BaseDaedalusListener) ExitLogAndExpression(ctx *LogAndExpressionContext) {}

// EnterArrayIndex is called when production arrayIndex is entered.
func (s *BaseDaedalusListener) EnterArrayIndex(ctx *ArrayIndexContext) {}

// ExitArrayIndex is called when production arrayIndex is exited.
func (s *BaseDaedalusListener) ExitArrayIndex(ctx *ArrayIndexContext) {}

// EnterArraySize is called when production arraySize is entered.
func (s *BaseDaedalusListener) EnterArraySize(ctx *ArraySizeContext) {}

// ExitArraySize is called when production arraySize is exited.
func (s *BaseDaedalusListener) ExitArraySize(ctx *ArraySizeContext) {}

// EnterIntegerLiteralValue is called when production integerLiteralValue is entered.
func (s *BaseDaedalusListener) EnterIntegerLiteralValue(ctx *IntegerLiteralValueContext) {}

// ExitIntegerLiteralValue is called when production integerLiteralValue is exited.
func (s *BaseDaedalusListener) ExitIntegerLiteralValue(ctx *IntegerLiteralValueContext) {}

// EnterFloatLiteralValue is called when production floatLiteralValue is entered.
func (s *BaseDaedalusListener) EnterFloatLiteralValue(ctx *FloatLiteralValueContext) {}

// ExitFloatLiteralValue is called when production floatLiteralValue is exited.
func (s *BaseDaedalusListener) ExitFloatLiteralValue(ctx *FloatLiteralValueContext) {}

// EnterStringLiteralValue is called when production stringLiteralValue is entered.
func (s *BaseDaedalusListener) EnterStringLiteralValue(ctx *StringLiteralValueContext) {}

// ExitStringLiteralValue is called when production stringLiteralValue is exited.
func (s *BaseDaedalusListener) ExitStringLiteralValue(ctx *StringLiteralValueContext) {}

// EnterNullLiteralValue is called when production nullLiteralValue is entered.
func (s *BaseDaedalusListener) EnterNullLiteralValue(ctx *NullLiteralValueContext) {}

// ExitNullLiteralValue is called when production nullLiteralValue is exited.
func (s *BaseDaedalusListener) ExitNullLiteralValue(ctx *NullLiteralValueContext) {}

// EnterNoFuncLiteralValue is called when production noFuncLiteralValue is entered.
func (s *BaseDaedalusListener) EnterNoFuncLiteralValue(ctx *NoFuncLiteralValueContext) {}

// ExitNoFuncLiteralValue is called when production noFuncLiteralValue is exited.
func (s *BaseDaedalusListener) ExitNoFuncLiteralValue(ctx *NoFuncLiteralValueContext) {}

// EnterFuncCallValue is called when production funcCallValue is entered.
func (s *BaseDaedalusListener) EnterFuncCallValue(ctx *FuncCallValueContext) {}

// ExitFuncCallValue is called when production funcCallValue is exited.
func (s *BaseDaedalusListener) ExitFuncCallValue(ctx *FuncCallValueContext) {}

// EnterReferenceValue is called when production referenceValue is entered.
func (s *BaseDaedalusListener) EnterReferenceValue(ctx *ReferenceValueContext) {}

// ExitReferenceValue is called when production referenceValue is exited.
func (s *BaseDaedalusListener) ExitReferenceValue(ctx *ReferenceValueContext) {}

// EnterAnyIdentifierValue is called when production anyIdentifierValue is entered.
func (s *BaseDaedalusListener) EnterAnyIdentifierValue(ctx *AnyIdentifierValueContext) {}

// ExitAnyIdentifierValue is called when production anyIdentifierValue is exited.
func (s *BaseDaedalusListener) ExitAnyIdentifierValue(ctx *AnyIdentifierValueContext) {}

// EnterReferenceAtom is called when production referenceAtom is entered.
func (s *BaseDaedalusListener) EnterReferenceAtom(ctx *ReferenceAtomContext) {}

// ExitReferenceAtom is called when production referenceAtom is exited.
func (s *BaseDaedalusListener) ExitReferenceAtom(ctx *ReferenceAtomContext) {}

// EnterReference is called when production reference is entered.
func (s *BaseDaedalusListener) EnterReference(ctx *ReferenceContext) {}

// ExitReference is called when production reference is exited.
func (s *BaseDaedalusListener) ExitReference(ctx *ReferenceContext) {}

// EnterTypeReference is called when production typeReference is entered.
func (s *BaseDaedalusListener) EnterTypeReference(ctx *TypeReferenceContext) {}

// ExitTypeReference is called when production typeReference is exited.
func (s *BaseDaedalusListener) ExitTypeReference(ctx *TypeReferenceContext) {}

// EnterAnyIdentifier is called when production anyIdentifier is entered.
func (s *BaseDaedalusListener) EnterAnyIdentifier(ctx *AnyIdentifierContext) {}

// ExitAnyIdentifier is called when production anyIdentifier is exited.
func (s *BaseDaedalusListener) ExitAnyIdentifier(ctx *AnyIdentifierContext) {}

// EnterNameNode is called when production nameNode is entered.
func (s *BaseDaedalusListener) EnterNameNode(ctx *NameNodeContext) {}

// ExitNameNode is called when production nameNode is exited.
func (s *BaseDaedalusListener) ExitNameNode(ctx *NameNodeContext) {}

// EnterParentReference is called when production parentReference is entered.
func (s *BaseDaedalusListener) EnterParentReference(ctx *ParentReferenceContext) {}

// ExitParentReference is called when production parentReference is exited.
func (s *BaseDaedalusListener) ExitParentReference(ctx *ParentReferenceContext) {}

// EnterAssignmentOperator is called when production assignmentOperator is entered.
func (s *BaseDaedalusListener) EnterAssignmentOperator(ctx *AssignmentOperatorContext) {}

// ExitAssignmentOperator is called when production assignmentOperator is exited.
func (s *BaseDaedalusListener) ExitAssignmentOperator(ctx *AssignmentOperatorContext) {}

// EnterAddOperator is called when production addOperator is entered.
func (s *BaseDaedalusListener) EnterAddOperator(ctx *AddOperatorContext) {}

// ExitAddOperator is called when production addOperator is exited.
func (s *BaseDaedalusListener) ExitAddOperator(ctx *AddOperatorContext) {}

// EnterBitMoveOperator is called when production bitMoveOperator is entered.
func (s *BaseDaedalusListener) EnterBitMoveOperator(ctx *BitMoveOperatorContext) {}

// ExitBitMoveOperator is called when production bitMoveOperator is exited.
func (s *BaseDaedalusListener) ExitBitMoveOperator(ctx *BitMoveOperatorContext) {}

// EnterCompOperator is called when production compOperator is entered.
func (s *BaseDaedalusListener) EnterCompOperator(ctx *CompOperatorContext) {}

// ExitCompOperator is called when production compOperator is exited.
func (s *BaseDaedalusListener) ExitCompOperator(ctx *CompOperatorContext) {}

// EnterEqOperator is called when production eqOperator is entered.
func (s *BaseDaedalusListener) EnterEqOperator(ctx *EqOperatorContext) {}

// ExitEqOperator is called when production eqOperator is exited.
func (s *BaseDaedalusListener) ExitEqOperator(ctx *EqOperatorContext) {}

// EnterOneArgOperator is called when production oneArgOperator is entered.
func (s *BaseDaedalusListener) EnterOneArgOperator(ctx *OneArgOperatorContext) {}

// ExitOneArgOperator is called when production oneArgOperator is exited.
func (s *BaseDaedalusListener) ExitOneArgOperator(ctx *OneArgOperatorContext) {}

// EnterMultOperator is called when production multOperator is entered.
func (s *BaseDaedalusListener) EnterMultOperator(ctx *MultOperatorContext) {}

// ExitMultOperator is called when production multOperator is exited.
func (s *BaseDaedalusListener) ExitMultOperator(ctx *MultOperatorContext) {}

// EnterBinAndOperator is called when production binAndOperator is entered.
func (s *BaseDaedalusListener) EnterBinAndOperator(ctx *BinAndOperatorContext) {}

// ExitBinAndOperator is called when production binAndOperator is exited.
func (s *BaseDaedalusListener) ExitBinAndOperator(ctx *BinAndOperatorContext) {}

// EnterBinOrOperator is called when production binOrOperator is entered.
func (s *BaseDaedalusListener) EnterBinOrOperator(ctx *BinOrOperatorContext) {}

// ExitBinOrOperator is called when production binOrOperator is exited.
func (s *BaseDaedalusListener) ExitBinOrOperator(ctx *BinOrOperatorContext) {}

// EnterLogAndOperator is called when production logAndOperator is entered.
func (s *BaseDaedalusListener) EnterLogAndOperator(ctx *LogAndOperatorContext) {}

// ExitLogAndOperator is called when production logAndOperator is exited.
func (s *BaseDaedalusListener) ExitLogAndOperator(ctx *LogAndOperatorContext) {}

// EnterLogOrOperator is called when production logOrOperator is entered.
func (s *BaseDaedalusListener) EnterLogOrOperator(ctx *LogOrOperatorContext) {}

// ExitLogOrOperator is called when production logOrOperator is exited.
func (s *BaseDaedalusListener) ExitLogOrOperator(ctx *LogOrOperatorContext) {}
