package Visitantes

import (
	"fmt"
	"primeraGramatica/Analizador/Ast"
	"primeraGramatica/Analizador/Errores"
	"primeraGramatica/parser"

	"github.com/colegno/arraylist"
)

type Visitador struct {
	*parser.BaseNparserListener
	Consola string
	Errores arraylist.List
}

func NewVisitor() *Visitador {
	return new(Visitador)
}

func (v *Visitador) ExitInicio(ctx *parser.InicioContext) {
	fmt.Println("Ingreso>>>>>>>>>>>>>>>")
	instrucciones := ctx.GetLista()
	//fmt.Println(instrucciones)
	//Desde aquí ya tengo todo el resultado del parser, listo para ejecutar

	/*
		1) Declarar el scope
		2) Primera pasada para declarar todas las variables
		3) Correr todas las instrucciones
	*/

	//Creando el scope global
	entornoGlobal := Ast.NewScope("global", nil)

	//Variables para las declaraciones
	var actual interface{}
	var tipo, _ interface{}
	var respuesta interface{}

	//Primera pasada para agregar todas las declaraciones de las variables
	for i := 0; i < instrucciones.Len(); i++ {
		actual = instrucciones.GetValue(i)
		if actual != nil {
			_, tipo = actual.(Ast.Abstracto).GetTipo()
		} else {
			continue
		}
		if tipo.(Ast.TipoDato) == Ast.DECLARACION {
			//Declarar variables globales
			respuesta = actual.(Ast.Instruccion).Run(&entornoGlobal)
			if respuesta.(Ast.TipoRetornado).Tipo == Ast.ERROR {
				//Hay error y agregarlo a la lista de errores
				v.Errores.Add(respuesta)
				v.Consola += respuesta.(Ast.TipoRetornado).Valor.(Errores.CustomSyntaxError).Msg + "\n"
			}
		}
	}

	//Ejecutar todas instrucciones siguientes
	for i := 0; i < instrucciones.Len(); i++ {
		actual = instrucciones.GetValue(i)
		if actual != nil {
			_, tipo = actual.(Ast.Abstracto).GetTipo()
		} else {
			continue
		}
		if tipo.(Ast.TipoDato) != Ast.DECLARACION {
			//Declarar variables globales
			respuesta = actual.(Ast.Instruccion).Run(&entornoGlobal)
			if respuesta.(Ast.TipoRetornado).Tipo == Ast.ERROR {
				//Hay error y agregarlo a la lista de errores
				v.Errores.Add(respuesta)
				v.Consola += respuesta.(Ast.TipoRetornado).Valor.(Errores.CustomSyntaxError).Msg + "\n"
			}
			//Verificar que sea un string para agregarlo a la consola
			if respuesta.(Ast.TipoRetornado).Tipo == Ast.STRING {
				v.Consola += respuesta.(Ast.TipoRetornado).Valor.(string) + "\n"
			}
		}
	}
}

func (v *Visitador) GetConsola() string {
	return v.Consola
}

func (v *Visitador) UpdateConsola(entrada string) {
	v.Consola += v.Consola + entrada + "\n"
}
