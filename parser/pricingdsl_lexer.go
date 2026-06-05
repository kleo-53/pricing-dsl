// Code generated from grammar/PricingDSL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type PricingDSLLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var PricingDSLLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func pricingdsllexerLexerInit() {
	staticData := &PricingDSLLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'.'", "'RULE'", "'IF'", "'THEN'", "'PRIORITY'", "'GROUP'", "",
		"'percent'", "'fixed'", "'final'", "'coupon'", "'promocode'", "'cert'",
		"'AND'", "'OR'", "'true'", "'=='", "'!='", "'>='", "'>'", "'<='", "'<'",
	}
	staticData.SymbolicNames = []string{
		"", "", "RULE", "IF", "THEN", "PRIORITY", "GROUP", "STATUS", "PERCENT",
		"FIXED", "FINAL", "COUPON", "PROMOCODE", "CERT", "AND", "OR", "TRUE",
		"EQ", "NE", "GE", "GT", "LE", "LT", "ID", "NUMBER", "STRING", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "RULE", "IF", "THEN", "PRIORITY", "GROUP", "STATUS", "PERCENT",
		"FIXED", "FINAL", "COUPON", "PROMOCODE", "CERT", "AND", "OR", "TRUE",
		"EQ", "NE", "GE", "GT", "LE", "LT", "ID", "NUMBER", "STRING", "ESCAPE",
		"WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 26, 212, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3,
		6, 101, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15,
		1, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1,
		19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 5, 22, 175,
		8, 22, 10, 22, 12, 22, 178, 9, 22, 1, 23, 4, 23, 181, 8, 23, 11, 23, 12,
		23, 182, 1, 23, 1, 23, 4, 23, 187, 8, 23, 11, 23, 12, 23, 188, 3, 23, 191,
		8, 23, 1, 24, 1, 24, 1, 24, 5, 24, 196, 8, 24, 10, 24, 12, 24, 199, 9,
		24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 26, 4, 26, 207, 8, 26, 11, 26,
		12, 26, 208, 1, 26, 1, 26, 0, 0, 27, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11,
		6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15,
		31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24,
		49, 25, 51, 0, 53, 26, 1, 0, 5, 2, 0, 65, 90, 97, 122, 4, 0, 48, 57, 65,
		90, 95, 95, 97, 122, 1, 0, 48, 57, 2, 0, 34, 34, 92, 92, 3, 0, 9, 10, 13,
		13, 32, 32, 218, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0,
		0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0,
		0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0,
		0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1,
		0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37,
		1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0,
		45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0,
		1, 55, 1, 0, 0, 0, 3, 57, 1, 0, 0, 0, 5, 62, 1, 0, 0, 0, 7, 65, 1, 0, 0,
		0, 9, 70, 1, 0, 0, 0, 11, 79, 1, 0, 0, 0, 13, 100, 1, 0, 0, 0, 15, 102,
		1, 0, 0, 0, 17, 110, 1, 0, 0, 0, 19, 116, 1, 0, 0, 0, 21, 122, 1, 0, 0,
		0, 23, 129, 1, 0, 0, 0, 25, 139, 1, 0, 0, 0, 27, 144, 1, 0, 0, 0, 29, 148,
		1, 0, 0, 0, 31, 151, 1, 0, 0, 0, 33, 156, 1, 0, 0, 0, 35, 159, 1, 0, 0,
		0, 37, 162, 1, 0, 0, 0, 39, 165, 1, 0, 0, 0, 41, 167, 1, 0, 0, 0, 43, 170,
		1, 0, 0, 0, 45, 172, 1, 0, 0, 0, 47, 180, 1, 0, 0, 0, 49, 192, 1, 0, 0,
		0, 51, 202, 1, 0, 0, 0, 53, 206, 1, 0, 0, 0, 55, 56, 5, 46, 0, 0, 56, 2,
		1, 0, 0, 0, 57, 58, 5, 82, 0, 0, 58, 59, 5, 85, 0, 0, 59, 60, 5, 76, 0,
		0, 60, 61, 5, 69, 0, 0, 61, 4, 1, 0, 0, 0, 62, 63, 5, 73, 0, 0, 63, 64,
		5, 70, 0, 0, 64, 6, 1, 0, 0, 0, 65, 66, 5, 84, 0, 0, 66, 67, 5, 72, 0,
		0, 67, 68, 5, 69, 0, 0, 68, 69, 5, 78, 0, 0, 69, 8, 1, 0, 0, 0, 70, 71,
		5, 80, 0, 0, 71, 72, 5, 82, 0, 0, 72, 73, 5, 73, 0, 0, 73, 74, 5, 79, 0,
		0, 74, 75, 5, 82, 0, 0, 75, 76, 5, 73, 0, 0, 76, 77, 5, 84, 0, 0, 77, 78,
		5, 89, 0, 0, 78, 10, 1, 0, 0, 0, 79, 80, 5, 71, 0, 0, 80, 81, 5, 82, 0,
		0, 81, 82, 5, 79, 0, 0, 82, 83, 5, 85, 0, 0, 83, 84, 5, 80, 0, 0, 84, 12,
		1, 0, 0, 0, 85, 86, 5, 68, 0, 0, 86, 87, 5, 73, 0, 0, 87, 88, 5, 83, 0,
		0, 88, 89, 5, 65, 0, 0, 89, 90, 5, 66, 0, 0, 90, 91, 5, 76, 0, 0, 91, 92,
		5, 69, 0, 0, 92, 101, 5, 68, 0, 0, 93, 94, 5, 69, 0, 0, 94, 95, 5, 78,
		0, 0, 95, 96, 5, 65, 0, 0, 96, 97, 5, 66, 0, 0, 97, 98, 5, 76, 0, 0, 98,
		99, 5, 69, 0, 0, 99, 101, 5, 68, 0, 0, 100, 85, 1, 0, 0, 0, 100, 93, 1,
		0, 0, 0, 101, 14, 1, 0, 0, 0, 102, 103, 5, 112, 0, 0, 103, 104, 5, 101,
		0, 0, 104, 105, 5, 114, 0, 0, 105, 106, 5, 99, 0, 0, 106, 107, 5, 101,
		0, 0, 107, 108, 5, 110, 0, 0, 108, 109, 5, 116, 0, 0, 109, 16, 1, 0, 0,
		0, 110, 111, 5, 102, 0, 0, 111, 112, 5, 105, 0, 0, 112, 113, 5, 120, 0,
		0, 113, 114, 5, 101, 0, 0, 114, 115, 5, 100, 0, 0, 115, 18, 1, 0, 0, 0,
		116, 117, 5, 102, 0, 0, 117, 118, 5, 105, 0, 0, 118, 119, 5, 110, 0, 0,
		119, 120, 5, 97, 0, 0, 120, 121, 5, 108, 0, 0, 121, 20, 1, 0, 0, 0, 122,
		123, 5, 99, 0, 0, 123, 124, 5, 111, 0, 0, 124, 125, 5, 117, 0, 0, 125,
		126, 5, 112, 0, 0, 126, 127, 5, 111, 0, 0, 127, 128, 5, 110, 0, 0, 128,
		22, 1, 0, 0, 0, 129, 130, 5, 112, 0, 0, 130, 131, 5, 114, 0, 0, 131, 132,
		5, 111, 0, 0, 132, 133, 5, 109, 0, 0, 133, 134, 5, 111, 0, 0, 134, 135,
		5, 99, 0, 0, 135, 136, 5, 111, 0, 0, 136, 137, 5, 100, 0, 0, 137, 138,
		5, 101, 0, 0, 138, 24, 1, 0, 0, 0, 139, 140, 5, 99, 0, 0, 140, 141, 5,
		101, 0, 0, 141, 142, 5, 114, 0, 0, 142, 143, 5, 116, 0, 0, 143, 26, 1,
		0, 0, 0, 144, 145, 5, 65, 0, 0, 145, 146, 5, 78, 0, 0, 146, 147, 5, 68,
		0, 0, 147, 28, 1, 0, 0, 0, 148, 149, 5, 79, 0, 0, 149, 150, 5, 82, 0, 0,
		150, 30, 1, 0, 0, 0, 151, 152, 5, 116, 0, 0, 152, 153, 5, 114, 0, 0, 153,
		154, 5, 117, 0, 0, 154, 155, 5, 101, 0, 0, 155, 32, 1, 0, 0, 0, 156, 157,
		5, 61, 0, 0, 157, 158, 5, 61, 0, 0, 158, 34, 1, 0, 0, 0, 159, 160, 5, 33,
		0, 0, 160, 161, 5, 61, 0, 0, 161, 36, 1, 0, 0, 0, 162, 163, 5, 62, 0, 0,
		163, 164, 5, 61, 0, 0, 164, 38, 1, 0, 0, 0, 165, 166, 5, 62, 0, 0, 166,
		40, 1, 0, 0, 0, 167, 168, 5, 60, 0, 0, 168, 169, 5, 61, 0, 0, 169, 42,
		1, 0, 0, 0, 170, 171, 5, 60, 0, 0, 171, 44, 1, 0, 0, 0, 172, 176, 7, 0,
		0, 0, 173, 175, 7, 1, 0, 0, 174, 173, 1, 0, 0, 0, 175, 178, 1, 0, 0, 0,
		176, 174, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 46, 1, 0, 0, 0, 178, 176,
		1, 0, 0, 0, 179, 181, 7, 2, 0, 0, 180, 179, 1, 0, 0, 0, 181, 182, 1, 0,
		0, 0, 182, 180, 1, 0, 0, 0, 182, 183, 1, 0, 0, 0, 183, 190, 1, 0, 0, 0,
		184, 186, 5, 46, 0, 0, 185, 187, 7, 2, 0, 0, 186, 185, 1, 0, 0, 0, 187,
		188, 1, 0, 0, 0, 188, 186, 1, 0, 0, 0, 188, 189, 1, 0, 0, 0, 189, 191,
		1, 0, 0, 0, 190, 184, 1, 0, 0, 0, 190, 191, 1, 0, 0, 0, 191, 48, 1, 0,
		0, 0, 192, 197, 5, 34, 0, 0, 193, 196, 3, 51, 25, 0, 194, 196, 8, 3, 0,
		0, 195, 193, 1, 0, 0, 0, 195, 194, 1, 0, 0, 0, 196, 199, 1, 0, 0, 0, 197,
		195, 1, 0, 0, 0, 197, 198, 1, 0, 0, 0, 198, 200, 1, 0, 0, 0, 199, 197,
		1, 0, 0, 0, 200, 201, 5, 34, 0, 0, 201, 50, 1, 0, 0, 0, 202, 203, 5, 92,
		0, 0, 203, 204, 7, 3, 0, 0, 204, 52, 1, 0, 0, 0, 205, 207, 7, 4, 0, 0,
		206, 205, 1, 0, 0, 0, 207, 208, 1, 0, 0, 0, 208, 206, 1, 0, 0, 0, 208,
		209, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0, 210, 211, 6, 26, 0, 0, 211, 54,
		1, 0, 0, 0, 9, 0, 100, 176, 182, 188, 190, 195, 197, 208, 1, 6, 0, 0,
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

// PricingDSLLexerInit initializes any static state used to implement PricingDSLLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewPricingDSLLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func PricingDSLLexerInit() {
	staticData := &PricingDSLLexerLexerStaticData
	staticData.once.Do(pricingdsllexerLexerInit)
}

// NewPricingDSLLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewPricingDSLLexer(input antlr.CharStream) *PricingDSLLexer {
	PricingDSLLexerInit()
	l := new(PricingDSLLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &PricingDSLLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "PricingDSL.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// PricingDSLLexer tokens.
const (
	PricingDSLLexerT__0      = 1
	PricingDSLLexerRULE      = 2
	PricingDSLLexerIF        = 3
	PricingDSLLexerTHEN      = 4
	PricingDSLLexerPRIORITY  = 5
	PricingDSLLexerGROUP     = 6
	PricingDSLLexerSTATUS    = 7
	PricingDSLLexerPERCENT   = 8
	PricingDSLLexerFIXED     = 9
	PricingDSLLexerFINAL     = 10
	PricingDSLLexerCOUPON    = 11
	PricingDSLLexerPROMOCODE = 12
	PricingDSLLexerCERT      = 13
	PricingDSLLexerAND       = 14
	PricingDSLLexerOR        = 15
	PricingDSLLexerTRUE      = 16
	PricingDSLLexerEQ        = 17
	PricingDSLLexerNE        = 18
	PricingDSLLexerGE        = 19
	PricingDSLLexerGT        = 20
	PricingDSLLexerLE        = 21
	PricingDSLLexerLT        = 22
	PricingDSLLexerID        = 23
	PricingDSLLexerNUMBER    = 24
	PricingDSLLexerSTRING    = 25
	PricingDSLLexerWS        = 26
)
