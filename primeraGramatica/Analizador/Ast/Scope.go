package Ast

import (
	"strings"
)

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

func (scope *Scope) Exist(ident string) bool {
	id := strings.ToUpper(ident)

	for scope_actual := scope; scope_actual != nil; scope_actual = scope_actual.prev {

		for key, _ := range scope_actual.tablaSimbolos {
			if key == id {
				return true
			}
		}
	}
	return false
}

func (scope *Scope) Exist_actual(ident string) bool {
	id := strings.ToUpper(ident)
	for key, _ := range scope.tablaSimbolos {
		if key == id {
			return true
		}
	}
	return false
}

func (scope *Scope) UpdateSimbolo(ident string, valorNuevo Simbolo) {
	id := strings.ToUpper(ident)
	for scope_actual := scope; scope_actual != nil; scope_actual = scope_actual.prev {

		for key, _ := range scope_actual.tablaSimbolos {
			if key == id {
				scope.tablaSimbolos[key] = valorNuevo
			}
		}
	}

}

func (scope *Scope) GetSimbolo(ident string) Simbolo {
	id := strings.ToUpper(ident)

	for scope_actual := scope; scope_actual != nil; scope_actual = scope_actual.prev {
		for key, simboloRetorno := range scope_actual.tablaSimbolos {
			if key == id {
				nsimbolo := simboloRetorno.(Simbolo)
				return nsimbolo
			}
		}
	}
	var simboloNull Simbolo
	return simboloNull
}
