package main

import "fmt"

func main() {

	// declaração de um array de inteiros, vazio com 5 posições
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))
}
