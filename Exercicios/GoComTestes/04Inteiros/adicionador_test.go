package inteiros

import "testing"

func TestAdicionador(t *testing.T) {

	t.Run("Testa se Adiciona(2, 2) é 4", func(t *testing.T) {
		soma := Adiciona(2, 2)
		esperado := 4

		if soma != esperado {
			t.Errorf("Esperado: %d, resultado '%d'", esperado, soma)
		}
	})

	t.Run("Testa se Adiciona(4, 5) é 9", func(t *testing.T) {
		soma := Adiciona(4, 5)
		esperado := 9

		if soma != esperado {
			t.Errorf("Esperado: %d, resultado '%d'", esperado, soma)
		}
	})
}
