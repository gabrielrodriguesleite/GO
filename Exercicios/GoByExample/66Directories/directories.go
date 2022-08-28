package main

import "os"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Go possui varias funções úteis para trabalhar com pastas no sistema de arquivos
func main() {

	// Criando um subdiretório no diretório atual.
	err := os.Mkdir("subdir", 0755)
	check(err)

	// Quando se está trabalhando com diretórios temporários é uma boa prática
	// agendar sua remoção. os.RemoveAll vai deletar uma árvore de diretório
	// inteiro similar ao "rm -rf".
	defer os.RemoveAll("subdir")

}
