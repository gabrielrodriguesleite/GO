package main

import (
	"fmt"
	"regexp"
)

// Go possui suporte nativo para expressões regulares.
// Neste código alguns exemplos communs relacionados
// a tarefas com expressões regulares.
func main() {

	// Este código testa se um padrão combina com uma string
	match, _ := regexp.MatchString("t([a-z]+)te", "tomate")
	fmt.Println(match) // true

	// No exemplo acima usamos um padrão de string diretamente
	// mas para outras tarefas envolvendo regexp vamos precisar
	// compilar uma estrutura otimizada de expressão regular
	r, _ := regexp.Compile("t([a-z]+te)")

	// Vários metodos estão disponíveis para essa estrutura.
	// Este teste é o mesmo que o visto anteriormente
	fmt.Println(r.MatchString("tomate")) // true

	// Este teste encontra a parte que combina com a regexp
	fmt.Println(r.FindString("tomate topete")) // tomate

	// Este encontra a primeira combinação mas retorna
	// os índices de início e fim do trecho ao invés do
	// próprio texto que combina.
	fmt.Println("idx:", r.FindStringIndex("tomate topete")) // idx: [0 6]

	// A variante Submatch inclui informação sobre ambos, o
	// padrão inteiro que combina e e as subcorrespondências
	// dentro dessas correspondências. Neste exemplo vai retornar
	// informações de ambos p([a-z]+ego) e ([a-z]+)
	fmt.Println(r.FindStringSubmatch("tomate topete")) // [tomate omate]

	// Assim como no exemplo anterior este método vai retornar informações
	// sobre os indices das correspondências e das subcorrespondências.
	fmt.Println(r.FindStringSubmatchIndex("tomate topete")) // [0 6 1 6]

	fmt.Println(r.FindAllString("tomate topete torniquete", -1)) // [tomate topete torniquete]

	// https://gobyexample.com/regular-expressions
}
