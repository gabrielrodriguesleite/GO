package main

import (
	"fmt"
	"time"
)

// Timeouts são importantes que se conectam com recursos externos
// ou que estejam vinculados de outra forma ao tempo de execução
// Implementar timeouts em  go é facil e elegante graças a channels e select
func main() {

	// Nesse exemplo suponha que nós estamos executando uma chamada externa
	// que retorna seu resultado no channel c1 após 2 segundos
	// Note que o channel é buffered então o envio na goroutine não é bloqueante
	// Este é um padrão comum para prevenir que goroutine vaze em caso do
	// channel nunca ser lido
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// Aqui o select implementa um timeout.
	// res := <-c1 aguarda o "result" e <-time.Afet aguarda um valor
	// ser enviado apos o timeout de 1s.
	// Já que o select continua com o primeiro recebimento que está pronto,
	// nos vamos receber o caso timeout se a operação levar mais que o 1s permitido
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
