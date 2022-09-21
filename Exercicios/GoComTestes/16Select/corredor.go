package main

import (
	"fmt"
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
+ select - sincronia de processos ("escuta vários canais")

*/

var limiteDeDezSegundos = 10 * time.Second

func Corredor(a, b string) (vencedor string, erro error) {
	return Configuravel(a, b, limiteDeDezSegundos)
}

func Configuravel(a, b string, tempoLimite time.Duration) (vencedor string, erro error) {
	select {
	case <-ping(a): // o primeiro canal (go routine) que enviar pelo canal primeiro é retornado
		return a, nil
	case <-ping(b): // time.After retorna um canal que recebe a mensagem quando o tempo é atingido
		return b, nil
	case <-time.After(tempoLimite): // caso para que a função não fique bloqueado para sempre
		return "", fmt.Errorf("tempo limite de espera excedido para %s e %s", a, b)
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
