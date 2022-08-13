package main

import "fmt"

// Podemos usar esta sintaxe para iterar sobre valores
// recebidos de channels

func main() {

	// Vamos iterar sobre 2 valores no channel queue
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// Este range itera sobre cada elemento conforme eles são
	// recebidos de queue.
	// Como fechamos o channel acima, a iteração termina após
	// receber os 2 elementos
	for elem := range queue {
		fmt.Println(elem)
	}

}

// Este exemplo mostra que é possível fechar um canal não vazio
// mas continuar recebendo os valores restantes.
