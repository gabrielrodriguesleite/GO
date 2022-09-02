package main

import (
	"fmt"
	"io"
	"os/exec"
)

// As vezes é necessário lançar outros processos de dentro da aplicação Go.
func main() {

	// Exemplo simples de um comando que não leva argumentos e que se captura a
	// saída para imprimir na tela.
	dateCmd := exec.Command("date")

	// O método "Output" roda o comando, aguarda o término e coleta a saída
	// padrão. Se nenhum erro ocorrer "dateOut" vai conter bytes com a
	// informação da data.
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(dateOut))

	// ------------

	// "Output" e outros métodos de "Command" vão retornar "*exec.Error" se
	// houver um erro na execução do comando, por exemplo um caminho errado,
	// ou "*exec.ExitError" se o comando rodar mas sair com um código de retorno
	// diferente de "0".
	_, err = exec.Command("date", "-x").Output()
	if err != nil {
		switch e := err.(type) {
		case *exec.Error:
			fmt.Println("failed executing:", err)
		case *exec.ExitError:
			fmt.Println("command exit rc =", e.ExitCode())
		default:
			panic(err)
		}
	}

	// ------------

	// Neste exemplo é usado pipe para para enviar dados para o comando no seu
	// "stdin" e então é usado pipe para coletar os dados de retorno do seu "stdout".
	grepCmd := exec.Command("grep", "hello")

	// Aqui foi escrito de forma explicita os passos que são seguidos:
	// Pegamos os pipes de entrada e saída,
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	// o processo é iniciado,
	grepCmd.Start()
	// dados são escritos na entrada padrão,
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	// dados resultantes são coletados da saída pradrão,
	// e por fim aguardamos o processo concluir.
	grepBytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait()

	// Neste exemplo o tratamento de erros foi omitido porém deve ser feido da forma
	// normal com "if err != nil". Aqui coletamos apenas o "StdoutPipe" mas é possível
	// coletar também "StderrPipe" da mesma maneira.
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// ------------

	//
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
