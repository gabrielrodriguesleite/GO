package main

import (
	"flag"
	"fmt"
	"os"
)

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

	// Para diferentes subcomandos definimos diferentes flags suportadas.
	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	// Os subcomandos são esperados como primeiros argumentos do programa.
	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println(" enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", fooCmd.Args())
	case "bar":

	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

}
