package main

import (
	"fmt"
	"sort"
)

// As vezes queremos ordenar uma coleção de alguma forma
// diferente da ordem natural.
// Por exemplo, supondo que queremos ordenar strings por
// seu comprimento ao invés da ordem alfabética.
// Aqui temos um exemplo desta função de ordenação
// personalizada em Go

// Para criar a função de ordenação personalizada em Go
// é necessário definirmos um tipo correspondente.
// Aqui creamos o tipo byLength que possui o nome da
// operação que desejamos e definimos o tipo dos dados.
type byLength []string

// Para usarmos a função genérica Sort do pacote sort precisamos
// implementar os os métodos da interface sort.Interface - Len,
// Less e Swap. Len e Swap vão ser similares para muitos tipos
// e Less vai conter a função personalizada de ordenação.
// Neste exemplo queremos ordenar de maneira crescente pelo
// comprimento da string, então usamos len(s[i]) e len(s[j])
func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
