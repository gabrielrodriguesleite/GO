package main

import (
	"fmt"
)

func main() {
	var a int
	var operador string
	var b int

	fmt.Println("Entre com um número então com um operador e então com outro número:")

	_, err := fmt.Scanln(&a, &operador, &b)

	if err != nil {
		fmt.Println(err)
	} else if operador == "+" {
		fmt.Println("Total:", a+b)
	} else if operador == "-" {
		fmt.Print("Total:", a-b)
	} else if operador == "*" {
		fmt.Println("Total:", a*b)
	} else if operador == "/" {
		if b == 0 {
			fmt.Println("Infinito. (Divisão por 0)")
		} else {
			fmt.Println("Total:", float64(a)/float64(b))
		}
	} else {
		fmt.Println("Operador inválido, use + - * ou /")
	}
}
