package main

import "fmt"

func main() {

	// declaração de um array de inteiros, vazio com 5 posições
	var a [5]int
	fmt.Println("emp:", a)

	// usando um indice podemos definir um valor e ler um valor
	// a função len() retorna o tamanho
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))
}
