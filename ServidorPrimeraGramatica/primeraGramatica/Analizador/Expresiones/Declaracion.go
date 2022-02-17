package Expresiones

import (
	//"fmt"
	"ServidorPrimeraGramatica/primeraGramatica/Analizador/Ast"
	"ServidorPrimeraGramatica/primeraGramatica/Analizador/Errores"
	"strconv"
)

type Declaracion struct {
	Id      string
	Tipo    Ast.TipoDato
	Valor   interface{}
	Fila    int
	Columna int
}

func NewDeclaracion(id string, tipo Ast.TipoDato, valor interface{}, fila int, columna int) Declaracion {
	nd := Declaracion{
		Id:      id,
		Tipo:    tipo,
		Valor:   valor,
		Fila:    fila,
		Columna: columna,
	}
	return nd
}

func (d Declaracion) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, Ast.DECLARACION
}

func (d Declaracion) Run(scope *Ast.Scope) interface{} {
	//Verificar que el id no exista
	existe := scope.Exist_actual(d.Id)
	//Verificar que los tipos sean correctos
	valor := d.Valor.(Ast.Expresion).GetValue(*scope)

	//Revisar si el retorno es un error
	if valor.Tipo == Ast.ERROR {
		return valor
	}

	if valor.Tipo == d.Tipo && !existe {
		//El tipo es correcto y no existe en el entorno actual
		//Crear símbolo y agregarlo a la tabla del entorno actual
		nSimbolo := Ast.Simbolo{
			Identificador: d.Id,
			Valor:         valor,
			Fila:          d.Fila,
			Columna:       d.Columna,
			Tipo:          d.Tipo,
		}
		scope.Add(nSimbolo)
	} else if existe {
		//Ya existe y generar error semántico
		//fmt.Println("Error, ese elemento ya existe en el ámbito local")
		msg := "Semantic error, the element \"" + d.Id + "\" already exist in this scope." +
			" -- Line:" + strconv.Itoa(d.Fila) + " Column: " + strconv.Itoa(d.Columna)
		nError := Errores.NewError(d.Fila, d.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	} else {
		//Error de tipos, generar error semántico
		//fmt.Println("Error, los tipos no coinciden en la declaración")
		msg := "Semantic error, can't assign " + Ast.ValorTipoDato[int(valor.Tipo)] +
			"type to " + Ast.ValorTipoDato[int(d.Tipo)] +
			"type. -- Line: " + strconv.Itoa(d.Fila) +
			" Column: " + strconv.Itoa(d.Columna)
		nError := Errores.NewError(d.Fila, d.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}
	return Ast.TipoRetornado{
		Tipo:  Ast.BOOLEAN,
		Valor: true,
	}
}
