// Code generated from Nparser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Nparser

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseNparserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseNparserVisitor) VisitProd_inicio(ctx *Prod_inicioContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNparserVisitor) VisitOp_arit(ctx *Op_aritContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNparserVisitor) VisitUnario(ctx *UnarioContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNparserVisitor) VisitProd_numero(ctx *Prod_numeroContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNparserVisitor) VisitProd_decimal(ctx *Prod_decimalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNparserVisitor) VisitOp_multi(ctx *Op_multiContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNparserVisitor) VisitOp_div(ctx *Op_divContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNparserVisitor) VisitOp_mod(ctx *Op_modContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNparserVisitor) VisitOp_sum(ctx *Op_sumContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseNparserVisitor) VisitOp_resta(ctx *Op_restaContext) interface{} {
	return v.VisitChildren(ctx)
}
