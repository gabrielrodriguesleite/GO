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
}
