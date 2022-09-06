package somatudo

import "testing"
import "reflect"

func TestSomaTudo(t *testing.T) {
	resultado := SomaTudo([]int{1, 2}, []int{0, 9})
	esperado := []int{3, 9}

	// Este teste falha por que go não permite comparar slice com algo além de nil.
	// if resultado != esperado {
	if !reflect.DeepEqual(resultado, esperado) {
		t.Errorf("Resultado: '%v', esperado '%v'", resultado, esperado)
	}
}
