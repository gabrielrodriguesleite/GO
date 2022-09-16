package main

import (
	"bytes"
	"testing"
)

//spy sleeper
type Sleeper interface {
	Sleep()
}

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
