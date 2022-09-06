package somatudo

func soma(a []int) int {
	return a[0] + a[1]
}

func SomaTudo(arr ...[]int) (resultados []int) {
	qtdNumeros := len(arr)
	resultados = make([]int, qtdNumeros)

	for i, n := range arr {
		resultados[i] = soma(n)
	}
	return
}
