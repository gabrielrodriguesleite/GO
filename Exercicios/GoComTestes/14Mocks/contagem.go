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
	fmt.Fprint(saida, "3")
}

func main() {
	Contagem(os.Stdout)
}
