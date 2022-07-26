package main

import "fmt"

func main() {

	// um slice
	nums := []int{2, 3, 4}
	sum := 0

	// range usado para somar numeros de um slice. (também funcionaria com array)
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum", sum)

	// range retorna o indice e o valor do slice ou array
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// range iterando em map retornando chave e valor
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// range retornando as chaves de um map
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// range interando string (rune) e retornando o indice e o valor para unicode
	// veja: https://gobyexample.com/strings-and-runes
	// rune é um inteiro que representa um código Unicode: veja: https://go.dev/blog/strings
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
