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
	ducacaoA := medirTempoDeResposta(a)
	ducacaoB := medirTempoDeResposta(b)

	if ducacaoA < ducacaoB {
		return a
	}
	return b
}

func medirTempoDeResposta(URL string) time.Duration {
	inicio := time.Now()
	http.Get(URL)
	return time.Since(inicio)
}
