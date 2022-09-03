package main

import "testing"

func TestOlaTu(t *testing.T) {

	resultado := OlaTu("Gabriel")
	esperado := "Olá, Gabriel"

	if resultado != esperado {
		t.ErrorF("Resultado: '%s'\nEsperado: '%s'", resultado, esperado)
	}
}
