package iteracao

import "testing"

func TestRepetir(t *testing.T) {
	repeticoes := Repetir("a")
	esperado := "aaaaa"

	if repeticoes != esperado {
		t.Errorf("Esperado: '%s' mas obteve: '%s'", esperado, repeticoes)
	}
}
