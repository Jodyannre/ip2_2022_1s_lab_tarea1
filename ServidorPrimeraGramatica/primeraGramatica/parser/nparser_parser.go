// Code generated from Nparser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Nparser

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import "ServidorPrimeraGramatica/primeraGramatica/Analizador/Ast"
import "ServidorPrimeraGramatica/primeraGramatica/Analizador/Expresiones"
import "ServidorPrimeraGramatica/primeraGramatica/Analizador/Instrucciones"
import "github.com/colegno/arraylist"
import "strings"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 46, 251,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 3, 2, 3, 2, 3, 2, 3, 3,
	7, 3, 37, 10, 3, 12, 3, 14, 3, 40, 11, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 56, 10, 4,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 74, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 5, 7, 106, 10, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	7, 7, 128, 10, 7, 12, 7, 14, 7, 131, 11, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 5, 8, 141, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 147,
	10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 153, 10, 10, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 165, 10, 11,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 215, 10, 12, 3, 13, 7,
	13, 218, 10, 13, 12, 13, 14, 13, 221, 11, 13, 3, 13, 3, 13, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 2, 3, 12, 17, 2, 4, 6, 8, 10, 12, 14, 16, 18,
	20, 22, 24, 26, 28, 30, 2, 6, 3, 2, 33, 34, 3, 2, 35, 36, 3, 2, 28, 29,
	3, 2, 4, 5, 2, 264, 2, 32, 3, 2, 2, 2, 4, 38, 3, 2, 2, 2, 6, 55, 3, 2,
	2, 2, 8, 73, 3, 2, 2, 2, 10, 75, 3, 2, 2, 2, 12, 105, 3, 2, 2, 2, 14, 140,
	3, 2, 2, 2, 16, 146, 3, 2, 2, 2, 18, 152, 3, 2, 2, 2, 20, 164, 3, 2, 2,
	2, 22, 214, 3, 2, 2, 2, 24, 219, 3, 2, 2, 2, 26, 224, 3, 2, 2, 2, 28, 235,
	3, 2, 2, 2, 30, 241, 3, 2, 2, 2, 32, 33, 5, 4, 3, 2, 33, 34, 8, 2, 1, 2,
	34, 3, 3, 2, 2, 2, 35, 37, 5, 6, 4, 2, 36, 35, 3, 2, 2, 2, 37, 40, 3, 2,
	2, 2, 38, 36, 3, 2, 2, 2, 38, 39, 3, 2, 2, 2, 39, 41, 3, 2, 2, 2, 40, 38,
	3, 2, 2, 2, 41, 42, 8, 3, 1, 2, 42, 5, 3, 2, 2, 2, 43, 44, 5, 22, 12, 2,
	44, 45, 8, 4, 1, 2, 45, 56, 3, 2, 2, 2, 46, 47, 5, 10, 6, 2, 47, 48, 8,
	4, 1, 2, 48, 56, 3, 2, 2, 2, 49, 50, 5, 8, 5, 2, 50, 51, 8, 4, 1, 2, 51,
	56, 3, 2, 2, 2, 52, 53, 5, 30, 16, 2, 53, 54, 8, 4, 1, 2, 54, 56, 3, 2,
	2, 2, 55, 43, 3, 2, 2, 2, 55, 46, 3, 2, 2, 2, 55, 49, 3, 2, 2, 2, 55, 52,
	3, 2, 2, 2, 56, 7, 3, 2, 2, 2, 57, 58, 7, 21, 2, 2, 58, 59, 7, 24, 2, 2,
	59, 60, 5, 14, 8, 2, 60, 61, 7, 31, 2, 2, 61, 62, 7, 30, 2, 2, 62, 63,
	5, 12, 7, 2, 63, 64, 8, 5, 1, 2, 64, 65, 7, 29, 2, 2, 65, 66, 7, 27, 2,
	2, 66, 74, 3, 2, 2, 2, 67, 68, 7, 21, 2, 2, 68, 69, 7, 24, 2, 2, 69, 70,
	5, 14, 8, 2, 70, 71, 8, 5, 1, 2, 71, 72, 7, 27, 2, 2, 72, 74, 3, 2, 2,
	2, 73, 57, 3, 2, 2, 2, 73, 67, 3, 2, 2, 2, 74, 9, 3, 2, 2, 2, 75, 76, 7,
	24, 2, 2, 76, 77, 7, 31, 2, 2, 77, 78, 7, 30, 2, 2, 78, 79, 5, 12, 7, 2,
	79, 80, 8, 6, 1, 2, 80, 81, 7, 29, 2, 2, 81, 82, 7, 27, 2, 2, 82, 11, 3,
	2, 2, 2, 83, 84, 8, 7, 1, 2, 84, 85, 7, 35, 2, 2, 85, 86, 5, 12, 7, 14,
	86, 87, 8, 7, 1, 2, 87, 106, 3, 2, 2, 2, 88, 89, 7, 22, 2, 2, 89, 106,
	8, 7, 1, 2, 90, 91, 7, 23, 2, 2, 91, 106, 8, 7, 1, 2, 92, 93, 7, 43, 2,
	2, 93, 106, 8, 7, 1, 2, 94, 95, 7, 13, 2, 2, 95, 106, 8, 7, 1, 2, 96, 97,
	7, 14, 2, 2, 97, 106, 8, 7, 1, 2, 98, 99, 7, 24, 2, 2, 99, 106, 8, 7, 1,
	2, 100, 101, 7, 37, 2, 2, 101, 102, 5, 12, 7, 2, 102, 103, 7, 38, 2, 2,
	103, 104, 8, 7, 1, 2, 104, 106, 3, 2, 2, 2, 105, 83, 3, 2, 2, 2, 105, 88,
	3, 2, 2, 2, 105, 90, 3, 2, 2, 2, 105, 92, 3, 2, 2, 2, 105, 94, 3, 2, 2,
	2, 105, 96, 3, 2, 2, 2, 105, 98, 3, 2, 2, 2, 105, 100, 3, 2, 2, 2, 106,
	129, 3, 2, 2, 2, 107, 108, 12, 13, 2, 2, 108, 109, 9, 2, 2, 2, 109, 110,
	5, 12, 7, 14, 110, 111, 8, 7, 1, 2, 111, 128, 3, 2, 2, 2, 112, 113, 12,
	12, 2, 2, 113, 114, 9, 3, 2, 2, 114, 115, 5, 12, 7, 13, 115, 116, 8, 7,
	1, 2, 116, 128, 3, 2, 2, 2, 117, 118, 12, 11, 2, 2, 118, 119, 9, 4, 2,
	2, 119, 120, 5, 12, 7, 12, 120, 121, 8, 7, 1, 2, 121, 128, 3, 2, 2, 2,
	122, 123, 12, 10, 2, 2, 123, 124, 9, 5, 2, 2, 124, 125, 5, 12, 7, 11, 125,
	126, 8, 7, 1, 2, 126, 128, 3, 2, 2, 2, 127, 107, 3, 2, 2, 2, 127, 112,
	3, 2, 2, 2, 127, 117, 3, 2, 2, 2, 127, 122, 3, 2, 2, 2, 128, 131, 3, 2,
	2, 2, 129, 127, 3, 2, 2, 2, 129, 130, 3, 2, 2, 2, 130, 13, 3, 2, 2, 2,
	131, 129, 3, 2, 2, 2, 132, 133, 7, 10, 2, 2, 133, 141, 8, 8, 1, 2, 134,
	135, 7, 11, 2, 2, 135, 141, 8, 8, 1, 2, 136, 137, 7, 12, 2, 2, 137, 141,
	8, 8, 1, 2, 138, 139, 7, 9, 2, 2, 139, 141, 8, 8, 1, 2, 140, 132, 3, 2,
	2, 2, 140, 134, 3, 2, 2, 2, 140, 136, 3, 2, 2, 2, 140, 138, 3, 2, 2, 2,
	141, 15, 3, 2, 2, 2, 142, 143, 7, 29, 2, 2, 143, 147, 8, 9, 1, 2, 144,
	145, 7, 28, 2, 2, 145, 147, 8, 9, 1, 2, 146, 142, 3, 2, 2, 2, 146, 144,
	3, 2, 2, 2, 147, 17, 3, 2, 2, 2, 148, 149, 7, 4, 2, 2, 149, 153, 8, 10,
	1, 2, 150, 151, 7, 5, 2, 2, 151, 153, 8, 10, 1, 2, 152, 148, 3, 2, 2, 2,
	152, 150, 3, 2, 2, 2, 153, 19, 3, 2, 2, 2, 154, 155, 7, 33, 2, 2, 155,
	165, 8, 11, 1, 2, 156, 157, 7, 34, 2, 2, 157, 165, 8, 11, 1, 2, 158, 159,
	7, 32, 2, 2, 159, 165, 8, 11, 1, 2, 160, 161, 7, 36, 2, 2, 161, 165, 8,
	11, 1, 2, 162, 163, 7, 35, 2, 2, 163, 165, 8, 11, 1, 2, 164, 154, 3, 2,
	2, 2, 164, 156, 3, 2, 2, 2, 164, 158, 3, 2, 2, 2, 164, 160, 3, 2, 2, 2,
	164, 162, 3, 2, 2, 2, 165, 21, 3, 2, 2, 2, 166, 167, 7, 6, 2, 2, 167, 168,
	7, 25, 2, 2, 168, 169, 7, 37, 2, 2, 169, 170, 5, 12, 7, 2, 170, 171, 7,
	38, 2, 2, 171, 172, 7, 41, 2, 2, 172, 173, 5, 4, 3, 2, 173, 174, 7, 42,
	2, 2, 174, 175, 5, 28, 15, 2, 175, 176, 7, 27, 2, 2, 176, 177, 8, 12, 1,
	2, 177, 215, 3, 2, 2, 2, 178, 179, 7, 6, 2, 2, 179, 180, 7, 25, 2, 2, 180,
	181, 7, 37, 2, 2, 181, 182, 5, 12, 7, 2, 182, 183, 7, 38, 2, 2, 183, 184,
	7, 41, 2, 2, 184, 185, 5, 4, 3, 2, 185, 186, 7, 42, 2, 2, 186, 187, 7,
	27, 2, 2, 187, 188, 8, 12, 1, 2, 188, 215, 3, 2, 2, 2, 189, 190, 7, 6,
	2, 2, 190, 191, 7, 25, 2, 2, 191, 192, 7, 37, 2, 2, 192, 193, 5, 12, 7,
	2, 193, 194, 7, 38, 2, 2, 194, 195, 7, 41, 2, 2, 195, 196, 5, 4, 3, 2,
	196, 197, 7, 42, 2, 2, 197, 198, 5, 24, 13, 2, 198, 199, 7, 27, 2, 2, 199,
	200, 8, 12, 1, 2, 200, 215, 3, 2, 2, 2, 201, 202, 7, 6, 2, 2, 202, 203,
	7, 25, 2, 2, 203, 204, 7, 37, 2, 2, 204, 205, 5, 12, 7, 2, 205, 206, 7,
	38, 2, 2, 206, 207, 7, 41, 2, 2, 207, 208, 5, 4, 3, 2, 208, 209, 7, 42,
	2, 2, 209, 210, 5, 24, 13, 2, 210, 211, 5, 28, 15, 2, 211, 212, 7, 27,
	2, 2, 212, 213, 8, 12, 1, 2, 213, 215, 3, 2, 2, 2, 214, 166, 3, 2, 2, 2,
	214, 178, 3, 2, 2, 2, 214, 189, 3, 2, 2, 2, 214, 201, 3, 2, 2, 2, 215,
	23, 3, 2, 2, 2, 216, 218, 5, 26, 14, 2, 217, 216, 3, 2, 2, 2, 218, 221,
	3, 2, 2, 2, 219, 217, 3, 2, 2, 2, 219, 220, 3, 2, 2, 2, 220, 222, 3, 2,
	2, 2, 221, 219, 3, 2, 2, 2, 222, 223, 8, 13, 1, 2, 223, 25, 3, 2, 2, 2,
	224, 225, 7, 7, 2, 2, 225, 226, 7, 6, 2, 2, 226, 227, 7, 25, 2, 2, 227,
	228, 7, 37, 2, 2, 228, 229, 5, 12, 7, 2, 229, 230, 7, 38, 2, 2, 230, 231,
	7, 41, 2, 2, 231, 232, 5, 4, 3, 2, 232, 233, 7, 42, 2, 2, 233, 234, 8,
	14, 1, 2, 234, 27, 3, 2, 2, 2, 235, 236, 7, 7, 2, 2, 236, 237, 7, 41, 2,
	2, 237, 238, 5, 4, 3, 2, 238, 239, 8, 15, 1, 2, 239, 240, 7, 42, 2, 2,
	240, 29, 3, 2, 2, 2, 241, 242, 7, 3, 2, 2, 242, 243, 7, 26, 2, 2, 243,
	244, 7, 8, 2, 2, 244, 245, 7, 37, 2, 2, 245, 246, 5, 12, 7, 2, 246, 247,
	8, 16, 1, 2, 247, 248, 7, 38, 2, 2, 248, 249, 7, 27, 2, 2, 249, 31, 3,
	2, 2, 2, 14, 38, 55, 73, 105, 127, 129, 140, 146, 152, 164, 214, 219,
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
	"inicio", "instrucciones", "instruccion", "declaracion", "asignacion",
	"expresion", "tipo", "operador_comparacion", "operador_logico", "operador_aritmetico",
	"control_IF_simple", "control_IF_entonces", "control_if_entonces", "control_else",
	"funcion_imprimir",
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
	NparserRULE_inicio               = 0
	NparserRULE_instrucciones        = 1
	NparserRULE_instruccion          = 2
	NparserRULE_declaracion          = 3
	NparserRULE_asignacion           = 4
	NparserRULE_expresion            = 5
	NparserRULE_tipo                 = 6
	NparserRULE_operador_comparacion = 7
	NparserRULE_operador_logico      = 8
	NparserRULE_operador_aritmetico  = 9
	NparserRULE_control_IF_simple    = 10
	NparserRULE_control_IF_entonces  = 11
	NparserRULE_control_if_entonces  = 12
	NparserRULE_control_else         = 13
	NparserRULE_funcion_imprimir     = 14
)

