// Code generated from c:\Projects\Go\src\primeraGramatica\parser\Nparser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Nparser

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import "primeraGramatica/Analizador/Ast"
import "primeraGramatica/Analizador/Expresiones"
import "primeraGramatica/Analizador/Instrucciones"
import "github.com/colegno/arraylist"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 46, 250,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 3, 2, 3, 2, 3, 2, 3, 3,
	7, 3, 37, 10, 3, 12, 3, 14, 3, 40, 11, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	5, 4, 59, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 77, 10, 5, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 109, 10, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 126,
	10, 7, 12, 7, 14, 7, 129, 11, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 5, 8, 139, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 146, 10,
	9, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 152, 10, 10, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 164, 10, 11, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 214, 10, 12, 3, 13, 7, 13,
	217, 10, 13, 12, 13, 14, 13, 220, 11, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 2, 3, 12, 17, 2, 4, 6, 8, 10, 12, 14, 16, 18,
	20, 22, 24, 26, 28, 30, 2, 2, 2, 263, 2, 32, 3, 2, 2, 2, 4, 38, 3, 2, 2,
	2, 6, 58, 3, 2, 2, 2, 8, 76, 3, 2, 2, 2, 10, 78, 3, 2, 2, 2, 12, 108, 3,
	2, 2, 2, 14, 138, 3, 2, 2, 2, 16, 145, 3, 2, 2, 2, 18, 151, 3, 2, 2, 2,
	20, 163, 3, 2, 2, 2, 22, 213, 3, 2, 2, 2, 24, 218, 3, 2, 2, 2, 26, 223,
	3, 2, 2, 2, 28, 234, 3, 2, 2, 2, 30, 240, 3, 2, 2, 2, 32, 33, 5, 4, 3,
	2, 33, 34, 8, 2, 1, 2, 34, 3, 3, 2, 2, 2, 35, 37, 5, 6, 4, 2, 36, 35, 3,
	2, 2, 2, 37, 40, 3, 2, 2, 2, 38, 36, 3, 2, 2, 2, 38, 39, 3, 2, 2, 2, 39,
	41, 3, 2, 2, 2, 40, 38, 3, 2, 2, 2, 41, 42, 8, 3, 1, 2, 42, 5, 3, 2, 2,
	2, 43, 44, 5, 12, 7, 2, 44, 45, 8, 4, 1, 2, 45, 59, 3, 2, 2, 2, 46, 47,
	5, 22, 12, 2, 47, 48, 8, 4, 1, 2, 48, 59, 3, 2, 2, 2, 49, 50, 5, 10, 6,
	2, 50, 51, 8, 4, 1, 2, 51, 59, 3, 2, 2, 2, 52, 53, 5, 8, 5, 2, 53, 54,
	8, 4, 1, 2, 54, 59, 3, 2, 2, 2, 55, 56, 5, 30, 16, 2, 56, 57, 8, 4, 1,
	2, 57, 59, 3, 2, 2, 2, 58, 43, 3, 2, 2, 2, 58, 46, 3, 2, 2, 2, 58, 49,
	3, 2, 2, 2, 58, 52, 3, 2, 2, 2, 58, 55, 3, 2, 2, 2, 59, 7, 3, 2, 2, 2,
	60, 61, 7, 21, 2, 2, 61, 62, 7, 24, 2, 2, 62, 63, 5, 14, 8, 2, 63, 64,
	7, 31, 2, 2, 64, 65, 7, 30, 2, 2, 65, 66, 5, 12, 7, 2, 66, 67, 8, 5, 1,
	2, 67, 68, 7, 29, 2, 2, 68, 69, 7, 27, 2, 2, 69, 77, 3, 2, 2, 2, 70, 71,
	7, 21, 2, 2, 71, 72, 7, 24, 2, 2, 72, 73, 5, 14, 8, 2, 73, 74, 8, 5, 1,
	2, 74, 75, 7, 27, 2, 2, 75, 77, 3, 2, 2, 2, 76, 60, 3, 2, 2, 2, 76, 70,
	3, 2, 2, 2, 77, 9, 3, 2, 2, 2, 78, 79, 7, 24, 2, 2, 79, 80, 7, 31, 2, 2,
	80, 81, 7, 30, 2, 2, 81, 82, 5, 12, 7, 2, 82, 83, 8, 6, 1, 2, 83, 84, 7,
	29, 2, 2, 84, 85, 7, 27, 2, 2, 85, 11, 3, 2, 2, 2, 86, 87, 8, 7, 1, 2,
	87, 88, 7, 35, 2, 2, 88, 89, 5, 12, 7, 13, 89, 90, 8, 7, 1, 2, 90, 109,
	3, 2, 2, 2, 91, 92, 7, 22, 2, 2, 92, 109, 8, 7, 1, 2, 93, 94, 7, 23, 2,
	2, 94, 109, 8, 7, 1, 2, 95, 96, 7, 43, 2, 2, 96, 109, 8, 7, 1, 2, 97, 98,
	7, 13, 2, 2, 98, 109, 8, 7, 1, 2, 99, 100, 7, 14, 2, 2, 100, 109, 8, 7,
	1, 2, 101, 102, 7, 24, 2, 2, 102, 109, 8, 7, 1, 2, 103, 104, 7, 37, 2,
	2, 104, 105, 5, 12, 7, 2, 105, 106, 7, 38, 2, 2, 106, 107, 8, 7, 1, 2,
	107, 109, 3, 2, 2, 2, 108, 86, 3, 2, 2, 2, 108, 91, 3, 2, 2, 2, 108, 93,
	3, 2, 2, 2, 108, 95, 3, 2, 2, 2, 108, 97, 3, 2, 2, 2, 108, 99, 3, 2, 2,
	2, 108, 101, 3, 2, 2, 2, 108, 103, 3, 2, 2, 2, 109, 127, 3, 2, 2, 2, 110,
	111, 12, 12, 2, 2, 111, 112, 5, 18, 10, 2, 112, 113, 5, 12, 7, 13, 113,
	114, 8, 7, 1, 2, 114, 126, 3, 2, 2, 2, 115, 116, 12, 11, 2, 2, 116, 117,
	5, 16, 9, 2, 117, 118, 5, 12, 7, 12, 118, 119, 8, 7, 1, 2, 119, 126, 3,
	2, 2, 2, 120, 121, 12, 10, 2, 2, 121, 122, 5, 20, 11, 2, 122, 123, 5, 12,
	7, 11, 123, 124, 8, 7, 1, 2, 124, 126, 3, 2, 2, 2, 125, 110, 3, 2, 2, 2,
	125, 115, 3, 2, 2, 2, 125, 120, 3, 2, 2, 2, 126, 129, 3, 2, 2, 2, 127,
	125, 3, 2, 2, 2, 127, 128, 3, 2, 2, 2, 128, 13, 3, 2, 2, 2, 129, 127, 3,
	2, 2, 2, 130, 131, 7, 10, 2, 2, 131, 139, 8, 8, 1, 2, 132, 133, 7, 11,
	2, 2, 133, 139, 8, 8, 1, 2, 134, 135, 7, 12, 2, 2, 135, 139, 8, 8, 1, 2,
	136, 137, 7, 9, 2, 2, 137, 139, 8, 8, 1, 2, 138, 130, 3, 2, 2, 2, 138,
	132, 3, 2, 2, 2, 138, 134, 3, 2, 2, 2, 138, 136, 3, 2, 2, 2, 139, 15, 3,
	2, 2, 2, 140, 141, 7, 29, 2, 2, 141, 146, 8, 9, 1, 2, 142, 143, 7, 29,
	2, 2, 143, 144, 7, 31, 2, 2, 144, 146, 8, 9, 1, 2, 145, 140, 3, 2, 2, 2,
	145, 142, 3, 2, 2, 2, 146, 17, 3, 2, 2, 2, 147, 148, 7, 4, 2, 2, 148, 152,
	8, 10, 1, 2, 149, 150, 7, 5, 2, 2, 150, 152, 8, 10, 1, 2, 151, 147, 3,
	2, 2, 2, 151, 149, 3, 2, 2, 2, 152, 19, 3, 2, 2, 2, 153, 154, 7, 33, 2,
	2, 154, 164, 8, 11, 1, 2, 155, 156, 7, 34, 2, 2, 156, 164, 8, 11, 1, 2,
	157, 158, 7, 32, 2, 2, 158, 164, 8, 11, 1, 2, 159, 160, 7, 36, 2, 2, 160,
	164, 8, 11, 1, 2, 161, 162, 7, 35, 2, 2, 162, 164, 8, 11, 1, 2, 163, 153,
	3, 2, 2, 2, 163, 155, 3, 2, 2, 2, 163, 157, 3, 2, 2, 2, 163, 159, 3, 2,
	2, 2, 163, 161, 3, 2, 2, 2, 164, 21, 3, 2, 2, 2, 165, 166, 7, 6, 2, 2,
	166, 167, 7, 25, 2, 2, 167, 168, 7, 37, 2, 2, 168, 169, 5, 12, 7, 2, 169,
	170, 7, 38, 2, 2, 170, 171, 7, 41, 2, 2, 171, 172, 5, 4, 3, 2, 172, 173,
	7, 42, 2, 2, 173, 174, 5, 28, 15, 2, 174, 175, 7, 27, 2, 2, 175, 176, 8,
	12, 1, 2, 176, 214, 3, 2, 2, 2, 177, 178, 7, 6, 2, 2, 178, 179, 7, 25,
	2, 2, 179, 180, 7, 37, 2, 2, 180, 181, 5, 12, 7, 2, 181, 182, 7, 38, 2,
	2, 182, 183, 7, 41, 2, 2, 183, 184, 5, 4, 3, 2, 184, 185, 7, 42, 2, 2,
	185, 186, 7, 27, 2, 2, 186, 187, 8, 12, 1, 2, 187, 214, 3, 2, 2, 2, 188,
	189, 7, 6, 2, 2, 189, 190, 7, 25, 2, 2, 190, 191, 7, 37, 2, 2, 191, 192,
	5, 12, 7, 2, 192, 193, 7, 38, 2, 2, 193, 194, 7, 41, 2, 2, 194, 195, 5,
	4, 3, 2, 195, 196, 7, 42, 2, 2, 196, 197, 5, 24, 13, 2, 197, 198, 7, 27,
	2, 2, 198, 199, 8, 12, 1, 2, 199, 214, 3, 2, 2, 2, 200, 201, 7, 6, 2, 2,
	201, 202, 7, 25, 2, 2, 202, 203, 7, 37, 2, 2, 203, 204, 5, 12, 7, 2, 204,
	205, 7, 38, 2, 2, 205, 206, 7, 41, 2, 2, 206, 207, 5, 4, 3, 2, 207, 208,
	7, 42, 2, 2, 208, 209, 5, 24, 13, 2, 209, 210, 5, 28, 15, 2, 210, 211,
	7, 27, 2, 2, 211, 212, 8, 12, 1, 2, 212, 214, 3, 2, 2, 2, 213, 165, 3,
	2, 2, 2, 213, 177, 3, 2, 2, 2, 213, 188, 3, 2, 2, 2, 213, 200, 3, 2, 2,
	2, 214, 23, 3, 2, 2, 2, 215, 217, 5, 26, 14, 2, 216, 215, 3, 2, 2, 2, 217,
	220, 3, 2, 2, 2, 218, 216, 3, 2, 2, 2, 218, 219, 3, 2, 2, 2, 219, 221,
	3, 2, 2, 2, 220, 218, 3, 2, 2, 2, 221, 222, 8, 13, 1, 2, 222, 25, 3, 2,
	2, 2, 223, 224, 7, 7, 2, 2, 224, 225, 7, 6, 2, 2, 225, 226, 7, 25, 2, 2,
	226, 227, 7, 37, 2, 2, 227, 228, 5, 12, 7, 2, 228, 229, 7, 38, 2, 2, 229,
	230, 7, 41, 2, 2, 230, 231, 5, 4, 3, 2, 231, 232, 7, 42, 2, 2, 232, 233,
	8, 14, 1, 2, 233, 27, 3, 2, 2, 2, 234, 235, 7, 7, 2, 2, 235, 236, 7, 41,
	2, 2, 236, 237, 5, 4, 3, 2, 237, 238, 8, 15, 1, 2, 238, 239, 7, 42, 2,
	2, 239, 29, 3, 2, 2, 2, 240, 241, 7, 3, 2, 2, 241, 242, 7, 26, 2, 2, 242,
	243, 7, 8, 2, 2, 243, 244, 7, 37, 2, 2, 244, 245, 5, 12, 7, 2, 245, 246,
	8, 16, 1, 2, 246, 247, 7, 38, 2, 2, 247, 248, 7, 27, 2, 2, 248, 31, 3,
	2, 2, 2, 14, 38, 58, 76, 108, 125, 127, 138, 145, 151, 163, 213, 218,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

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
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type Nparser struct {
	*antlr.BaseParser
}

