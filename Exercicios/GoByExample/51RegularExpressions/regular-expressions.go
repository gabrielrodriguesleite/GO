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
	fmt.Println(match)

	// No exemplo acima usamos um padrão de string diretamente
	// mas para outras tarefas envolvendo regexp vamos precisar
	// compilar uma estrutura otimizada de expressão regular
	r, _ := regexp.Compile("p([a-z]+ego)")

	fmt.Println(r.MatchString("pessego"))

}
