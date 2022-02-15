// Code generated from Nparser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Nparser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseNparserListener is a complete listener for a parse tree produced by Nparser.
type BaseNparserListener struct{}

var _ NparserListener = &BaseNparserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseNparserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseNparserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseNparserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseNparserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterInicio is called when production inicio is entered.
func (s *BaseNparserListener) EnterInicio(ctx *InicioContext) {}

// ExitInicio is called when production inicio is exited.
func (s *BaseNparserListener) ExitInicio(ctx *InicioContext) {}

// EnterInstrucciones is called when production instrucciones is entered.
func (s *BaseNparserListener) EnterInstrucciones(ctx *InstruccionesContext) {}

// ExitInstrucciones is called when production instrucciones is exited.
func (s *BaseNparserListener) ExitInstrucciones(ctx *InstruccionesContext) {}

// EnterInstruccion is called when production instruccion is entered.
func (s *BaseNparserListener) EnterInstruccion(ctx *InstruccionContext) {}

// ExitInstruccion is called when production instruccion is exited.
func (s *BaseNparserListener) ExitInstruccion(ctx *InstruccionContext) {}

// EnterDeclaracion is called when production declaracion is entered.
func (s *BaseNparserListener) EnterDeclaracion(ctx *DeclaracionContext) {}

// ExitDeclaracion is called when production declaracion is exited.
func (s *BaseNparserListener) ExitDeclaracion(ctx *DeclaracionContext) {}

// EnterAsignacion is called when production asignacion is entered.
func (s *BaseNparserListener) EnterAsignacion(ctx *AsignacionContext) {}

// ExitAsignacion is called when production asignacion is exited.
func (s *BaseNparserListener) ExitAsignacion(ctx *AsignacionContext) {}

// EnterExpresion is called when production expresion is entered.
func (s *BaseNparserListener) EnterExpresion(ctx *ExpresionContext) {}

// ExitExpresion is called when production expresion is exited.
func (s *BaseNparserListener) ExitExpresion(ctx *ExpresionContext) {}

// EnterTipo is called when production tipo is entered.
func (s *BaseNparserListener) EnterTipo(ctx *TipoContext) {}

// ExitTipo is called when production tipo is exited.
func (s *BaseNparserListener) ExitTipo(ctx *TipoContext) {}

// EnterOperador_comparacion is called when production operador_comparacion is entered.
func (s *BaseNparserListener) EnterOperador_comparacion(ctx *Operador_comparacionContext) {}

// ExitOperador_comparacion is called when production operador_comparacion is exited.
func (s *BaseNparserListener) ExitOperador_comparacion(ctx *Operador_comparacionContext) {}

// EnterOperador_logico is called when production operador_logico is entered.
func (s *BaseNparserListener) EnterOperador_logico(ctx *Operador_logicoContext) {}

// ExitOperador_logico is called when production operador_logico is exited.
func (s *BaseNparserListener) ExitOperador_logico(ctx *Operador_logicoContext) {}

// EnterOperador_aritmetico is called when production operador_aritmetico is entered.
func (s *BaseNparserListener) EnterOperador_aritmetico(ctx *Operador_aritmeticoContext) {}

// ExitOperador_aritmetico is called when production operador_aritmetico is exited.
func (s *BaseNparserListener) ExitOperador_aritmetico(ctx *Operador_aritmeticoContext) {}

// EnterControl_IF_simple is called when production control_IF_simple is entered.
func (s *BaseNparserListener) EnterControl_IF_simple(ctx *Control_IF_simpleContext) {}

// ExitControl_IF_simple is called when production control_IF_simple is exited.
func (s *BaseNparserListener) ExitControl_IF_simple(ctx *Control_IF_simpleContext) {}

// EnterControl_IF_entonces is called when production control_IF_entonces is entered.
func (s *BaseNparserListener) EnterControl_IF_entonces(ctx *Control_IF_entoncesContext) {}

// ExitControl_IF_entonces is called when production control_IF_entonces is exited.
func (s *BaseNparserListener) ExitControl_IF_entonces(ctx *Control_IF_entoncesContext) {}

// EnterControl_if_entonces is called when production control_if_entonces is entered.
func (s *BaseNparserListener) EnterControl_if_entonces(ctx *Control_if_entoncesContext) {}

// ExitControl_if_entonces is called when production control_if_entonces is exited.
func (s *BaseNparserListener) ExitControl_if_entonces(ctx *Control_if_entoncesContext) {}

// EnterControl_else is called when production control_else is entered.
func (s *BaseNparserListener) EnterControl_else(ctx *Control_elseContext) {}

// ExitControl_else is called when production control_else is exited.
func (s *BaseNparserListener) ExitControl_else(ctx *Control_elseContext) {}

// EnterFuncion_imprimir is called when production funcion_imprimir is entered.
func (s *BaseNparserListener) EnterFuncion_imprimir(ctx *Funcion_imprimirContext) {}

// ExitFuncion_imprimir is called when production funcion_imprimir is exited.
func (s *BaseNparserListener) ExitFuncion_imprimir(ctx *Funcion_imprimirContext) {}
