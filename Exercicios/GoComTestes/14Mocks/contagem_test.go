package main

import (
	"bytes"
	"testing"
)

//spy sleeper
//mock
type SleeperSpy struct {
	Chamadas int
}

func (s *SleeperSpy) Sleep() {
	s.Chamadas++
}

// Spies são um tipo de mock no qual podemos gravar informações como por
// exemplo: quantas vezes a função foi chamada. Também podem gravar os
// argumentos definidos, etc.

const pausa = "pausa"
const escrita = "escrita"

type SpyContagemOperacoes struct {
	Chamadas []string
}

func (s *SpyContagemOperacoes) Sleep() {
	s.Chamadas = append(s.Chamadas, pausa)
}

func (s *SpyContagemOperacoes) Write(p []byte) (n int, err error) {
	s.Chamadas = append(s.Chamadas, escrita)
	return
}

func Test(t *testing.T) {
	t.Run("teste simples", func(t *testing.T) {

		buffer := &bytes.Buffer{}
		sleeperSpy := &SleeperSpy{}

		Contagem(buffer, sleeperSpy)

		resultado := buffer.String()
		esperado := `3
		2
		1
		Vai!`

		if resultado != esperado {
			t.Errorf("resultado: '%s', esperado '%s'", resultado, esperado)
		}

		if sleeperSpy.Chamadas != 4 {
			t.Errorf("não houve chamadas suficientes de sleep, esperando 4, resultado %d", sleeperSpy.Chamadas)
		}
	})
}
