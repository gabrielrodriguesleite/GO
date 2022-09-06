package soma

import "testing"

func TestSoma(t *testing.T) {
	t.Run("Coleção de 5 números", func(t *testing.T) {
		numeros := []int{1, 2, 3, 4, 5}
		resultado := Soma(numeros)
		esperado := 15
		if esperado != resultado {
			t.Errorf("Resultado: '%d', esperado: '%d', dado: '%v'", resultado, esperado, numeros)
		}
	})

	t.Run("Coleção de qualquer número", func(t *testing.T) {
		numeros := []int{3, 4, 5}
		resultado := Soma(numeros)
		esperado := 12
		if esperado != resultado {
			t.Errorf("Resultado: '%d', esperado: '%d', dado: '%v'", resultado, esperado, numeros)
		}
	})

}

// SAÍDA ESPERADA PARA O TESTE COVERAGE

/* go test -cover 06Soma/*
ok      command-line-arguments  0.004s  coverage: 100.0% of statements
*/
