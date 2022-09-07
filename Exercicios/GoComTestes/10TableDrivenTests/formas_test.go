package formas

import (
	"testing"
)

func TestArea(t *testing.T) {

	// struct anônima usada pra declara a "tabela"
	testesArea := []struct {
		forma    Forma
		esperado float64
	}{
		{Retangulo{12, 6}, 72.0},
		{Circulo{10}, 314.1592653589793},
		{Triangulo{12, 6}, 36.0},
	}

	for _, tt := range testesArea {
		resultado := tt.forma.Area()
		if resultado != tt.esperado {
			t.Errorf("Resultado: '%.2f' esperado '%.2f'", resultado, tt.esperado)
		}
	}
}