// IInicioContext is an interface to support dynamic dispatch.
type IInicioContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetLista returns the lista attribute.
	GetLista() *arraylist.List

	// SetLista sets the lista attribute.
	SetLista(*arraylist.List)

	// IsInicioContext differentiates from other interfaces.
	IsInicioContext()
}

type InicioContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	lista          *arraylist.List
	_instrucciones IInstruccionesContext
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

func (s *InicioContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *InicioContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *InicioContext) GetLista() *arraylist.List { return s.lista }

func (s *InicioContext) SetLista(v *arraylist.List) { s.lista = v }

func (s *InicioContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *InicioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InicioContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InicioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterInicio(s)
	}
}

func (s *InicioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitInicio(s)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(30)

		var _x = p.Instrucciones()

		localctx.(*InicioContext)._instrucciones = _x
	}

	localctx.(*InicioContext).lista = localctx.(*InicioContext).Get_instrucciones().GetList()
	//fmt.Println(localctx.(*InicioContext).lista)

	return localctx
}

// IInstruccionesContext is an interface to support dynamic dispatch.
type IInstruccionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instruccion returns the _instruccion rule contexts.
	Get_instruccion() IInstruccionContext

	// Set_instruccion sets the _instruccion rule contexts.
	Set_instruccion(IInstruccionContext)

	// GetE returns the e rule context list.
	GetE() []IInstruccionContext

	// SetE sets the e rule context list.
	SetE([]IInstruccionContext)

	// GetList returns the list attribute.
	GetList() *arraylist.List

	// SetList sets the list attribute.
	SetList(*arraylist.List)

	// IsInstruccionesContext differentiates from other interfaces.
	IsInstruccionesContext()
}

type InstruccionesContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	list         *arraylist.List
	_instruccion IInstruccionContext
	e            []IInstruccionContext
}

func NewEmptyInstruccionesContext() *InstruccionesContext {
	var p = new(InstruccionesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_instrucciones
	return p
}

func (*InstruccionesContext) IsInstruccionesContext() {}

func NewInstruccionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionesContext {
	var p = new(InstruccionesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_instrucciones

	return p
}

func (s *InstruccionesContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionesContext) Get_instruccion() IInstruccionContext { return s._instruccion }

func (s *InstruccionesContext) Set_instruccion(v IInstruccionContext) { s._instruccion = v }

func (s *InstruccionesContext) GetE() []IInstruccionContext { return s.e }

func (s *InstruccionesContext) SetE(v []IInstruccionContext) { s.e = v }

func (s *InstruccionesContext) GetList() *arraylist.List { return s.list }

func (s *InstruccionesContext) SetList(v *arraylist.List) { s.list = v }

func (s *InstruccionesContext) AllInstruccion() []IInstruccionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstruccionContext)(nil)).Elem())
	var tst = make([]IInstruccionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstruccionContext)
		}
	}

	return tst
}

func (s *InstruccionesContext) Instruccion(i int) IInstruccionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstruccionContext)
}

func (s *InstruccionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstruccionesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterInstrucciones(s)
	}
}

func (s *InstruccionesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitInstrucciones(s)
	}
}

func (p *Nparser) Instrucciones() (localctx IInstruccionesContext) {
	this := p
	_ = this

	localctx = NewInstruccionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, NparserRULE_instrucciones)

	localctx.(*InstruccionesContext).list = arraylist.New()

	var _la int

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

	p.EnterOuterAlt(localctx, 1)
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NparserSENTENCIA)|(1<<NparserIF)|(1<<NparserDECLARAR)|(1<<NparserID))) != 0 {
		{
			p.SetState(33)

			var _x = p.Instruccion()

			localctx.(*InstruccionesContext)._instruccion = _x
		}
		localctx.(*InstruccionesContext).e = append(localctx.(*InstruccionesContext).e, localctx.(*InstruccionesContext)._instruccion)

		p.SetState(38)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listInt := localctx.(*InstruccionesContext).GetE()
	for _, e := range listInt {
		localctx.(*InstruccionesContext).list.Add(e.GetEx())
	}
	//fmt.Printf("tipo %T",localctx.(*InstruccionesContext).GetE())

	return localctx
}

// IInstruccionContext is an interface to support dynamic dispatch.
type IInstruccionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_control_IF_simple returns the _control_IF_simple rule contexts.
	Get_control_IF_simple() IControl_IF_simpleContext

	// Get_asignacion returns the _asignacion rule contexts.
	Get_asignacion() IAsignacionContext

	// Get_declaracion returns the _declaracion rule contexts.
	Get_declaracion() IDeclaracionContext

	// Get_funcion_imprimir returns the _funcion_imprimir rule contexts.
	Get_funcion_imprimir() IFuncion_imprimirContext

	// Set_control_IF_simple sets the _control_IF_simple rule contexts.
	Set_control_IF_simple(IControl_IF_simpleContext)

	// Set_asignacion sets the _asignacion rule contexts.
	Set_asignacion(IAsignacionContext)

	// Set_declaracion sets the _declaracion rule contexts.
	Set_declaracion(IDeclaracionContext)

	// Set_funcion_imprimir sets the _funcion_imprimir rule contexts.
	Set_funcion_imprimir(IFuncion_imprimirContext)

	// GetEx returns the ex attribute.
	GetEx() interface{}

	// SetEx sets the ex attribute.
	SetEx(interface{})

	// IsInstruccionContext differentiates from other interfaces.
	IsInstruccionContext()
}

type InstruccionContext struct {
	*antlr.BaseParserRuleContext
	parser             antlr.Parser
	ex                 interface{}
	_control_IF_simple IControl_IF_simpleContext
	_asignacion        IAsignacionContext
	_declaracion       IDeclaracionContext
	_funcion_imprimir  IFuncion_imprimirContext
}

func NewEmptyInstruccionContext() *InstruccionContext {
	var p = new(InstruccionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_instruccion
	return p
}

func (*InstruccionContext) IsInstruccionContext() {}

func NewInstruccionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionContext {
	var p = new(InstruccionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_instruccion

	return p
}

func (s *InstruccionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionContext) Get_control_IF_simple() IControl_IF_simpleContext {
	return s._control_IF_simple
}

func (s *InstruccionContext) Get_asignacion() IAsignacionContext { return s._asignacion }

func (s *InstruccionContext) Get_declaracion() IDeclaracionContext { return s._declaracion }

func (s *InstruccionContext) Get_funcion_imprimir() IFuncion_imprimirContext {
	return s._funcion_imprimir
}

func (s *InstruccionContext) Set_control_IF_simple(v IControl_IF_simpleContext) {
	s._control_IF_simple = v
}

func (s *InstruccionContext) Set_asignacion(v IAsignacionContext) { s._asignacion = v }

func (s *InstruccionContext) Set_declaracion(v IDeclaracionContext) { s._declaracion = v }

func (s *InstruccionContext) Set_funcion_imprimir(v IFuncion_imprimirContext) {
	s._funcion_imprimir = v
}

func (s *InstruccionContext) GetEx() interface{} { return s.ex }

func (s *InstruccionContext) SetEx(v interface{}) { s.ex = v }

func (s *InstruccionContext) Control_IF_simple() IControl_IF_simpleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IControl_IF_simpleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IControl_IF_simpleContext)
}

func (s *InstruccionContext) Asignacion() IAsignacionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAsignacionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAsignacionContext)
}

func (s *InstruccionContext) Declaracion() IDeclaracionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclaracionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclaracionContext)
}

func (s *InstruccionContext) Funcion_imprimir() IFuncion_imprimirContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncion_imprimirContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncion_imprimirContext)
}

func (s *InstruccionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstruccionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterInstruccion(s)
	}
}

func (s *InstruccionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitInstruccion(s)
	}
}

func (p *Nparser) Instruccion() (localctx IInstruccionContext) {
	this := p
	_ = this

	localctx = NewInstruccionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, NparserRULE_instruccion)

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

	p.SetState(53)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NparserIF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(41)

			var _x = p.Control_IF_simple()

			localctx.(*InstruccionContext)._control_IF_simple = _x
		}
		localctx.(*InstruccionContext).ex = localctx.(*InstruccionContext).Get_control_IF_simple().GetEx()

	case NparserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(44)

			var _x = p.Asignacion()

			localctx.(*InstruccionContext)._asignacion = _x
		}
		localctx.(*InstruccionContext).ex = localctx.(*InstruccionContext).Get_asignacion().GetEx()

	case NparserDECLARAR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(47)

			var _x = p.Declaracion()

			localctx.(*InstruccionContext)._declaracion = _x
		}
		localctx.(*InstruccionContext).ex = localctx.(*InstruccionContext).Get_declaracion().GetEx()

	case NparserSENTENCIA:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(50)

			var _x = p.Funcion_imprimir()

			localctx.(*InstruccionContext)._funcion_imprimir = _x
		}
		localctx.(*InstruccionContext).ex = localctx.(*InstruccionContext).Get_funcion_imprimir().GetEx()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDeclaracionContext is an interface to support dynamic dispatch.
type IDeclaracionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Get_MENOR returns the _MENOR token.
	Get_MENOR() antlr.Token

	// Get_DECLARAR returns the _DECLARAR token.
	Get_DECLARAR() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Set_MENOR sets the _MENOR token.
	Set_MENOR(antlr.Token)

	// Set_DECLARAR sets the _DECLARAR token.
	Set_DECLARAR(antlr.Token)

	// Get_tipo returns the _tipo rule contexts.
	Get_tipo() ITipoContext

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// Set_tipo sets the _tipo rule contexts.
	Set_tipo(ITipoContext)

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// GetEx returns the ex attribute.
	GetEx() Ast.Instruccion

	// SetEx sets the ex attribute.
	SetEx(Ast.Instruccion)

	// IsDeclaracionContext differentiates from other interfaces.
	IsDeclaracionContext()
}

type DeclaracionContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	ex         Ast.Instruccion
	_ID        antlr.Token
	_tipo      ITipoContext
	_MENOR     antlr.Token
	_expresion IExpresionContext
	_DECLARAR  antlr.Token
}

func NewEmptyDeclaracionContext() *DeclaracionContext {
	var p = new(DeclaracionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_declaracion
	return p
}

func (*DeclaracionContext) IsDeclaracionContext() {}

func NewDeclaracionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclaracionContext {
	var p = new(DeclaracionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_declaracion

	return p
}

func (s *DeclaracionContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclaracionContext) Get_ID() antlr.Token { return s._ID }

func (s *DeclaracionContext) Get_MENOR() antlr.Token { return s._MENOR }

func (s *DeclaracionContext) Get_DECLARAR() antlr.Token { return s._DECLARAR }

func (s *DeclaracionContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *DeclaracionContext) Set_MENOR(v antlr.Token) { s._MENOR = v }

func (s *DeclaracionContext) Set_DECLARAR(v antlr.Token) { s._DECLARAR = v }

func (s *DeclaracionContext) Get_tipo() ITipoContext { return s._tipo }

func (s *DeclaracionContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *DeclaracionContext) Set_tipo(v ITipoContext) { s._tipo = v }

func (s *DeclaracionContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *DeclaracionContext) GetEx() Ast.Instruccion { return s.ex }

func (s *DeclaracionContext) SetEx(v Ast.Instruccion) { s.ex = v }

func (s *DeclaracionContext) DECLARAR() antlr.TerminalNode {
	return s.GetToken(NparserDECLARAR, 0)
}

func (s *DeclaracionContext) ID() antlr.TerminalNode {
	return s.GetToken(NparserID, 0)
}

func (s *DeclaracionContext) Tipo() ITipoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITipoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *DeclaracionContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(NparserIGUAL, 0)
}

func (s *DeclaracionContext) MENOR() antlr.TerminalNode {
	return s.GetToken(NparserMENOR, 0)
}

func (s *DeclaracionContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *DeclaracionContext) MAYOR() antlr.TerminalNode {
	return s.GetToken(NparserMAYOR, 0)
}

func (s *DeclaracionContext) PUNTOCOMA() antlr.TerminalNode {
	return s.GetToken(NparserPUNTOCOMA, 0)
}

func (s *DeclaracionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclaracionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclaracionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterDeclaracion(s)
	}
}

func (s *DeclaracionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitDeclaracion(s)
	}
}

func (p *Nparser) Declaracion() (localctx IDeclaracionContext) {
	this := p
	_ = this

	localctx = NewDeclaracionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, NparserRULE_declaracion)

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

	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(55)
			p.Match(NparserDECLARAR)
		}
		{
			p.SetState(56)

			var _m = p.Match(NparserID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(57)

			var _x = p.Tipo()

			localctx.(*DeclaracionContext)._tipo = _x
		}
		{
			p.SetState(58)
			p.Match(NparserIGUAL)
		}
		{
			p.SetState(59)

			var _m = p.Match(NparserMENOR)

			localctx.(*DeclaracionContext)._MENOR = _m
		}
		{
			p.SetState(60)

			var _x = p.expresion(0)

			localctx.(*DeclaracionContext)._expresion = _x
		}

		fila := (func() int {
			if localctx.(*DeclaracionContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*DeclaracionContext).Get_ID().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*DeclaracionContext).Get_MENOR() == nil {
				return 0
			} else {
				return localctx.(*DeclaracionContext).Get_MENOR().GetColumn()
			}
		}())
		columna++
		localctx.(*DeclaracionContext).ex = Expresiones.NewDeclaracion((func() string {
			if localctx.(*DeclaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclaracionContext).Get_ID().GetText()
			}
		}()), localctx.(*DeclaracionContext).Get_tipo().GetEx(), localctx.(*DeclaracionContext).Get_expresion().GetEx(), fila, columna)

		{
			p.SetState(62)
			p.Match(NparserMAYOR)
		}
		{
			p.SetState(63)
			p.Match(NparserPUNTOCOMA)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(65)

			var _m = p.Match(NparserDECLARAR)

			localctx.(*DeclaracionContext)._DECLARAR = _m
		}
		{
			p.SetState(66)

			var _m = p.Match(NparserID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(67)

			var _x = p.Tipo()

			localctx.(*DeclaracionContext)._tipo = _x
		}

		fila := (func() int {
			if localctx.(*DeclaracionContext).Get_ID() == nil {
				return 0
			} else {
				return localctx.(*DeclaracionContext).Get_ID().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*DeclaracionContext).Get_DECLARAR() == nil {
				return 0
			} else {
				return localctx.(*DeclaracionContext).Get_DECLARAR().GetColumn()
			}
		}())
		columna++
		localctx.(*DeclaracionContext).ex = Expresiones.NewDeclaracion((func() string {
			if localctx.(*DeclaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclaracionContext).Get_ID().GetText()
			}
		}()), localctx.(*DeclaracionContext).Get_tipo().GetEx(), nil, fila, columna)

		{
			p.SetState(69)
			p.Match(NparserPUNTOCOMA)
		}

	}

	return localctx
}

// IAsignacionContext is an interface to support dynamic dispatch.
type IAsignacionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Get_MENOR returns the _MENOR token.
	Get_MENOR() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Set_MENOR sets the _MENOR token.
	Set_MENOR(antlr.Token)

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// GetEx returns the ex attribute.
	GetEx() Ast.Instruccion

	// SetEx sets the ex attribute.
	SetEx(Ast.Instruccion)

	// IsAsignacionContext differentiates from other interfaces.
	IsAsignacionContext()
}

type AsignacionContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	ex         Ast.Instruccion
	_ID        antlr.Token
	_MENOR     antlr.Token
	_expresion IExpresionContext
}

func NewEmptyAsignacionContext() *AsignacionContext {
	var p = new(AsignacionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_asignacion
	return p
}

func (*AsignacionContext) IsAsignacionContext() {}

func NewAsignacionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsignacionContext {
	var p = new(AsignacionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_asignacion

	return p
}

func (s *AsignacionContext) GetParser() antlr.Parser { return s.parser }

func (s *AsignacionContext) Get_ID() antlr.Token { return s._ID }

func (s *AsignacionContext) Get_MENOR() antlr.Token { return s._MENOR }

func (s *AsignacionContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *AsignacionContext) Set_MENOR(v antlr.Token) { s._MENOR = v }

func (s *AsignacionContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *AsignacionContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *AsignacionContext) GetEx() Ast.Instruccion { return s.ex }

func (s *AsignacionContext) SetEx(v Ast.Instruccion) { s.ex = v }

func (s *AsignacionContext) ID() antlr.TerminalNode {
	return s.GetToken(NparserID, 0)
}

func (s *AsignacionContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(NparserIGUAL, 0)
}

func (s *AsignacionContext) MENOR() antlr.TerminalNode {
	return s.GetToken(NparserMENOR, 0)
}

func (s *AsignacionContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *AsignacionContext) MAYOR() antlr.TerminalNode {
	return s.GetToken(NparserMAYOR, 0)
}

func (s *AsignacionContext) PUNTOCOMA() antlr.TerminalNode {
	return s.GetToken(NparserPUNTOCOMA, 0)
}

func (s *AsignacionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsignacionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterAsignacion(s)
	}
}

func (s *AsignacionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitAsignacion(s)
	}
}

func (p *Nparser) Asignacion() (localctx IAsignacionContext) {
	this := p
	_ = this

	localctx = NewAsignacionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, NparserRULE_asignacion)

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(73)

		var _m = p.Match(NparserID)

		localctx.(*AsignacionContext)._ID = _m
	}
	{
		p.SetState(74)
		p.Match(NparserIGUAL)
	}
	{
		p.SetState(75)

		var _m = p.Match(NparserMENOR)

		localctx.(*AsignacionContext)._MENOR = _m
	}
	{
		p.SetState(76)

		var _x = p.expresion(0)

		localctx.(*AsignacionContext)._expresion = _x
	}

	fila := (func() int {
		if localctx.(*AsignacionContext).Get_ID() == nil {
			return 0
		} else {
			return localctx.(*AsignacionContext).Get_ID().GetLine()
		}
	}())
	columna := (func() int {
		if localctx.(*AsignacionContext).Get_MENOR() == nil {
			return 0
		} else {
			return localctx.(*AsignacionContext).Get_MENOR().GetColumn()
		}
	}())
	columna++
	localctx.(*AsignacionContext).ex = Expresiones.NewAsignacion((func() string {
		if localctx.(*AsignacionContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*AsignacionContext).Get_ID().GetText()
		}
	}()), localctx.(*AsignacionContext).Get_expresion().GetEx(), fila, columna)

	{
		p.SetState(78)
		p.Match(NparserMAYOR)
	}
	{
		p.SetState(79)
		p.Match(NparserPUNTOCOMA)
	}

	return localctx
}

// IExpresionContext is an interface to support dynamic dispatch.
type IExpresionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_RESTA returns the _RESTA token.
	Get_RESTA() antlr.Token

	// Get_NUMERO returns the _NUMERO token.
	Get_NUMERO() antlr.Token

	// Get_DECIMAL returns the _DECIMAL token.
	Get_DECIMAL() antlr.Token

	// Get_CADENA returns the _CADENA token.
	Get_CADENA() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// GetOperador returns the operador token.
	GetOperador() antlr.Token

	// Set_RESTA sets the _RESTA token.
	Set_RESTA(antlr.Token)

	// Set_NUMERO sets the _NUMERO token.
	Set_NUMERO(antlr.Token)

	// Set_DECIMAL sets the _DECIMAL token.
	Set_DECIMAL(antlr.Token)

	// Set_CADENA sets the _CADENA token.
	Set_CADENA(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// SetOperador sets the operador token.
	SetOperador(antlr.Token)

	// GetOp_left returns the op_left rule contexts.
	GetOp_left() IExpresionContext

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// GetOp_right returns the op_right rule contexts.
	GetOp_right() IExpresionContext

	// SetOp_left sets the op_left rule contexts.
	SetOp_left(IExpresionContext)

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// SetOp_right sets the op_right rule contexts.
	SetOp_right(IExpresionContext)

	// GetEx returns the ex attribute.
	GetEx() Ast.Expresion

	// SetEx sets the ex attribute.
	SetEx(Ast.Expresion)

	// IsExpresionContext differentiates from other interfaces.
	IsExpresionContext()
}

type ExpresionContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	ex         Ast.Expresion
	op_left    IExpresionContext
	_RESTA     antlr.Token
	_expresion IExpresionContext
	_NUMERO    antlr.Token
	_DECIMAL   antlr.Token
	_CADENA    antlr.Token
	_ID        antlr.Token
	operador   antlr.Token
	op_right   IExpresionContext
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

func (s *ExpresionContext) Get_RESTA() antlr.Token { return s._RESTA }

func (s *ExpresionContext) Get_NUMERO() antlr.Token { return s._NUMERO }

