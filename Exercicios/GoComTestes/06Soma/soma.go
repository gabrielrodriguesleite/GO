package soma

func Soma(arr [5]int) (resultado int) {
	soma := 0
	for _, numero := range arr {
		soma += numero
	}
	return soma
}
