package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Pausa()
}

type SleeperPadrao struct{}

func (d *SleeperPadrao) Pausa() {
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

// Parece que está tudo correto porém o teste ainda não confirma que
// as chamadas da função sleeper ocorrem na ordem correta com as impressões.
// De fora intercalada atrazo, impressão.

type SleeperConfiguravel struct {
	duracao time.Duration
	pausa   func(time.Duration)
}

func (s *SleeperConfiguravel) Pausa() {
	s.pausa(s.duracao)
}

const ultimaPalavra = "Vai!"
const inicioContagem = 3

func Contagem(saida io.Writer, sleeper Sleeper) {
	for i := inicioContagem; i > 0; i-- {
		sleeper.Pausa()
		fmt.Fprintln(saida, i)
	}

	sleeper.Pausa()
	fmt.Fprint(saida, ultimaPalavra)
}

func main() {
	sleeper := &SleeperConfiguravel{duracao: 1 * time.Second}
	Contagem(os.Stdout, sleeper)
}
