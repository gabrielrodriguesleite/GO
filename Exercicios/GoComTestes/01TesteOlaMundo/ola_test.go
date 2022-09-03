package main

import "testing"

func TestOla(t *testing.T) {
	resultado := Ola()
	esperado := "Olá, mundo"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}

// SAÍDA ESPERADA PARA ESTE CÓDIGO:

// UM CÓDIGO QUE TESTA OK
/*go test 01TesteOlaMundo/*
ok      command-line-arguments  0.003s
*/

// UM CÓDIGO COM FALHA IRIA IMPRIMIR ALGO ASSIM:
/* go test 01TesteOlaMundo/*
--- FAIL: TestOla (0.00s)
    ola_test.go:10: resultado 'Olá, mundo', esperado 'Olá, mndo'
FAIL
FAIL    command-line-arguments  0.003s
FAIL
*/