func NewNparser(input antlr.TokenStream) *Nparser {
	this := new(Nparser)

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

func (p *Nparser) Inicio() (localctx IInicioContext) {
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

func (p *Nparser) Instrucciones() (localctx IInstruccionesContext) {
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

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NparserSENTENCIA)|(1<<NparserIF)|(1<<NparserTRUE)|(1<<NparserFALSE)|(1<<NparserDECLARAR)|(1<<NparserNUMERO)|(1<<NparserDECIMAL)|(1<<NparserID))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(NparserRESTA-33))|(1<<(NparserPAR_IZQ-33))|(1<<(NparserCADENA-33)))) != 0) {
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

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// Get_control_IF_simple returns the _control_IF_simple rule contexts.
	Get_control_IF_simple() IControl_IF_simpleContext

	// Get_asignacion returns the _asignacion rule contexts.
	Get_asignacion() IAsignacionContext

	// Get_declaracion returns the _declaracion rule contexts.
	Get_declaracion() IDeclaracionContext

	// Get_funcion_imprimir returns the _funcion_imprimir rule contexts.
	Get_funcion_imprimir() IFuncion_imprimirContext

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

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
	_expresion         IExpresionContext
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

func (s *InstruccionContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *InstruccionContext) Get_control_IF_simple() IControl_IF_simpleContext {
	return s._control_IF_simple
}

func (s *InstruccionContext) Get_asignacion() IAsignacionContext { return s._asignacion }

func (s *InstruccionContext) Get_declaracion() IDeclaracionContext { return s._declaracion }

func (s *InstruccionContext) Get_funcion_imprimir() IFuncion_imprimirContext {
	return s._funcion_imprimir
}

func (s *InstruccionContext) Set_expresion(v IExpresionContext) { s._expresion = v }

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

func (s *InstruccionContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

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

func (p *Nparser) Instruccion() (localctx IInstruccionContext) {
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

	p.SetState(56)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(41)

			var _x = p.expresion(0)

			localctx.(*InstruccionContext)._expresion = _x
		}
		localctx.(*InstruccionContext).ex = localctx.(*InstruccionContext).Get_expresion().GetEx()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(44)

			var _x = p.Control_IF_simple()

			localctx.(*InstruccionContext)._control_IF_simple = _x
		}
		localctx.(*InstruccionContext).ex = localctx.(*InstruccionContext).Get_control_IF_simple().GetEx()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(47)

			var _x = p.Asignacion()

			localctx.(*InstruccionContext)._asignacion = _x
		}
		localctx.(*InstruccionContext).ex = localctx.(*InstruccionContext).Get_asignacion().GetEx()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(50)

			var _x = p.Declaracion()

			localctx.(*InstruccionContext)._declaracion = _x
		}
		localctx.(*InstruccionContext).ex = localctx.(*InstruccionContext).Get_declaracion().GetEx()

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(53)

			var _x = p.Funcion_imprimir()

			localctx.(*InstruccionContext)._funcion_imprimir = _x
		}
		localctx.(*InstruccionContext).ex = localctx.(*InstruccionContext).Get_funcion_imprimir().GetEx()

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

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

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
	_expresion IExpresionContext
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

