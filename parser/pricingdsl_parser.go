// Code generated from grammar/PricingDSL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // PricingDSL
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

type PricingDSLParser struct {
	*antlr.BaseParser
}

var PricingDSLParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func pricingdslParserInit() {
	staticData := &PricingDSLParserStaticData
	staticData.LiteralNames = []string{
		"", "'.'", "'RULE'", "'IF'", "'THEN'", "'PRIORITY'", "'GROUP'", "'percent'",
		"'fixed'", "'final'", "'coupon'", "'promocode'", "'cert'", "'AND'",
		"'OR'", "'true'", "'=='", "'!='", "'>='", "'>'", "'<='", "'<'",
	}
	staticData.SymbolicNames = []string{
		"", "", "RULE", "IF", "THEN", "PRIORITY", "GROUP", "PERCENT", "FIXED",
		"FINAL", "COUPON", "PROMOCODE", "CERT", "AND", "OR", "TRUE", "EQ", "NE",
		"GE", "GT", "LE", "LT", "ID", "NUMBER", "STRING", "WS",
	}
	staticData.RuleNames = []string{
		"start", "rule", "modifier", "modifierType", "groupType", "expression",
		"comparison", "comparator", "identifier", "value",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 25, 76, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 1, 0, 4,
		0, 22, 8, 0, 11, 0, 12, 0, 23, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 3, 1, 36, 8, 1, 1, 1, 1, 1, 3, 1, 40, 8, 1, 1, 2, 1,
		2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 5, 5, 52, 8, 5, 10,
		5, 12, 5, 55, 9, 5, 1, 5, 3, 5, 58, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7,
		1, 7, 1, 8, 1, 8, 1, 8, 5, 8, 69, 8, 8, 10, 8, 12, 8, 72, 9, 8, 1, 9, 1,
		9, 1, 9, 0, 0, 10, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 0, 5, 1, 0, 7, 9,
		1, 0, 10, 12, 1, 0, 13, 14, 1, 0, 16, 21, 1, 0, 23, 24, 71, 0, 21, 1, 0,
		0, 0, 2, 27, 1, 0, 0, 0, 4, 41, 1, 0, 0, 0, 6, 44, 1, 0, 0, 0, 8, 46, 1,
		0, 0, 0, 10, 57, 1, 0, 0, 0, 12, 59, 1, 0, 0, 0, 14, 63, 1, 0, 0, 0, 16,
		65, 1, 0, 0, 0, 18, 73, 1, 0, 0, 0, 20, 22, 3, 2, 1, 0, 21, 20, 1, 0, 0,
		0, 22, 23, 1, 0, 0, 0, 23, 21, 1, 0, 0, 0, 23, 24, 1, 0, 0, 0, 24, 25,
		1, 0, 0, 0, 25, 26, 5, 0, 0, 1, 26, 1, 1, 0, 0, 0, 27, 28, 5, 2, 0, 0,
		28, 29, 5, 24, 0, 0, 29, 30, 5, 3, 0, 0, 30, 31, 3, 10, 5, 0, 31, 32, 5,
		4, 0, 0, 32, 35, 3, 4, 2, 0, 33, 34, 5, 5, 0, 0, 34, 36, 5, 23, 0, 0, 35,
		33, 1, 0, 0, 0, 35, 36, 1, 0, 0, 0, 36, 39, 1, 0, 0, 0, 37, 38, 5, 6, 0,
		0, 38, 40, 3, 8, 4, 0, 39, 37, 1, 0, 0, 0, 39, 40, 1, 0, 0, 0, 40, 3, 1,
		0, 0, 0, 41, 42, 3, 6, 3, 0, 42, 43, 5, 23, 0, 0, 43, 5, 1, 0, 0, 0, 44,
		45, 7, 0, 0, 0, 45, 7, 1, 0, 0, 0, 46, 47, 7, 1, 0, 0, 47, 9, 1, 0, 0,
		0, 48, 53, 3, 12, 6, 0, 49, 50, 7, 2, 0, 0, 50, 52, 3, 12, 6, 0, 51, 49,
		1, 0, 0, 0, 52, 55, 1, 0, 0, 0, 53, 51, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0,
		54, 58, 1, 0, 0, 0, 55, 53, 1, 0, 0, 0, 56, 58, 5, 15, 0, 0, 57, 48, 1,
		0, 0, 0, 57, 56, 1, 0, 0, 0, 58, 11, 1, 0, 0, 0, 59, 60, 3, 16, 8, 0, 60,
		61, 3, 14, 7, 0, 61, 62, 3, 18, 9, 0, 62, 13, 1, 0, 0, 0, 63, 64, 7, 3,
		0, 0, 64, 15, 1, 0, 0, 0, 65, 70, 5, 22, 0, 0, 66, 67, 5, 1, 0, 0, 67,
		69, 5, 22, 0, 0, 68, 66, 1, 0, 0, 0, 69, 72, 1, 0, 0, 0, 70, 68, 1, 0,
		0, 0, 70, 71, 1, 0, 0, 0, 71, 17, 1, 0, 0, 0, 72, 70, 1, 0, 0, 0, 73, 74,
		7, 4, 0, 0, 74, 19, 1, 0, 0, 0, 6, 23, 35, 39, 53, 57, 70,
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

// PricingDSLParserInit initializes any static state used to implement PricingDSLParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPricingDSLParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PricingDSLParserInit() {
	staticData := &PricingDSLParserStaticData
	staticData.once.Do(pricingdslParserInit)
}

// NewPricingDSLParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPricingDSLParser(input antlr.TokenStream) *PricingDSLParser {
	PricingDSLParserInit()
	this := new(PricingDSLParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &PricingDSLParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "PricingDSL.g4"

	return this
}

// PricingDSLParser tokens.
const (
	PricingDSLParserEOF       = antlr.TokenEOF
	PricingDSLParserT__0      = 1
	PricingDSLParserRULE      = 2
	PricingDSLParserIF        = 3
	PricingDSLParserTHEN      = 4
	PricingDSLParserPRIORITY  = 5
	PricingDSLParserGROUP     = 6
	PricingDSLParserPERCENT   = 7
	PricingDSLParserFIXED     = 8
	PricingDSLParserFINAL     = 9
	PricingDSLParserCOUPON    = 10
	PricingDSLParserPROMOCODE = 11
	PricingDSLParserCERT      = 12
	PricingDSLParserAND       = 13
	PricingDSLParserOR        = 14
	PricingDSLParserTRUE      = 15
	PricingDSLParserEQ        = 16
	PricingDSLParserNE        = 17
	PricingDSLParserGE        = 18
	PricingDSLParserGT        = 19
	PricingDSLParserLE        = 20
	PricingDSLParserLT        = 21
	PricingDSLParserID        = 22
	PricingDSLParserNUMBER    = 23
	PricingDSLParserSTRING    = 24
	PricingDSLParserWS        = 25
)

// PricingDSLParser rules.
const (
	PricingDSLParserRULE_start        = 0
	PricingDSLParserRULE_rule         = 1
	PricingDSLParserRULE_modifier     = 2
	PricingDSLParserRULE_modifierType = 3
	PricingDSLParserRULE_groupType    = 4
	PricingDSLParserRULE_expression   = 5
	PricingDSLParserRULE_comparison   = 6
	PricingDSLParserRULE_comparator   = 7
	PricingDSLParserRULE_identifier   = 8
	PricingDSLParserRULE_value        = 9
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllRule_() []IRuleContext
	Rule_(i int) IRuleContext

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PricingDSLParserRULE_start
	return p
}

func InitEmptyStartContext(p *StartContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PricingDSLParserRULE_start
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PricingDSLParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(PricingDSLParserEOF, 0)
}

func (s *StartContext) AllRule_() []IRuleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRuleContext); ok {
			len++
		}
	}

	tst := make([]IRuleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRuleContext); ok {
			tst[i] = t.(IRuleContext)
			i++
		}
	}

	return tst
}

