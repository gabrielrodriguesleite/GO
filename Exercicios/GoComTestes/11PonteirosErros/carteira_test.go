package carteira

import "testing"

func TestCarteira(t *testing.T) {
	carteira := Carteira{}

	carteira.Depositar(Bitcoin(10))

	resultado := carteira.Saldo()

	esperado := Bitcoin(20)

	if resultado != esperado {
		t.Errorf("Resultado: '%s', esperado: '%s'", resultado, esperado)
	}

}
