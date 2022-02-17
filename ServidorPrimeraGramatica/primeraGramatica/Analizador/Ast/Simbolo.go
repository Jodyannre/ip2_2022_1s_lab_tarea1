package Ast

type Simbolo struct {
	Identificador string
	Valor         interface{}
	Fila          int
	Columna       int
	Tipo          TipoDato
}

func NewSimbolo(identificador string, valor interface{}, fila int, columna int, tipo TipoDato) Simbolo {
	simbolo := Simbolo{
		Identificador: identificador,
		Valor:         valor,
		Fila:          fila,
		Columna:       columna,
		Tipo:          tipo,
	}
	return simbolo
}
