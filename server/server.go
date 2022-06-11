package server

import (
	"fmt"
	"log"
	"net/http"

	petname "github.com/dustinkirkland/golang-petname"
)

func Server() {
	log.Println("Starting Server")

	http.HandleFunc("/", home)
	http.HandleFunc("/names", petNames)

	err := http.ListenAndServe(":8080", nil)
	log.Println("ERRO: listen and serve: ", err)

}

func home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Welcome")

}
func petNames(w http.ResponseWriter, req *http.Request) {
	words := 2
	separator := "-"

	fmt.Fprint(w, petname.Generate(words, separator))

}
