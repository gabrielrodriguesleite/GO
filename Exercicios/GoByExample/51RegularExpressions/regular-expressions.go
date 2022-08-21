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
	match, _ := regexp.MatchString("p([a-z]+)ego", "pessego")
	fmt.Println(match) // true

	// No exemplo acima usamos um padrão de string diretamente
	// mas para outras tarefas envolvendo regexp vamos precisar
	// compilar uma estrutura otimizada de expressão regular
	r, _ := regexp.Compile("p([a-z]+ego)")

	// Vários metodos estão disponíveis para essa estrutura.
	// Este teste é o mesmo que o visto anteriormente
	fmt.Println(r.MatchString("pessego")) // true

	// Este teste encontra a parte que combina com a regexp
	fmt.Println(r.FindString("pessego passado")) // pessego

	// Este encontra a primeira combinação mas retorna
	// os índices de início e fim do trecho ao invés do
	// próprio texto que combina.
	fmt.Println("idx:", r.FindStringIndex("pessego passado")) // idx: [0 7]

}
