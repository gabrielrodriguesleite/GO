package inteiros

import (
	"fmt"
	"testing"
)

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

func ExampleAdiciona() {
	// Example funciona como teste, para mostrar os detalhes
	// rode com "go test -v 04Inteiros/*"
	soma := Adiciona(1, 5)
	fmt.Println(soma)
	// Output: 4
}
