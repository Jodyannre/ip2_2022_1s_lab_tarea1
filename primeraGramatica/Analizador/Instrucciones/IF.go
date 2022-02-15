package Instrucciones

import (
	"fmt"
	"primeraGramatica/Analizador/Ast"

	"github.com/colegno/arraylist"
)

type IF struct {
	Condicion     Ast.Expresion
	Instrucciones *arraylist.List
	Lista_if_else *arraylist.List
	Tipo          Ast.TipoDato
}

func NewIF(condicion Ast.Expresion, instrucciones *arraylist.List, lista *arraylist.List, tipo Ast.TipoDato) IF {
	nif := IF{
		Condicion:     condicion,
		Instrucciones: instrucciones,
		Lista_if_else: lista,
		Tipo:          tipo,
	}
	return nif
}

func (i IF) GetTipo() Ast.TipoDato {
	return Ast.INSTRUCCION
}

func (i IF) Run(scope *Ast.Scope) interface{} {
	//Crear el nuevo scope
	newScope := Ast.NewScope("if", scope)
	//Ejecutar la instrucción if
	GetResultado(i, &newScope, -1)
	return true
}

func GetResultado(i IF, scope *Ast.Scope, pos int) Ast.TipoRetornado {
	var condicion1 Ast.TipoRetornado
	if pos == -1 {
		condicion1 = i.Condicion.GetValue(*scope)
	} else {
		//Conseguir el if/entonces
		elemento := i.Lista_if_else.GetValue(pos).(IF)
		//Evaluar si es if o entonces
		if elemento.Tipo == Ast.IF_ENTONCES {
			//Es un if
			condicion1 = elemento.Condicion.GetValue(*scope)
		} else if elemento.Tipo == Ast.ENTONCES {
			//Es un else
			condicion1 = Ast.TipoRetornado{
				Tipo:  Ast.BOOLEAN,
				Valor: true,
			}
		}
	}

	if condicion1.Tipo == Ast.BOOLEAN {
		if condicion1.Valor.(bool) {
			//Es verdadera; ejecutar las instrucciones
			n := 0
			for n < i.Instrucciones.Len() {
				elemento := i.Instrucciones.GetValue(n).(Ast.Abstracto)
				if elemento.GetTipo() == Ast.EXPRESION {
					instruccion := i.Instrucciones.GetValue(n).(Ast.Expresion)
					instruccion.GetValue(*scope)
				} else if elemento.GetTipo() == Ast.INSTRUCCION {
					instruccion := i.Instrucciones.GetValue(n).(Ast.Instruccion)
					instruccion.Run(scope)
				}
				n++
			}
		} else {
			//Es falsa, buscar en la lista si hay otras
			//Llamada recursiva o fin
			if pos >= i.Lista_if_else.Len() {
				//Llego al final
				return Ast.TipoRetornado{
					Valor: true,
					Tipo:  Ast.BOOLEAN,
				}
			}
			return GetResultado(i, scope, pos+1)
		}
	} else {
		//No es booleano, entonces generar un error semántico
		fmt.Println("Error semántico, la expresión no es un booleano")
	}
	//Se acabo todo y retornar un true de finalizado
	return Ast.TipoRetornado{
		Valor: true,
		Tipo:  Ast.BOOLEAN,
	}
}