func (s *ExpresionContext) Get_DECIMAL() antlr.Token { return s._DECIMAL }

func (s *ExpresionContext) Get_CADENA() antlr.Token { return s._CADENA }

func (s *ExpresionContext) Get_ID() antlr.Token { return s._ID }

func (s *ExpresionContext) GetOperador() antlr.Token { return s.operador }

func (s *ExpresionContext) Set_RESTA(v antlr.Token) { s._RESTA = v }

func (s *ExpresionContext) Set_NUMERO(v antlr.Token) { s._NUMERO = v }

func (s *ExpresionContext) Set_DECIMAL(v antlr.Token) { s._DECIMAL = v }

func (s *ExpresionContext) Set_CADENA(v antlr.Token) { s._CADENA = v }

func (s *ExpresionContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ExpresionContext) SetOperador(v antlr.Token) { s.operador = v }

func (s *ExpresionContext) GetOp_left() IExpresionContext { return s.op_left }

func (s *ExpresionContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *ExpresionContext) GetOp_right() IExpresionContext { return s.op_right }

func (s *ExpresionContext) SetOp_left(v IExpresionContext) { s.op_left = v }

func (s *ExpresionContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *ExpresionContext) SetOp_right(v IExpresionContext) { s.op_right = v }

func (s *ExpresionContext) GetEx() Ast.Expresion { return s.ex }

func (s *ExpresionContext) SetEx(v Ast.Expresion) { s.ex = v }

func (s *ExpresionContext) RESTA() antlr.TerminalNode {
	return s.GetToken(NparserRESTA, 0)
}

func (s *ExpresionContext) AllExpresion() []IExpresionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpresionContext)(nil)).Elem())
	var tst = make([]IExpresionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpresionContext)
		}
	}

	return tst
}

func (s *ExpresionContext) Expresion(i int) IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ExpresionContext) NUMERO() antlr.TerminalNode {
	return s.GetToken(NparserNUMERO, 0)
}

func (s *ExpresionContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(NparserDECIMAL, 0)
}

func (s *ExpresionContext) CADENA() antlr.TerminalNode {
	return s.GetToken(NparserCADENA, 0)
}

func (s *ExpresionContext) TRUE() antlr.TerminalNode {
	return s.GetToken(NparserTRUE, 0)
}

func (s *ExpresionContext) FALSE() antlr.TerminalNode {
	return s.GetToken(NparserFALSE, 0)
}

func (s *ExpresionContext) ID() antlr.TerminalNode {
	return s.GetToken(NparserID, 0)
}

func (s *ExpresionContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(NparserPAR_IZQ, 0)
}

func (s *ExpresionContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(NparserPAR_DER, 0)
}

func (s *ExpresionContext) MULTIPLICACION() antlr.TerminalNode {
	return s.GetToken(NparserMULTIPLICACION, 0)
}

func (s *ExpresionContext) DIVISION() antlr.TerminalNode {
	return s.GetToken(NparserDIVISION, 0)
}

func (s *ExpresionContext) SUMA() antlr.TerminalNode {
	return s.GetToken(NparserSUMA, 0)
}

func (s *ExpresionContext) MAYOR() antlr.TerminalNode {
	return s.GetToken(NparserMAYOR, 0)
}

func (s *ExpresionContext) MAYOR_I() antlr.TerminalNode {
	return s.GetToken(NparserMAYOR_I, 0)
}

func (s *ExpresionContext) AND() antlr.TerminalNode {
	return s.GetToken(NparserAND, 0)
}

func (s *ExpresionContext) OR() antlr.TerminalNode {
	return s.GetToken(NparserOR, 0)
}

func (s *ExpresionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpresionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpresionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterExpresion(s)
	}
}

func (s *ExpresionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitExpresion(s)
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
	_startState := 10
	p.EnterRecursionRule(localctx, 10, NparserRULE_expresion, _p)
	var _la int

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
	p.SetState(103)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NparserRESTA:
		{
			p.SetState(82)

			var _m = p.Match(NparserRESTA)

			localctx.(*ExpresionContext)._RESTA = _m
		}
		{
			p.SetState(83)

			var _x = p.expresion(12)

			localctx.(*ExpresionContext).op_left = _x
			localctx.(*ExpresionContext)._expresion = _x
		}

		fila := (func() int {
			if localctx.(*ExpresionContext).Get_RESTA() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).Get_RESTA().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*ExpresionContext).Get_RESTA() == nil {
				return 0
			} else {
				return localctx.(*ExpresionContext).Get_RESTA().GetColumn()
			}
		}())
		columna++
		operador := (func() string {
			if localctx.(*ExpresionContext).Get_RESTA() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).Get_RESTA().GetText()
			}
		}())
		localctx.(*ExpresionContext).ex = Expresiones.NewOperation(localctx.(*ExpresionContext).GetOp_left().GetEx(), operador, nil, true, fila, columna)

	case NparserNUMERO:
		{
			p.SetState(86)

			var _m = p.Match(NparserNUMERO)

			localctx.(*ExpresionContext)._NUMERO = _m
		}

		num, err := strconv.Atoi((func() string {
			if localctx.(*ExpresionContext).Get_NUMERO() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).Get_NUMERO().GetText()
			}
		}()))
		if err != nil {
			fmt.Println("Hay un error en el get número")
			fmt.Println(err)
		}
		localctx.(*ExpresionContext).ex = Expresiones.NewPrimitivo(num, Ast.INTEGER)

	case NparserDECIMAL:
		{
			p.SetState(88)

			var _m = p.Match(NparserDECIMAL)

			localctx.(*ExpresionContext)._DECIMAL = _m
		}

		num, err := strconv.ParseFloat((func() string {
			if localctx.(*ExpresionContext).Get_DECIMAL() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).Get_DECIMAL().GetText()
			}
		}()), 64)
		if err != nil {
			fmt.Println("Hay un error en el get número")
			fmt.Println(err)
		}
		localctx.(*ExpresionContext).ex = Expresiones.NewPrimitivo(num, Ast.REAL)

	case NparserCADENA:
		{
			p.SetState(90)

			var _m = p.Match(NparserCADENA)

			localctx.(*ExpresionContext)._CADENA = _m
		}

		texto := (func() string {
			if localctx.(*ExpresionContext).Get_CADENA() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).Get_CADENA().GetText()
			}
		}())
		texto = texto[1 : len(texto)-1]
		localctx.(*ExpresionContext).ex = Expresiones.NewPrimitivo(texto, Ast.STRING)

	case NparserTRUE:
		{
			p.SetState(92)
			p.Match(NparserTRUE)
		}

		valor := true
		localctx.(*ExpresionContext).ex = Expresiones.NewPrimitivo(valor, Ast.BOOLEAN)

	case NparserFALSE:
		{
			p.SetState(94)
			p.Match(NparserFALSE)
		}

		valor := false
		localctx.(*ExpresionContext).ex = Expresiones.NewPrimitivo(valor, Ast.BOOLEAN)

	case NparserID:
		{
			p.SetState(96)

			var _m = p.Match(NparserID)

			localctx.(*ExpresionContext)._ID = _m
		}

		id := (func() string {
			if localctx.(*ExpresionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).Get_ID().GetText()
			}
		}())
		localctx.(*ExpresionContext).ex = Expresiones.NewIdentificador(id, Ast.IDENTIFICADOR)

	case NparserPAR_IZQ:
		{
			p.SetState(98)
			p.Match(NparserPAR_IZQ)
		}
		{
			p.SetState(99)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext)._expresion = _x
		}
		{
			p.SetState(100)
			p.Match(NparserPAR_DER)
		}

		localctx.(*ExpresionContext).ex = localctx.(*ExpresionContext).Get_expresion().GetEx()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(125)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).op_left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, NparserRULE_expresion)
				p.SetState(105)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(106)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpresionContext).operador = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == NparserMULTIPLICACION || _la == NparserDIVISION) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpresionContext).operador = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(107)

					var _x = p.expresion(12)

					localctx.(*ExpresionContext).op_right = _x
					localctx.(*ExpresionContext)._expresion = _x
				}

				fila := (func() int {
					if localctx.(*ExpresionContext).GetOperador() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).GetOperador().GetLine()
					}
				}())
				columna := (func() int {
					if localctx.(*ExpresionContext).GetOperador() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).GetOperador().GetColumn()
					}
				}())
				columna--
				localctx.(*ExpresionContext).ex = Expresiones.NewOperation(localctx.(*ExpresionContext).GetOp_left().GetEx(), (func() string {
					if localctx.(*ExpresionContext).GetOperador() == nil {
						return ""
					} else {
						return localctx.(*ExpresionContext).GetOperador().GetText()
					}
				}()), localctx.(*ExpresionContext).GetOp_right().GetEx(), false, fila, columna)

			case 2:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).op_left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, NparserRULE_expresion)
				p.SetState(110)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(111)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpresionContext).operador = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == NparserRESTA || _la == NparserSUMA) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpresionContext).operador = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(112)

					var _x = p.expresion(11)

					localctx.(*ExpresionContext).op_right = _x
					localctx.(*ExpresionContext)._expresion = _x
				}

				fila := (func() int {
					if localctx.(*ExpresionContext).GetOperador() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).GetOperador().GetLine()
					}
				}())
				columna := (func() int {
					if localctx.(*ExpresionContext).GetOperador() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).GetOperador().GetColumn()
					}
				}())
				columna--
				localctx.(*ExpresionContext).ex = Expresiones.NewOperation(localctx.(*ExpresionContext).GetOp_left().GetEx(), (func() string {
					if localctx.(*ExpresionContext).GetOperador() == nil {
						return ""
					} else {
						return localctx.(*ExpresionContext).GetOperador().GetText()
					}
				}()), localctx.(*ExpresionContext).GetOp_right().GetEx(), false, fila, columna)

			case 3:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).op_left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, NparserRULE_expresion)
				p.SetState(115)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(116)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpresionContext).operador = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == NparserMAYOR_I || _la == NparserMAYOR) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpresionContext).operador = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(117)

					var _x = p.expresion(10)

					localctx.(*ExpresionContext).op_right = _x
					localctx.(*ExpresionContext)._expresion = _x
				}

				fila := (func() int {
					if localctx.(*ExpresionContext).GetOperador() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).GetOperador().GetLine()
					}
				}())
				columna := (func() int {
					if localctx.(*ExpresionContext).GetOperador() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).GetOperador().GetColumn()
					}
				}())
				columna--
				localctx.(*ExpresionContext).ex = Expresiones.NewOperation(localctx.(*ExpresionContext).GetOp_left().GetEx(), (func() string {
					if localctx.(*ExpresionContext).GetOperador() == nil {
						return ""
					} else {
						return localctx.(*ExpresionContext).GetOperador().GetText()
					}
				}()), localctx.(*ExpresionContext).GetOp_right().GetEx(), false, fila, columna)

			case 4:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).op_left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, NparserRULE_expresion)
				p.SetState(120)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(121)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpresionContext).operador = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == NparserAND || _la == NparserOR) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpresionContext).operador = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(122)

					var _x = p.expresion(9)

					localctx.(*ExpresionContext).op_right = _x
					localctx.(*ExpresionContext)._expresion = _x
				}

				fila := (func() int {
					if localctx.(*ExpresionContext).GetOperador() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).GetOperador().GetLine()
					}
				}())
				columna := (func() int {
					if localctx.(*ExpresionContext).GetOperador() == nil {
						return 0
					} else {
						return localctx.(*ExpresionContext).GetOperador().GetColumn()
					}
				}())
				columna--
				operacion := strings.ToUpper((func() string {
					if localctx.(*ExpresionContext).GetOperador() == nil {
						return ""
					} else {
						return localctx.(*ExpresionContext).GetOperador().GetText()
					}
				}()))
				localctx.(*ExpresionContext).ex = Expresiones.NewOperation(localctx.(*ExpresionContext).GetOp_left().GetEx(), operacion, localctx.(*ExpresionContext).GetOp_right().GetEx(), false, fila, columna)

			}

		}
		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// ITipoContext is an interface to support dynamic dispatch.
type ITipoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetEx returns the ex attribute.
	GetEx() Ast.TipoDato

	// SetEx sets the ex attribute.
	SetEx(Ast.TipoDato)

	// IsTipoContext differentiates from other interfaces.
	IsTipoContext()
}

type TipoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	ex     Ast.TipoDato
}

func NewEmptyTipoContext() *TipoContext {
	var p = new(TipoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_tipo
	return p
}

func (*TipoContext) IsTipoContext() {}

func NewTipoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TipoContext {
	var p = new(TipoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_tipo

	return p
}

func (s *TipoContext) GetParser() antlr.Parser { return s.parser }

func (s *TipoContext) GetEx() Ast.TipoDato { return s.ex }

func (s *TipoContext) SetEx(v Ast.TipoDato) { s.ex = v }

func (s *TipoContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(NparserINTEGER, 0)
}

func (s *TipoContext) REAL() antlr.TerminalNode {
	return s.GetToken(NparserREAL, 0)
}

func (s *TipoContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(NparserBOOLEAN, 0)
}

func (s *TipoContext) STRING() antlr.TerminalNode {
	return s.GetToken(NparserSTRING, 0)
}

func (s *TipoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TipoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TipoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterTipo(s)
	}
}

func (s *TipoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitTipo(s)
	}
}

func (p *Nparser) Tipo() (localctx ITipoContext) {
	this := p
	_ = this

	localctx = NewTipoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, NparserRULE_tipo)

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

	p.SetState(138)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NparserINTEGER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(130)
			p.Match(NparserINTEGER)
		}
		localctx.(*TipoContext).ex = Ast.INTEGER

	case NparserREAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(132)
			p.Match(NparserREAL)
		}
		localctx.(*TipoContext).ex = Ast.REAL

	case NparserBOOLEAN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(134)
			p.Match(NparserBOOLEAN)
		}
		localctx.(*TipoContext).ex = Ast.BOOLEAN

	case NparserSTRING:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(136)
			p.Match(NparserSTRING)
		}
		localctx.(*TipoContext).ex = Ast.STRING

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IOperador_comparacionContext is an interface to support dynamic dispatch.
type IOperador_comparacionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_MAYOR returns the _MAYOR token.
	Get_MAYOR() antlr.Token

	// Get_MAYOR_I returns the _MAYOR_I token.
	Get_MAYOR_I() antlr.Token

	// Set_MAYOR sets the _MAYOR token.
	Set_MAYOR(antlr.Token)

	// Set_MAYOR_I sets the _MAYOR_I token.
	Set_MAYOR_I(antlr.Token)

	// GetOperador returns the operador attribute.
	GetOperador() string

	// SetOperador sets the operador attribute.
	SetOperador(string)

	// IsOperador_comparacionContext differentiates from other interfaces.
	IsOperador_comparacionContext()
}

type Operador_comparacionContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	operador string
	_MAYOR   antlr.Token
	_MAYOR_I antlr.Token
}

func NewEmptyOperador_comparacionContext() *Operador_comparacionContext {
	var p = new(Operador_comparacionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_operador_comparacion
	return p
}

func (*Operador_comparacionContext) IsOperador_comparacionContext() {}

func NewOperador_comparacionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Operador_comparacionContext {
	var p = new(Operador_comparacionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_operador_comparacion

	return p
}

func (s *Operador_comparacionContext) GetParser() antlr.Parser { return s.parser }

func (s *Operador_comparacionContext) Get_MAYOR() antlr.Token { return s._MAYOR }

func (s *Operador_comparacionContext) Get_MAYOR_I() antlr.Token { return s._MAYOR_I }

func (s *Operador_comparacionContext) Set_MAYOR(v antlr.Token) { s._MAYOR = v }

func (s *Operador_comparacionContext) Set_MAYOR_I(v antlr.Token) { s._MAYOR_I = v }

func (s *Operador_comparacionContext) GetOperador() string { return s.operador }

func (s *Operador_comparacionContext) SetOperador(v string) { s.operador = v }

func (s *Operador_comparacionContext) MAYOR() antlr.TerminalNode {
	return s.GetToken(NparserMAYOR, 0)
}

func (s *Operador_comparacionContext) MAYOR_I() antlr.TerminalNode {
	return s.GetToken(NparserMAYOR_I, 0)
}

func (s *Operador_comparacionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Operador_comparacionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Operador_comparacionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterOperador_comparacion(s)
	}
}

func (s *Operador_comparacionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitOperador_comparacion(s)
	}
}

func (p *Nparser) Operador_comparacion() (localctx IOperador_comparacionContext) {
	this := p
	_ = this

	localctx = NewOperador_comparacionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, NparserRULE_operador_comparacion)

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

	p.SetState(144)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NparserMAYOR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(140)

			var _m = p.Match(NparserMAYOR)

			localctx.(*Operador_comparacionContext)._MAYOR = _m
		}
		localctx.(*Operador_comparacionContext).operador = (func() string {
			if localctx.(*Operador_comparacionContext).Get_MAYOR() == nil {
				return ""
			} else {
				return localctx.(*Operador_comparacionContext).Get_MAYOR().GetText()
			}
		}())

	case NparserMAYOR_I:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(142)

			var _m = p.Match(NparserMAYOR_I)

			localctx.(*Operador_comparacionContext)._MAYOR_I = _m
		}
		localctx.(*Operador_comparacionContext).operador = (func() string {
			if localctx.(*Operador_comparacionContext).Get_MAYOR_I() == nil {
				return ""
			} else {
				return localctx.(*Operador_comparacionContext).Get_MAYOR_I().GetText()
			}
		}())

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IOperador_logicoContext is an interface to support dynamic dispatch.
type IOperador_logicoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_AND returns the _AND token.
	Get_AND() antlr.Token

	// Get_OR returns the _OR token.
	Get_OR() antlr.Token

	// Set_AND sets the _AND token.
	Set_AND(antlr.Token)

	// Set_OR sets the _OR token.
	Set_OR(antlr.Token)

	// GetOperador returns the operador attribute.
	GetOperador() string

	// SetOperador sets the operador attribute.
	SetOperador(string)

	// IsOperador_logicoContext differentiates from other interfaces.
	IsOperador_logicoContext()
}

type Operador_logicoContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	operador string
	_AND     antlr.Token
	_OR      antlr.Token
}

func NewEmptyOperador_logicoContext() *Operador_logicoContext {
	var p = new(Operador_logicoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_operador_logico
	return p
}

func (*Operador_logicoContext) IsOperador_logicoContext() {}

func NewOperador_logicoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Operador_logicoContext {
	var p = new(Operador_logicoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_operador_logico

	return p
}

func (s *Operador_logicoContext) GetParser() antlr.Parser { return s.parser }

func (s *Operador_logicoContext) Get_AND() antlr.Token { return s._AND }

func (s *Operador_logicoContext) Get_OR() antlr.Token { return s._OR }

func (s *Operador_logicoContext) Set_AND(v antlr.Token) { s._AND = v }

func (s *Operador_logicoContext) Set_OR(v antlr.Token) { s._OR = v }

func (s *Operador_logicoContext) GetOperador() string { return s.operador }

func (s *Operador_logicoContext) SetOperador(v string) { s.operador = v }

func (s *Operador_logicoContext) AND() antlr.TerminalNode {
	return s.GetToken(NparserAND, 0)
}

func (s *Operador_logicoContext) OR() antlr.TerminalNode {
	return s.GetToken(NparserOR, 0)
}

func (s *Operador_logicoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Operador_logicoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Operador_logicoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterOperador_logico(s)
	}
}

func (s *Operador_logicoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitOperador_logico(s)
	}
}

func (p *Nparser) Operador_logico() (localctx IOperador_logicoContext) {
	this := p
	_ = this

	localctx = NewOperador_logicoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, NparserRULE_operador_logico)

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

	p.SetState(150)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NparserAND:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(146)

			var _m = p.Match(NparserAND)

			localctx.(*Operador_logicoContext)._AND = _m
		}
		localctx.(*Operador_logicoContext).operador = (func() string {
			if localctx.(*Operador_logicoContext).Get_AND() == nil {
				return ""
			} else {
				return localctx.(*Operador_logicoContext).Get_AND().GetText()
			}
		}())

	case NparserOR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(148)

			var _m = p.Match(NparserOR)

			localctx.(*Operador_logicoContext)._OR = _m
		}
		localctx.(*Operador_logicoContext).operador = (func() string {
			if localctx.(*Operador_logicoContext).Get_OR() == nil {
				return ""
			} else {
				return localctx.(*Operador_logicoContext).Get_OR().GetText()
			}
		}())

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IOperador_aritmeticoContext is an interface to support dynamic dispatch.
type IOperador_aritmeticoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_MULTIPLICACION returns the _MULTIPLICACION token.
	Get_MULTIPLICACION() antlr.Token

	// Get_DIVISION returns the _DIVISION token.
	Get_DIVISION() antlr.Token

	// Get_MODULO returns the _MODULO token.
	Get_MODULO() antlr.Token

	// Get_SUMA returns the _SUMA token.
	Get_SUMA() antlr.Token

	// Get_RESTA returns the _RESTA token.
	Get_RESTA() antlr.Token

	// Set_MULTIPLICACION sets the _MULTIPLICACION token.
	Set_MULTIPLICACION(antlr.Token)

	// Set_DIVISION sets the _DIVISION token.
	Set_DIVISION(antlr.Token)

	// Set_MODULO sets the _MODULO token.
	Set_MODULO(antlr.Token)

	// Set_SUMA sets the _SUMA token.
	Set_SUMA(antlr.Token)

	// Set_RESTA sets the _RESTA token.
	Set_RESTA(antlr.Token)

	// GetOperador returns the operador attribute.
	GetOperador() string

	// SetOperador sets the operador attribute.
	SetOperador(string)

	// IsOperador_aritmeticoContext differentiates from other interfaces.
	IsOperador_aritmeticoContext()
}

type Operador_aritmeticoContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	operador        string
	_MULTIPLICACION antlr.Token
	_DIVISION       antlr.Token
	_MODULO         antlr.Token
	_SUMA           antlr.Token
	_RESTA          antlr.Token
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

func (s *Operador_aritmeticoContext) Get_MULTIPLICACION() antlr.Token { return s._MULTIPLICACION }

func (s *Operador_aritmeticoContext) Get_DIVISION() antlr.Token { return s._DIVISION }

func (s *Operador_aritmeticoContext) Get_MODULO() antlr.Token { return s._MODULO }

func (s *Operador_aritmeticoContext) Get_SUMA() antlr.Token { return s._SUMA }

func (s *Operador_aritmeticoContext) Get_RESTA() antlr.Token { return s._RESTA }

func (s *Operador_aritmeticoContext) Set_MULTIPLICACION(v antlr.Token) { s._MULTIPLICACION = v }

func (s *Operador_aritmeticoContext) Set_DIVISION(v antlr.Token) { s._DIVISION = v }

func (s *Operador_aritmeticoContext) Set_MODULO(v antlr.Token) { s._MODULO = v }

func (s *Operador_aritmeticoContext) Set_SUMA(v antlr.Token) { s._SUMA = v }

func (s *Operador_aritmeticoContext) Set_RESTA(v antlr.Token) { s._RESTA = v }

func (s *Operador_aritmeticoContext) GetOperador() string { return s.operador }

