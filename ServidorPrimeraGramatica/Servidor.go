package main

import (
	"ServidorPrimeraGramatica/primeraGramatica"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mensaje desde GET")
}

func PostUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mensaje desde POST")
}

func main() {
	/*
		StrictSlash es para que la slash del final no afecte a la ruta
	*/

	//Crear el servidor
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/users", GetUsers).Methods("GET")
	r.HandleFunc("/api/users", PostUsers).Methods("POST")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./Web/")))
	server := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	primeraGramatica.Principal()
	fmt.Println("Escuchando.......")
	server.ListenAndServe()

}
