package Expresiones

import (
	"fmt"
	"primeraGramatica/Analizador/Ast"
)

type Declaracion struct {
	Id      string
	Tipo    Ast.TipoDato
	Valor   interface{}
	Linea   int
	Columna int
}

func NewDeclaracion(linea int, columna int, id string, tipo Ast.TipoDato, valor interface{}) Declaracion {
	nd := Declaracion{
		Id:      id,
		Tipo:    tipo,
		Valor:   valor,
		Linea:   linea,
		Columna: columna,
	}
	return nd
}

func (d *Declaracion) Run(scope Ast.Scope) {
	//Verificar que el id no exista
	existe := scope.Exist_actual(d.Id)
	//Verificar que los tipos sean correctos
	valor := d.Valor.(Ast.Expresion).GetValue(scope)

	if valor.Tipo == d.Tipo && !existe {
		//El tipo es correcto y no existe en el entorno actual
		//Crear símbolo y agregarlo a la tabla del entorno actual
		nSimbolo := Ast.Simbolo{
			Identificador: d.Id,
			Valor:         valor,
			Linea:         d.Linea,
			Columna:       d.Columna,
			Tipo:          d.Tipo,
		}
		scope.Add(nSimbolo)
	} else if existe {
		//Ya existe y generar error semántico
		fmt.Println("Error, ese elemento ya existe en el ámbito local")
	} else {
		//Error de tipos, generar error semántico
		fmt.Println("Error, los tipos no coinciden en la declaración")
	}

}