func (s *DeclaracionContext) Set_ID(v antlr.Token) { s._ID = v }

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

func (p *Nparser) Declaracion() (localctx IDeclaracionContext) {
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

	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(58)
			p.Match(NparserDECLARAR)
		}
		{
			p.SetState(59)

			var _m = p.Match(NparserID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(60)

			var _x = p.Tipo()

			localctx.(*DeclaracionContext)._tipo = _x
		}
		{
			p.SetState(61)
			p.Match(NparserIGUAL)
		}
		{
			p.SetState(62)
			p.Match(NparserMENOR)
		}
		{
			p.SetState(63)

			var _x = p.expresion(0)

			localctx.(*DeclaracionContext)._expresion = _x
		}

		localctx.(*DeclaracionContext).ex = Expresiones.NewDeclaracion((func() string {
			if localctx.(*DeclaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclaracionContext).Get_ID().GetText()
			}
		}()), localctx.(*DeclaracionContext).Get_tipo().GetEx(), localctx.(*DeclaracionContext).Get_expresion().GetEx())

		{
			p.SetState(65)
			p.Match(NparserMAYOR)
		}
		{
			p.SetState(66)
			p.Match(NparserPUNTOCOMA)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(68)
			p.Match(NparserDECLARAR)
		}
		{
			p.SetState(69)

			var _m = p.Match(NparserID)

			localctx.(*DeclaracionContext)._ID = _m
		}
		{
			p.SetState(70)

			var _x = p.Tipo()

			localctx.(*DeclaracionContext)._tipo = _x
		}

		localctx.(*DeclaracionContext).ex = Expresiones.NewDeclaracion((func() string {
			if localctx.(*DeclaracionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*DeclaracionContext).Get_ID().GetText()
			}
		}()), localctx.(*DeclaracionContext).Get_tipo().GetEx(), nil)

		{
			p.SetState(72)
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

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

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

func (s *AsignacionContext) Set_ID(v antlr.Token) { s._ID = v }

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

func (p *Nparser) Asignacion() (localctx IAsignacionContext) {
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
		p.SetState(76)

		var _m = p.Match(NparserID)

		localctx.(*AsignacionContext)._ID = _m
	}
	{
		p.SetState(77)
		p.Match(NparserIGUAL)
	}
	{
		p.SetState(78)
		p.Match(NparserMENOR)
	}
	{
		p.SetState(79)

		var _x = p.expresion(0)

		localctx.(*AsignacionContext)._expresion = _x
	}

	//linea = strconv.Atoi((func() int { if localctx.(*AsignacionContext).Get_ID() == nil { return 0 } else { return localctx.(*AsignacionContext).Get_ID().GetLine() }}()))
	//columna = strconv.Atoi((func() int { if localctx.(*AsignacionContext).Get_ID() == nil { return 0 } else { return localctx.(*AsignacionContext).Get_ID().GetColumn() }}()))
	localctx.(*AsignacionContext).ex = Expresiones.NewAsignacion((func() string {
		if localctx.(*AsignacionContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*AsignacionContext).Get_ID().GetText()
		}
	}()), localctx.(*AsignacionContext).Get_expresion().GetEx())

	{
		p.SetState(81)
		p.Match(NparserMAYOR)
	}
	{
		p.SetState(82)
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

	// GetOp_left returns the op_left rule contexts.
	GetOp_left() IExpresionContext

	// Get_expresion returns the _expresion rule contexts.
	Get_expresion() IExpresionContext

	// Get_operador_logico returns the _operador_logico rule contexts.
	Get_operador_logico() IOperador_logicoContext

	// GetOp_right returns the op_right rule contexts.
	GetOp_right() IExpresionContext

	// Get_operador_comparacion returns the _operador_comparacion rule contexts.
	Get_operador_comparacion() IOperador_comparacionContext

	// Get_operador_aritmetico returns the _operador_aritmetico rule contexts.
	Get_operador_aritmetico() IOperador_aritmeticoContext

	// SetOp_left sets the op_left rule contexts.
	SetOp_left(IExpresionContext)

	// Set_expresion sets the _expresion rule contexts.
	Set_expresion(IExpresionContext)

	// Set_operador_logico sets the _operador_logico rule contexts.
	Set_operador_logico(IOperador_logicoContext)

	// SetOp_right sets the op_right rule contexts.
	SetOp_right(IExpresionContext)

	// Set_operador_comparacion sets the _operador_comparacion rule contexts.
	Set_operador_comparacion(IOperador_comparacionContext)

	// Set_operador_aritmetico sets the _operador_aritmetico rule contexts.
	Set_operador_aritmetico(IOperador_aritmeticoContext)

	// GetEx returns the ex attribute.
	GetEx() Ast.Expresion

	// SetEx sets the ex attribute.
	SetEx(Ast.Expresion)

	// IsExpresionContext differentiates from other interfaces.
	IsExpresionContext()
}

