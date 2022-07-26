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

	// Neste exemplo será executado o comando "ls".
	// Go precisa do caminho exato do binário então vamos usar LookPath para
	// encontrar este caminho. (provavelmente "/bin/ls")
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// Exec também precisa dos argumentos em forma de slice (ao invés de uma
	// string grande). Entregamos alguns argumentos. Note que o primeiro
	// deve ser o nome do comando.
	args := []string{"ls", "-a", "-l", "-h"}

	// Também precisamos fornecer as variáveis de ambiente.
	// Neste caso usamos as atuais.
	env := os.Environ()

	// Aqui finalmente chamamos "syscall.Exec". Caso a chamada dê sucesso
	// este processo termina aqui e é substituido pelo "/bin/ls -a -l -h".
	// Se houver um erro teremos um valor no retorno.
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}

// SAÍDA ESPERADA PARA ESTE CÓDIGO

// Quando rodamos o programa ele é substituido pelo comando "ls".
/* go run 78ExecingProcesses/execing-processes.go
total 344K
drwxrwxr-x 80 k_co k_co  16K set  1 23:18 .
drwxrwxr-x  3 k_co k_co 4,0K ago 28 22:50 ..
drwxrwxr-x  2 k_co k_co 4,0K ago 28 22:50 01HelloWorld
drwxrwxr-x  2 k_co k_co 4,0K ago 28 22:50 02Values
	...
*/

// Note que Go não oferece a um "fork" clássico do Unix.
// Normalmente isso não é problema já que iniciar uma goroutine, lançar um
// processo e executar um processo cobrem a maioria dos casos de uso de um "fork".
