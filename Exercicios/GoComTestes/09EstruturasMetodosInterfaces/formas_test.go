package formas

import (
	"testing"
)

func TestPerimetro(t *testing.T) {
	ret := Retangulo{10.0, 10.0}
	res := Perimetro(ret)
	esp := 40.0

	if res != esp {
		t.Errorf("Resultado: '%.2f' esperado: '%.2f'", res, esp)
	}
}

func TestArea(t *testing.T) {
	verificaArea := func(t *testing.T, forma Forma, esp float64) {
		t.Helper()
		res := forma.Area()

		if res != esp {
			t.Errorf("Resultado: '%.2f' esperado: '%.2f'", res, esp)
		}
	}

	t.Run("do retângulo:", func(t *testing.T) {

		ret := Retangulo{12.0, 6.0}
		esp := 72.0

		verificaArea(t, ret, esp)
	})

	t.Run("do círculo:", func(t *testing.T) {

		cir := Circulo{10}
		esp := 314.1592653589793

		verificaArea(t, cir, esp)
	})
}
