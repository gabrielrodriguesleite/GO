package main

import "testing"

func TestOlaTu(t *testing.T) {

	// Função auxiliar para verificação dos resultados (asserção).
	verificaMensagemCorreta := func(t testing.TB, resultado, esperado string) {
		// É muito importante indicar uma função aulixiliar de testes usando
		// "t.Helper()", assim caso o teste falhe, go irá indicar a linha do código
		// que falhou ao invés de indicar estas linhas.
		t.Helper()
		if resultado != esperado {
			t.Errorf("Resultado: '%s'\nEsperado: '%s'", resultado, esperado)
		}
	}

	// TDD RESULTADO 01:

	// Teste implementado:
	// R: Fail = função undefined

	/* go test 02TesteOlaTu/*
	# command-line-arguments [command-line-arguments.test]
	02TesteOlaTu/ola-tu_test.go:7:15: undefined: OlaTu
	02TesteOlaTu/ola-tu_test.go:11:5: t.ErrorF undefined (type *testing.T has no field or method ErrorF)
	FAIL    command-line-arguments [build failed]
	FAIL
	*/

	// EXEMPLO DE ERRO QUANDO É PASSADO UM PARAMETRO A MAIS:
	/* go test 02TesteOlaTu/*
	# command-line-arguments [command-line-arguments.test]
	02TesteOlaTu/ola-tu_test.go:7:32: too many arguments in call to OlaTu
					have (string, string)
					want (string)
	FAIL    command-line-arguments [build failed]
	FAIL*/
	t.Run("Teste se diz olá para as pessoas", func(t *testing.T) {

		resultado := OlaTu("Gabriel")
		esperado := "Olá, Gabriel"

		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("Teste se diz 'Olá, mundo' quando uma string vazia for passada", func(t *testing.T) {
		resultado := OlaTu("")
		esperado := "Olá, mundo"

		verificaMensagemCorreta(t, resultado, esperado)
	})
}
