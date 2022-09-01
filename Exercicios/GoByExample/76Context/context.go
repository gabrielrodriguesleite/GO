package main

import "net/http"

// No exemplo anterior foi implementado um simples server.
// Servers são úteis para demonstrar o controle de cancelamento com
// "context.Context". Um "Context" carrega deadlines, sinais de cancelamento
// e outros valores no escopo de requisições que cruzam entre API e rotinas go.

func hello(w http.ResponseWriter, req *http.Request) {

}

func main() {

	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)

}
