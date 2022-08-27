package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

// O pacote filepath possui funções para analisar e contruir caminhos de arquivos
// de forma portável entre sistemas operacionais. Por exemplo dir/file no linux
// e dir\file no windows.
func main() {

	// Join deve ser usado quando se constroem paths de forma portável.
	// Aceita n número de argumentos e constroe a hierarquia do caminho à partir destes.
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p) // p: dir1/dir2/filename

	// É importante sempre dar preferência ao Join ao invés de concatenar caminhos
	// manualmente com /s ou \s. Além de prover portabilidade, Join também normaliza
	// os caminhos por remover separadores e mudanças de pastas desnecessárias.
	fmt.Println(filepath.Join("dir1//", "filename"))       // dir1/filename
	fmt.Println(filepath.Join("dir1/../dir1", "filename")) // dir1/filename

	// Dir e Base podem ser usados para separar um caminho em pasta e arquivo.
	// De maneira alternativa, Split irá retornar ambos na mesma chamada:
	fmt.Println("Dir(p):", filepath.Dir(p))   // Dir(p): dir1/dir2
	fmt.Println("Base(p):", filepath.Base(p)) // Base(p): filename

	f, d := filepath.Split(p)
	fmt.Printf("Split(p) → Dir: %s, Base: %s\n", f, d) // Split(p) → Dir: dir1/dir2/, Base: filename

	// É possível também testar se um caminho é relativo ou absoluto.
	fmt.Println(filepath.IsAbs("dir/file"))  // false
	fmt.Println(filepath.IsAbs("/dir/file")) // true

	// Alguns arquivos possuem nomes com a extenção indicada após o ponto.
	// É possível dividir a extenção e capturar ela com Ext:
	filename := "config.json"
	ext := filepath.Ext(filename)
	fmt.Println(ext) // .json

	// strings.TrimSuffix é útil para remover a extenção do nome do arquivo.
	fmt.Println(strings.TrimSuffix(filename, ext))

	// Rel procura um caminho relativo entre uma base e um target.
	// Retorna um erro se o target não possuir um caminho relativo a base.
	rel, err := filepath.Rel("a/b/", "a/b/t/file")
	if err != nil {
		panic(err)
	}

	fmt.Println(rel) // t/file

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}

	fmt.Println(rel) // ../c/t/file

}

/* SAÍDA ESPERADA PARA O CÓDIGO
p: dir1/dir2/filename
dir1/filename
dir1/filename
Dir(p): dir1/dir2
Base(p): filename
Split(p) → Dir: dir1/dir2/, Base: filename
false
true
.json
config
t/file
../c/t/file
*/