func (s *StartContext) Rule_(i int) IRuleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRuleContext); ok {
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

	return t.(IRuleContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PricingDSLListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PricingDSLListener); ok {
		listenerT.ExitStart(s)
	}
}

func (s *StartContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PricingDSLVisitor:
		return t.VisitStart(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PricingDSLParser) Start_() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PricingDSLParserRULE_start)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PricingDSLParserRULE {
		{
			p.SetState(20)
			p.Rule_()
		}

		p.SetState(23)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(25)
		p.Match(PricingDSLParserEOF)
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

// IRuleContext is an interface to support dynamic dispatch.
type IRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RULE() antlr.TerminalNode
	STRING() antlr.TerminalNode
	IF() antlr.TerminalNode
	Expression() IExpressionContext
	THEN() antlr.TerminalNode
	Modifier() IModifierContext
	PRIORITY() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	GROUP() antlr.TerminalNode
	GroupType() IGroupTypeContext

	// IsRuleContext differentiates from other interfaces.
	IsRuleContext()
}

type RuleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleContext() *RuleContext {
	var p = new(RuleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PricingDSLParserRULE_rule
	return p
}

func InitEmptyRuleContext(p *RuleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PricingDSLParserRULE_rule
}

func (*RuleContext) IsRuleContext() {}

func NewRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleContext {
	var p = new(RuleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PricingDSLParserRULE_rule

	return p
}

func (s *RuleContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleContext) RULE() antlr.TerminalNode {
	return s.GetToken(PricingDSLParserRULE, 0)
}

func (s *RuleContext) STRING() antlr.TerminalNode {
	return s.GetToken(PricingDSLParserSTRING, 0)
}

func (s *RuleContext) IF() antlr.TerminalNode {
	return s.GetToken(PricingDSLParserIF, 0)
}

func (s *RuleContext) Expression() IExpressionContext {
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

func (s *RuleContext) THEN() antlr.TerminalNode {
	return s.GetToken(PricingDSLParserTHEN, 0)
}

func (s *RuleContext) Modifier() IModifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IModifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IModifierContext)
}

func (s *RuleContext) PRIORITY() antlr.TerminalNode {
	return s.GetToken(PricingDSLParserPRIORITY, 0)
}

func (s *RuleContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(PricingDSLParserNUMBER, 0)
}

func (s *RuleContext) GROUP() antlr.TerminalNode {
	return s.GetToken(PricingDSLParserGROUP, 0)
}

func (s *RuleContext) GroupType() IGroupTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupTypeContext)
}

