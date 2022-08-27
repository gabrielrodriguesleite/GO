package main

// Ler e escrever arquivos são tarefas básicas porém necessárias para muitos
// programas Go. Primeiramente veremos exemplos de leitura.

// Ler arquivos envolve verificar por erros com frequência. Esta função
// irá ajudar nessa tarefa.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

}
