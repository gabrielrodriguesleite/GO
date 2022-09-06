package somatodoresto

import "testing"

func TestSomaTodoResto(t *testing.T) {
	res := SomaTodoResto([]int{1, 2}, []int{0, 9})
	esp := []int{2, 9}

	if !reflect.DeepEqual(res, esp) {
		t.Errorf("Resultado: '%v', esperado: '%v", res, esp)
	}

}
