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

	// Podemos criar um novo modelo de texto e traduzir seu corpo a
	// partir de uma string. Modelos são uma mistura de texto estático
	// e "ações" enclausuradas em {{...}} estas que são usadas para
	// inserir conteúdo dinâmico.
	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}

	// Como alternativa, podemos usar temlate.Must para lançar um panic
	// no caso do parse retornar um erro. O que é especialmente útil
	// para modelos inicializados no escopo global.
	t1 = template.Must(t1.Parse("Value: {{.}}\n"))

	// Executar o modelo geramos texto com valores especificados pelas
	// "actions". As actions {{.}} são substituídas pelos valores
	// passados por parâmetro para o Execute
	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"Go", "Rust", "C++", "C#",
	})

	// ---------

	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	t2 := Create("t2", "Name: {{.Name}}\n")

	t2.Execute(os.Stdout, struct {
		Name string
	}{"Jane Doe"})

	t2.Execute(os.Stdout, map[string]string{
		"Name": "Mickey Mouse",
	})

	t3 := Create("t3", "{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	t4 := Create("t4", "Range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout, []string{"Go", "Rust", "C++", "C#"})
}
