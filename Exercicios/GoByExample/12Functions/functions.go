package main

import "fmt"

// função basica que pega 2 inteiros e retorna um inteiro com a soma dos dois
func plus(a int, b int) int {

	// go não retorna automaticamente o valor da expressão
	return a + b

}

func plusPlus(a, b, c int) int {

	return a + b + c

}

func main() {

	res := plus(1, 2)
	fmt.Println("1 + 2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1 + 2 + 3 =", res)

}
