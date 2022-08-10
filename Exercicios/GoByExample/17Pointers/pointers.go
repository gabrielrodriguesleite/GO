package main

import "fmt"

// pointers permitem passar e trabalhar com referências a valores

// essa função recebe um inteiro, então os argumentos serão passados como valor
// zeroval receberá uma cópia do valor com o qual ela for chamada
func zeroval(ival int) {
	ival = 0
}

// zeroptr possui um parametro do tipo ponteiro para inteiro
// dessa forma zoroptr receberá um endereço de memória da variável com a qual for chamada
// permitindo assim alterar o valor que essa memória/variável armazena
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i) // 1 - Valor inicial

	zeroval(i)
	fmt.Println("zeroval:", i) // 1 - passar por valor não altera a original

	// a sintaxe &i passa o endereço da variável i, ou seja um ponteiro para i
	zeroptr(&i)
	fmt.Println("zeroptr:", i) // 0 - passar por referência altera a original

	// pointers podem ser impressos:
	fmt.Println("pointer:", &i) // 0x0000160d8 - o endereço em memória (referência ou ponteiro)

	// zeroval não altera o valor de i em main
	// por outro lado zeroptr altera pois recebe o endereço em memória da variável
}
