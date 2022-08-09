package main

import "fmt"

func main() {

	var a = "initial"
	fmt.Println(a)

	// varios valores atribuidos ao mesmo tempo
	var b, c int = 1, 2
	fmt.Println(b, c)

	// o tipo pode ser inferido
	var d = true
	fmt.Println(d)

	// por padrão int recebe 0
	var e int
	fmt.Println(e)

	// := declara e inicia a variavel
	f := "apple"
	fmt.Println(f)

}
