package main

import (
	"fmt"
	s "strings"
)

// O pacote padrão strings prove muitas funções úteis para trabalhar
// com strings. Aqui temos alguns exemplos:

// Apelidamos fmt.Println para um nome mais curto como será muito usado,
// o mesmo é feito com stings
var p = fmt.Println

// Uma amostra das funções disponíveis no pacote padrão strings.
// Como são funções do pacote e não métodos de string, é necessário
// passar a string em questão como primero argumento da função.
// Para mais exemplos acesse: https://pkg.go.dev/strings
func main() {

	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))

}
