package main

import (
	"os"
	"os/exec"
	"syscall"
)

// No exemplo anterior vimos como lançar processos externos.
// Fazemos isso quando precisamos um processo externo acessível para o
// processo go que está rodando. Mas as vezes apenas queremos substituir
// o processo Go com outro processo (não necessariamente um Go).
// Para isso Go possui uma implementação para a clássica função "exec".
func main() {

	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{"ls", "-a", "-l", "-h"}
	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
