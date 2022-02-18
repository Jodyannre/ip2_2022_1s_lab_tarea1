package Expresiones

import (
	dato_interface "ServidorPrimeraGramatica/primeraGramatica/Analizador/Ast"
	"ServidorPrimeraGramatica/primeraGramatica/Analizador/Errores"
	"fmt"
	"math"
	"strconv"
)

var suma_dominante = [5][5]dato_interface.TipoDato{
	{dato_interface.INTEGER, dato_interface.REAL, dato_interface.STRING, dato_interface.NULL, dato_interface.NULL},
	{dato_interface.REAL, dato_interface.REAL, dato_interface.STRING, dato_interface.NULL, dato_interface.NULL},
	{dato_interface.STRING, dato_interface.STRING, dato_interface.STRING, dato_interface.STRING, dato_interface.NULL},
	{dato_interface.NULL, dato_interface.NULL, dato_interface.STRING, dato_interface.NULL, dato_interface.NULL},
	{dato_interface.NULL, dato_interface.NULL, dato_interface.NULL, dato_interface.NULL, dato_interface.NULL},
}

var resta_dominante = [5][5]dato_interface.TipoDato{
	{dato_interface.INTEGER, dato_interface.REAL, dato_interface.NULL, dato_interface.NULL, dato_interface.NULL},
	{dato_interface.REAL, dato_interface.REAL, dato_interface.NULL, dato_interface.NULL, dato_interface.NULL},
	{dato_interface.NULL, dato_interface.NULL, dato_interface.NULL, dato_interface.NULL, dato_interface.NULL},
	{dato_interface.NULL, dato_interface.NULL, dato_interface.NULL, dato_interface.NULL, dato_interface.NULL},
	{dato_interface.NULL, dato_interface.NULL, dato_interface.NULL, dato_interface.NULL, dato_interface.NULL},
}

var mul_div_dominante = [5][5]dato_interface.TipoDato{
	{dato_interface.INTEGER, dato_interface.REAL, dato_interface.NULL, dato_interface.NULL, dato_interface.NULL},
	{dato_interface.REAL, dato_interface.REAL, dato_interface.NULL, dato_interface.NULL, dato_interface.NULL},
	{dato_interface.NULL, dato_interface.NULL, dato_interface.NULL, dato_interface.NULL, dato_interface.NULL},
	{dato_interface.NULL, dato_interface.NULL, dato_interface.NULL, dato_interface.NULL, dato_interface.NULL},
	{dato_interface.NULL, dato_interface.NULL, dato_interface.NULL, dato_interface.NULL, dato_interface.NULL},
}

type Operacion struct {
	operando_der dato_interface.Expresion
	operador     string
	operando_izq dato_interface.Expresion
	unario       bool
	Fila         int
	Columna      int
}

func NewOperation(op_izq dato_interface.Expresion, operador string,
	op_der dato_interface.Expresion, unario bool, fila, columna int) Operacion {
	nuevo := Operacion{
		operando_der: op_der,
		operando_izq: op_izq,
		operador:     operador,
		unario:       unario,
		Fila:         fila,
		Columna:      columna,
	}
	return nuevo
}

func (op Operacion) GetTipo() (dato_interface.TipoDato, dato_interface.TipoDato) {
	return dato_interface.EXPRESION, dato_interface.PRIMITIVO
}

