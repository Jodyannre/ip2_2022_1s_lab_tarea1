// Code generated from Nparser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Nparser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by Nparser.
type NparserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by Nparser#prod_inicio.
	VisitProd_inicio(ctx *Prod_inicioContext) interface{}

	// Visit a parse tree produced by Nparser#op_arit.
	VisitOp_arit(ctx *Op_aritContext) interface{}

	// Visit a parse tree produced by Nparser#unario.
	VisitUnario(ctx *UnarioContext) interface{}

	// Visit a parse tree produced by Nparser#prod_numero.
	VisitProd_numero(ctx *Prod_numeroContext) interface{}

	// Visit a parse tree produced by Nparser#prod_decimal.
	VisitProd_decimal(ctx *Prod_decimalContext) interface{}

	// Visit a parse tree produced by Nparser#op_multi.
	VisitOp_multi(ctx *Op_multiContext) interface{}

	// Visit a parse tree produced by Nparser#op_div.
	VisitOp_div(ctx *Op_divContext) interface{}

	// Visit a parse tree produced by Nparser#op_mod.
	VisitOp_mod(ctx *Op_modContext) interface{}

	// Visit a parse tree produced by Nparser#op_sum.
	VisitOp_sum(ctx *Op_sumContext) interface{}

	// Visit a parse tree produced by Nparser#op_resta.
	VisitOp_resta(ctx *Op_restaContext) interface{}
}
