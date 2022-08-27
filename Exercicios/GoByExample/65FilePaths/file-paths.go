package main

import (
	"fmt"
	"path/filepath"
)

// O pacote filepath possui funções para analisar e contruir caminhos de arquivos
// de forma portável entre sistemas operacionais. Por exemplo dir/file no linux
// e dir\file no windows.
func main() {
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

}
