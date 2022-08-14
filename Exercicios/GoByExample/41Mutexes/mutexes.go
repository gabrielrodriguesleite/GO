package main

import (
	"fmt"
	"sync"
)

// Outra forma de acessar e alterar estado de dentro de goroutines
// de forma segura é pelo uso de mutex.

// Container possui um map de couters, neste exemplo o intuito é
// atualizar isto de forma concorrente por meio de goroutines.
// Incluimos um Mutex para sincronizar o acesso.
// Mutexes não devem ser copiados então usamos essa estrutura no
// código por meio de um ponteiro
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

// O procedimento deve ser bloquear o acesso com mutex.Lock()
// Então usar um defer para desbloquear no final da função com
// mutex.Unlock()
func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {

	// Não é necessário inicializar um mutex como seu valor zerado
	// já está pronto para o uso.
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	// Esta função incrementa em um loop um contador nomeado
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	// Rodando várias goroutines concorrentes.
	// Todas acessam o mesmo Container, e duas delas acesam o
	// mesmo contador.
	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	// Aguardar todas goroutines finalizarem
	wg.Wait()

	fmt.Println(c.counters)
}
