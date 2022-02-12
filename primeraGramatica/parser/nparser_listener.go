// Code generated from Nparser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Nparser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// NparserListener is a complete listener for a parse tree produced by Nparser.
type NparserListener interface {
	antlr.ParseTreeListener

	// EnterProd_inicio is called when entering the prod_inicio production.
	EnterProd_inicio(c *Prod_inicioContext)

	// EnterOp_arit is called when entering the op_arit production.
	EnterOp_arit(c *Op_aritContext)

	// EnterUnario is called when entering the unario production.
	EnterUnario(c *UnarioContext)

	// EnterProd_numero is called when entering the prod_numero production.
	EnterProd_numero(c *Prod_numeroContext)

	// EnterProd_decimal is called when entering the prod_decimal production.
	EnterProd_decimal(c *Prod_decimalContext)

	// EnterOp_multi is called when entering the op_multi production.
	EnterOp_multi(c *Op_multiContext)

	// EnterOp_div is called when entering the op_div production.
	EnterOp_div(c *Op_divContext)

	// EnterOp_mod is called when entering the op_mod production.
	EnterOp_mod(c *Op_modContext)

	// EnterOp_sum is called when entering the op_sum production.
	EnterOp_sum(c *Op_sumContext)

	// EnterOp_resta is called when entering the op_resta production.
	EnterOp_resta(c *Op_restaContext)

	// ExitProd_inicio is called when exiting the prod_inicio production.
	ExitProd_inicio(c *Prod_inicioContext)

	// ExitOp_arit is called when exiting the op_arit production.
	ExitOp_arit(c *Op_aritContext)

	// ExitUnario is called when exiting the unario production.
	ExitUnario(c *UnarioContext)

	// ExitProd_numero is called when exiting the prod_numero production.
	ExitProd_numero(c *Prod_numeroContext)

	// ExitProd_decimal is called when exiting the prod_decimal production.
	ExitProd_decimal(c *Prod_decimalContext)

	// ExitOp_multi is called when exiting the op_multi production.
	ExitOp_multi(c *Op_multiContext)

	// ExitOp_div is called when exiting the op_div production.
	ExitOp_div(c *Op_divContext)

	// ExitOp_mod is called when exiting the op_mod production.
	ExitOp_mod(c *Op_modContext)

	// ExitOp_sum is called when exiting the op_sum production.
	ExitOp_sum(c *Op_sumContext)

	// ExitOp_resta is called when exiting the op_resta production.
	ExitOp_resta(c *Op_restaContext)
}