func (s *Operador_aritmeticoContext) SetOperador(v string) { s.operador = v }

func (s *Operador_aritmeticoContext) MULTIPLICACION() antlr.TerminalNode {
	return s.GetToken(NparserMULTIPLICACION, 0)
}

func (s *Operador_aritmeticoContext) DIVISION() antlr.TerminalNode {
	return s.GetToken(NparserDIVISION, 0)
}

func (s *Operador_aritmeticoContext) MODULO() antlr.TerminalNode {
	return s.GetToken(NparserMODULO, 0)
}

func (s *Operador_aritmeticoContext) SUMA() antlr.TerminalNode {
	return s.GetToken(NparserSUMA, 0)
}

func (s *Operador_aritmeticoContext) RESTA() antlr.TerminalNode {
	return s.GetToken(NparserRESTA, 0)
}

func (s *Operador_aritmeticoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Operador_aritmeticoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Operador_aritmeticoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterOperador_aritmetico(s)
	}
}

func (s *Operador_aritmeticoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitOperador_aritmetico(s)
	}
}

func (p *Nparser) Operador_aritmetico() (localctx IOperador_aritmeticoContext) {
	this := p
	_ = this

	localctx = NewOperador_aritmeticoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, NparserRULE_operador_aritmetico)

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

	p.SetState(162)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NparserMULTIPLICACION:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(152)

			var _m = p.Match(NparserMULTIPLICACION)

			localctx.(*Operador_aritmeticoContext)._MULTIPLICACION = _m
		}
		localctx.(*Operador_aritmeticoContext).operador = (func() string {
			if localctx.(*Operador_aritmeticoContext).Get_MULTIPLICACION() == nil {
				return ""
			} else {
				return localctx.(*Operador_aritmeticoContext).Get_MULTIPLICACION().GetText()
			}
		}())

	case NparserDIVISION:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(154)

			var _m = p.Match(NparserDIVISION)

			localctx.(*Operador_aritmeticoContext)._DIVISION = _m
		}
		localctx.(*Operador_aritmeticoContext).operador = (func() string {
			if localctx.(*Operador_aritmeticoContext).Get_DIVISION() == nil {
				return ""
			} else {
				return localctx.(*Operador_aritmeticoContext).Get_DIVISION().GetText()
			}
		}())

	case NparserMODULO:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(156)

			var _m = p.Match(NparserMODULO)

			localctx.(*Operador_aritmeticoContext)._MODULO = _m
		}
		localctx.(*Operador_aritmeticoContext).operador = (func() string {
			if localctx.(*Operador_aritmeticoContext).Get_MODULO() == nil {
				return ""
			} else {
				return localctx.(*Operador_aritmeticoContext).Get_MODULO().GetText()
			}
		}())

	case NparserSUMA:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(158)

			var _m = p.Match(NparserSUMA)

			localctx.(*Operador_aritmeticoContext)._SUMA = _m
		}
		localctx.(*Operador_aritmeticoContext).operador = (func() string {
			if localctx.(*Operador_aritmeticoContext).Get_SUMA() == nil {
				return ""
			} else {
				return localctx.(*Operador_aritmeticoContext).Get_SUMA().GetText()
			}
		}())

	case NparserRESTA:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(160)

			var _m = p.Match(NparserRESTA)

			localctx.(*Operador_aritmeticoContext)._RESTA = _m
		}
		localctx.(*Operador_aritmeticoContext).operador = (func() string {
			if localctx.(*Operador_aritmeticoContext).Get_RESTA() == nil {
				return ""
			} else {
				return localctx.(*Operador_aritmeticoContext).Get_RESTA().GetText()
			}
		}())

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IControl_IF_simpleContext is an interface to support dynamic dispatch.
type IControl_IF_simpleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_IF returns the _IF token.
	Get_IF() antlr.Token

	// Get_PAR_IZQ returns the _PAR_IZQ token.
	Get_PAR_IZQ() antlr.Token

	// Set_IF sets the _IF token.
	Set_IF(antlr.Token)

	// Set_PAR_IZQ sets the _PAR_IZQ token.
	Set_PAR_IZQ(antlr.Token)

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Get_control_else returns the _control_else rule contexts.
	Get_control_else() IControl_elseContext

	// Get_control_IF_entonces returns the _control_IF_entonces rule contexts.
	Get_control_IF_entonces() IControl_IF_entoncesContext

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// Set_control_else sets the _control_else rule contexts.
	Set_control_else(IControl_elseContext)

	// Set_control_IF_entonces sets the _control_IF_entonces rule contexts.
	Set_control_IF_entonces(IControl_IF_entoncesContext)

	// GetEx returns the ex attribute.
	GetEx() Ast.Instruccion

	// SetEx sets the ex attribute.
	SetEx(Ast.Instruccion)

	// IsControl_IF_simpleContext differentiates from other interfaces.
	IsControl_IF_simpleContext()
}

type Control_IF_simpleContext struct {
	*antlr.BaseParserRuleContext
	parser               antlr.Parser
	ex                   Ast.Instruccion
	_IF                  antlr.Token
	_PAR_IZQ             antlr.Token
	_expresion           IExpresionContext
	_instrucciones       IInstruccionesContext
	_control_else        IControl_elseContext
	_control_IF_entonces IControl_IF_entoncesContext
}

func NewEmptyControl_IF_simpleContext() *Control_IF_simpleContext {
	var p = new(Control_IF_simpleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_control_IF_simple
	return p
}

func (*Control_IF_simpleContext) IsControl_IF_simpleContext() {}

func NewControl_IF_simpleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Control_IF_simpleContext {
	var p = new(Control_IF_simpleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_control_IF_simple

	return p
}

func (s *Control_IF_simpleContext) GetParser() antlr.Parser { return s.parser }

func (s *Control_IF_simpleContext) Get_IF() antlr.Token { return s._IF }

func (s *Control_IF_simpleContext) Get_PAR_IZQ() antlr.Token { return s._PAR_IZQ }

func (s *Control_IF_simpleContext) Set_IF(v antlr.Token) { s._IF = v }

func (s *Control_IF_simpleContext) Set_PAR_IZQ(v antlr.Token) { s._PAR_IZQ = v }

func (s *Control_IF_simpleContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *Control_IF_simpleContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *Control_IF_simpleContext) Get_control_else() IControl_elseContext { return s._control_else }

func (s *Control_IF_simpleContext) Get_control_IF_entonces() IControl_IF_entoncesContext {
	return s._control_IF_entonces
}

func (s *Control_IF_simpleContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *Control_IF_simpleContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *Control_IF_simpleContext) Set_control_else(v IControl_elseContext) { s._control_else = v }

func (s *Control_IF_simpleContext) Set_control_IF_entonces(v IControl_IF_entoncesContext) {
	s._control_IF_entonces = v
}

func (s *Control_IF_simpleContext) GetEx() Ast.Instruccion { return s.ex }

func (s *Control_IF_simpleContext) SetEx(v Ast.Instruccion) { s.ex = v }

func (s *Control_IF_simpleContext) IF() antlr.TerminalNode {
	return s.GetToken(NparserIF, 0)
}

func (s *Control_IF_simpleContext) DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(NparserDOSPUNTOS, 0)
}

func (s *Control_IF_simpleContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(NparserPAR_IZQ, 0)
}

func (s *Control_IF_simpleContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Control_IF_simpleContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(NparserPAR_DER, 0)
}

func (s *Control_IF_simpleContext) CORCHETE_IZQ() antlr.TerminalNode {
	return s.GetToken(NparserCORCHETE_IZQ, 0)
}

func (s *Control_IF_simpleContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Control_IF_simpleContext) CORCHETE_DER() antlr.TerminalNode {
	return s.GetToken(NparserCORCHETE_DER, 0)
}

func (s *Control_IF_simpleContext) Control_else() IControl_elseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IControl_elseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IControl_elseContext)
}

func (s *Control_IF_simpleContext) PUNTOCOMA() antlr.TerminalNode {
	return s.GetToken(NparserPUNTOCOMA, 0)
}

func (s *Control_IF_simpleContext) Control_IF_entonces() IControl_IF_entoncesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IControl_IF_entoncesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IControl_IF_entoncesContext)
}

func (s *Control_IF_simpleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Control_IF_simpleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Control_IF_simpleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterControl_IF_simple(s)
	}
}

func (s *Control_IF_simpleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitControl_IF_simple(s)
	}
}

