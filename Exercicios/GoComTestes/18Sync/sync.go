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

func NovoContador() *Contador {
	return &Contador{}
}

// Mutex cria uma trava no nosso contador.
/* + Mutex é uma trava de exclusão mútua. O valor zero de um Mutex é um
Mutex destravado.
*/

/* UM EXEMPLO ERRADO SERIA:
type Contador struct {
    sync.Mutex
    valor int
}

func (c *Contador) Incrementa() {
    c.Lock()
    defer c.Unlock()
    c.valor++
}

Dessa forma o tipo embutido se torna parte da inteface pública
já que os métodos Lock e Unlock ficarão expostos.

Quando devemos usar Mutex e quando devemos usar Channel?
MAIS SOBRE: https://github.com/golang/go/wiki/MutexOrChannel
*/

// go vet - analisa o código.
// Para ser usado em roteiros de construção na busca de problemas no código.
