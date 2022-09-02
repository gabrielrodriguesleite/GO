package main

import (
	"fmt"
	"os/exec"
)

// As vezes é necessário lançar outros processos de dentro da aplicação Go.
func main() {

	dateCmd := exec.Command("date")
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))
}