func (p *Nparser) Control_IF_simple() (localctx IControl_IF_simpleContext) {
	this := p
	_ = this

	localctx = NewControl_IF_simpleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, NparserRULE_control_IF_simple)

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

	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(164)

			var _m = p.Match(NparserIF)

			localctx.(*Control_IF_simpleContext)._IF = _m
		}
		{
			p.SetState(165)
			p.Match(NparserDOSPUNTOS)
		}
		{
			p.SetState(166)

			var _m = p.Match(NparserPAR_IZQ)

			localctx.(*Control_IF_simpleContext)._PAR_IZQ = _m
		}
		{
			p.SetState(167)

			var _x = p.expresion(0)

			localctx.(*Control_IF_simpleContext)._expresion = _x
		}
		{
			p.SetState(168)
			p.Match(NparserPAR_DER)
		}
		{
			p.SetState(169)
			p.Match(NparserCORCHETE_IZQ)
		}
		{
			p.SetState(170)

			var _x = p.Instrucciones()

			localctx.(*Control_IF_simpleContext)._instrucciones = _x
		}
		{
			p.SetState(171)
			p.Match(NparserCORCHETE_DER)
		}
		{
			p.SetState(172)

			var _x = p.Control_else()

			localctx.(*Control_IF_simpleContext)._control_else = _x
		}
		{
			p.SetState(173)
			p.Match(NparserPUNTOCOMA)
		}

		fila := (func() int {
			if localctx.(*Control_IF_simpleContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*Control_IF_simpleContext).Get_IF().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*Control_IF_simpleContext).Get_PAR_IZQ() == nil {
				return 0
			} else {
				return localctx.(*Control_IF_simpleContext).Get_PAR_IZQ().GetColumn()
			}
		}())
		columna++
		lista_entonces := arraylist.New()
		lista_entonces.Add(localctx.(*Control_IF_simpleContext).Get_control_else().GetEx())
		localctx.(*Control_IF_simpleContext).ex = Instrucciones.NewIF(localctx.(*Control_IF_simpleContext).Get_expresion().GetEx(), localctx.(*Control_IF_simpleContext).Get_instrucciones().GetList(), lista_entonces, Ast.IF, fila, columna)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(176)

			var _m = p.Match(NparserIF)

			localctx.(*Control_IF_simpleContext)._IF = _m
		}
		{
			p.SetState(177)
			p.Match(NparserDOSPUNTOS)
		}
		{
			p.SetState(178)

			var _m = p.Match(NparserPAR_IZQ)

			localctx.(*Control_IF_simpleContext)._PAR_IZQ = _m
		}
		{
			p.SetState(179)

			var _x = p.expresion(0)

			localctx.(*Control_IF_simpleContext)._expresion = _x
		}
		{
			p.SetState(180)
			p.Match(NparserPAR_DER)
		}
		{
			p.SetState(181)
			p.Match(NparserCORCHETE_IZQ)
		}
		{
			p.SetState(182)

			var _x = p.Instrucciones()

			localctx.(*Control_IF_simpleContext)._instrucciones = _x
		}
		{
			p.SetState(183)
			p.Match(NparserCORCHETE_DER)
		}
		{
			p.SetState(184)
			p.Match(NparserPUNTOCOMA)
		}

		fila := (func() int {
			if localctx.(*Control_IF_simpleContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*Control_IF_simpleContext).Get_IF().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*Control_IF_simpleContext).Get_PAR_IZQ() == nil {
				return 0
			} else {
				return localctx.(*Control_IF_simpleContext).Get_PAR_IZQ().GetColumn()
			}
		}())
		columna++
		lista_null := arraylist.New()
		localctx.(*Control_IF_simpleContext).ex = Instrucciones.NewIF(localctx.(*Control_IF_simpleContext).Get_expresion().GetEx(), localctx.(*Control_IF_simpleContext).Get_instrucciones().GetList(), lista_null, Ast.IF, fila, columna)

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(187)

			var _m = p.Match(NparserIF)

			localctx.(*Control_IF_simpleContext)._IF = _m
		}
		{
			p.SetState(188)
			p.Match(NparserDOSPUNTOS)
		}
		{
			p.SetState(189)

			var _m = p.Match(NparserPAR_IZQ)

			localctx.(*Control_IF_simpleContext)._PAR_IZQ = _m
		}
		{
			p.SetState(190)

			var _x = p.expresion(0)

			localctx.(*Control_IF_simpleContext)._expresion = _x
		}
		{
			p.SetState(191)
			p.Match(NparserPAR_DER)
		}
		{
			p.SetState(192)
			p.Match(NparserCORCHETE_IZQ)
		}
		{
			p.SetState(193)

			var _x = p.Instrucciones()

			localctx.(*Control_IF_simpleContext)._instrucciones = _x
		}
		{
			p.SetState(194)
			p.Match(NparserCORCHETE_DER)
		}
		{
			p.SetState(195)

			var _x = p.Control_IF_entonces()

			localctx.(*Control_IF_simpleContext)._control_IF_entonces = _x
		}
		{
			p.SetState(196)
			p.Match(NparserPUNTOCOMA)
		}

		fila := (func() int {
			if localctx.(*Control_IF_simpleContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*Control_IF_simpleContext).Get_IF().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*Control_IF_simpleContext).Get_PAR_IZQ() == nil {
				return 0
			} else {
				return localctx.(*Control_IF_simpleContext).Get_PAR_IZQ().GetColumn()
			}
		}())
		columna++
		lista_entonces := localctx.(*Control_IF_simpleContext).Get_control_IF_entonces().GetList()
		localctx.(*Control_IF_simpleContext).ex = Instrucciones.NewIF(localctx.(*Control_IF_simpleContext).Get_expresion().GetEx(), localctx.(*Control_IF_simpleContext).Get_instrucciones().GetList(), lista_entonces, Ast.IF, fila, columna)

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(199)

			var _m = p.Match(NparserIF)

			localctx.(*Control_IF_simpleContext)._IF = _m
		}
		{
			p.SetState(200)
			p.Match(NparserDOSPUNTOS)
		}
		{
			p.SetState(201)

			var _m = p.Match(NparserPAR_IZQ)

			localctx.(*Control_IF_simpleContext)._PAR_IZQ = _m
		}
		{
			p.SetState(202)

			var _x = p.expresion(0)

			localctx.(*Control_IF_simpleContext)._expresion = _x
		}
		{
			p.SetState(203)
			p.Match(NparserPAR_DER)
		}
		{
			p.SetState(204)
			p.Match(NparserCORCHETE_IZQ)
		}
		{
			p.SetState(205)

			var _x = p.Instrucciones()

			localctx.(*Control_IF_simpleContext)._instrucciones = _x
		}
		{
			p.SetState(206)
			p.Match(NparserCORCHETE_DER)
		}
		{
			p.SetState(207)

			var _x = p.Control_IF_entonces()

			localctx.(*Control_IF_simpleContext)._control_IF_entonces = _x
		}
		{
			p.SetState(208)

			var _x = p.Control_else()

			localctx.(*Control_IF_simpleContext)._control_else = _x
		}
		{
			p.SetState(209)
			p.Match(NparserPUNTOCOMA)
		}

		fila := (func() int {
			if localctx.(*Control_IF_simpleContext).Get_IF() == nil {
				return 0
			} else {
				return localctx.(*Control_IF_simpleContext).Get_IF().GetLine()
			}
		}())
		columna := (func() int {
			if localctx.(*Control_IF_simpleContext).Get_PAR_IZQ() == nil {
				return 0
			} else {
				return localctx.(*Control_IF_simpleContext).Get_PAR_IZQ().GetColumn()
			}
		}())
		columna++
		lista_entonces := localctx.(*Control_IF_simpleContext).Get_control_IF_entonces().GetList()
		lista_entonces.Add(localctx.(*Control_IF_simpleContext).Get_control_else().GetEx())
		localctx.(*Control_IF_simpleContext).ex = Instrucciones.NewIF(localctx.(*Control_IF_simpleContext).Get_expresion().GetEx(), localctx.(*Control_IF_simpleContext).Get_instrucciones().GetList(), lista_entonces, Ast.IF, fila, columna)

	}

	return localctx
}

// IControl_IF_entoncesContext is an interface to support dynamic dispatch.
type IControl_IF_entoncesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_control_if_entonces returns the _control_if_entonces rule contexts.
	Get_control_if_entonces() IControl_if_entoncesContext

	// Set_control_if_entonces sets the _control_if_entonces rule contexts.
	Set_control_if_entonces(IControl_if_entoncesContext)

	// GetE returns the e rule context list.
	GetE() []IControl_if_entoncesContext

	// SetE sets the e rule context list.
	SetE([]IControl_if_entoncesContext)

	// GetList returns the list attribute.
	GetList() *arraylist.List

	// SetList sets the list attribute.
	SetList(*arraylist.List)

	// IsControl_IF_entoncesContext differentiates from other interfaces.
	IsControl_IF_entoncesContext()
}

type Control_IF_entoncesContext struct {
	*antlr.BaseParserRuleContext
	parser               antlr.Parser
	list                 *arraylist.List
	_control_if_entonces IControl_if_entoncesContext
	e                    []IControl_if_entoncesContext
}

func NewEmptyControl_IF_entoncesContext() *Control_IF_entoncesContext {
	var p = new(Control_IF_entoncesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_control_IF_entonces
	return p
}

func (*Control_IF_entoncesContext) IsControl_IF_entoncesContext() {}

func NewControl_IF_entoncesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Control_IF_entoncesContext {
	var p = new(Control_IF_entoncesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_control_IF_entonces

	return p
}

func (s *Control_IF_entoncesContext) GetParser() antlr.Parser { return s.parser }

func (s *Control_IF_entoncesContext) Get_control_if_entonces() IControl_if_entoncesContext {
	return s._control_if_entonces
}

func (s *Control_IF_entoncesContext) Set_control_if_entonces(v IControl_if_entoncesContext) {
	s._control_if_entonces = v
}

func (s *Control_IF_entoncesContext) GetE() []IControl_if_entoncesContext { return s.e }

func (s *Control_IF_entoncesContext) SetE(v []IControl_if_entoncesContext) { s.e = v }

func (s *Control_IF_entoncesContext) GetList() *arraylist.List { return s.list }

func (s *Control_IF_entoncesContext) SetList(v *arraylist.List) { s.list = v }

func (s *Control_IF_entoncesContext) AllControl_if_entonces() []IControl_if_entoncesContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IControl_if_entoncesContext)(nil)).Elem())
	var tst = make([]IControl_if_entoncesContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IControl_if_entoncesContext)
		}
	}

	return tst
}

func (s *Control_IF_entoncesContext) Control_if_entonces(i int) IControl_if_entoncesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IControl_if_entoncesContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IControl_if_entoncesContext)
}

func (s *Control_IF_entoncesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Control_IF_entoncesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Control_IF_entoncesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterControl_IF_entonces(s)
	}
}

func (s *Control_IF_entoncesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitControl_IF_entonces(s)
	}
}

func (p *Nparser) Control_IF_entonces() (localctx IControl_IF_entoncesContext) {
	this := p
	_ = this

	localctx = NewControl_IF_entoncesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, NparserRULE_control_IF_entonces)

	localctx.(*Control_IF_entoncesContext).list = arraylist.New()

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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(217)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(214)

				var _x = p.Control_if_entonces()

				localctx.(*Control_IF_entoncesContext)._control_if_entonces = _x
			}
			localctx.(*Control_IF_entoncesContext).e = append(localctx.(*Control_IF_entoncesContext).e, localctx.(*Control_IF_entoncesContext)._control_if_entonces)

		}
		p.SetState(219)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
	}

	listInt := localctx.(*Control_IF_entoncesContext).GetE()
	for _, e := range listInt {
		localctx.(*Control_IF_entoncesContext).list.Add(e.GetEx())
	}
	//fmt.Printf("tipo %T",localctx.(*Control_IF_entoncesContext).GetE())

	return localctx
}

// IControl_if_entoncesContext is an interface to support dynamic dispatch.
type IControl_if_entoncesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ENTONCES returns the _ENTONCES token.
	Get_ENTONCES() antlr.Token

	// Get_PAR_IZQ returns the _PAR_IZQ token.
	Get_PAR_IZQ() antlr.Token

	// Set_ENTONCES sets the _ENTONCES token.
	Set_ENTONCES(antlr.Token)

	// Set_PAR_IZQ sets the _PAR_IZQ token.
	Set_PAR_IZQ(antlr.Token)

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetEx returns the ex attribute.
	GetEx() Ast.Instruccion

	// SetEx sets the ex attribute.
	SetEx(Ast.Instruccion)

	// IsControl_if_entoncesContext differentiates from other interfaces.
	IsControl_if_entoncesContext()
}

type Control_if_entoncesContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	ex             Ast.Instruccion
	_ENTONCES      antlr.Token
	_PAR_IZQ       antlr.Token
	_expresion     IExpresionContext
	_instrucciones IInstruccionesContext
}

func NewEmptyControl_if_entoncesContext() *Control_if_entoncesContext {
	var p = new(Control_if_entoncesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_control_if_entonces
	return p
}

func (*Control_if_entoncesContext) IsControl_if_entoncesContext() {}

func NewControl_if_entoncesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Control_if_entoncesContext {
	var p = new(Control_if_entoncesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_control_if_entonces

	return p
}

func (s *Control_if_entoncesContext) GetParser() antlr.Parser { return s.parser }

func (s *Control_if_entoncesContext) Get_ENTONCES() antlr.Token { return s._ENTONCES }

func (s *Control_if_entoncesContext) Get_PAR_IZQ() antlr.Token { return s._PAR_IZQ }

func (s *Control_if_entoncesContext) Set_ENTONCES(v antlr.Token) { s._ENTONCES = v }

func (s *Control_if_entoncesContext) Set_PAR_IZQ(v antlr.Token) { s._PAR_IZQ = v }

func (s *Control_if_entoncesContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *Control_if_entoncesContext) Get_instrucciones() IInstruccionesContext {
	return s._instrucciones
}

func (s *Control_if_entoncesContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *Control_if_entoncesContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *Control_if_entoncesContext) GetEx() Ast.Instruccion { return s.ex }

func (s *Control_if_entoncesContext) SetEx(v Ast.Instruccion) { s.ex = v }

func (s *Control_if_entoncesContext) ENTONCES() antlr.TerminalNode {
	return s.GetToken(NparserENTONCES, 0)
}

func (s *Control_if_entoncesContext) IF() antlr.TerminalNode {
	return s.GetToken(NparserIF, 0)
}

func (s *Control_if_entoncesContext) DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(NparserDOSPUNTOS, 0)
}

func (s *Control_if_entoncesContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(NparserPAR_IZQ, 0)
}

func (s *Control_if_entoncesContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Control_if_entoncesContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(NparserPAR_DER, 0)
}

func (s *Control_if_entoncesContext) CORCHETE_IZQ() antlr.TerminalNode {
	return s.GetToken(NparserCORCHETE_IZQ, 0)
}

func (s *Control_if_entoncesContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Control_if_entoncesContext) CORCHETE_DER() antlr.TerminalNode {
	return s.GetToken(NparserCORCHETE_DER, 0)
}

func (s *Control_if_entoncesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Control_if_entoncesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Control_if_entoncesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterControl_if_entonces(s)
	}
}

