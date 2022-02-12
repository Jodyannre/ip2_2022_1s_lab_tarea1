// Code generated from Nparser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Nparser

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 46, 34, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 5, 3, 16, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 22, 10, 3, 12, 3, 14,
	3, 25, 11, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 32, 10, 4, 3, 4, 2, 3,
	4, 5, 2, 4, 6, 2, 2, 2, 37, 2, 8, 3, 2, 2, 2, 4, 15, 3, 2, 2, 2, 6, 31,
	3, 2, 2, 2, 8, 9, 5, 4, 3, 2, 9, 3, 3, 2, 2, 2, 10, 11, 8, 3, 1, 2, 11,
	12, 7, 35, 2, 2, 12, 16, 5, 4, 3, 6, 13, 16, 7, 22, 2, 2, 14, 16, 7, 23,
	2, 2, 15, 10, 3, 2, 2, 2, 15, 13, 3, 2, 2, 2, 15, 14, 3, 2, 2, 2, 16, 23,
	3, 2, 2, 2, 17, 18, 12, 5, 2, 2, 18, 19, 5, 6, 4, 2, 19, 20, 5, 4, 3, 6,
	20, 22, 3, 2, 2, 2, 21, 17, 3, 2, 2, 2, 22, 25, 3, 2, 2, 2, 23, 21, 3,
	2, 2, 2, 23, 24, 3, 2, 2, 2, 24, 5, 3, 2, 2, 2, 25, 23, 3, 2, 2, 2, 26,
	32, 7, 33, 2, 2, 27, 32, 7, 34, 2, 2, 28, 32, 7, 32, 2, 2, 29, 32, 7, 36,
	2, 2, 30, 32, 7, 35, 2, 2, 31, 26, 3, 2, 2, 2, 31, 27, 3, 2, 2, 2, 31,
	28, 3, 2, 2, 2, 31, 29, 3, 2, 2, 2, 31, 30, 3, 2, 2, 2, 32, 7, 3, 2, 2,
	2, 5, 15, 23, 31,
}
var literalNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "':'", "'.'", "';'", "'>='", "'>'", "'<'", "'='", "'%'",
	"'*'", "'/'", "'-'", "'+'", "'('", "')'", "'{'", "'}'", "'['", "']'",
}
var symbolicNames = []string{
	"", "SENTENCIA", "AND", "OR", "IF", "ENTONCES", "CONSOLA", "STRING", "INTEGER",
	"REAL", "BOOLEAN", "TRUE", "FALSE", "PUBLICO", "PRIVADO", "MAIN", "PRINCIPAL",
	"METODO", "CLASE", "DECLARAR", "NUMERO", "DECIMAL", "ID", "DOSPUNTOS",
	"PUNTO", "PUNTOCOMA", "MAYOR_I", "MAYOR", "MENOR", "IGUAL", "MODULO", "MULTIPLICACION",
	"DIVISION", "RESTA", "SUMA", "PAR_IZQ", "PAR_DER", "LLAVE_IZQ", "LLAVE_DER",
	"CORCHETE_IZQ", "CORCHETE_DER", "CADENA", "WHITESPACE", "COMMENT", "LINE_COMMENT",
}

var ruleNames = []string{
	"inicio", "expresion", "operador_aritmetico",
}

type Nparser struct {
	*antlr.BaseParser
}

// NewNparser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *Nparser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewNparser(input antlr.TokenStream) *Nparser {
	this := new(Nparser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Nparser.g4"

	return this
}

