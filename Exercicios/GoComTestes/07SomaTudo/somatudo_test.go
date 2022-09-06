package somatudo

import "testing"

func TestSomaTudo(t *testing.T) {
	resultado := SomaTudo([]int{1, 2}, []int{0, 9})
	esperado := []int{3, 9}

	if resultado != esperado {
		t.Errorf("Resultado: '%v', esperado '%v'", resultado, esperado)
	}
}
