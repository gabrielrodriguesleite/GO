package main

import "fmt"

// o recurso de retornar multiplos valores é suportado por Go
// a assinatura da função mostra (int, int) indicando o retorno de 2 ints
func vals() (int, int) {
	return 3, 7
}

func main() {

	// multipla atribuição de multiplos retornos
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// retornos podem ser omitidos usando o _ blank identifier
	_, c := vals()
	fmt.Println(c)

}

// é bastante comum retornar o valor e um valor de erro
