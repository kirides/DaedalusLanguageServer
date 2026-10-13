// Code generated from Daedalus.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Daedalus

import "github.com/antlr4-go/antlr/v4"

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

// EnterNamespaceDef is called when production namespaceDef is entered.
func (s *BaseDaedalusListener) EnterNamespaceDef(ctx *NamespaceDefContext) {}

// ExitNamespaceDef is called when production namespaceDef is exited.
func (s *BaseDaedalusListener) ExitNamespaceDef(ctx *NamespaceDefContext) {}

// EnterMainBlock is called when production mainBlock is entered.
func (s *BaseDaedalusListener) EnterMainBlock(ctx *MainBlockContext) {}

// ExitMainBlock is called when production mainBlock is exited.
func (s *BaseDaedalusListener) ExitMainBlock(ctx *MainBlockContext) {}

// EnterContentBlock is called when production contentBlock is entered.
func (s *BaseDaedalusListener) EnterContentBlock(ctx *ContentBlockContext) {}

// ExitContentBlock is called when production contentBlock is exited.
func (s *BaseDaedalusListener) ExitContentBlock(ctx *ContentBlockContext) {}

// EnterVarDecl is called when production varDecl is entered.
func (s *BaseDaedalusListener) EnterVarDecl(ctx *VarDeclContext) {}

// ExitVarDecl is called when production varDecl is exited.
func (s *BaseDaedalusListener) ExitVarDecl(ctx *VarDeclContext) {}

// EnterMetaValue is called when production metaValue is entered.
func (s *BaseDaedalusListener) EnterMetaValue(ctx *MetaValueContext) {}

// ExitMetaValue is called when production metaValue is exited.
func (s *BaseDaedalusListener) ExitMetaValue(ctx *MetaValueContext) {}

// EnterZParserExtenderMeta is called when production zParserExtenderMeta is entered.
func (s *BaseDaedalusListener) EnterZParserExtenderMeta(ctx *ZParserExtenderMetaContext) {}

// ExitZParserExtenderMeta is called when production zParserExtenderMeta is exited.
func (s *BaseDaedalusListener) ExitZParserExtenderMeta(ctx *ZParserExtenderMetaContext) {}

// EnterZParserExtenderMetaBlock is called when production zParserExtenderMetaBlock is entered.
func (s *BaseDaedalusListener) EnterZParserExtenderMetaBlock(ctx *ZParserExtenderMetaBlockContext) {}

// ExitZParserExtenderMetaBlock is called when production zParserExtenderMetaBlock is exited.
func (s *BaseDaedalusListener) ExitZParserExtenderMetaBlock(ctx *ZParserExtenderMetaBlockContext) {}

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

// EnterBitMoveExpression is called when production bitMoveExpression is entered.
func (s *BaseDaedalusListener) EnterBitMoveExpression(ctx *BitMoveExpressionContext) {}

// ExitBitMoveExpression is called when production bitMoveExpression is exited.
func (s *BaseDaedalusListener) ExitBitMoveExpression(ctx *BitMoveExpressionContext) {}

// EnterReferenceExpression is called when production referenceExpression is entered.
func (s *BaseDaedalusListener) EnterReferenceExpression(ctx *ReferenceExpressionContext) {}

// ExitReferenceExpression is called when production referenceExpression is exited.
func (s *BaseDaedalusListener) ExitReferenceExpression(ctx *ReferenceExpressionContext) {}

// EnterAssignmentExpression is called when production assignmentExpression is entered.
func (s *BaseDaedalusListener) EnterAssignmentExpression(ctx *AssignmentExpressionContext) {}

// ExitAssignmentExpression is called when production assignmentExpression is exited.
func (s *BaseDaedalusListener) ExitAssignmentExpression(ctx *AssignmentExpressionContext) {}

// EnterBracketExpression is called when production bracketExpression is entered.
func (s *BaseDaedalusListener) EnterBracketExpression(ctx *BracketExpressionContext) {}

// ExitBracketExpression is called when production bracketExpression is exited.
func (s *BaseDaedalusListener) ExitBracketExpression(ctx *BracketExpressionContext) {}

// EnterUnaryOperation is called when production unaryOperation is entered.
func (s *BaseDaedalusListener) EnterUnaryOperation(ctx *UnaryOperationContext) {}

// ExitUnaryOperation is called when production unaryOperation is exited.
func (s *BaseDaedalusListener) ExitUnaryOperation(ctx *UnaryOperationContext) {}

// EnterFloatLiteralExpression is called when production floatLiteralExpression is entered.
func (s *BaseDaedalusListener) EnterFloatLiteralExpression(ctx *FloatLiteralExpressionContext) {}

