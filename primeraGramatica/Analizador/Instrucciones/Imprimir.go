package Instrucciones

import (
	Abstract "primeraGramatica/Analizador/Ast"
)

type Imprimir struct {
	Expresiones Abstract.Expresion
}

func NewImprimir(val Abstract.Expresion) Imprimir {
	nuevo := Imprimir{Expresiones: val}
	return nuevo
}

func (i Imprimir) GetTipo() Abstract.TipoDato {
	return Abstract.INSTRUCCION
}

func (i Imprimir) Run(scope *Abstract.Scope) interface{} {
	resultado_expresion := i.Expresiones.GetValue(*scope)
	return resultado_expresion.Valor
}
