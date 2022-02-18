package primeraGramatica

import (

	//"primeraGramatica/Analizador/Ast"

	//"fmt"
	"ServidorPrimeraGramatica/primeraGramatica/Analizador/Ast"
	"ServidorPrimeraGramatica/primeraGramatica/Analizador/Errores"
	"ServidorPrimeraGramatica/primeraGramatica/Analizador/Visitantes"
	"ServidorPrimeraGramatica/primeraGramatica/parser"
	"fmt"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func Principal(input string) string {
	/*
			var input string = `declarar variable integer = <5>;
			declarar variable2 real = <5.5>;
			`


		var input string = `
					PUBLICO MAIN :clase [ principal( STRING { } ) :  metodo [
				##
				variable2 declarar real = <6.6>;
				declarar variable1 integer = <5>;
				declarar cadena1 string = <"hola">;
				if: (3>=5)[
					variable1 = <16>;
					variable1 = <variable1 + 1>;
					variable4 = <"hola2">;
					sentencia.consola(true + 1);
					sentencia.consola("prueba de un string");
				] entonces if:(true) [
					declarar cadena2 string = <"mundo">;
					sentencia.consola(cadena1+cadena2);
				] entonces [
					sentencia.consola("Estoy dentro del else");
				];
				sentencia.consola(cadena2);
				]]
			`
	*/

	//Obteniendo el input
	cadena_entrada := antlr.NewInputStream(input)

	//Creando los costum errors
	errores_sintacticos := &Errores.CustomErrorListener{}
	errores_lexicos := &Errores.CustomErrorListener{}

	//Creando el lexer
	lexer := parser.NewNlexer(cadena_entrada)

	//Modificando el listener de los errores
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errores_lexicos)

	//Obteniendo los tokens
	streamTokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	//Creando el parser
	nparser := parser.NewNparser(streamTokens)

	//Modificando el listener de los errores
	nparser.RemoveErrorListeners()
	nparser.AddErrorListener(errores_sintacticos)
	nparser.BuildParseTrees = true

	//Leyendo la producción de inicio
	tree := nparser.Inicio()

	//fmt.Printf("Errores léxicos \n %v", errores_lexicos.Errores)
	//fmt.Printf("\nErrores sintáticos")
	//fmt.Printf("\n%v", errores_sintacticos.Errores)

	//Visitador para recorrer el árbol
	var nvisitor = Visitantes.NewVisitor()

	//Agregar errores léxicos y sintácticos al array de errores en el nVisitor

	//Recorrer el array de errores léxicos
	for i := 0; i < len(errores_lexicos.Errores); i++ {
		//fmt.Printf("Error #"+strconv.Itoa(i)+" %v \n", errores_lexicos.Errores[i])
		errores_lexicos.Errores[i].Tipo = Ast.ERROR_LEXICO
		nvisitor.Errores.Add(errores_lexicos.Errores[i])
		//Actualizar la consola
		nvisitor.UpdateConsola("Lexical error, " + errores_lexicos.Errores[i].Msg + " -- Line: " +
			strconv.Itoa(errores_lexicos.Errores[i].Linea) + " Column: " +
			strconv.Itoa(errores_lexicos.Errores[i].Columna))
	}

	//Recorrer el array de errores sintácticos
	for i := 0; i < len(errores_sintacticos.Errores); i++ {
		//fmt.Printf("Error #"+strconv.Itoa(i)+" %v \n", errores_sintacticos.Errores[i])
		errores_sintacticos.Errores[i].Tipo = Ast.ERROR_SINTACTICO
		nvisitor.Errores.Add(errores_sintacticos.Errores[i])
		//Actualizar la consola
		nvisitor.UpdateConsola("Syntax error , " + errores_sintacticos.Errores[i].Msg + " -- Line: " +
			strconv.Itoa(errores_sintacticos.Errores[i].Linea) + " Column: " +
			strconv.Itoa(errores_sintacticos.Errores[i].Columna))
	}

	antlr.ParseTreeWalkerDefault.Walk(nvisitor, tree)
	fmt.Println(nvisitor.GetConsola())
	return nvisitor.GetConsola()
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
