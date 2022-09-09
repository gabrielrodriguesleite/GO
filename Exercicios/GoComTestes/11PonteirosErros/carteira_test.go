package carteira

import "testing"

func TestCarteira(t *testing.T) {
	t.Run("Depositar", func(t *testing.T) {

		carteira := Carteira{}

		carteira.Depositar(Bitcoin(10))
		carteira.Depositar(Bitcoin(10))

		resultado := carteira.Saldo()

		esperado := Bitcoin(20)

		if resultado != esperado {
			t.Errorf("Resultado: '%s', esperado: '%s'", resultado, esperado)
		}
	})

	t.Run("Retirar", func(t *testing.T) {
		carteira := Carteira{saldo: Bitcoin(30)}

		carteira.Retirar(Bitcoin(14))

		resultado := carteira.Saldo()

		esperado := Bitcoin(16)

		if resultado != esperado {
			t.Errorf("Resultado: '%s', esperado: '%s'", resultado, esperado)
		}
	})
}
