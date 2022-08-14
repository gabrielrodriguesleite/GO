package main

import (
	"fmt"
	"time"
)

// Rate limiting é um importante mecanismo para controlar
// uso de recursos e manter a qualidade do serviço.
// Elegantemente go suporta rate limiting com goroutines,
// channels e tickers.
func main() {

	// Primeiro vemos um rate limiting básico.
	// Supondo que queremos limitar nosso manipulador de requisições
	// de entrada. Nós vamos servir estas requisições por canais de
	// mesmo nome.
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
