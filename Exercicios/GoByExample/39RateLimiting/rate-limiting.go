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

	fmt.Println("Sistema de limite de taxa (1/200ms)")
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

	// ----------------------------------------------------------
	fmt.Println("Com sistema que lida com surto de até 3 (2+1/200ms)")

	// Também podemos querer permitir surtos de requisições no
	// esquema limitante preservando os limites totais.
	// É possível conseguir isso bufferizando nosso canal limitante.
	// Este canal burtyLimiter vai permitir surtos de até 3 eventos
	burstyLimiter := make(chan time.Time, 3)

	// Lotando o canal para simular um surto controlado.
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// A cada 200ms vamos tentar incluir um novo valor ao burstyLimiter
	// até o limite máximo de 3.
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// Agora simulando 5 requisições de entrada. As 3 primeiras delas
	// vão se beneficiar da capacidade de absorver surto de burstyLimiter
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

// Rodando a aplicação vemos como a mesma lida com a primeira leva de
// requisições, uma a cada 200ms em média como desejado.

// Para a segunda leva de requisições servimos as 3 primeiras
// imediatamente por conta do sistema que lida com surto em seguida
// são servidas as últimas 2 com ~200ms de atrazo cada.
