package main

import (
	"fmt"
	"os"
)

// Durante toda execução do programa, muitas vezes queremos
// criar dados que não são necessários após a finalização do
// programa. Arquivos e diretórios temporários são úteis para
// este propósito desde que não poluam o sistema de arquivos
// com o passar do tempo.

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.CreateTemp("", "sample")
	check(err)
	fmt.Println("Temp file name:", f.Name())
}
