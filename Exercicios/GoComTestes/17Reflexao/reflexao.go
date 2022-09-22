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

	for i := 0; i < valor.NumField(); i++ { // NumField
		campo := valor.Field(i) // Campo

		if campo.Kind() == reflect.String { // Kind
			fn(campo.String())
		}
	}
}

// Neste momento o código está frágil e pouco segura mas o importante nesse momento do TDD
// é apenas fazê-lo passar. Daí então refatorar para remover as lacunas.

// O pacote reflect possui uma função ValueOf que retorna um Value de uma variável. Isso
// permite inspecionar um valor inclusive seus campos usados nas próximas linhas.

// Pode-se presumir sobre o valor passado:

// + É possível procurar pelo primeiro e único campo, porém pode não haver nenhum campo
// o que causaria um pânico.

// + Em seguida é possível chamar String() que retorna o valor subjacente como string,
// mas retornará um erro se o campo for de outro tipo.

// ==========

// "valor" também possui um método chamado NumField que retorna a quantidade de campos no
// valor. Isso permite iterar sobre os campos e chamar fn, fazendo passar o teste.

// A próxima falha desse código é presumir o tipo string. O teste irá verificar isso.
