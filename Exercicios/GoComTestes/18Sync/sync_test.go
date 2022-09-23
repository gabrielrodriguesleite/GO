package main

import "testing"

// PROBLEMÁTICA: O objetivo é criar um contador seguro para concorrência.
// Primeiramente será desenvolvido um contador não seguro para verificar
// o funcionamento em um ambiente com apenas 1 linha de execução.
// Então serão utilizados várias go routines para expor a falha e então
// consertar.
func TestContador(t *testing.T) {
	t.Run("incrementar o contador 3 vezes resulta no valor 3", func(t *testing.T) {
		contador := Contador{}
		contador.Incrementa()
		contador.Incrementa()
		contador.Incrementa()

		verificaContador(t, contador, 3)
	})
}

func verificaContador(t *testing.T, resultado Contador, esperado int) {
	t.Helper()
	if resultado.Valor() != esperado {
		t.Errorf("resultado %d, esperado %d", resultado.Valor(), 3)
	}
}
