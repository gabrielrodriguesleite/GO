package main

import "testing"

func TestOlaTu(t *testing.T) {

	resultado := OlaTu("Gabriel")
	esperado := "Olá, Gabriel"

	if resultado != esperado {
		t.ErrorF("Resultado: '%s'\nEsperado: '%s'", resultado, esperado)
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