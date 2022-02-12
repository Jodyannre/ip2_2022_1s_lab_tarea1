package Ast

type Simbolo struct {
	Identificador string
	Valor         interface{}
	Linea         int
	Columna       int
	Tipo          TipoDato
}

func NewSimbolo(Identificador string, Valor interface{}, Linea int, Columna int, Tipo TipoDato) Simbolo {
	simbolo := Simbolo{
		Identificador: Identificador,
		Valor:         Valor,
		Linea:         Linea,
		Columna:       Columna,
		Tipo:          Tipo,
	}
	return simbolo
}
