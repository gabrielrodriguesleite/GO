package main

import "flag"

// Algumas ferramentas de linha de comando como ferramentas de"go" ou "git" possuem
// muitos _subcomandos_ cada um com suas próprias flags.
// Por exemplo, "go build" e "go get" são dois subcomandos diferentes das
// ferramentas de "go". O pacote flag permite definir de maneira fácil
// subcomandos com suas próprias flags.

func main() {

	// Aqui são declarados subcomandos usando "NewFlagSet" e então são
	// definidas flags específicas para este comando.
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

}