// ExitFloatLiteralExpression is called when production floatLiteralExpression is exited.
func (s *BaseDaedalusListener) ExitFloatLiteralExpression(ctx *FloatLiteralExpressionContext) {}

// EnterLogAndExpression is called when production logAndExpression is entered.
func (s *BaseDaedalusListener) EnterLogAndExpression(ctx *LogAndExpressionContext) {}

// ExitLogAndExpression is called when production logAndExpression is exited.
func (s *BaseDaedalusListener) ExitLogAndExpression(ctx *LogAndExpressionContext) {}

// EnterFuncCallExpression is called when production funcCallExpression is entered.
func (s *BaseDaedalusListener) EnterFuncCallExpression(ctx *FuncCallExpressionContext) {}

// ExitFuncCallExpression is called when production funcCallExpression is exited.
func (s *BaseDaedalusListener) ExitFuncCallExpression(ctx *FuncCallExpressionContext) {}

// EnterStringLiteralExpression is called when production stringLiteralExpression is entered.
func (s *BaseDaedalusListener) EnterStringLiteralExpression(ctx *StringLiteralExpressionContext) {}

// ExitStringLiteralExpression is called when production stringLiteralExpression is exited.
func (s *BaseDaedalusListener) ExitStringLiteralExpression(ctx *StringLiteralExpressionContext) {}

// EnterEqExpression is called when production eqExpression is entered.
func (s *BaseDaedalusListener) EnterEqExpression(ctx *EqExpressionContext) {}

// ExitEqExpression is called when production eqExpression is exited.
func (s *BaseDaedalusListener) ExitEqExpression(ctx *EqExpressionContext) {}

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

// EnterIntegerLiteralExpression is called when production integerLiteralExpression is entered.
func (s *BaseDaedalusListener) EnterIntegerLiteralExpression(ctx *IntegerLiteralExpressionContext) {}

// ExitIntegerLiteralExpression is called when production integerLiteralExpression is exited.
func (s *BaseDaedalusListener) ExitIntegerLiteralExpression(ctx *IntegerLiteralExpressionContext) {}

// EnterBinOrExpression is called when production binOrExpression is entered.
func (s *BaseDaedalusListener) EnterBinOrExpression(ctx *BinOrExpressionContext) {}

// ExitBinOrExpression is called when production binOrExpression is exited.
func (s *BaseDaedalusListener) ExitBinOrExpression(ctx *BinOrExpressionContext) {}

// EnterMultExpression is called when production multExpression is entered.
func (s *BaseDaedalusListener) EnterMultExpression(ctx *MultExpressionContext) {}

// ExitMultExpression is called when production multExpression is exited.
func (s *BaseDaedalusListener) ExitMultExpression(ctx *MultExpressionContext) {}

// EnterNullLiteralExpression is called when production nullLiteralExpression is entered.
func (s *BaseDaedalusListener) EnterNullLiteralExpression(ctx *NullLiteralExpressionContext) {}

// ExitNullLiteralExpression is called when production nullLiteralExpression is exited.
func (s *BaseDaedalusListener) ExitNullLiteralExpression(ctx *NullLiteralExpressionContext) {}

// EnterArrayIndex is called when production arrayIndex is entered.
func (s *BaseDaedalusListener) EnterArrayIndex(ctx *ArrayIndexContext) {}

// ExitArrayIndex is called when production arrayIndex is exited.
func (s *BaseDaedalusListener) ExitArrayIndex(ctx *ArrayIndexContext) {}

// EnterArraySize is called when production arraySize is entered.
func (s *BaseDaedalusListener) EnterArraySize(ctx *ArraySizeContext) {}

// ExitArraySize is called when production arraySize is exited.
func (s *BaseDaedalusListener) ExitArraySize(ctx *ArraySizeContext) {}

// EnterTypeReference is called when production typeReference is entered.
func (s *BaseDaedalusListener) EnterTypeReference(ctx *TypeReferenceContext) {}

// ExitTypeReference is called when production typeReference is exited.
func (s *BaseDaedalusListener) ExitTypeReference(ctx *TypeReferenceContext) {}

// EnterNameNode is called when production nameNode is entered.
func (s *BaseDaedalusListener) EnterNameNode(ctx *NameNodeContext) {}

// ExitNameNode is called when production nameNode is exited.
func (s *BaseDaedalusListener) ExitNameNode(ctx *NameNodeContext) {}

// EnterParentReference is called when production parentReference is entered.
func (s *BaseDaedalusListener) EnterParentReference(ctx *ParentReferenceContext) {}

// ExitParentReference is called when production parentReference is exited.
func (s *BaseDaedalusListener) ExitParentReference(ctx *ParentReferenceContext) {}
