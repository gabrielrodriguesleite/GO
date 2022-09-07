package carteira

import "testing"

func TestCarteira(t *testing.T) {
	carteira := Carteira{}

	carteira.Depositar(10)

	resultado := carteira.Saldo()

	esperado := 10

	if resultado != esperado {
		t.Errorf("Resultado: '%d', esperado: '%d'", resultado, esperado)
	}

}
