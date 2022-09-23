package main

import "testing"

// PROBLEMÁTICA: O objetivo é criar um contador seguro para concorrência.
// Primeiramente será desenvolvido um contador não seguro para verificar
// o funcionamento em um ambiente com apenas 1 linha de execução.
// Então serão utilizados várias go routines para expor a falha e então
// consertar.
func TestContador(t *testing.T) {
	t.Run("", func(t *testing.T) {

	})
}
