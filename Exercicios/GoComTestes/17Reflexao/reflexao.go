package main

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
	fn("Com essa chamada da função o teste deve funcionar")
}

// Nesse momento é importante partir para uma verificação do que está sendo chamado
// dentro da função fn.