func (s *RuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PricingDSLListener); ok {
		listenerT.EnterRule(s)
	}
}

func (s *RuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PricingDSLListener); ok {
		listenerT.ExitRule(s)
	}
}

func (s *RuleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PricingDSLVisitor:
		return t.VisitRule(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PricingDSLParser) Rule_() (localctx IRuleContext) {
	localctx = NewRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PricingDSLParserRULE_rule)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(27)
		p.Match(PricingDSLParserRULE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(28)
		p.Match(PricingDSLParserSTRING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(29)
		p.Match(PricingDSLParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(30)
		p.Expression()
	}
	{
		p.SetState(31)
		p.Match(PricingDSLParserTHEN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(32)
		p.Modifier()
	}
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PricingDSLParserPRIORITY {
		{
			p.SetState(33)
			p.Match(PricingDSLParserPRIORITY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(34)
			p.Match(PricingDSLParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PricingDSLParserGROUP {
		{
			p.SetState(37)
			p.Match(PricingDSLParserGROUP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(38)
			p.GroupType()
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

// IModifierContext is an interface to support dynamic dispatch.
type IModifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ModifierType() IModifierTypeContext
	NUMBER() antlr.TerminalNode

	// IsModifierContext differentiates from other interfaces.
	IsModifierContext()
}

type ModifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModifierContext() *ModifierContext {
	var p = new(ModifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PricingDSLParserRULE_modifier
	return p
}

func InitEmptyModifierContext(p *ModifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PricingDSLParserRULE_modifier
}

func (*ModifierContext) IsModifierContext() {}

func NewModifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModifierContext {
	var p = new(ModifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PricingDSLParserRULE_modifier

	return p
}

func (s *ModifierContext) GetParser() antlr.Parser { return s.parser }

func (s *ModifierContext) ModifierType() IModifierTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IModifierTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IModifierTypeContext)
}

func (s *ModifierContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(PricingDSLParserNUMBER, 0)
}

func (s *ModifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PricingDSLListener); ok {
		listenerT.EnterModifier(s)
	}
}

func (s *ModifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PricingDSLListener); ok {
		listenerT.ExitModifier(s)
	}
}

func (s *ModifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PricingDSLVisitor:
		return t.VisitModifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PricingDSLParser) Modifier() (localctx IModifierContext) {
	localctx = NewModifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PricingDSLParserRULE_modifier)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(41)
		p.ModifierType()
	}
	{
		p.SetState(42)
		p.Match(PricingDSLParserNUMBER)
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

// IModifierTypeContext is an interface to support dynamic dispatch.
type IModifierTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PERCENT() antlr.TerminalNode
	FIXED() antlr.TerminalNode
	FINAL() antlr.TerminalNode

	// IsModifierTypeContext differentiates from other interfaces.
	IsModifierTypeContext()
}

type ModifierTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModifierTypeContext() *ModifierTypeContext {
	var p = new(ModifierTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PricingDSLParserRULE_modifierType
	return p
}

func InitEmptyModifierTypeContext(p *ModifierTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PricingDSLParserRULE_modifierType
}

func (*ModifierTypeContext) IsModifierTypeContext() {}

func NewModifierTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModifierTypeContext {
	var p = new(ModifierTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PricingDSLParserRULE_modifierType

	return p
}

func (s *ModifierTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ModifierTypeContext) PERCENT() antlr.TerminalNode {
	return s.GetToken(PricingDSLParserPERCENT, 0)
}

func (s *ModifierTypeContext) FIXED() antlr.TerminalNode {
	return s.GetToken(PricingDSLParserFIXED, 0)
}

func (s *ModifierTypeContext) FINAL() antlr.TerminalNode {
	return s.GetToken(PricingDSLParserFINAL, 0)
}

func (s *ModifierTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModifierTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModifierTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PricingDSLListener); ok {
		listenerT.EnterModifierType(s)
	}
}

func (s *ModifierTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PricingDSLListener); ok {
		listenerT.ExitModifierType(s)
	}
}

func (s *ModifierTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PricingDSLVisitor:
		return t.VisitModifierType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PricingDSLParser) ModifierType() (localctx IModifierTypeContext) {
	localctx = NewModifierTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, PricingDSLParserRULE_modifierType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&896) != 0) {
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

// IGroupTypeContext is an interface to support dynamic dispatch.
type IGroupTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COUPON() antlr.TerminalNode
	PROMOCODE() antlr.TerminalNode
	CERT() antlr.TerminalNode

	// IsGroupTypeContext differentiates from other interfaces.
	IsGroupTypeContext()
}

type GroupTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupTypeContext() *GroupTypeContext {
	var p = new(GroupTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PricingDSLParserRULE_groupType
	return p
}

func InitEmptyGroupTypeContext(p *GroupTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PricingDSLParserRULE_groupType
}

func (*GroupTypeContext) IsGroupTypeContext() {}

func NewGroupTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupTypeContext {
	var p = new(GroupTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PricingDSLParserRULE_groupType

	return p
}

func (s *GroupTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupTypeContext) COUPON() antlr.TerminalNode {
	return s.GetToken(PricingDSLParserCOUPON, 0)
}

func (s *GroupTypeContext) PROMOCODE() antlr.TerminalNode {
	return s.GetToken(PricingDSLParserPROMOCODE, 0)
}

func (s *GroupTypeContext) CERT() antlr.TerminalNode {
	return s.GetToken(PricingDSLParserCERT, 0)
}

func (s *GroupTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PricingDSLListener); ok {
		listenerT.EnterGroupType(s)
	}
}

func (s *GroupTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PricingDSLListener); ok {
		listenerT.ExitGroupType(s)
	}
}

func (s *GroupTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PricingDSLVisitor:
		return t.VisitGroupType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PricingDSLParser) GroupType() (localctx IGroupTypeContext) {
	localctx = NewGroupTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, PricingDSLParserRULE_groupType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(46)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7168) != 0) {
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

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllComparison() []IComparisonContext
	Comparison(i int) IComparisonContext
	AllAND() []antlr.TerminalNode
	AND(i int) antlr.TerminalNode
	AllOR() []antlr.TerminalNode
	OR(i int) antlr.TerminalNode
	TRUE() antlr.TerminalNode

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
	p.RuleIndex = PricingDSLParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PricingDSLParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PricingDSLParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) AllComparison() []IComparisonContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IComparisonContext); ok {
			len++
		}
	}

	tst := make([]IComparisonContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IComparisonContext); ok {
			tst[i] = t.(IComparisonContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) Comparison(i int) IComparisonContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparisonContext); ok {
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

	return t.(IComparisonContext)
}

func (s *ExpressionContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(PricingDSLParserAND)
}

func (s *ExpressionContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(PricingDSLParserAND, i)
}

func (s *ExpressionContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(PricingDSLParserOR)
}

func (s *ExpressionContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(PricingDSLParserOR, i)
}

func (s *ExpressionContext) TRUE() antlr.TerminalNode {
	return s.GetToken(PricingDSLParserTRUE, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PricingDSLListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PricingDSLListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PricingDSLVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PricingDSLParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, PricingDSLParserRULE_expression)
	var _la int

	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PricingDSLParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(48)
			p.Comparison()
		}
		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PricingDSLParserAND || _la == PricingDSLParserOR {
			{
				p.SetState(49)
				_la = p.GetTokenStream().LA(1)

				if !(_la == PricingDSLParserAND || _la == PricingDSLParserOR) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(50)
				p.Comparison()
			}

			p.SetState(55)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case PricingDSLParserTRUE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(56)
			p.Match(PricingDSLParserTRUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IComparisonContext is an interface to support dynamic dispatch.
type IComparisonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Identifier() IIdentifierContext
	Comparator() IComparatorContext
	Value() IValueContext

	// IsComparisonContext differentiates from other interfaces.
	IsComparisonContext()
}

type ComparisonContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonContext() *ComparisonContext {
	var p = new(ComparisonContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PricingDSLParserRULE_comparison
	return p
}

func InitEmptyComparisonContext(p *ComparisonContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PricingDSLParserRULE_comparison
}

func (*ComparisonContext) IsComparisonContext() {}

func NewComparisonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonContext {
	var p = new(ComparisonContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PricingDSLParserRULE_comparison

	return p
}

func (s *ComparisonContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ComparisonContext) Comparator() IComparatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparatorContext)
}

func (s *ComparisonContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PricingDSLListener); ok {
		listenerT.EnterComparison(s)
	}
}

func (s *ComparisonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PricingDSLListener); ok {
		listenerT.ExitComparison(s)
	}
}

func (s *ComparisonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PricingDSLVisitor:
		return t.VisitComparison(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PricingDSLParser) Comparison() (localctx IComparisonContext) {
	localctx = NewComparisonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, PricingDSLParserRULE_comparison)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(59)
		p.Identifier()
	}
	{
		p.SetState(60)
		p.Comparator()
	}
	{
		p.SetState(61)
		p.Value()
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

// IComparatorContext is an interface to support dynamic dispatch.
type IComparatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EQ() antlr.TerminalNode
	NE() antlr.TerminalNode
	GE() antlr.TerminalNode
	GT() antlr.TerminalNode
	LE() antlr.TerminalNode
	LT() antlr.TerminalNode

	// IsComparatorContext differentiates from other interfaces.
	IsComparatorContext()
}

type ComparatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparatorContext() *ComparatorContext {
	var p = new(ComparatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PricingDSLParserRULE_comparator
	return p
}

func InitEmptyComparatorContext(p *ComparatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PricingDSLParserRULE_comparator
}

func (*ComparatorContext) IsComparatorContext() {}

func NewComparatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparatorContext {
	var p = new(ComparatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PricingDSLParserRULE_comparator

	return p
}

func (s *ComparatorContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparatorContext) EQ() antlr.TerminalNode {
	return s.GetToken(PricingDSLParserEQ, 0)
}

func (s *ComparatorContext) NE() antlr.TerminalNode {
	return s.GetToken(PricingDSLParserNE, 0)
}

func (s *ComparatorContext) GE() antlr.TerminalNode {
	return s.GetToken(PricingDSLParserGE, 0)
}

func (s *ComparatorContext) GT() antlr.TerminalNode {
	return s.GetToken(PricingDSLParserGT, 0)
}

func (s *ComparatorContext) LE() antlr.TerminalNode {
	return s.GetToken(PricingDSLParserLE, 0)
}

func (s *ComparatorContext) LT() antlr.TerminalNode {
	return s.GetToken(PricingDSLParserLT, 0)
}

func (s *ComparatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PricingDSLListener); ok {
		listenerT.EnterComparator(s)
	}
}

func (s *ComparatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PricingDSLListener); ok {
		listenerT.ExitComparator(s)
	}
}

func (s *ComparatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PricingDSLVisitor:
		return t.VisitComparator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PricingDSLParser) Comparator() (localctx IComparatorContext) {
	localctx = NewComparatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, PricingDSLParserRULE_comparator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(63)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4128768) != 0) {
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

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PricingDSLParserRULE_identifier
	return p
}

func InitEmptyIdentifierContext(p *IdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PricingDSLParserRULE_identifier
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PricingDSLParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(PricingDSLParserID)
}

func (s *IdentifierContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(PricingDSLParserID, i)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PricingDSLListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PricingDSLListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (s *IdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PricingDSLVisitor:
		return t.VisitIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PricingDSLParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, PricingDSLParserRULE_identifier)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(65)
		p.Match(PricingDSLParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PricingDSLParserT__0 {
		{
			p.SetState(66)
			p.Match(PricingDSLParserT__0)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(67)
			p.Match(PricingDSLParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(72)
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

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING() antlr.TerminalNode
	NUMBER() antlr.TerminalNode

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PricingDSLParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PricingDSLParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PricingDSLParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(PricingDSLParserSTRING, 0)
}

func (s *ValueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(PricingDSLParserNUMBER, 0)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PricingDSLListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PricingDSLListener); ok {
		listenerT.ExitValue(s)
	}
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case PricingDSLVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *PricingDSLParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, PricingDSLParserRULE_value)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(73)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PricingDSLParserNUMBER || _la == PricingDSLParserSTRING) {
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
