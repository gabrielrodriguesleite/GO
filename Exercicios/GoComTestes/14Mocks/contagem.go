package main

import (
	"fmt"
	"io"
	"os"
)

/* DEFINIÇÃO DO PROBLEMA:
Escrever um programa que conta regressivamente a partir de 3 e então
imprime "Vai!" (sem imprimir o zero)
*/

func Contagem(saida io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(saida, i)
	}
	fmt.Fprint(saida, "Vai!")
}

func main() {
	Contagem(os.Stdout)
}
