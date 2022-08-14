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

	var ops uint64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("ops:", ops)
}
