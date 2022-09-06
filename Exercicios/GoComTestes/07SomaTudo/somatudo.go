package somatudo

func soma(a []int) int {
	return a[0] + a[1]
}

func SomaTudo(arr ...[]int) (resultados []int) {
	for _, n := range arr {
		resultados = append(resultados, soma(n))
		// append retorna um novo slice que corresponde a 2 slices
		// anexados.
	}
	return
}
