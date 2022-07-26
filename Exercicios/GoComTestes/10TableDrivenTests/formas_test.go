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
		// "O teste é lido de forma mais clara como se fosse a 
		// afirmação da verdade, não uma sequência de operações." - Kent Beck
		{forma: Retangulo {Largura: 12, Altura: 6}, esperado: 72.0},
		{forma: Circulo   {Raio: 10},               esperado: 314.1592653589793},
		{forma: Triangulo {Base: 12, Altura: 6},    esperado: 36.0},
	}

	for _, tt := range testesArea {
		resultado := tt.forma.Area()
		if resultado != tt.esperado {
			t.Errorf("Resultado: '%.2f' esperado '%.2f'", resultado, tt.esperado)
		}
	}
}
