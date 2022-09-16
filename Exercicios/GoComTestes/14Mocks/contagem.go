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

const ultimaPalavra = "Vai!"
const inicioContagem = 3

func Contagem(saida io.Writer) {
	for i := inicioContagem; i > 0; i-- {
		fmt.Fprintln(saida, i)
	}
	fmt.Fprint(saida, ultimaPalavra)
}

func main() {
	Contagem(os.Stdout)
}
