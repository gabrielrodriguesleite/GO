package main

import "fmt"

func main() {

	// slices são arrays com super poderes
	// Para criar um slice vazio com tamanho diferente de zero se usa a palavra make
	// aqui criamos um slice com 3 posições do tipo string
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// É possível atribuir e ler o valor e saber o tamanho assim como arrays
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len", len(s))

	// append retorna um **novo** slice contendo um ou mais valores passados por parametro.
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

}
