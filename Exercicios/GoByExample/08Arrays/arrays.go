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

	// declaração e inicialização de array em uma linha
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}
