package somatodoresto

func Soma(a []int) (s int) {
	for _, n := range a {
		s += n
	}
	return
}

func SomaTodoResto(arr ...[]int) (res []int) {
	for _, num := range arr {
		resto := num[1:]
		// slice pode ser fatiado
		res = append(res, Soma(resto))
	}
	return
}
