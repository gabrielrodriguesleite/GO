package main

import "sync"

type Contador struct {
	mu    sync.Mutex
	valor int
}

func (c *Contador) Incrementa() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.valor++
}

func (c *Contador) Valor() (retorno int) {
	return c.valor
}

// Mutex cria uma trava no nosso contador.
/* + Mutex é uma trava de exclusão mútua. O valor zero de um Mutex é um
Mutex destravado.
*/
