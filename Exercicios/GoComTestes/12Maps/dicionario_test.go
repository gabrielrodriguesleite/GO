package main

import "testing"

func TestBusca(t *testing.T) {
	dicionario := Dicionario{"teste": "isso é apenas um teste"}
	resultado := Busca(dicionario, "teste")
	esperado := "isso é apenas um teste"

	comparaString(t, resultado, esperado)

}

func comparaString(t *testing.T, resultado, esperado string) {
	t.Helper()

	if resultado != esperado {
		t.Errorf("Resultado: '%s', esperado '%s', dado: '%s'", resultado, esperado, "teste")
	}
}
