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

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		i := i

		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()

	fmt.Println("Fim em:", time.Now())

}
