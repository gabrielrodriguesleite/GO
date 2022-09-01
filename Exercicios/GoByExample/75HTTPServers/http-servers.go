package main

import (
	"fmt"
	"net/http"
)

// Escrever um servido http básico é uma tarefa fácil usando o pacote
// net/http que vem com Go

/*
	Um conceito fundamental em servidores net/http é manipuladores (handlers).

Um manipulador é um objeto que implementa a interface http.Handler.
Uma forma comum de escrever manipuladores é usando a função http.HandlerFunc
que adapta a função com a assinatura apropriada.
*/
func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
	/*
	   	Funções que servem com manipuladores recebem um http.ResponseWriter e

	   um http.Request como argumentos. O ResponseWrites é usado para preencher o
	   corpo da Response. Neste exemplo a response leva apenas um "hello\n."
	*/
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}

// SAÍDA ESPERADA PARA ESTE CÓDIGO

// go run 75HTTPServers/http-servers.go

/* curl localhost:8090/hello
hello
*/

/*curl localhost:8090/headers
User-Agent: curl/7.68.0
Accept: * /*
*/

/* curl localhost:8090/
404 page not found
*/
