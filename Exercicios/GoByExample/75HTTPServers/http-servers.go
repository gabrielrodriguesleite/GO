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
