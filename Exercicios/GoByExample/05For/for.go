package main

import "fmt"

func main() {

	// basico
	i := 1
	for i < 3 {
		fmt.Println(i)
		i = i + 1
	}

	// classico inicio/condição/postumo
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// como um while
	for {
		fmt.Println("loop")
		break
	}

	// uso do continue para pular interação
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}
