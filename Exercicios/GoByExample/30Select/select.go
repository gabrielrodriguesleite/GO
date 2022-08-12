package main

import (
	"fmt"
	"time"
)

// Selet permite aguradar a multiplas operações de channels
// Combinar goroutines e channels com select é uma ferramenta poderosa de go
func main() {

	// Neste exemplo vamos usar select entre 2 channels
	c1 := make(chan string)
	c2 := make(chan string)

	// Cada channel vai receber um valor apos um tempo para
	// simular por exemplo uma operação bloqueante RPC
	// executando em goroutines concorrentes
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// Usa-se select para aguardar abos os valores simultaneamente
	// imprimido cada um assim que chegue
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}

	// recebemos o valor "one" e então "two" como esperado

	// o tempo total da execução deve estar em torno de 2s
	//  como ambos Sleep's executam concorrentes

}
