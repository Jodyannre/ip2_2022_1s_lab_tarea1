// Code generated from Nparser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Nparser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// NparserListener is a complete listener for a parse tree produced by Nparser.
type NparserListener interface {
	antlr.ParseTreeListener

	// EnterInicio is called when entering the inicio production.
	EnterInicio(c *InicioContext)

	// EnterInstrucciones is called when entering the instrucciones production.
	EnterInstrucciones(c *InstruccionesContext)

	// EnterInstruccion is called when entering the instruccion production.
	EnterInstruccion(c *InstruccionContext)

	// EnterDeclaracion is called when entering the declaracion production.
	EnterDeclaracion(c *DeclaracionContext)

	// EnterAsignacion is called when entering the asignacion production.
	EnterAsignacion(c *AsignacionContext)

	// EnterExpresion is called when entering the expresion production.
	EnterExpresion(c *ExpresionContext)

	// EnterTipo is called when entering the tipo production.
	EnterTipo(c *TipoContext)

	// EnterOperador_comparacion is called when entering the operador_comparacion production.
	EnterOperador_comparacion(c *Operador_comparacionContext)

	// EnterOperador_logico is called when entering the operador_logico production.
	EnterOperador_logico(c *Operador_logicoContext)

	// EnterOperador_aritmetico is called when entering the operador_aritmetico production.
	EnterOperador_aritmetico(c *Operador_aritmeticoContext)

	// EnterControl_IF_simple is called when entering the control_IF_simple production.
	EnterControl_IF_simple(c *Control_IF_simpleContext)

	// EnterControl_IF_entonces is called when entering the control_IF_entonces production.
	EnterControl_IF_entonces(c *Control_IF_entoncesContext)

	// EnterControl_if_entonces is called when entering the control_if_entonces production.
	EnterControl_if_entonces(c *Control_if_entoncesContext)

	// EnterControl_else is called when entering the control_else production.
	EnterControl_else(c *Control_elseContext)

	// EnterFuncion_imprimir is called when entering the funcion_imprimir production.
	EnterFuncion_imprimir(c *Funcion_imprimirContext)

	// ExitInicio is called when exiting the inicio production.
	ExitInicio(c *InicioContext)

	// ExitInstrucciones is called when exiting the instrucciones production.
	ExitInstrucciones(c *InstruccionesContext)

	// ExitInstruccion is called when exiting the instruccion production.
	ExitInstruccion(c *InstruccionContext)

	// ExitDeclaracion is called when exiting the declaracion production.
	ExitDeclaracion(c *DeclaracionContext)

	// ExitAsignacion is called when exiting the asignacion production.
	ExitAsignacion(c *AsignacionContext)

	// ExitExpresion is called when exiting the expresion production.
	ExitExpresion(c *ExpresionContext)

	// ExitTipo is called when exiting the tipo production.
	ExitTipo(c *TipoContext)

	// ExitOperador_comparacion is called when exiting the operador_comparacion production.
	ExitOperador_comparacion(c *Operador_comparacionContext)

	// ExitOperador_logico is called when exiting the operador_logico production.
	ExitOperador_logico(c *Operador_logicoContext)

	// ExitOperador_aritmetico is called when exiting the operador_aritmetico production.
	ExitOperador_aritmetico(c *Operador_aritmeticoContext)

	// ExitControl_IF_simple is called when exiting the control_IF_simple production.
	ExitControl_IF_simple(c *Control_IF_simpleContext)

	// ExitControl_IF_entonces is called when exiting the control_IF_entonces production.
	ExitControl_IF_entonces(c *Control_IF_entoncesContext)

	// ExitControl_if_entonces is called when exiting the control_if_entonces production.
	ExitControl_if_entonces(c *Control_if_entoncesContext)

	// ExitControl_else is called when exiting the control_else production.
	ExitControl_else(c *Control_elseContext)

	// ExitFuncion_imprimir is called when exiting the funcion_imprimir production.
	ExitFuncion_imprimir(c *Funcion_imprimirContext)
}
