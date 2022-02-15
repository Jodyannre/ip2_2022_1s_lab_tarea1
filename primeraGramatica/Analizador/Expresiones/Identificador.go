package Expresiones

import (
	Ast "primeraGramatica/Analizador/Ast"
)

type Identificador struct {
	Tipo  Ast.TipoDato
	Valor string
}

func (p Identificador) GetTipo() Ast.TipoDato {
	return Ast.EXPRESION
}

func (p Identificador) GetValue(scope Ast.Scope) Ast.TipoRetornado {
	//Buscar el símbolo en la tabla de símbolos y retornar el valor
	//Verificar que el id no exista
	if scope.Exist_actual(p.Valor) {
		//Existe el identificar y retornar el valor
		simbolo := scope.GetSimbolo(p.Valor)
		return simbolo.Valor.(Ast.TipoRetornado)
	} else {
		//No existe el identificador, retornar error semantico
		return Ast.TipoRetornado{Valor: nil, Tipo: Ast.NULL}
	}
}

func NewIdentificador(val string, tipo Ast.TipoDato) Identificador {
	return Identificador{Tipo: tipo, Valor: val}
}
