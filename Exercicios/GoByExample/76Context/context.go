package main

import (
	"fmt"
	"net/http"
	"time"
)

// No exemplo anterior foi implementado um simples server.
// Servers são úteis para demonstrar o controle de cancelamento com
// "context.Context". Um "Context" carrega deadlines, sinais de cancelamento
// e outros valores no escopo de requisições que cruzam entre API e rotinas go.

func hello(w http.ResponseWriter, req *http.Request) {

	// Um "context.Context" é criado para cada requisição pelo mecanismo do
	// pacote net/http, e está disponível com o método "Context".
	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {

	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)

}

// SAÍDA ESPERADA PARA O CÓDIGO

// go run 76Context/context.go &

// COMPORTAMENTO AGUARDANDO 10s P/ GET
/* curl localhost:8090/hello
server: hello handler started
server: hello handler ended
hello
*/

// COMPORTAMENTO CANCELANDO A REQUISIÇÃO GET
/* curl localhost:8090/hello
server: hello handler started
^C
server: context canceled
server: hello handler ended
*/
