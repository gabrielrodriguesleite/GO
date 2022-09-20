package main

import (
	"net/http"
	"time"
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
	inicioA := time.Now()
	http.Get(a)
	ducacaoA := time.Since(inicioA)

	inicioB := time.Now()
	http.Get(b)
	ducacaoB := time.Since(inicioB)

	if ducacaoA < ducacaoB {
		return a
	}
	return b
}