type ExpresionContext struct {
	*antlr.BaseParserRuleContext
	parser                antlr.Parser
	ex                    Ast.Expresion
	op_left               IExpresionContext
	_RESTA                antlr.Token
	_expresion            IExpresionContext
	_NUMERO               antlr.Token
	_DECIMAL              antlr.Token
	_CADENA               antlr.Token
	_ID                   antlr.Token
	_operador_logico      IOperador_logicoContext
	op_right              IExpresionContext
	_operador_comparacion IOperador_comparacionContext
	_operador_aritmetico  IOperador_aritmeticoContext
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

func (s *ExpresionContext) Set_RESTA(v antlr.Token) { s._RESTA = v }

func (s *ExpresionContext) Set_NUMERO(v antlr.Token) { s._NUMERO = v }

func (s *ExpresionContext) Set_DECIMAL(v antlr.Token) { s._DECIMAL = v }

func (s *ExpresionContext) Set_CADENA(v antlr.Token) { s._CADENA = v }

func (s *ExpresionContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ExpresionContext) GetOp_left() IExpresionContext { return s.op_left }

func (s *ExpresionContext) Get_expresion() IExpresionContext { return s._expresion }

func (s *ExpresionContext) Get_operador_logico() IOperador_logicoContext { return s._operador_logico }

func (s *ExpresionContext) GetOp_right() IExpresionContext { return s.op_right }

func (s *ExpresionContext) Get_operador_comparacion() IOperador_comparacionContext {
	return s._operador_comparacion
}

func (s *ExpresionContext) Get_operador_aritmetico() IOperador_aritmeticoContext {
	return s._operador_aritmetico
}

func (s *ExpresionContext) SetOp_left(v IExpresionContext) { s.op_left = v }

func (s *ExpresionContext) Set_expresion(v IExpresionContext) { s._expresion = v }

func (s *ExpresionContext) Set_operador_logico(v IOperador_logicoContext) { s._operador_logico = v }

func (s *ExpresionContext) SetOp_right(v IExpresionContext) { s.op_right = v }

func (s *ExpresionContext) Set_operador_comparacion(v IOperador_comparacionContext) {
	s._operador_comparacion = v
}

func (s *ExpresionContext) Set_operador_aritmetico(v IOperador_aritmeticoContext) {
	s._operador_aritmetico = v
}

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

func (s *ExpresionContext) Operador_logico() IOperador_logicoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperador_logicoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperador_logicoContext)
}

