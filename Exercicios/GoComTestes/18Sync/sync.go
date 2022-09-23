package main

type Contador struct {
	valor int
}

func (c *Contador) Incrementa() {
	c.valor++
}

func (c *Contador) Valor() (retorno int) {
	return c.valor
}
