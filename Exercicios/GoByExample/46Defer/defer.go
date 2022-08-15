package main

import (
	"fmt"
	"os"
)

// defer = adiar/postergar
// defer é usado para garantir que uma função é chamada como último
// passo de uma função que está finalizando, geralmente com o intuito
// de limpeza.
// defer é também usado onde ex: ensure e finally seriam usadas em
// outras linguagens.

// Supondo que queremos criar um arquivo, escrever nele e então
// fechá-lo quando terminar. Aqui temos um exemplo desses passos
// utilizando defer.
func main() {

	// Imediatamente após pegar um objeto arquivo com createFile
	// chamamos defer para fechar o arquivo com closeFile
	// Com isso garantimos que closeFile será chamado após writeFile
	// ter finalizado e antes do fechamento de main.
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)

}

func createFile(p string) *os.File {
	fmt.Println("criando")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("escrevendo")
	fmt.Fprintln(f, "dados")
}

// É sempre importante verificar por erros quando fechamos arquivos
// mesmo em funções defer
func closeFile(f *os.File) {
	fmt.Println("fechando")
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

// Rodando a aplicação podemos confirmar que o arquivo é fechado
// após a escrita.
