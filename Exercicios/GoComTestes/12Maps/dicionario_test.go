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
		_, resultado := dicionario.Busca("desconhecida")

		comparaErro(t, resultado, ErrNaoEncontrado)
	})
}

func TestAdiciona(t *testing.T) {
	dicionario := Dicionario{}
	dicionario.Adiciona("teste", "isto é apenas um teste")
	esperado := "isto é apenas um teste"
	resultado, err := dicionario.Busca("teste")
	if err != nil {
		t.Fatal("não foi possível encontrar a palavra adicionada: ", err)
	}
	if esperado != resultado {
		t.Errorf("Resultado: '%s', esperado '%s'", resultado, esperado)
	}
}

func comparaErro(t *testing.T, resultado, esperado error) {
	t.Helper()
	if resultado != esperado {
		t.Errorf("Resultado: '%s', esperado '%s'", resultado, esperado)
	}
}

func comparaString(t *testing.T, resultado, esperado string) {
	t.Helper()

	if resultado != esperado {
		t.Errorf("Resultado: '%s', esperado '%s', dado: '%s'", resultado, esperado, "teste")
	}
}
