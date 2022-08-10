package main

import "fmt"

// o recurso de retornar multiplos valores é suportado por Go
func vals() (int, int) {
	return 3, 7
}

func main() {

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

}

// é bastante comum retornar o valor e um valor de erro
