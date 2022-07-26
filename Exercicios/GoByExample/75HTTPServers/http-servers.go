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
	// Este manipulador faz algo mais sofistiado pois lê o cabeçalho da requisição
	// http e devolve ele no corpo da requisição.
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	// Então são registrados os manipuladores para suas respectivas rotas usando
	// http.HandleFunc. Esta função configura o roteador padrão do pacote net/http
	// e leva uma função como argumento.
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// Por fim se chama ListenAndServe com a porta e o manipulador.
	// "nil" significa que iremos utilizar o roteador padrão que foi configurado
	// acima.
	http.ListenAndServe(":8090", nil)
}

// SAÍDA ESPERADA PARA ESTE CÓDIGO

// RODAR O SERVIDOR DESACOPLADO
// go run 75HTTPServers/http-servers.go &

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
