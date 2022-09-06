package soma

func Soma(arr []int) (resultado int) {
	soma := 0
	for _, numero := range arr {
		soma += numero
	}
	return soma
}