func (op Operacion) GetValue(entorno dato_interface.Scope) dato_interface.TipoRetornado {
	var tipo_izq dato_interface.TipoRetornado
	var tipo_der dato_interface.TipoRetornado
	var result_dominante dato_interface.TipoDato

	if op.unario {
		tipo_izq = op.operando_izq.GetValue(entorno)
	} else {
		tipo_izq = op.operando_izq.GetValue(entorno)
		tipo_der = op.operando_der.GetValue(entorno)
	}

	switch op.operador {
	case "+":
		result_dominante = suma_dominante[tipo_izq.Tipo][tipo_der.Tipo]

		if result_dominante == dato_interface.INTEGER {
			return dato_interface.TipoRetornado{
				Tipo:  result_dominante,
				Valor: tipo_izq.Valor.(int) + tipo_der.Valor.(int),
			}
		} else if result_dominante == dato_interface.REAL {

			if tipo_izq.Tipo == dato_interface.INTEGER {
				tipo_izq = dato_interface.TipoRetornado{
					Valor: float64(tipo_izq.Valor.(int)),
					Tipo:  dato_interface.REAL,
				}
			}
			if tipo_der.Tipo == dato_interface.INTEGER {
				tipo_der = dato_interface.TipoRetornado{
					Valor: float64(tipo_der.Valor.(int)),
					Tipo:  dato_interface.REAL,
				}
			}

			return dato_interface.TipoRetornado{
				Tipo:  result_dominante,
				Valor: tipo_izq.Valor.(float64) + tipo_der.Valor.(float64),
			}
		} else if result_dominante == dato_interface.STRING {
			cadena_izq := fmt.Sprintf("%v", tipo_izq.Valor)
			cadena_der := fmt.Sprintf("%v", tipo_der.Valor)
			return dato_interface.TipoRetornado{
				Tipo:  result_dominante,
				Valor: cadena_izq + cadena_der,
			}
		} else if result_dominante == dato_interface.NULL {
			/*
				return dato_interface.TipoRetornado{
					Tipo:  result_dominante,
					Valor: nil,
				}
			*/
			msg := "Semantic error, can't add " + dato_interface.ValorTipoDato[tipo_izq.Tipo] +
				" type to " + dato_interface.ValorTipoDato[tipo_der.Tipo] +
				" type. -- Line: " + strconv.Itoa(op.Fila) +
				" Column: " + strconv.Itoa(op.Columna)
			nError := Errores.NewError(op.Fila, op.Columna, msg)
			nError.Tipo = dato_interface.ERROR_SEMANTICO
			return dato_interface.TipoRetornado{
				Tipo:  dato_interface.ERROR,
				Valor: nError,
			}
		}

	case "-":
		if op.unario {
			//Es una operación unaria
			if tipo_izq.Tipo == dato_interface.REAL {
				//Es un real
				return dato_interface.TipoRetornado{
					Tipo:  dato_interface.REAL,
					Valor: tipo_izq.Valor.(float64) * -1,
				}
			} else {
				//Es un int
				return dato_interface.TipoRetornado{
					Tipo:  dato_interface.INTEGER,
					Valor: tipo_izq.Valor.(int) * -1,
				}
			}

		}

		result_dominante = resta_dominante[tipo_izq.Tipo][tipo_der.Tipo]

		if result_dominante == dato_interface.INTEGER {
			return dato_interface.TipoRetornado{
				Tipo:  result_dominante,
				Valor: tipo_izq.Valor.(int) - tipo_der.Valor.(int),
			}
		} else if result_dominante == dato_interface.REAL {

			if tipo_izq.Tipo == dato_interface.INTEGER {
				tipo_izq = dato_interface.TipoRetornado{
					Valor: float64(tipo_izq.Valor.(int)),
					Tipo:  dato_interface.REAL,
				}
			}
			if tipo_der.Tipo == dato_interface.INTEGER {
				tipo_der = dato_interface.TipoRetornado{
					Valor: float64(tipo_der.Valor.(int)),
					Tipo:  dato_interface.REAL,
				}
			}

			return dato_interface.TipoRetornado{
				Tipo:  result_dominante,
				Valor: tipo_izq.Valor.(float64) - tipo_der.Valor.(float64),
			}
		} else if result_dominante == dato_interface.NULL {
			/*
				return dato_interface.TipoRetornado{
					Tipo:  result_dominante,
					Valor: nil,
				}
			*/
			msg := "Semantic error, can't subtract " + dato_interface.ValorTipoDato[tipo_izq.Tipo] +
				" type to " + dato_interface.ValorTipoDato[tipo_der.Tipo] +
				" type. -- Line: " + strconv.Itoa(op.Fila) +
				" Column: " + strconv.Itoa(op.Columna)
			nError := Errores.NewError(op.Fila, op.Columna, msg)
			nError.Tipo = dato_interface.ERROR_SEMANTICO
			return dato_interface.TipoRetornado{
				Tipo:  dato_interface.ERROR,
				Valor: nError,
			}
		}

	case "*":
		result_dominante = mul_div_dominante[tipo_izq.Tipo][tipo_der.Tipo]
		if result_dominante == dato_interface.INTEGER {
			return dato_interface.TipoRetornado{
				Tipo:  result_dominante,
				Valor: tipo_izq.Valor.(int) * tipo_der.Valor.(int),
			}

		} else if result_dominante == dato_interface.REAL {
			if tipo_izq.Tipo == dato_interface.INTEGER {
				tipo_izq = dato_interface.TipoRetornado{
					Valor: float64(tipo_izq.Valor.(int)),
					Tipo:  dato_interface.REAL,
				}
			}
			if tipo_der.Tipo == dato_interface.INTEGER {
				tipo_der = dato_interface.TipoRetornado{
					Valor: float64(tipo_der.Valor.(int)),
					Tipo:  dato_interface.REAL,
				}
			}
			return dato_interface.TipoRetornado{
				Tipo:  result_dominante,
				Valor: tipo_izq.Valor.(float64) * tipo_der.Valor.(float64),
			}

		} else if result_dominante == dato_interface.NULL {
			/*
				return dato_interface.TipoRetornado{
					Tipo:  result_dominante,
					Valor: nil,
				}
			*/
			msg := "Semantic error, can't multiply " + dato_interface.ValorTipoDato[tipo_izq.Tipo] +
				" type to " + dato_interface.ValorTipoDato[tipo_der.Tipo] +
				" type. -- Line: " + strconv.Itoa(op.Fila) +
				" Column: " + strconv.Itoa(op.Columna)
			nError := Errores.NewError(op.Fila, op.Columna, msg)
			nError.Tipo = dato_interface.ERROR_SEMANTICO
			return dato_interface.TipoRetornado{
				Tipo:  dato_interface.ERROR,
				Valor: nError,
			}

		}
	case "/":
		result_dominante = mul_div_dominante[tipo_izq.Tipo][tipo_der.Tipo]
		if result_dominante == dato_interface.INTEGER {
			return dato_interface.TipoRetornado{
				Tipo:  result_dominante,
				Valor: tipo_izq.Valor.(int) / tipo_der.Valor.(int),
			}

		} else if result_dominante == dato_interface.REAL {
			if tipo_izq.Tipo == dato_interface.INTEGER {
				tipo_izq = dato_interface.TipoRetornado{
					Valor: float64(tipo_izq.Valor.(int)),
					Tipo:  dato_interface.REAL,
				}
			}
			if tipo_der.Tipo == dato_interface.INTEGER {
				tipo_der = dato_interface.TipoRetornado{
					Valor: float64(tipo_der.Valor.(int)),
					Tipo:  dato_interface.REAL,
				}
			}
			return dato_interface.TipoRetornado{
				Tipo:  result_dominante,
				Valor: tipo_izq.Valor.(float64) / tipo_der.Valor.(float64),
			}

		} else if result_dominante == dato_interface.NULL {
			/*
				return dato_interface.TipoRetornado{
					Tipo:  result_dominante,
					Valor: nil,
				}
			*/
			msg := "Semantic error, can't divide " + dato_interface.ValorTipoDato[tipo_izq.Tipo] +
				" type to " + dato_interface.ValorTipoDato[tipo_der.Tipo] +
				" type. -- Line: " + strconv.Itoa(op.Fila) +
				" Column: " + strconv.Itoa(op.Columna)
			nError := Errores.NewError(op.Fila, op.Columna, msg)
			nError.Tipo = dato_interface.ERROR_SEMANTICO
			return dato_interface.TipoRetornado{
				Tipo:  dato_interface.ERROR,
				Valor: nError,
			}

		}
	case "%":
		result_dominante = mul_div_dominante[tipo_izq.Tipo][tipo_der.Tipo]
		if result_dominante == dato_interface.INTEGER {
			return dato_interface.TipoRetornado{
				Tipo:  result_dominante,
				Valor: tipo_izq.Valor.(int) % tipo_der.Valor.(int),
			}

		} else if result_dominante == dato_interface.REAL {
			if tipo_izq.Tipo == dato_interface.INTEGER {
				tipo_izq = dato_interface.TipoRetornado{
					Valor: float64(tipo_izq.Valor.(int)),
					Tipo:  dato_interface.REAL,
				}
			}
			if tipo_der.Tipo == dato_interface.INTEGER {
				tipo_der = dato_interface.TipoRetornado{
					Valor: float64(tipo_der.Valor.(int)),
					Tipo:  dato_interface.REAL,
				}
			}
			return dato_interface.TipoRetornado{
				Tipo:  result_dominante,
				Valor: math.Mod(tipo_izq.Valor.(float64), tipo_der.Valor.(float64)),
			}

		} else if result_dominante == dato_interface.NULL {
			/*
				return dato_interface.TipoRetornado{
					Tipo:  result_dominante,
					Valor: nil,
				}
			*/
			msg := "Semantic error, can't divide " + dato_interface.ValorTipoDato[tipo_izq.Tipo] +
				" type to " + dato_interface.ValorTipoDato[tipo_der.Tipo] +
				" type. -- Line: " + strconv.Itoa(op.Fila) +
				" Column: " + strconv.Itoa(op.Columna)
			nError := Errores.NewError(op.Fila, op.Columna, msg)
			nError.Tipo = dato_interface.ERROR_SEMANTICO
			return dato_interface.TipoRetornado{
				Tipo:  dato_interface.ERROR,
				Valor: nError,
			}

		}
	case "AND":
		if tipo_der.Tipo != dato_interface.BOOLEAN || tipo_izq.Tipo != dato_interface.BOOLEAN {
			msg := "Semantic error, can't logically multiply " + dato_interface.ValorTipoDato[tipo_izq.Tipo] +
				" type to " + dato_interface.ValorTipoDato[tipo_der.Tipo] +
				" type. -- Line: " + strconv.Itoa(op.Fila) +
				" Column: " + strconv.Itoa(op.Columna)
			nError := Errores.NewError(op.Fila, op.Columna, msg)
			nError.Tipo = dato_interface.ERROR_SEMANTICO
			return dato_interface.TipoRetornado{
				Tipo:  dato_interface.ERROR,
				Valor: nError,
			}

		}
		return dato_interface.TipoRetornado{
			Tipo:  dato_interface.BOOLEAN,
			Valor: tipo_izq.Valor.(bool) && tipo_der.Valor.(bool),
		}
	case "OR":
		if tipo_der.Tipo != dato_interface.BOOLEAN || tipo_izq.Tipo != dato_interface.BOOLEAN {
			msg := "Semantic error, can't logically add " + dato_interface.ValorTipoDato[tipo_izq.Tipo] +
				" type to " + dato_interface.ValorTipoDato[tipo_der.Tipo] +
				" type. -- Line: " + strconv.Itoa(op.Fila) +
				" Column: " + strconv.Itoa(op.Columna)
			nError := Errores.NewError(op.Fila, op.Columna, msg)
			nError.Tipo = dato_interface.ERROR_SEMANTICO
			return dato_interface.TipoRetornado{
				Tipo:  dato_interface.ERROR,
				Valor: nError,
			}

		}
		return dato_interface.TipoRetornado{
			Tipo:  dato_interface.BOOLEAN,
			Valor: tipo_izq.Valor.(bool) || tipo_der.Valor.(bool),
		}

	case ">":
		var val_der interface{}
		var val_izq interface{}
		result_dominante = suma_dominante[tipo_izq.Tipo][tipo_der.Tipo]
		if result_dominante == dato_interface.INTEGER || result_dominante == dato_interface.REAL {

			if tipo_izq.Tipo == dato_interface.REAL || tipo_der.Tipo == dato_interface.REAL {

				if tipo_izq.Tipo == dato_interface.REAL {
					val_izq = tipo_izq.Valor.(float64)
				} else {
					val_izq = (float64)(tipo_izq.Valor.(int))
				}

				if tipo_izq.Tipo == dato_interface.REAL {
					val_der = tipo_der.Valor.(float64)
				} else {
					val_der = (float64)(tipo_der.Valor.(int))
				}

				return dato_interface.TipoRetornado{
					Tipo:  dato_interface.BOOLEAN,
					Valor: val_izq.(float64) > val_der.(float64),
				}

			} else {
				val_izq = tipo_izq.Valor.(int)
				val_der = tipo_der.Valor.(int)
				return dato_interface.TipoRetornado{
					Tipo:  dato_interface.BOOLEAN,
					Valor: val_izq.(int) > val_der.(int),
				}
			}
		}
		msg := "Semantic error, can't compare " + dato_interface.ValorTipoDato[tipo_izq.Tipo] +
			" with " + dato_interface.ValorTipoDato[tipo_der.Tipo] +
			" type. -- Line: " + strconv.Itoa(op.Fila) +
			" Column: " + strconv.Itoa(op.Columna)
		nError := Errores.NewError(op.Fila, op.Columna, msg)
		nError.Tipo = dato_interface.ERROR_SEMANTICO
		return dato_interface.TipoRetornado{
			Tipo:  dato_interface.ERROR,
			Valor: nError,
		}
	case ">=":
		var val_der interface{}
		var val_izq interface{}
		result_dominante = suma_dominante[tipo_izq.Tipo][tipo_der.Tipo]
		if result_dominante == dato_interface.INTEGER || result_dominante == dato_interface.REAL {

			if tipo_izq.Tipo == dato_interface.REAL || tipo_der.Tipo == dato_interface.REAL {

				if tipo_izq.Tipo == dato_interface.REAL {
					val_izq = tipo_izq.Valor.(float64)
				} else {
					val_izq = (float64)(tipo_izq.Valor.(int))
				}

				if tipo_izq.Tipo == dato_interface.REAL {
					val_der = tipo_der.Valor.(float64)
				} else {
					val_der = (float64)(tipo_der.Valor.(int))
				}

				return dato_interface.TipoRetornado{
					Tipo:  dato_interface.BOOLEAN,
					Valor: val_izq.(float64) >= val_der.(float64),
				}

			} else {
				val_izq = tipo_izq.Valor.(int)
				val_der = tipo_der.Valor.(int)
				return dato_interface.TipoRetornado{
					Tipo:  dato_interface.BOOLEAN,
					Valor: val_izq.(int) >= val_der.(int),
				}
			}
		}
		msg := "Semantic error, can't compare " + dato_interface.ValorTipoDato[tipo_izq.Tipo] +
			" with " + dato_interface.ValorTipoDato[tipo_der.Tipo] +
			" type. -- Line: " + strconv.Itoa(op.Fila) +
			" Column: " + strconv.Itoa(op.Columna)
		nError := Errores.NewError(op.Fila, op.Columna, msg)
		nError.Tipo = dato_interface.ERROR_SEMANTICO
		return dato_interface.TipoRetornado{
			Tipo:  dato_interface.ERROR,
			Valor: nError,
		}

	}
	return dato_interface.TipoRetornado{Tipo: dato_interface.NULL, Valor: nil}
}
