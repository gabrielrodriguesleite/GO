package main

import "fmt"

// Para testar é bom serpara o "domínio"(regras de negócio) do resto do mundo(efeitos colaterais).
// A função fmt.Println é um efeito colateral(está imprimindo o valor na saída padrão do terminal).
// A string que estamos usando dentro dela é o domínio.

func Ola() string {
	// Domínio = regra de negócio
	return "Olá, mundo"
}

func main() {
	// Efeito colateral - resto do mundo.
	fmt.Println(Ola())
}
