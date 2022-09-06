package soma

import "testing"

func TestSoma(t *testing.T) {
	t.Run("Coleção de 5 números", func(t *testing.T) {
		numeros := [5]int{1, 2, 3, 4, 5}
		resultado := Soma(numeros)
		esperado := 15
		if esperado != resultado {
			t.Errorf("Resultado: '%d', esperado: '%d', dado: '%v'", resultado, esperado, numeros)
		}
	})
}
