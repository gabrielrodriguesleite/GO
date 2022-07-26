package main

import "fmt"

// essa função chama ela mesma até que atinja o caso base de fact(0)
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {

	fmt.Println(fact(7)) // 5040

	// closures podem ser recursivas, mas para isso é necessário declara previamente
	// com um tipo var explicitamente antes de defini-la
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		// como fib foi declarada previamente, go sabe onde chamar a função fib
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}
