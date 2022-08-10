package main

import "fmt"

// Go suporta funções anônimas e estas podem formar "closures"
// funções anônimas são úteis quando precisamos definir uma função em linha sem ter que nomear ela.

// intSeq retorna outra função que definimos no seu corpo
// a função retornada envolve (enclausura) i para formar um closure

// MAIS AQUI: https://medium.com/@viniazvd/closures-bef1a912c506
// closure: é uma função interior que tem acessa a variábeis de uma função exterior
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	// atribuimos a função a variável nextInt, esta função captura a sua própria variável i
	nextInt := intSeq()

	// Sua variável i atualizada da cada chamada de nextInt()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// newInts terá sua própria variável i
	newInts := intSeq()
	fmt.Println(newInts())
}
