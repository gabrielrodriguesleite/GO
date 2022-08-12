package main

import "fmt"

// É possível definir a finalidade do channel quando definimos uma função.
// Se deve ser de entrada ou de saída. Isso melhora a tipagem da aplicação.

// Esta função apenas caeita um channel que envie valores
// causando um erro de compilação ao tentar receber um valor por meio deste
// channel
func ping(pings chan<- string, msg string) {

	pings <- msg

}

// Esta função aceita um channel que recebe valores (pings)
// e um segundo que envia (pongs)
func pong(pings <-chan string, pongs chan<- string) {

	msg := <-pings
	pongs <- msg

}

func main() {

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)

}
