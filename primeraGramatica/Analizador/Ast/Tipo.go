package Ast

type TipoDato int

const (
	INTEGER TipoDato = iota
	REAL
	STRING
	BOOLEAN
	NULL
)

type TipoRetornado struct {
	Tipo  TipoDato
	Valor interface{}
}
