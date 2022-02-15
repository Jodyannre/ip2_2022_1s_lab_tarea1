package Expresiones

import (
	"fmt"
	"primeraGramatica/Analizador/Ast"
)

type Asignacion struct {
	Id    string
	Valor interface{}
}

func NewAsignacion(id string, valor interface{}) Asignacion {
	na := Asignacion{
		Id:    id,
		Valor: valor,
	}
	return na
}

func (a Asignacion) GetTipo() Ast.TipoDato {
	return Ast.INSTRUCCION
}

func (a Asignacion) Run(scope *Ast.Scope) interface{} {
	//Verificar que el id  exista
	existe := scope.Exist_actual(a.Id)
	//Obtener el valor del id
	simbolo_id := scope.GetSimbolo(a.Id)
	//Verificar que los tipos sean correctos
	valor := a.Valor.(Ast.Expresion).GetValue(*scope)

	if existe {
		//Existe, ahora verificar los tipos
		if simbolo_id.Valor.(Ast.TipoRetornado).Tipo == valor.Tipo {
			//Los tipos son correctos, actualizar el símbolo
			simbolo_id.Valor = valor
			scope.UpdateSimbolo(a.Id, simbolo_id)
		} else {
			//Error de tipos, generar un error semántico
			fmt.Println("Erro de tipos")
		}
	} else {
		//No existe, generar un error semántico
		fmt.Println("Error, no existe ese id")
	}
	return true
}
