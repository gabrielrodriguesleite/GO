package main

import (
	// Ao importar o pacote "embed" use algum identificador de exportação
	// do pacote ou faça um importe em branco com "_"
	"embed"
)

// "go:embed" é uma diretiva de compilador que permite um programa incluir
// arquivos e diretórios no binário na hora da contrução.
// Mais aqui: https://pkg.go.dev/embed
// E aqui também: https://pkg.go.dev/cmd/compile#hdr-Compiler_Directives

// A diretiva embed aceita caminhos relativos para o diretório contendo o
// arquivo fonte. Esta diretiva irá imbutir o arquivo dentro da variável
// do tipo string que segue a diretiva.

// go:embed folder/single_file.txt
var fileString string

// Também é possível embutir o arquivo em um array de bytes.

// go:embed folder/single_file.txt
var fileByte []byte

// É possível embutir vários arquivos usando coringas.
// Este exemplo usa uma variável do tipo embed.FS.type que
// implementa um sistema de arquivos virtual.
// Mais em: https://pkg.go.dev/embed#FS

// go:embed folder/single_file.txt
// go:embed folder/*.hash
var folder embed.FS

func main() {

	// Imprimindo o conteúdo de single_file.txt
	println(fileString)
	println(string(fileByte))

	content1, _ := folder.ReadFile("folder/file1.hash")
	println(string(content1))

	content2, _ := folder.ReadFile("folder/file2.hash")
	println(string(content2))

	println("Fim")

}
