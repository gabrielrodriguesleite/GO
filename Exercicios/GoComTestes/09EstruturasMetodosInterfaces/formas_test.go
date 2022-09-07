package formas

import (
	"testing"
)

func TestPerimetro(t *testing.T) {
	res := Perimetro(10.0, 10.0)
	esp := 40.0

	if res != esp {
		t.Errorf("Resultado: '%.2f' esperado: '%.2f'", res, esp)
	}
}

func TestArea(t *testing.T) {
	res := Area(12.0, 6.0)
	esp := 72.0

	if res != esp {
		t.Errorf("Resultado: '%.2f' esperado: '%.2f'", res, esp)
	}
}
