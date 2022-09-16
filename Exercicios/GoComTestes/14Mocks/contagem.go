package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type SleeperPadrao struct{}

func (d *SleeperPadrao) Sleep() {
	time.Sleep(1 * time.Second)
}

/* DEFINIÇÃO DO PROBLEMA:
Escrever um programa que conta regressivamente a partir de 3 e então
imprime "Vai!" (sem imprimir o zero)
*/

/*
Este exemplo deixa claro que o sleep irá atrasar os testes.
Neste caso é essencial usar a injeção de dependência pra
contornar este problema.
*/

const ultimaPalavra = "Vai!"
const inicioContagem = 3

func Contagem(saida io.Writer, sleeper Sleeper) {
	for i := inicioContagem; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(saida, i)
	}

	time.Sleep(1 * time.Second)
	fmt.Fprint(saida, ultimaPalavra)
}

func main() {
	sleeper := &SleeperPadrao{}
	Contagem(os.Stdout, sleeper)
}
