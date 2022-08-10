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

}