// Nparser tokens.
const (
	NparserEOF            = antlr.TokenEOF
	NparserSENTENCIA      = 1
	NparserAND            = 2
	NparserOR             = 3
	NparserIF             = 4
	NparserENTONCES       = 5
	NparserCONSOLA        = 6
	NparserSTRING         = 7
	NparserINTEGER        = 8
	NparserREAL           = 9
	NparserBOOLEAN        = 10
	NparserTRUE           = 11
	NparserFALSE          = 12
	NparserPUBLICO        = 13
	NparserPRIVADO        = 14
	NparserMAIN           = 15
	NparserPRINCIPAL      = 16
	NparserMETODO         = 17
	NparserCLASE          = 18
	NparserDECLARAR       = 19
	NparserNUMERO         = 20
	NparserDECIMAL        = 21
	NparserID             = 22
	NparserDOSPUNTOS      = 23
	NparserPUNTO          = 24
	NparserPUNTOCOMA      = 25
	NparserMAYOR_I        = 26
	NparserMAYOR          = 27
	NparserMENOR          = 28
	NparserIGUAL          = 29
	NparserMODULO         = 30
	NparserMULTIPLICACION = 31
	NparserDIVISION       = 32
	NparserRESTA          = 33
	NparserSUMA           = 34
	NparserPAR_IZQ        = 35
	NparserPAR_DER        = 36
	NparserLLAVE_IZQ      = 37
	NparserLLAVE_DER      = 38
	NparserCORCHETE_IZQ   = 39
	NparserCORCHETE_DER   = 40
	NparserCADENA         = 41
	NparserWHITESPACE     = 42
	NparserCOMMENT        = 43
	NparserLINE_COMMENT   = 44
)

// Nparser rules.
const (
	NparserRULE_inicio              = 0
	NparserRULE_expresion           = 1
	NparserRULE_operador_aritmetico = 2
)

// IInicioContext is an interface to support dynamic dispatch.
type IInicioContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInicioContext differentiates from other interfaces.
	IsInicioContext()
}

type InicioContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInicioContext() *InicioContext {
	var p = new(InicioContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_inicio
	return p
}

func (*InicioContext) IsInicioContext() {}

func NewInicioContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InicioContext {
	var p = new(InicioContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_inicio

	return p
}

func (s *InicioContext) GetParser() antlr.Parser { return s.parser }

func (s *InicioContext) CopyFrom(ctx *InicioContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *InicioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InicioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Prod_inicioContext struct {
	*InicioContext
}

func NewProd_inicioContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Prod_inicioContext {
	var p = new(Prod_inicioContext)

	p.InicioContext = NewEmptyInicioContext()
	p.parser = parser
	p.CopyFrom(ctx.(*InicioContext))

	return p
}

func (s *Prod_inicioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Prod_inicioContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Prod_inicioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterProd_inicio(s)
	}
}

func (s *Prod_inicioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitProd_inicio(s)
	}
}

func (s *Prod_inicioContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NparserVisitor:
		return t.VisitProd_inicio(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Nparser) Inicio() (localctx IInicioContext) {
	this := p
	_ = this

	localctx = NewInicioContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, NparserRULE_inicio)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewProd_inicioContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(6)
		p.expresion(0)
	}

	return localctx
}

// IExpresionContext is an interface to support dynamic dispatch.
type IExpresionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpresionContext differentiates from other interfaces.
	IsExpresionContext()
}

type ExpresionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpresionContext() *ExpresionContext {
	var p = new(ExpresionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_expresion
	return p
}

func (*ExpresionContext) IsExpresionContext() {}

func NewExpresionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpresionContext {
	var p = new(ExpresionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_expresion

	return p
}

func (s *ExpresionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpresionContext) CopyFrom(ctx *ExpresionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpresionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpresionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Op_aritContext struct {
	*ExpresionContext
	op_left  IExpresionContext
	op_right IExpresionContext
}

func NewOp_aritContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Op_aritContext {
	var p = new(Op_aritContext)

	p.ExpresionContext = NewEmptyExpresionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpresionContext))

	return p
}

func (s *Op_aritContext) GetOp_left() IExpresionContext { return s.op_left }

func (s *Op_aritContext) GetOp_right() IExpresionContext { return s.op_right }

func (s *Op_aritContext) SetOp_left(v IExpresionContext) { s.op_left = v }

func (s *Op_aritContext) SetOp_right(v IExpresionContext) { s.op_right = v }

func (s *Op_aritContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_aritContext) Operador_aritmetico() IOperador_aritmeticoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperador_aritmeticoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperador_aritmeticoContext)
}

func (s *Op_aritContext) AllExpresion() []IExpresionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpresionContext)(nil)).Elem())
	var tst = make([]IExpresionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpresionContext)
		}
	}

	return tst
}

func (s *Op_aritContext) Expresion(i int) IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Op_aritContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterOp_arit(s)
	}
}

func (s *Op_aritContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitOp_arit(s)
	}
}

