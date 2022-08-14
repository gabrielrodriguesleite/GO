package main

import (
	"fmt"
	"sync"
	"time"
)

// Podemos usar Wait Group para aguradar multiplas
// goroutines finalizarem.

// Esta função será executada a cada goroutine
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	// Sleep por 1s para simular uma tarefa pesada
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	fmt.Println("Inicio em:", time.Now())

	// WaitGroup é usado para aguardar por todas goroutines lançadas
	// finalizarem.
	// *Se uma WaitGroup é explicitamente passada para funções, isso
	// deve ser feito por ponteiro (by pointer)
	var wg sync.WaitGroup

	// Lançar 5 goroutines e incrementar o contador do WaitGroup
	// uma vez por goroutine lançada.
	for i := 1; i <= 5; i++ {
		wg.Add(1)

		// Para evitar o uso do mesmo valor de i em cada goroutine fechamento
		// Para saber mais: https://golang.org/doc/faq#closures_and_goroutines
		i := i

		// Envolver a chamada do worker para certificar que WaitGroup será avisado
		// que este worker finalizou.
		// Desta forma o worker não tem qe ser avisado sobre os concorrentes primitivos
		// involvidos na sua execução.
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	// Bloqueia até que o contador do WaitGroup volte a zero.
	// Todos os workers notificaram que foram completados.
	wg.Wait()

	fmt.Println("Fim em:", time.Now())

}

// Esta abordagem não implementou um via direta para propagar erros dos workers.
// Para saber sobre casos avançados de uso considere usar o pacote errgroup
// Mais aqui: https://pkg.go.dev/golang.org/x/sync/errgroup
