package main

import (
	"fmt"
	"strconv"
)

// Analizar/interpretar números à partir de strings é uma tarefa
// tanto comum como básica para muitas aplicações.
// Aqui temos exemplos disso em Go:

func main() {

	// O pacote integrado strconv provê intrerpretação de números.

	// Com ParseFloat, convertemos um ponto flutuante, 64 define a
	// precisão desejada.
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// Para ParseInt, 0 significa inferir a base 0 e 64 solicita
	// que o resultado caiba em 64 bits.
	i, _ := strconv.ParseInt("567", 0, 64)
	fmt.Println(i)

	// ParseInt vai reconhecer o formato hexadecimal.
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// Existe também a opção de converter um número não negativo
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// Atoi é uma forma conveniente de converter um inteiro com base 10
	k, _ := strconv.Atoi("123")
	fmt.Println(k)

	_, e := strconv.Atoi("wat")
	fmt.Println(e)

}
