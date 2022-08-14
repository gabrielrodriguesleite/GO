package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// O mecanismo primario para manegar estado em Go é a
// comunicação por meio dos canais. Vimos um exemplo disso
// em worker pools. Existem outras opções para manejar o estado.
// Aqui vemos o uso do pacote sync/atomic para atomic counters
// acessados por multimplas goroutines
func main() {

	// Usaremos um inteiro sem sinal para representar nosso contador.
	var ops uint64

	// Um WaitGroup auxiliara a aguardar todas as goroutines finalizarem
	// seu trabalho
	var wg sync.WaitGroup

	// Iniciamos 50 goroutines onde cada uma incrementar o contador
	// exatamente 1000 vezes.
	for i := 0; i < 50; i++ {
		wg.Add(1)

		// Para atomicamente incrementar o contador usamos AddUint64,
		// passando o endereço do nosso contador ops com a sintaxe &
		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}

	// Aguardamos até que todas as goroutines tenham feito seu trabalho
	wg.Wait()

	fmt.Println("ops:", ops)
}
