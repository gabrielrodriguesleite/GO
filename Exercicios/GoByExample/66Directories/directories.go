package main

import "os"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Go possui varias funções úteis para trabalhar com pastas no sistema de arquivos
func main() {

	err := os.Mkdir("subdir", 0755)
	check(err)

	defer os.RemoveAll("subdir")

}
