package carteira

import "fmt"

// Para criarmos Bitcoin usamos a sintaxe Bitcoin(123)
type Bitcoin int

type Carteira struct {
	saldo Bitcoin
}

func (c *Carteira) Depositar(quantidade Bitcoin) {
	c.saldo += quantidade
}

func (c *Carteira) Saldo() Bitcoin {
	return c.saldo
}

func (c *Carteira) Retirar(quantidade Bitcoin) {
	c.saldo -= quantidade
}

// ----------------------------

type Stringer interface {
	String() string
}

// Com o uso da interface a invés de imprimir o valor do int ao usar o print com %d
// é possível usar print com %s e receber a string resposta da implementação abaixo
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