func (s *ExpresionContext) Operador_comparacion() IOperador_comparacionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperador_comparacionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperador_comparacionContext)
}

func (s *ExpresionContext) Operador_aritmetico() IOperador_aritmeticoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperador_aritmeticoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperador_aritmeticoContext)
}

func (s *ExpresionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpresionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *Nparser) Expresion() (localctx IExpresionContext) {
	return p.expresion(0)
}

func (p *Nparser) expresion(_p int) (localctx IExpresionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpresionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpresionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, NparserRULE_expresion, _p)

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
	p.SetState(106)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NparserRESTA:
		{
			p.SetState(85)

			var _m = p.Match(NparserRESTA)

			localctx.(*ExpresionContext)._RESTA = _m
		}
		{
			p.SetState(86)

			var _x = p.expresion(11)

			localctx.(*ExpresionContext).op_left = _x
			localctx.(*ExpresionContext)._expresion = _x
		}

		operador := (func() string {
			if localctx.(*ExpresionContext).Get_RESTA() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).Get_RESTA().GetText()
			}
		}())
		localctx.(*ExpresionContext).ex = Expresiones.NewOperation(localctx.(*ExpresionContext).GetOp_left().GetEx(), operador, nil, true)

	case NparserNUMERO:
		{
			p.SetState(89)

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
			p.SetState(91)

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
			p.SetState(93)

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
			p.SetState(95)
			p.Match(NparserTRUE)
		}

		valor := true
		localctx.(*ExpresionContext).ex = Expresiones.NewPrimitivo(valor, Ast.BOOLEAN)

	case NparserFALSE:
		{
			p.SetState(97)
			p.Match(NparserFALSE)
		}

		valor := false
		localctx.(*ExpresionContext).ex = Expresiones.NewPrimitivo(valor, Ast.BOOLEAN)

	case NparserID:
		{
			p.SetState(99)

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
			p.SetState(101)
			p.Match(NparserPAR_IZQ)
		}
		{
			p.SetState(102)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext)._expresion = _x
		}
		{
			p.SetState(103)
			p.Match(NparserPAR_DER)
		}

		localctx.(*ExpresionContext).ex = localctx.(*ExpresionContext).Get_expresion().GetEx()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(123)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).op_left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, NparserRULE_expresion)
				p.SetState(108)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(109)

					var _x = p.Operador_logico()

					localctx.(*ExpresionContext)._operador_logico = _x
				}
				{
					p.SetState(110)

					var _x = p.expresion(11)

					localctx.(*ExpresionContext).op_right = _x
					localctx.(*ExpresionContext)._expresion = _x
				}

				operador := localctx.(*ExpresionContext).Get_operador_logico().GetOperador()
				localctx.(*ExpresionContext).ex = Expresiones.NewOperation(localctx.(*ExpresionContext).GetOp_left().GetEx(), operador, localctx.(*ExpresionContext).GetOp_right().GetEx(), false)

			case 2:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).op_left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, NparserRULE_expresion)
				p.SetState(113)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(114)

					var _x = p.Operador_comparacion()

					localctx.(*ExpresionContext)._operador_comparacion = _x
				}
				{
					p.SetState(115)

					var _x = p.expresion(10)

					localctx.(*ExpresionContext).op_right = _x
					localctx.(*ExpresionContext)._expresion = _x
				}

				operador := localctx.(*ExpresionContext).Get_operador_comparacion().GetOperador()
				localctx.(*ExpresionContext).ex = Expresiones.NewOperation(localctx.(*ExpresionContext).GetOp_left().GetEx(), operador, localctx.(*ExpresionContext).GetOp_right().GetEx(), false)

			case 3:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).op_left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, NparserRULE_expresion)
				p.SetState(118)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(119)

					var _x = p.Operador_aritmetico()

					localctx.(*ExpresionContext)._operador_aritmetico = _x
				}
				{
					p.SetState(120)

					var _x = p.expresion(9)

					localctx.(*ExpresionContext).op_right = _x
					localctx.(*ExpresionContext)._expresion = _x
				}

				operador := localctx.(*ExpresionContext).Get_operador_aritmetico().GetOperador()
				localctx.(*ExpresionContext).ex = Expresiones.NewOperation(localctx.(*ExpresionContext).GetOp_left().GetEx(), operador, localctx.(*ExpresionContext).GetOp_right().GetEx(), false)

			}

		}
		p.SetState(127)
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

