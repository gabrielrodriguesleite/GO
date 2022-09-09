package carteira

import "testing"

func TestCarteira(t *testing.T) {

	confirmaSaldo := func(t *testing.T, carteira Carteira, esperado Bitcoin) {
		t.Helper()
		resultado := carteira.Saldo()
		if resultado != esperado {
			t.Errorf("Resultado: '%s', esperado: '%s'", resultado, esperado)
		}
	}

	confirmaErro := func(t *testing.T, valor error, esperado string) {
		t.Helper()
		if valor == nil {
			t.Fatal("Esperava um erro mas nenhum ocorreu")
		}
		resultado := valor.Error()
		if resultado != esperado {
			t.Errorf("Resultado: '%s', esperado: '%s'", resultado, esperado)
		}
	}

	t.Run("Depositar", func(t *testing.T) {

		carteira := Carteira{}

		carteira.Depositar(Bitcoin(10))
		carteira.Depositar(Bitcoin(10))

		esperado := Bitcoin(20)

		confirmaSaldo(t, carteira, esperado)
	})

	t.Run("Retirar", func(t *testing.T) {
		carteira := Carteira{saldo: Bitcoin(30)}

		carteira.Retirar(Bitcoin(14))

		esperado := Bitcoin(16)

		confirmaSaldo(t, carteira, esperado)

	})

	t.Run("Retirar um saldo insuficiente", func(t *testing.T) {
		saldoInicial := Bitcoin(40)

		carteira := Carteira{saldo: saldoInicial}

		erro := carteira.Retirar(Bitcoin(100))

		confirmaSaldo(t, carteira, saldoInicial)

		confirmaErro(t, erro, "impossível retirar: saldo insuficiente")

	})
}
