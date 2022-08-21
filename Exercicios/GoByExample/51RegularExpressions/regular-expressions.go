package main

import (
	"fmt"
	"regexp"
)

// Go possui suporte nativo para expressões regulares.
// Neste código alguns exemplos communs relacionados
// a tarefas com expressões regulares.
func main() {

	match, _ := regexp.MatchString("p([a-z]+)ego", "pessego")
	fmt.Println(match)

}