func (p *Nparser) Tipo() (localctx ITipoContext) {
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

	p.SetState(136)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NparserINTEGER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(128)
			p.Match(NparserINTEGER)
		}
		localctx.(*TipoContext).ex = Ast.INTEGER

	case NparserREAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(130)
			p.Match(NparserREAL)
		}
		localctx.(*TipoContext).ex = Ast.REAL

	case NparserBOOLEAN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(132)
			p.Match(NparserBOOLEAN)
		}
		localctx.(*TipoContext).ex = Ast.BOOLEAN

	case NparserSTRING:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(134)
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

	// Get_IGUAL returns the _IGUAL token.
	Get_IGUAL() antlr.Token

	// Set_MAYOR sets the _MAYOR token.
	Set_MAYOR(antlr.Token)

	// Set_IGUAL sets the _IGUAL token.
	Set_IGUAL(antlr.Token)

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
	_IGUAL   antlr.Token
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

func (s *Operador_comparacionContext) Get_IGUAL() antlr.Token { return s._IGUAL }

func (s *Operador_comparacionContext) Set_MAYOR(v antlr.Token) { s._MAYOR = v }

func (s *Operador_comparacionContext) Set_IGUAL(v antlr.Token) { s._IGUAL = v }

func (s *Operador_comparacionContext) GetOperador() string { return s.operador }

func (s *Operador_comparacionContext) SetOperador(v string) { s.operador = v }

func (s *Operador_comparacionContext) MAYOR() antlr.TerminalNode {
	return s.GetToken(NparserMAYOR, 0)
}

func (s *Operador_comparacionContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(NparserIGUAL, 0)
}

func (s *Operador_comparacionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Operador_comparacionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *Nparser) Operador_comparacion() (localctx IOperador_comparacionContext) {
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

	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(138)

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

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(140)

			var _m = p.Match(NparserMAYOR)

			localctx.(*Operador_comparacionContext)._MAYOR = _m
		}
		{
			p.SetState(141)

			var _m = p.Match(NparserIGUAL)

			localctx.(*Operador_comparacionContext)._IGUAL = _m
		}
		localctx.(*Operador_comparacionContext).operador = (func() string {
			if localctx.(*Operador_comparacionContext).Get_MAYOR() == nil {
				return ""
			} else {
				return localctx.(*Operador_comparacionContext).Get_MAYOR().GetText()
			}
		}()) + (func() string {
			if localctx.(*Operador_comparacionContext).Get_IGUAL() == nil {
				return ""
			} else {
				return localctx.(*Operador_comparacionContext).Get_IGUAL().GetText()
			}
		}())

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

