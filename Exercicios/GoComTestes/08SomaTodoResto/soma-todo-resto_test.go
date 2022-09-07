package somatodoresto

import "testing"
import "reflect"

func TestSomaTodoResto(t *testing.T) {

	verificaSomas := func(t *testing.T, res, esp []int) {
		t.Helper()
		if !reflect.DeepEqual(res, esp) {
			t.Errorf("Resultado: '%v', esperado: '%v", res, esp)
		}
	}

	t.Run("Soma alguns slices.", func(t *testing.T) {

		res := SomaTodoResto([]int{1, 2}, []int{0, 9})
		esp := []int{2, 9}

		verificaSomas(t, res, esp)
	})

	t.Run("Soma slices vazios de forma segura.", func(t *testing.T) {

		res := SomaTodoResto([]int{}, []int{3, 4, 5})
		esp := []int{0, 9}

		verificaSomas(t, res, esp)
	})

}
