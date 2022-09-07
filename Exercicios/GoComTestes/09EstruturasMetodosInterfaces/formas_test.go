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
	t.Run("do retângulo:", func(t *testing.T) {

		ret := Retangulo{12.0, 6.0}
		res := Area(ret)
		esp := 72.0

		if res != esp {
			t.Errorf("Resultado: '%.2f' esperado: '%.2f'", res, esp)
		}
	})

	t.Run("do círculo:", func(t *testing.T) {

		cir := Circulo{10}
		res := Area(cir)
		esp := 314.1592653589793

		if res != esp {
			t.Errorf("Resultado: '%.2f' esperado: '%.2f'", res, esp)
		}
	})
}
