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
}
