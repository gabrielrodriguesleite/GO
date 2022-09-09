package carteira

import (
	"errors"
	"fmt"
)

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

func (c *Carteira) Retirar(quantidade Bitcoin) error {
	if quantidade > c.saldo {
		return errors.New("impossível retirar: saldo insuficiente")
	}
	c.saldo -= quantidade
	return nil
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

// PONTEIROS

/*
+ GO copia os valores quando são passados para funções/métodos. Então, se estiver
escrevendo uma função que precise mudar o estado, você precisará de um ponteiro para
o valor que você quer mudar.

+ O fato de que Go pega um cópia dos valores é muito útil na maior parte do tempo,
mas às vezes você não vai querer que o seu sistema faça cópia de alguma coisa. Nesse
caso, você precisa passar uma referência. Podemos, por exemplo, ter dados muito
grandes, ou coisas que você talvez pretenda ter apenas uma instância (como conexões
a banco de dados).

+ Ponteiros podem ser "nil". "Nil" também é útil para descrever algo que está faltando.

*/

// ERROS

// "Não apenas verifique por erros, mas trate-os graciosamente:"
// Main em: https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully

// TIPOS
/*
+ Tipos são úteis para adicionar domínios mais específicos a valores.

+ Tipos também permitem implementar interfaces.
*/
