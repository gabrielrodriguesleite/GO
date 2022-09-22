package main

import "testing"

func TestPercorre(t *testing.T) {

	esperado := "Leite"

	var resultado []string

	x := struct {
		Nome string
	}{esperado}

	percorre(x, func(entrada string) {
		resultado = append(resultado, entrada)
	})

	if len(resultado) != 1 {
		t.Errorf("número incorreto de chamadas da função: resultado: %d, esperado %d", len(resultado), 1)
	}
}