package main

import "reflect"

// Desafio: escrever uma função percorre(x interface{}, fn func(entrada string))
// que recebe uma struct x e chama fn para todos os campos string encontrados
// dentro dela.

// Nesse desafio será utilizado o conceito de reflection (reflexão).

// A reflexão em computação é a habilidade de um programa examinar sua]
// própria estrutura, particularmente através de tipos; é uma forma de
// metaprogramação. Também é uma ótima fonte de confusão. -- The Go Blog: Reflection
// fonte: https://blog.golang.org/laws-of-reflection

// Nota: Este conteúdo deve ser pré introdução de Generics em Go. Vide parágrafo inicial
// em The Go Blog: Reflection -- The Laws of Reflection

func percorre(x interface{}, fn func(entrada string)) {
	// fn("Com essa chamada da função o teste deve funcionar")
	valor := reflect.ValueOf(x) // ValorDe
	campo := valor.Field(0)     // Campo
	fn(campo.String())
}

// Neste momento o código está frágil e pouco segura mas o importante nesse momento do TDD
// é apenas fazê-lo passar. Daí então refatorar para remover as lacunas.
