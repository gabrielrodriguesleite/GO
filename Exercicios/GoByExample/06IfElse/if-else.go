package main

import "fmt"

func main() {

	// basico if else
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// if sem else
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// declaração pode preceder condição, a variável definida aqui está disponível em todos os braços
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "ha multiple digits")
	}

	// parentesis não são obrigatórios mas chaves são

	// não existe ternario então é preciso uma declaração completa inclusive para condições basicas

}
