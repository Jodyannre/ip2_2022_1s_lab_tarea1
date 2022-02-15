package main

import (

	//"primeraGramatica/Analizador/Ast"

	"fmt"
	"primeraGramatica/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Visitador struct {
	*parser.BaseNparserListener
	consola string
}

func NewVisitor() *Visitador {
	return new(Visitador)
}

func (v *Visitador) ExitInicio(ctx *parser.InicioContext) {
	fmt.Println("Ingreso>>>>>>>>>>>>>>>")
	resultado := ctx.GetLista()
	fmt.Println(resultado)
	//Desde aquí ya tengo todo el resultado del parser, listo para ejecutar
}

type Uno struct {
	valor int
}

func (u *Uno) Modificar() {
	u.valor += 1
}

func main() {
	/*
		var input string = `declarar variable integer = <5>;
		declarar variable2 real = <5.5>;
		`*/
	var input string = `If: (5>6) [declarar variable1 integer=<5>; 
	declarar variable1 integer=<5>;
	declarar variable4 real=<5.5>;];`

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
