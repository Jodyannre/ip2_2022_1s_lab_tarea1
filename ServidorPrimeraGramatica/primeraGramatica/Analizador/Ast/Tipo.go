package Ast

type TipoDato int

var ValorTipoDato = [19]string{
	"Integer",
	"Real",
	"String",
	"Boolean",
	"Null",
	"Identificador",
	"Declaración",
	"If",
	"Entonces",
	"If Entonces",
	"Asignación",
	"Imprimir",
	"Expresión",
	"Instrucción",
	"Error",
	"Primitivo",
	"Error Léxico",
	"Error Sintáctico",
	"Error Semántico",
}

const (
	INTEGER TipoDato = iota
	REAL
	STRING
	BOOLEAN
	NULL
	IDENTIFICADOR
	DECLARACION
	IF
	ENTONCES
	IF_ENTONCES
	ASIGNACION
	IMPRIMIR
	EXPRESION
	INSTRUCCION
	ERROR
	PRIMITIVO
	ERROR_LEXICO
	ERROR_SINTACTICO
	ERROR_SEMANTICO
	RESPUESTA
)

type TipoRetornado struct {
	Tipo  TipoDato
	Valor interface{}
}
