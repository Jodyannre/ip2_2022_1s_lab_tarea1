package Ast

type TipoDato int

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
)

type TipoRetornado struct {
	Tipo  TipoDato
	Valor interface{}
}