func (s *Op_aritContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NparserVisitor:
		return t.VisitOp_arit(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnarioContext struct {
	*ExpresionContext
}

func NewUnarioContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnarioContext {
	var p = new(UnarioContext)

	p.ExpresionContext = NewEmptyExpresionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpresionContext))

	return p
}

func (s *UnarioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnarioContext) RESTA() antlr.TerminalNode {
	return s.GetToken(NparserRESTA, 0)
}

func (s *UnarioContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *UnarioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterUnario(s)
	}
}

func (s *UnarioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitUnario(s)
	}
}

func (s *UnarioContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NparserVisitor:
		return t.VisitUnario(s)

	default:
		return t.VisitChildren(s)
	}
}

type Prod_numeroContext struct {
	*ExpresionContext
}

func NewProd_numeroContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Prod_numeroContext {
	var p = new(Prod_numeroContext)

	p.ExpresionContext = NewEmptyExpresionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpresionContext))

	return p
}

func (s *Prod_numeroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Prod_numeroContext) NUMERO() antlr.TerminalNode {
	return s.GetToken(NparserNUMERO, 0)
}

func (s *Prod_numeroContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterProd_numero(s)
	}
}

func (s *Prod_numeroContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitProd_numero(s)
	}
}

func (s *Prod_numeroContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NparserVisitor:
		return t.VisitProd_numero(s)

	default:
		return t.VisitChildren(s)
	}
}

type Prod_decimalContext struct {
	*ExpresionContext
}

func NewProd_decimalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Prod_decimalContext {
	var p = new(Prod_decimalContext)

	p.ExpresionContext = NewEmptyExpresionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpresionContext))

	return p
}

func (s *Prod_decimalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Prod_decimalContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(NparserDECIMAL, 0)
}

func (s *Prod_decimalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterProd_decimal(s)
	}
}

func (s *Prod_decimalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitProd_decimal(s)
	}
}

func (s *Prod_decimalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NparserVisitor:
		return t.VisitProd_decimal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Nparser) Expresion() (localctx IExpresionContext) {
	return p.expresion(0)
}

func (p *Nparser) expresion(_p int) (localctx IExpresionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpresionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpresionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, NparserRULE_expresion, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(13)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NparserRESTA:
		localctx = NewUnarioContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(9)
			p.Match(NparserRESTA)
		}
		{
			p.SetState(10)
			p.expresion(4)
		}

	case NparserNUMERO:
		localctx = NewProd_numeroContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(11)
			p.Match(NparserNUMERO)
		}

	case NparserDECIMAL:
		localctx = NewProd_decimalContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(12)
			p.Match(NparserDECIMAL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewOp_aritContext(p, NewExpresionContext(p, _parentctx, _parentState))
			localctx.(*Op_aritContext).op_left = _prevctx

			p.PushNewRecursionContext(localctx, _startState, NparserRULE_expresion)
			p.SetState(15)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
			}
			{
				p.SetState(16)
				p.Operador_aritmetico()
			}
			{
				p.SetState(17)

				var _x = p.expresion(4)

				localctx.(*Op_aritContext).op_right = _x
			}

		}
		p.SetState(23)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}

	return localctx
}

// IOperador_aritmeticoContext is an interface to support dynamic dispatch.
type IOperador_aritmeticoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperador_aritmeticoContext differentiates from other interfaces.
	IsOperador_aritmeticoContext()
}

type Operador_aritmeticoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperador_aritmeticoContext() *Operador_aritmeticoContext {
	var p = new(Operador_aritmeticoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_operador_aritmetico
	return p
}

func (*Operador_aritmeticoContext) IsOperador_aritmeticoContext() {}

func NewOperador_aritmeticoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Operador_aritmeticoContext {
	var p = new(Operador_aritmeticoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_operador_aritmetico

	return p
}

func (s *Operador_aritmeticoContext) GetParser() antlr.Parser { return s.parser }

func (s *Operador_aritmeticoContext) CopyFrom(ctx *Operador_aritmeticoContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Operador_aritmeticoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Operador_aritmeticoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Op_restaContext struct {
	*Operador_aritmeticoContext
}

func NewOp_restaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Op_restaContext {
	var p = new(Op_restaContext)

	p.Operador_aritmeticoContext = NewEmptyOperador_aritmeticoContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Operador_aritmeticoContext))

	return p
}

func (s *Op_restaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_restaContext) RESTA() antlr.TerminalNode {
	return s.GetToken(NparserRESTA, 0)
}

func (s *Op_restaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterOp_resta(s)
	}
}

