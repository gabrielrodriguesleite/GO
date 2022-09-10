package main

import "testing"

func TestBusca(t *testing.T) {
	dicionario := Dicionario{"teste": "isso é apenas um teste"}

	t.Run("palavra conhecida", func(t *testing.T) {
		resultado, _ := dicionario.Busca("teste")
		esperado := "isso é apenas um teste"
		comparaString(t, resultado, esperado)
	})

	t.Run("palavra desconhecida", func(t *testing.T) {
		_, err := dicionario.Busca("desconhecida")
		if err == nil {
			t.Fatal("é esperado que um erro seja obtido")
		}
	})
}

func comparaString(t *testing.T, resultado, esperado string) {
	t.Helper()

	if resultado != esperado {
		t.Errorf("Resultado: '%s', esperado '%s', dado: '%s'", resultado, esperado, "teste")
	}
}
