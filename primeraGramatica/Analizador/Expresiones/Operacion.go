package Expresiones

import (
	"fmt"
	dato_interface "primeraGramatica/Analizador/Ast"
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
}

func NewOperation(op_izq dato_interface.Expresion, operador string, op_der dato_interface.Expresion, unario bool) Operacion {
	nuevo := Operacion{
		operando_der: op_der,
		operando_izq: op_izq,
		operador:     operador,
		unario:       unario,
	}
	return nuevo
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
			return dato_interface.TipoRetornado{
				Tipo:  result_dominante,
				Valor: nil,
			}
		}

	case "-":
		result_dominante = resta_dominante[tipo_izq.Tipo][tipo_der.Tipo]

		if result_dominante == dato_interface.INTEGER {
			return dato_interface.TipoRetornado{
				Tipo:  result_dominante,
				Valor: tipo_izq.Valor.(int) - tipo_der.Valor.(int),
			}
		} else if result_dominante == dato_interface.REAL {

			return dato_interface.TipoRetornado{
				Tipo:  result_dominante,
				Valor: tipo_izq.Valor.(float64) - tipo_der.Valor.(float64),
			}
		} else if result_dominante == dato_interface.NULL {
			return dato_interface.TipoRetornado{
				Tipo:  result_dominante,
				Valor: nil,
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
			return dato_interface.TipoRetornado{
				Tipo:  result_dominante,
				Valor: tipo_izq.Valor.(float64) * tipo_der.Valor.(float64),
			}

		} else if result_dominante == dato_interface.NULL {
			return dato_interface.TipoRetornado{
				Tipo:  result_dominante,
				Valor: nil,
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
			return dato_interface.TipoRetornado{
				Tipo:  result_dominante,
				Valor: tipo_izq.Valor.(float64) / tipo_der.Valor.(float64),
			}

		} else if result_dominante == dato_interface.NULL {
			return dato_interface.TipoRetornado{
				Tipo:  result_dominante,
				Valor: nil,
			}

		}
	}

	return dato_interface.TipoRetornado{Tipo: dato_interface.NULL, Valor: nil}
}
