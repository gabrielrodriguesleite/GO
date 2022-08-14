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

	// Nesse momento é seguro acessar ops pois sabemos que nenhuma goroutine
	// está escrevendo ele.
	// Ler atomics de maneira segura é possível usando funções como
	// atomic.LoadUint64
	fmt.Println("ops:", ops)

}

// Experamos ter 50000 operações.
// Se tivesse sido usada a forma não atomica para incrementar ex: ops++
// A cada execução teríamos um valor diferente por que cada goroutine
// estaria interferindo na execução da outra.
// Além disso temos falhas de data race quando rodamos com a flag -race