func (p *Nparser) Operador_logico() (localctx IOperador_logicoContext) {
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

	p.SetState(149)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NparserAND:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(145)

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
			p.SetState(147)

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

func (p *Nparser) Operador_aritmetico() (localctx IOperador_aritmeticoContext) {
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

	p.SetState(161)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NparserMULTIPLICACION:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(151)

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
			p.SetState(153)

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
			p.SetState(155)

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
			p.SetState(157)

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
			p.SetState(159)

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

func (p *Nparser) Control_IF_simple() (localctx IControl_IF_simpleContext) {
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

	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(163)
			p.Match(NparserIF)
		}
		{
			p.SetState(164)
			p.Match(NparserDOSPUNTOS)
		}
		{
			p.SetState(165)
			p.Match(NparserPAR_IZQ)
		}
		{
			p.SetState(166)

			var _x = p.expresion(0)

			localctx.(*Control_IF_simpleContext)._expresion = _x
		}
		{
			p.SetState(167)
			p.Match(NparserPAR_DER)
		}
		{
			p.SetState(168)
			p.Match(NparserCORCHETE_IZQ)
		}
		{
			p.SetState(169)

			var _x = p.Instrucciones()

			localctx.(*Control_IF_simpleContext)._instrucciones = _x
		}
		{
			p.SetState(170)
			p.Match(NparserCORCHETE_DER)
		}
		{
			p.SetState(171)

			var _x = p.Control_else()

			localctx.(*Control_IF_simpleContext)._control_else = _x
		}
		{
			p.SetState(172)
			p.Match(NparserPUNTOCOMA)
		}

		lista_entonces := arraylist.New()
		lista_entonces.Add(localctx.(*Control_IF_simpleContext).Get_control_else().GetEx())
		localctx.(*Control_IF_simpleContext).ex = Instrucciones.NewIF(localctx.(*Control_IF_simpleContext).Get_expresion().GetEx(), localctx.(*Control_IF_simpleContext).Get_instrucciones().GetList(), lista_entonces, Ast.IF)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(175)
			p.Match(NparserIF)
		}
		{
			p.SetState(176)
			p.Match(NparserDOSPUNTOS)
		}
		{
			p.SetState(177)
			p.Match(NparserPAR_IZQ)
		}
		{
			p.SetState(178)

			var _x = p.expresion(0)

			localctx.(*Control_IF_simpleContext)._expresion = _x
		}
		{
			p.SetState(179)
			p.Match(NparserPAR_DER)
		}
		{
			p.SetState(180)
			p.Match(NparserCORCHETE_IZQ)
		}
		{
			p.SetState(181)

			var _x = p.Instrucciones()

			localctx.(*Control_IF_simpleContext)._instrucciones = _x
		}
		{
			p.SetState(182)
			p.Match(NparserCORCHETE_DER)
		}
		{
			p.SetState(183)
			p.Match(NparserPUNTOCOMA)
		}

		lista_null := arraylist.New()
		localctx.(*Control_IF_simpleContext).ex = Instrucciones.NewIF(localctx.(*Control_IF_simpleContext).Get_expresion().GetEx(), localctx.(*Control_IF_simpleContext).Get_instrucciones().GetList(), lista_null, Ast.IF)

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(186)
			p.Match(NparserIF)
		}
		{
			p.SetState(187)
			p.Match(NparserDOSPUNTOS)
		}
		{
			p.SetState(188)
			p.Match(NparserPAR_IZQ)
		}
		{
			p.SetState(189)

			var _x = p.expresion(0)

			localctx.(*Control_IF_simpleContext)._expresion = _x
		}
		{
			p.SetState(190)
			p.Match(NparserPAR_DER)
		}
		{
			p.SetState(191)
			p.Match(NparserCORCHETE_IZQ)
		}
		{
			p.SetState(192)

			var _x = p.Instrucciones()

			localctx.(*Control_IF_simpleContext)._instrucciones = _x
		}
		{
			p.SetState(193)
			p.Match(NparserCORCHETE_DER)
		}
		{
			p.SetState(194)

			var _x = p.Control_IF_entonces()

			localctx.(*Control_IF_simpleContext)._control_IF_entonces = _x
		}
		{
			p.SetState(195)
			p.Match(NparserPUNTOCOMA)
		}

		lista_entonces := localctx.(*Control_IF_simpleContext).Get_control_IF_entonces().GetList()
		localctx.(*Control_IF_simpleContext).ex = Instrucciones.NewIF(localctx.(*Control_IF_simpleContext).Get_expresion().GetEx(), localctx.(*Control_IF_simpleContext).Get_instrucciones().GetList(), lista_entonces, Ast.IF)

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(198)
			p.Match(NparserIF)
		}
		{
			p.SetState(199)
			p.Match(NparserDOSPUNTOS)
		}
		{
			p.SetState(200)
			p.Match(NparserPAR_IZQ)
		}
		{
			p.SetState(201)

			var _x = p.expresion(0)

			localctx.(*Control_IF_simpleContext)._expresion = _x
		}
		{
			p.SetState(202)
			p.Match(NparserPAR_DER)
		}
		{
			p.SetState(203)
			p.Match(NparserCORCHETE_IZQ)
		}
		{
			p.SetState(204)

			var _x = p.Instrucciones()

			localctx.(*Control_IF_simpleContext)._instrucciones = _x
		}
		{
			p.SetState(205)
			p.Match(NparserCORCHETE_DER)
		}
		{
			p.SetState(206)

			var _x = p.Control_IF_entonces()

			localctx.(*Control_IF_simpleContext)._control_IF_entonces = _x
		}
		{
			p.SetState(207)

			var _x = p.Control_else()

			localctx.(*Control_IF_simpleContext)._control_else = _x
		}
		{
			p.SetState(208)
			p.Match(NparserPUNTOCOMA)
		}

		lista_entonces := localctx.(*Control_IF_simpleContext).Get_control_IF_entonces().GetList()
		lista_entonces.Add(localctx.(*Control_IF_simpleContext).Get_control_else().GetEx())
		localctx.(*Control_IF_simpleContext).ex = Instrucciones.NewIF(localctx.(*Control_IF_simpleContext).Get_expresion().GetEx(), localctx.(*Control_IF_simpleContext).Get_instrucciones().GetList(), lista_entonces, Ast.IF)

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

