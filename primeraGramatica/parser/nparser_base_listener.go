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

// EnterProd_inicio is called when production prod_inicio is entered.
func (s *BaseNparserListener) EnterProd_inicio(ctx *Prod_inicioContext) {}

// ExitProd_inicio is called when production prod_inicio is exited.
func (s *BaseNparserListener) ExitProd_inicio(ctx *Prod_inicioContext) {}

// EnterOp_arit is called when production op_arit is entered.
func (s *BaseNparserListener) EnterOp_arit(ctx *Op_aritContext) {}

// ExitOp_arit is called when production op_arit is exited.
func (s *BaseNparserListener) ExitOp_arit(ctx *Op_aritContext) {}

// EnterUnario is called when production unario is entered.
func (s *BaseNparserListener) EnterUnario(ctx *UnarioContext) {}

// ExitUnario is called when production unario is exited.
func (s *BaseNparserListener) ExitUnario(ctx *UnarioContext) {}

// EnterProd_numero is called when production prod_numero is entered.
func (s *BaseNparserListener) EnterProd_numero(ctx *Prod_numeroContext) {}

// ExitProd_numero is called when production prod_numero is exited.
func (s *BaseNparserListener) ExitProd_numero(ctx *Prod_numeroContext) {}

// EnterProd_decimal is called when production prod_decimal is entered.
func (s *BaseNparserListener) EnterProd_decimal(ctx *Prod_decimalContext) {}

// ExitProd_decimal is called when production prod_decimal is exited.
func (s *BaseNparserListener) ExitProd_decimal(ctx *Prod_decimalContext) {}

// EnterOp_multi is called when production op_multi is entered.
func (s *BaseNparserListener) EnterOp_multi(ctx *Op_multiContext) {}

// ExitOp_multi is called when production op_multi is exited.
func (s *BaseNparserListener) ExitOp_multi(ctx *Op_multiContext) {}

// EnterOp_div is called when production op_div is entered.
func (s *BaseNparserListener) EnterOp_div(ctx *Op_divContext) {}

// ExitOp_div is called when production op_div is exited.
func (s *BaseNparserListener) ExitOp_div(ctx *Op_divContext) {}

// EnterOp_mod is called when production op_mod is entered.
func (s *BaseNparserListener) EnterOp_mod(ctx *Op_modContext) {}

// ExitOp_mod is called when production op_mod is exited.
func (s *BaseNparserListener) ExitOp_mod(ctx *Op_modContext) {}

// EnterOp_sum is called when production op_sum is entered.
func (s *BaseNparserListener) EnterOp_sum(ctx *Op_sumContext) {}

// ExitOp_sum is called when production op_sum is exited.
func (s *BaseNparserListener) ExitOp_sum(ctx *Op_sumContext) {}

// EnterOp_resta is called when production op_resta is entered.
func (s *BaseNparserListener) EnterOp_resta(ctx *Op_restaContext) {}

// ExitOp_resta is called when production op_resta is exited.
func (s *BaseNparserListener) ExitOp_resta(ctx *Op_restaContext) {}
