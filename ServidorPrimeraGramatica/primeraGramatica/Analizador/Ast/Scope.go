package Ast

import (
	"strings"

	"github.com/colegno/arraylist"
)

type Scope struct {
	Nombre        string
	prev          *Scope
	tablaSimbolos map[string]interface{}
	Errores       *arraylist.List
	Consola       string
}

func NewScope(name string, prev *Scope) Scope {

	nuevo := Scope{Nombre: name, prev: prev}
	nuevo.Errores = arraylist.New()
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
				scope_actual.tablaSimbolos[key] = valorNuevo
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

func (s *Scope) UpdateScopeGlobal() {
	//Obtener el scope global
	var scope_global *Scope
	if s.prev != nil {
		for scope_global = s; scope_global.prev != nil; scope_global = scope_global.prev {
			//Buscando el scope global
		}
	} else {
		scope_global = s
	}
	scope_global.Consola += s.Consola
	for i := 0; i < s.Errores.Len(); i++ {
		elemento := s.Errores.GetValue(i)
		scope_global.Errores.Add(elemento)
	}
}
