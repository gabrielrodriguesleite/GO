package main

import (
	"bytes"
	"testing"
)

//spy sleeper
type Sleeper interface {
	Sleep()
}

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

func Test(t *testing.T) {
	buffer := &bytes.Buffer{}

	Contagem(buffer)

	resultado := buffer.String()
	esperado := `3
2
1
Vai!`

	if resultado != esperado {
		t.Errorf("resultado: '%s', esperado '%s'", resultado, esperado)
	}
}
