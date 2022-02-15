package Expresiones

import (
	Ast "primeraGramatica/Analizador/Ast"
)

type Primitivo struct {
	Tipo  Ast.TipoDato
	Valor interface{}
}

func (p Primitivo) GetValue(entorno Ast.Scope) Ast.TipoRetornado {
	return Ast.TipoRetornado{
		Tipo:  p.Tipo,
		Valor: p.Valor,
	}
}

func (p Primitivo) GetTipo() Ast.TipoDato {
	return Ast.EXPRESION
}

func NewPrimitivo(val interface{}, tipo Ast.TipoDato) Primitivo {
	nuevo := Primitivo{Tipo: tipo, Valor: val}
	return nuevo
}
