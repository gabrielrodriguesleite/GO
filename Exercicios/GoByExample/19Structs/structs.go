package main

import "fmt"

// structs são coleções de campos tipados
// são úteis para agrupar dados juntos e formar records
// esta estrutura possui nome e idade
type person struct {
	name string
	age  int
}

// essa função constroi uma nova estrutura com o nome passado como argumento
// passar um pointer faz com que a variável local seja mantina através do escopo
// (o garbage collector não vai limpar já que ainda existe uma referência)
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	// essa sintaxe cria uma nova estrutura
	fmt.Println(person{"Bob", 20}) // {Fred 0}

	// os campos podem ser nomeados na inicialização
	fmt.Println(person{name: "Alice", age: 30}) // {Alice 30}

	// valores omitidos iniciam zerados
	fmt.Println(person{name: "Fred"}) // {Fred 0}

	// um prefixo & produz um ponteiro para a estrutura
	fmt.Println(&person{name: "Ann", age: 40}) // &{Ann 40}

	// uma função construtora que encapsula a criação de uma estrutura
	fmt.Println(newPerson("Jon")) // &{Jon 42}

	// acesso aos campos usando o .
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name) // Sean

	// o acesso por . aos campos funciona também nos apontadores de estruturas,
	// o apontador é desrefenciado automaticamente
	sp := &s
	fmt.Println(sp.age) // 50

	// estruturas são mutáveis
	sp.age = 51
	fmt.Println(sp.age) // 51

}
