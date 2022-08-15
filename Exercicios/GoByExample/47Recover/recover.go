package main

import "fmt"

// É possível recuperar de um panic com a função recover.
// recover pode ser usado para evitar que a aplicação seja abortada
// por um panic e permitir que a aplicação continue rodando

func mayPanic() {
	panic("a problem")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error: \n", r)
		}
	}()

	mayPanic()

	fmt.Println("After mayPanic()")
}
