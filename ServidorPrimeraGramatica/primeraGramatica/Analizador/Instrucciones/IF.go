package Instrucciones

import (
	"ServidorPrimeraGramatica/primeraGramatica/Analizador/Ast"
	"ServidorPrimeraGramatica/primeraGramatica/Analizador/Errores"
	"strconv"

	"github.com/colegno/arraylist"
)

type IF struct {
	Condicion     Ast.Expresion
	Instrucciones *arraylist.List
	Lista_if_else *arraylist.List
	Tipo          Ast.TipoDato
	Fila          int
	Columna       int
}

func NewIF(condicion Ast.Expresion, instrucciones *arraylist.List, lista *arraylist.List, tipo Ast.TipoDato,
	fila, columna int) IF {
	nif := IF{
		Condicion:     condicion,
		Instrucciones: instrucciones,
		Lista_if_else: lista,
		Tipo:          tipo,
		Fila:          fila,
		Columna:       columna,
	}
	return nif
}

func (i IF) GetTipo() (Ast.TipoDato, Ast.TipoDato) {
	return Ast.INSTRUCCION, Ast.IF
}

func (i IF) Run(scope *Ast.Scope) interface{} {
	//Crear el nuevo scope
	newScope := Ast.NewScope("if", scope)
	//Inicializar la lista de respuestas
	//Ejecutar la instrucción if
	GetResultado(i, &newScope, -1)
	//actualizar el scope global con los resultados
	newScope.UpdateScopeGlobal()
	return Ast.TipoRetornado{Valor: true, Tipo: Ast.BOOLEAN}
}

func GetResultado(i IF, scope *Ast.Scope, pos int) Ast.TipoRetornado {
	var condicion1 Ast.TipoRetornado
	if pos == -1 {
		condicion1 = i.Condicion.GetValue(*scope)
	} else {
		//Conseguir el if/entonces
		//elemento := i.Lista_if_else.GetValue(pos).(IF)
		//Evaluar si es if o entonces
		if i.Tipo == Ast.IF_ENTONCES {
			//Es un if
			//i = elemento
			scope.Nombre = "If entonces"
			condicion1 = i.Condicion.GetValue(*scope)
		} else if i.Tipo == Ast.ENTONCES {
			//Es un else
			//i = elemento
			scope.Nombre = "Entonces"
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
				//Verificar que la instrucción no sea null
				if i.Instrucciones.GetValue(n) == nil {
					n++
					continue
				}
				elemento2 := i.Instrucciones.GetValue(n).(Ast.Abstracto)
				tipo_abstracto, _ := elemento2.GetTipo()
				if tipo_abstracto == Ast.EXPRESION {
					instruccion := i.Instrucciones.GetValue(n).(Ast.Expresion)
					instruccion.GetValue(*scope)
				} else if tipo_abstracto == Ast.INSTRUCCION {
					instruccion := i.Instrucciones.GetValue(n).(Ast.Instruccion)
					resultado := instruccion.Run(scope)
					//Verificar si el resultado es un bool o un string
					if resultado.(Ast.TipoRetornado).Tipo == Ast.STRING {
						//Agregar a la consola
						scope.Consola += resultado.(Ast.TipoRetornado).Valor.(string) + "\n"
					}
					if resultado.(Ast.TipoRetornado).Tipo == Ast.ERROR {
						//Agregar a errores
						scope.Errores.Add(resultado.(Ast.TipoRetornado).Valor)
						scope.Consola += resultado.(Ast.TipoRetornado).Valor.(Errores.CustomSyntaxError).Msg + "\n"
					}
				}
				n++
			}
		} else {
			//Es falsa, buscar en la lista si hay otras
			//Llamada recursiva o fin
			//Recorrer la lista de ifs y else
			for j := 0; j < i.Lista_if_else.Len(); j++ {
				newScope := Ast.NewScope("if", scope)
				resultado := GetResultado(i.Lista_if_else.GetValue(j).(IF), &newScope, 0)
				if resultado.Valor == true {
					newScope.UpdateScopeGlobal()
					return Ast.TipoRetornado{
						Valor: true,
						Tipo:  Ast.BOOLEAN,
					}
				}
				if resultado.Tipo == Ast.ERROR {
					newScope.Errores.Add(resultado.Valor)
					newScope.Consola += resultado.Valor.(Errores.CustomSyntaxError).Msg + "\n"
					newScope.UpdateScopeGlobal()
					return Ast.TipoRetornado{
						Valor: true,
						Tipo:  Ast.BOOLEAN,
					}
				}
			}
			return Ast.TipoRetornado{
				Valor: false,
				Tipo:  Ast.BOOLEAN,
			}
		}
	} else {
		//No es booleano, entonces generar un error semántico
		//fmt.Println("Error semántico, la expresión no es un booleano")
		msg := "Semantic error, the condition of the expression is not a boolean expression." +
			" -- Line:" + strconv.Itoa(i.Fila) + " Column: " + strconv.Itoa(i.Columna)
		nError := Errores.NewError(i.Fila, i.Columna, msg)
		nError.Tipo = Ast.ERROR_SEMANTICO
		return Ast.TipoRetornado{
			Tipo:  Ast.ERROR,
			Valor: nError,
		}
	}
	//Se acabo todo y retornar un true de finalizado
	return Ast.TipoRetornado{
		Valor: true,
		Tipo:  Ast.BOOLEAN,
	}
}
