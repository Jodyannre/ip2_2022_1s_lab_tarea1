package main

import (
	"fmt"
	"primeraGramatica/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Visitor struct {
	*parser.BaseNparserListener
}

func NewVisitor() *Visitor {
	return new(Visitor)
}

func (v *Visitor) ExitProd_inicio(ctx *parser.Prod_inicioContext) {
	/*fmt.Println("Ingreso>>>>>>>>>>>>>>>")
	data, err := json.MarshalIndent(ctx, "", "  ")
	if err != nil {
		panic(err)
	}
	stringEsQuery := string(data)
	fmt.Println(stringEsQuery)
	*/
}

func (v *Visitor) VisitProd_numero(ctx *parser.Prod_numeroContext) int {
	fmt.Println(ctx.GetText())
	fmt.Println("Entro aquí")

	//n, _ := strconv.ParseInt(ctx.GetText(), 0, 64)
	return 5
}

func (v *Visitor) ExitOp_arit(ctx *parser.Op_aritContext) {
	/*
		izq := ctx.GetOp_left().GetText()
		der := ctx.GetOp_right().GetText()
		fmt.Println("Estos son números")
		fmt.Println(izq)
		fmt.Println(der)
	*/

}

func main() {
	var input string = "5+5"

	//Obteniendo el input
	cadena_entrada := antlr.NewInputStream(input)

	//Creando el lexer
	lexer := parser.NewNlexer(cadena_entrada)

	streamTokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	nparser := parser.NewNparser(streamTokens)
	tree := nparser.Inicio()

	var nvisitor = NewVisitor()

	antlr.ParseTreeWalkerDefault.Walk(nvisitor, tree)

	/*
		//Leyendo tokens del lexer e imprimiendolos
		for {
			t := lexer.NextToken()
			if t.GetTokenType() == antlr.TokenEOF {
				break
			}
			fmt.Printf("%s (%q)\n",
				lexer.SymbolicNames[t.GetTokenType()], t.GetText())
		}
	*/

	//Stream para el parser
	//parser_stream := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)

	//Creando el parser
	//nparser := parser.NewNparser(parser_stream)

}
