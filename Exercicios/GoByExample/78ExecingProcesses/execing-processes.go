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

// SAÍDA PADRÃO PARA ESTE CÓDIGO

/* go run 78ExecingProcesses/execing-processes.go
total 344K
drwxrwxr-x 80 k_co k_co  16K set  1 23:18 .
drwxrwxr-x  3 k_co k_co 4,0K ago 28 22:50 ..
drwxrwxr-x  2 k_co k_co 4,0K ago 28 22:50 01HelloWorld
drwxrwxr-x  2 k_co k_co 4,0K ago 28 22:50 02Values
	...
*/
