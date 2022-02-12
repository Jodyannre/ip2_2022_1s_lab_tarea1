package Expresiones

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

func (i Imprimir) Ejecutar(entorno Abstract.Scope) interface{} {
	resultado_expresion := i.Expresiones.GetValue(entorno)
	return resultado_expresion.Valor
}
