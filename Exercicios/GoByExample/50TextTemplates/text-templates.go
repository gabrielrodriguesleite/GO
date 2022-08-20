package main

import (
	"os"
	"text/template"
)

// Go possui na biblioteca padrão suporte para criação de conteúdo
// dinâmico ou para apresesentar uma saída personalizada para o
// usuário com o pacote text/template. Um pacote irmão de nome
// html/template prove a mesma API mas possui funcionalidades
// adicionais de segurança e é usada para gerar HTML

func main() {

	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}

	t1 = template.Must(t1.Parse("Value: {{.}}\n"))

	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"Go", "Rust", "C++", "C#",
	})
}