func (s *Control_if_entoncesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitControl_if_entonces(s)
	}
}

func (p *Nparser) Control_if_entonces() (localctx IControl_if_entoncesContext) {
	this := p
	_ = this

	localctx = NewControl_if_entoncesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, NparserRULE_control_if_entonces)

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(222)

		var _m = p.Match(NparserENTONCES)

		localctx.(*Control_if_entoncesContext)._ENTONCES = _m
	}
	{
		p.SetState(223)
		p.Match(NparserIF)
	}
	{
		p.SetState(224)
		p.Match(NparserDOSPUNTOS)
	}
	{
		p.SetState(225)

		var _m = p.Match(NparserPAR_IZQ)

		localctx.(*Control_if_entoncesContext)._PAR_IZQ = _m
	}
	{
		p.SetState(226)

		var _x = p.expresion(0)

		localctx.(*Control_if_entoncesContext)._expresion = _x
	}
	{
		p.SetState(227)
		p.Match(NparserPAR_DER)
	}
	{
		p.SetState(228)
		p.Match(NparserCORCHETE_IZQ)
	}
	{
		p.SetState(229)

		var _x = p.Instrucciones()

		localctx.(*Control_if_entoncesContext)._instrucciones = _x
	}
	{
		p.SetState(230)
		p.Match(NparserCORCHETE_DER)
	}

	fila := (func() int {
		if localctx.(*Control_if_entoncesContext).Get_ENTONCES() == nil {
			return 0
		} else {
			return localctx.(*Control_if_entoncesContext).Get_ENTONCES().GetLine()
		}
	}())
	columna := (func() int {
		if localctx.(*Control_if_entoncesContext).Get_PAR_IZQ() == nil {
			return 0
		} else {
			return localctx.(*Control_if_entoncesContext).Get_PAR_IZQ().GetColumn()
		}
	}())
	columna++
	lista_null := arraylist.New()
	localctx.(*Control_if_entoncesContext).ex = Instrucciones.NewIF(localctx.(*Control_if_entoncesContext).Get_expresion().GetEx(), localctx.(*Control_if_entoncesContext).Get_instrucciones().GetList(), lista_null, Ast.IF_ENTONCES, fila, columna)

	return localctx
}

// IControl_elseContext is an interface to support dynamic dispatch.
type IControl_elseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ENTONCES returns the _ENTONCES token.
	Get_ENTONCES() antlr.Token

	// Get_CORCHETE_IZQ returns the _CORCHETE_IZQ token.
	Get_CORCHETE_IZQ() antlr.Token

	// Set_ENTONCES sets the _ENTONCES token.
	Set_ENTONCES(antlr.Token)

	// Set_CORCHETE_IZQ sets the _CORCHETE_IZQ token.
	Set_CORCHETE_IZQ(antlr.Token)

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetEx returns the ex attribute.
	GetEx() Ast.Instruccion

	// SetEx sets the ex attribute.
	SetEx(Ast.Instruccion)

	// IsControl_elseContext differentiates from other interfaces.
	IsControl_elseContext()
}

type Control_elseContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	ex             Ast.Instruccion
	_ENTONCES      antlr.Token
	_CORCHETE_IZQ  antlr.Token
	_instrucciones IInstruccionesContext
}

func NewEmptyControl_elseContext() *Control_elseContext {
	var p = new(Control_elseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_control_else
	return p
}

func (*Control_elseContext) IsControl_elseContext() {}

func NewControl_elseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Control_elseContext {
	var p = new(Control_elseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_control_else

	return p
}

func (s *Control_elseContext) GetParser() antlr.Parser { return s.parser }

func (s *Control_elseContext) Get_ENTONCES() antlr.Token { return s._ENTONCES }

func (s *Control_elseContext) Get_CORCHETE_IZQ() antlr.Token { return s._CORCHETE_IZQ }

func (s *Control_elseContext) Set_ENTONCES(v antlr.Token) { s._ENTONCES = v }

func (s *Control_elseContext) Set_CORCHETE_IZQ(v antlr.Token) { s._CORCHETE_IZQ = v }

func (s *Control_elseContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *Control_elseContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *Control_elseContext) GetEx() Ast.Instruccion { return s.ex }

func (s *Control_elseContext) SetEx(v Ast.Instruccion) { s.ex = v }

func (s *Control_elseContext) ENTONCES() antlr.TerminalNode {
	return s.GetToken(NparserENTONCES, 0)
}

func (s *Control_elseContext) CORCHETE_IZQ() antlr.TerminalNode {
	return s.GetToken(NparserCORCHETE_IZQ, 0)
}

func (s *Control_elseContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Control_elseContext) CORCHETE_DER() antlr.TerminalNode {
	return s.GetToken(NparserCORCHETE_DER, 0)
}

func (s *Control_elseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Control_elseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Control_elseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterControl_else(s)
	}
}

func (s *Control_elseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitControl_else(s)
	}
}

func (p *Nparser) Control_else() (localctx IControl_elseContext) {
	this := p
	_ = this

	localctx = NewControl_elseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, NparserRULE_control_else)

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(233)

		var _m = p.Match(NparserENTONCES)

		localctx.(*Control_elseContext)._ENTONCES = _m
	}
	{
		p.SetState(234)

		var _m = p.Match(NparserCORCHETE_IZQ)

		localctx.(*Control_elseContext)._CORCHETE_IZQ = _m
	}
	{
		p.SetState(235)

		var _x = p.Instrucciones()

		localctx.(*Control_elseContext)._instrucciones = _x
	}

	fila := (func() int {
		if localctx.(*Control_elseContext).Get_ENTONCES() == nil {
			return 0
		} else {
			return localctx.(*Control_elseContext).Get_ENTONCES().GetLine()
		}
	}())
	columna := (func() int {
		if localctx.(*Control_elseContext).Get_CORCHETE_IZQ() == nil {
			return 0
		} else {
			return localctx.(*Control_elseContext).Get_CORCHETE_IZQ().GetColumn()
		}
	}())
	columna++
	expresionNull := Expresiones.NewPrimitivo(nil, Ast.NULL)
	lista_null := arraylist.New()
	localctx.(*Control_elseContext).ex = Instrucciones.NewIF(expresionNull, localctx.(*Control_elseContext).Get_instrucciones().GetList(), lista_null, Ast.ENTONCES, fila, columna)

	{
		p.SetState(237)
		p.Match(NparserCORCHETE_DER)
	}

	return localctx
}

// IFuncion_imprimirContext is an interface to support dynamic dispatch.
type IFuncion_imprimirContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_SENTENCIA returns the _SENTENCIA token.
	Get_SENTENCIA() antlr.Token

	// Get_PAR_IZQ returns the _PAR_IZQ token.
	Get_PAR_IZQ() antlr.Token

	// Set_SENTENCIA sets the _SENTENCIA token.
	Set_SENTENCIA(antlr.Token)

	// Set_PAR_IZQ sets the _PAR_IZQ token.
	Set_PAR_IZQ(antlr.Token)

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// GetEx returns the ex attribute.
	GetEx() Ast.Instruccion

	// SetEx sets the ex attribute.
	SetEx(Ast.Instruccion)

	// IsFuncion_imprimirContext differentiates from other interfaces.
	IsFuncion_imprimirContext()
}

type Funcion_imprimirContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	ex         Ast.Instruccion
	_SENTENCIA antlr.Token
	_PAR_IZQ   antlr.Token
	_expresion IExpresionContext
}

func NewEmptyFuncion_imprimirContext() *Funcion_imprimirContext {
	var p = new(Funcion_imprimirContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NparserRULE_funcion_imprimir
	return p
}

func (*Funcion_imprimirContext) IsFuncion_imprimirContext() {}

func NewFuncion_imprimirContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Funcion_imprimirContext {
	var p = new(Funcion_imprimirContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NparserRULE_funcion_imprimir

	return p
}

func (s *Funcion_imprimirContext) GetParser() antlr.Parser { return s.parser }

func (s *Funcion_imprimirContext) Get_SENTENCIA() antlr.Token { return s._SENTENCIA }

func (s *Funcion_imprimirContext) Get_PAR_IZQ() antlr.Token { return s._PAR_IZQ }

func (s *Funcion_imprimirContext) Set_SENTENCIA(v antlr.Token) { s._SENTENCIA = v }

func (s *Funcion_imprimirContext) Set_PAR_IZQ(v antlr.Token) { s._PAR_IZQ = v }

func (s *Funcion_imprimirContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *Funcion_imprimirContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *Funcion_imprimirContext) GetEx() Ast.Instruccion { return s.ex }

func (s *Funcion_imprimirContext) SetEx(v Ast.Instruccion) { s.ex = v }

func (s *Funcion_imprimirContext) SENTENCIA() antlr.TerminalNode {
	return s.GetToken(NparserSENTENCIA, 0)
}

func (s *Funcion_imprimirContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(NparserPUNTO, 0)
}

func (s *Funcion_imprimirContext) CONSOLA() antlr.TerminalNode {
	return s.GetToken(NparserCONSOLA, 0)
}

func (s *Funcion_imprimirContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(NparserPAR_IZQ, 0)
}

func (s *Funcion_imprimirContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Funcion_imprimirContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(NparserPAR_DER, 0)
}

func (s *Funcion_imprimirContext) PUNTOCOMA() antlr.TerminalNode {
	return s.GetToken(NparserPUNTOCOMA, 0)
}

func (s *Funcion_imprimirContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Funcion_imprimirContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Funcion_imprimirContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.EnterFuncion_imprimir(s)
	}
}

func (s *Funcion_imprimirContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NparserListener); ok {
		listenerT.ExitFuncion_imprimir(s)
	}
}

func (p *Nparser) Funcion_imprimir() (localctx IFuncion_imprimirContext) {
	this := p
	_ = this

	localctx = NewFuncion_imprimirContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, NparserRULE_funcion_imprimir)

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(239)

		var _m = p.Match(NparserSENTENCIA)

		localctx.(*Funcion_imprimirContext)._SENTENCIA = _m
	}
	{
		p.SetState(240)
		p.Match(NparserPUNTO)
	}
	{
		p.SetState(241)
		p.Match(NparserCONSOLA)
	}
	{
		p.SetState(242)

		var _m = p.Match(NparserPAR_IZQ)

		localctx.(*Funcion_imprimirContext)._PAR_IZQ = _m
	}
	{
		p.SetState(243)

		var _x = p.expresion(0)

		localctx.(*Funcion_imprimirContext)._expresion = _x
	}

	fila := (func() int {
		if localctx.(*Funcion_imprimirContext).Get_SENTENCIA() == nil {
			return 0
		} else {
			return localctx.(*Funcion_imprimirContext).Get_SENTENCIA().GetLine()
		}
	}())
	columna := (func() int {
		if localctx.(*Funcion_imprimirContext).Get_PAR_IZQ() == nil {
			return 0
		} else {
			return localctx.(*Funcion_imprimirContext).Get_PAR_IZQ().GetColumn()
		}
	}())
	columna++
	localctx.(*Funcion_imprimirContext).ex = Instrucciones.NewImprimir(localctx.(*Funcion_imprimirContext).Get_expresion().GetEx(), fila, columna)

	{
		p.SetState(245)
		p.Match(NparserPAR_DER)
	}
	{
		p.SetState(246)
		p.Match(NparserPUNTOCOMA)
	}

	return localctx
}

func (p *Nparser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
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
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 8)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}