func (s *Op_restaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitOp_resta(s)
	}
}

func (s *Op_restaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NparserVisitor:
		return t.VisitOp_resta(s)

	default:
		return t.VisitChildren(s)
	}
}

type Op_divContext struct {
	*Operador_aritmeticoContext
}

func NewOp_divContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Op_divContext {
	var p = new(Op_divContext)

	p.Operador_aritmeticoContext = NewEmptyOperador_aritmeticoContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Operador_aritmeticoContext))

	return p
}

func (s *Op_divContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_divContext) DIVISION() antlr.TerminalNode {
	return s.GetToken(NparserDIVISION, 0)
}

func (s *Op_divContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterOp_div(s)
	}
}

func (s *Op_divContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitOp_div(s)
	}
}

func (s *Op_divContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NparserVisitor:
		return t.VisitOp_div(s)

	default:
		return t.VisitChildren(s)
	}
}

type Op_multiContext struct {
	*Operador_aritmeticoContext
}

func NewOp_multiContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Op_multiContext {
	var p = new(Op_multiContext)

	p.Operador_aritmeticoContext = NewEmptyOperador_aritmeticoContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Operador_aritmeticoContext))

	return p
}

func (s *Op_multiContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_multiContext) MULTIPLICACION() antlr.TerminalNode {
	return s.GetToken(NparserMULTIPLICACION, 0)
}

func (s *Op_multiContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterOp_multi(s)
	}
}

func (s *Op_multiContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitOp_multi(s)
	}
}

func (s *Op_multiContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NparserVisitor:
		return t.VisitOp_multi(s)

	default:
		return t.VisitChildren(s)
	}
}

type Op_modContext struct {
	*Operador_aritmeticoContext
}

func NewOp_modContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Op_modContext {
	var p = new(Op_modContext)

	p.Operador_aritmeticoContext = NewEmptyOperador_aritmeticoContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Operador_aritmeticoContext))

	return p
}

func (s *Op_modContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_modContext) MODULO() antlr.TerminalNode {
	return s.GetToken(NparserMODULO, 0)
}

func (s *Op_modContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterOp_mod(s)
	}
}

func (s *Op_modContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitOp_mod(s)
	}
}

func (s *Op_modContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NparserVisitor:
		return t.VisitOp_mod(s)

	default:
		return t.VisitChildren(s)
	}
}

type Op_sumContext struct {
	*Operador_aritmeticoContext
}

func NewOp_sumContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Op_sumContext {
	var p = new(Op_sumContext)

	p.Operador_aritmeticoContext = NewEmptyOperador_aritmeticoContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Operador_aritmeticoContext))

	return p
}

func (s *Op_sumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Op_sumContext) SUMA() antlr.TerminalNode {
	return s.GetToken(NparserSUMA, 0)
}

func (s *Op_sumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterOp_sum(s)
	}
}

func (s *Op_sumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitOp_sum(s)
	}
}

func (s *Op_sumContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NparserVisitor:
		return t.VisitOp_sum(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Nparser) Operador_aritmetico() (localctx IOperador_aritmeticoContext) {
	this := p
	_ = this

	localctx = NewOperador_aritmeticoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, NparserRULE_operador_aritmetico)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(29)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NparserMULTIPLICACION:
		localctx = NewOp_multiContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(24)
			p.Match(NparserMULTIPLICACION)
		}

	case NparserDIVISION:
		localctx = NewOp_divContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(25)
			p.Match(NparserDIVISION)
		}

	case NparserMODULO:
		localctx = NewOp_modContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(26)
			p.Match(NparserMODULO)
		}

	case NparserSUMA:
		localctx = NewOp_sumContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(27)
			p.Match(NparserSUMA)
		}

	case NparserRESTA:
		localctx = NewOp_restaContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(28)
			p.Match(NparserRESTA)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *Nparser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExpresionContext = nil
		if localctx != nil {
			t = localctx.(*ExpresionContext)
		}
		return p.Expresion_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *Nparser) Expresion_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
