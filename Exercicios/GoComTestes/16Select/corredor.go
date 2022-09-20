package main

import (
	"net/http"
)

// Problemática:
// Fazer uma função chamada Corredor que recebe duas URLs que
// "competirão" entre si através de uma chamada HTTP GET onde a
// primeira URL a responder será retornada.
// Caso nenhuma responda dentro de 10 segundos a função deve
// retornar uma erro.

/*

+ net/http - chamadas http
+ net/http/httptest - testes http
+ goroutines - paralelismo
+ select - sincronia de processos

*/

func Corredor(a, b string) (vencedor string) {
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}

func ping(URL string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(URL)
		ch <- true
	}()
	return ch
}
