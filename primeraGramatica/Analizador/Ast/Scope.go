package Ast

import "strings"

type Scope struct {
	name          string
	prev          *Scope
	tablaSimbolos map[string]interface{}
}

func NewScope(name string, prev *Scope) Scope {

	nuevo := Scope{name: name, prev: prev}
	nuevo.tablaSimbolos = make(map[string]interface{})
	return nuevo
}

func (scope *Scope) Add(simbolo Simbolo) {
	id := strings.ToUpper(simbolo.Identificador)
	scope.tablaSimbolos[id] = simbolo
}

func (scope *Scope) Exist(simbolo Simbolo) bool {
	id := strings.ToUpper(simbolo.Identificador)

	for scope_actual := scope; scope_actual != nil; scope_actual = scope_actual.prev {

		for key, _ := range scope_actual.tablaSimbolos {
			if key == id {
				return true
			}
		}
	}
	return false
}

func (scope *Scope) getSimbolo(ident string) Simbolo {
	id := strings.ToUpper(ident)

	for scope_actual := scope; scope_actual != nil; scope_actual = scope_actual.prev {

		for key, simboloRetorno := range scope_actual.tablaSimbolos {
			if key == id {
				return simboloRetorno.(Simbolo)
			}
		}
	}
	var simboloNull Simbolo
	return simboloNull
}

func (scope *Scope) updateSimbolo(ident string, valorNuevo Simbolo) {
	id := strings.ToUpper(ident)

	for scope_actual := scope; scope_actual != nil; scope_actual = scope_actual.prev {

		for key, _ := range scope_actual.tablaSimbolos {
			if key == id {
				scope.tablaSimbolos[key] = valorNuevo
			}
		}
	}

}
