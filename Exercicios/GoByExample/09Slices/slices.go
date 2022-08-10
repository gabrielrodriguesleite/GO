package main

import "fmt"

func main() {

	// slices são arrays com super poderes
	// Para criar um slice vazio com tamanho diferente de zero se usa a palavra make
	// aqui criamos um slice com 3 posições do tipo string
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len", len(s))

}
