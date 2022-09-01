package main

import (
	"fmt"
	"net/http"
)

// Escrever um servido http básico é uma tarefa fácil usando o pacote
// net/http que vem com Go

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func main() {

}
