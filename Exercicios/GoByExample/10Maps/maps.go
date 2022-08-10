package main

import "fmt"

func main() {

	// maps também chamadas de hashes ou dicts em outras linguagens

	// para criar um map vazio com make:
	// make(map[key-type]val-type)
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	// atribuir uma valor de uma chave do map:
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// o metodo len retorna o numero de pares key/value quando chamado com um map
	fmt.Println("len:", len(m))

	// o metodo delete remove um par key/value quando invocado com o map e a chave
	delete(m, "k2")
	fmt.Println("map:", m)

	// o segundo parametro retornado indica se o valor está presente e serve pra distinguir um 0 ou '' de um inexistente
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// map também pode ser declarado e iniciado na mesma linha
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	// maps são impressos no formato map[k:v k:v]
}
