package main

import (
	"fmt"
	"sort"
)

// O pacote sort implementa funções de ordenação para tipos
// prontos e para tipos definidos pelo usuário.
// Primeiro vejamos ordenação para os tipos prontos.

func main() {

	// Métodos sort são especificos para cada tipo.
	// Aqui um exemplo que ordena strings.
	// Note que a ordenação é in-place ou seja modifica o slice
	// original e não devolve um novo.
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// Um exemplo de ordenação de inteiros.
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("ints:   ", ints)

	// Também é possível verificar se um slice está em ordem.
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}

// Rodando a aplicação vemos strings ordenadas
// Inteiros ordenados
// O resultado da verificação para ordenação de inteiros.