func (p *Nparser) Control_IF_entonces() (localctx IControl_IF_entoncesContext) {
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
	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(213)

				var _x = p.Control_if_entonces()

				localctx.(*Control_IF_entoncesContext)._control_if_entonces = _x
			}
			localctx.(*Control_IF_entoncesContext).e = append(localctx.(*Control_IF_entoncesContext).e, localctx.(*Control_IF_entoncesContext)._control_if_entonces)

		}
		p.SetState(218)
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

func (p *Nparser) Control_if_entonces() (localctx IControl_if_entoncesContext) {
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
		p.SetState(221)
		p.Match(NparserENTONCES)
	}
	{
		p.SetState(222)
		p.Match(NparserIF)
	}
	{
		p.SetState(223)
		p.Match(NparserDOSPUNTOS)
	}
	{
		p.SetState(224)
		p.Match(NparserPAR_IZQ)
	}
	{
		p.SetState(225)

		var _x = p.expresion(0)

		localctx.(*Control_if_entoncesContext)._expresion = _x
	}
	{
		p.SetState(226)
		p.Match(NparserPAR_DER)
	}
	{
		p.SetState(227)
		p.Match(NparserCORCHETE_IZQ)
	}
	{
		p.SetState(228)

		var _x = p.Instrucciones()

		localctx.(*Control_if_entoncesContext)._instrucciones = _x
	}
	{
		p.SetState(229)
		p.Match(NparserCORCHETE_DER)
	}

	lista_null := arraylist.New()
	localctx.(*Control_if_entoncesContext).ex = Instrucciones.NewIF(localctx.(*Control_if_entoncesContext).Get_expresion().GetEx(), localctx.(*Control_if_entoncesContext).Get_instrucciones().GetList(), lista_null, Ast.IF_ENTONCES)

	return localctx
}

// IControl_elseContext is an interface to support dynamic dispatch.
type IControl_elseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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

func (p *Nparser) Control_else() (localctx IControl_elseContext) {
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
		p.SetState(232)
		p.Match(NparserENTONCES)
	}
	{
		p.SetState(233)
		p.Match(NparserCORCHETE_IZQ)
	}
	{
		p.SetState(234)

		var _x = p.Instrucciones()

		localctx.(*Control_elseContext)._instrucciones = _x
	}

	expresionNull := Expresiones.NewPrimitivo(nil, Ast.NULL)
	lista_null := arraylist.New()
	localctx.(*Control_elseContext).ex = Instrucciones.NewIF(expresionNull, localctx.(*Control_elseContext).Get_instrucciones().GetList(), lista_null, Ast.ENTONCES)

	{
		p.SetState(236)
		p.Match(NparserCORCHETE_DER)
	}

	return localctx
}

// IFuncion_imprimirContext is an interface to support dynamic dispatch.
type IFuncion_imprimirContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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

func (p *Nparser) Funcion_imprimir() (localctx IFuncion_imprimirContext) {
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
		p.SetState(238)
		p.Match(NparserSENTENCIA)
	}
	{
		p.SetState(239)
		p.Match(NparserPUNTO)
	}
	{
		p.SetState(240)
		p.Match(NparserCONSOLA)
	}
	{
		p.SetState(241)
		p.Match(NparserPAR_IZQ)
	}
	{
		p.SetState(242)

		var _x = p.expresion(0)

		localctx.(*Funcion_imprimirContext)._expresion = _x
	}

	localctx.(*Funcion_imprimirContext).ex = Instrucciones.NewImprimir(localctx.(*Funcion_imprimirContext).Get_expresion().GetEx())

	{
		p.SetState(244)
		p.Match(NparserPAR_DER)
	}
	{
		p.SetState(245)
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
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 8)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
