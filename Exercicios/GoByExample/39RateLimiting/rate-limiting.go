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

	// Este channel limiter vai receber um valor a cada 200ms.
	// Este é o regulador do nosso esquema limitante.
	limiter := time.Tick(200 * time.Millisecond)

	// Por bloquear na recepção do canal limiter antes de servir
	// cada requisição, limitamos os sistema a 1 requisição a
	// cada 200ms
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// Também podemos querer permitir surtos de requisições no
	// esquema limitante preservando os limites totais.
	// É possível conseguir isso bufferizando nosso canal limitante.
	// Este canal burtyLimiter vai permitir surtos de até 3 eventos
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
