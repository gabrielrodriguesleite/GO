package main

// Escrever arquivo usando Go é uma tarefa parecida com a de leitura
// demonstrada no exemplo anterior

// Novamente é muito útil ter uma função para tratar a verificação por erros
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

}
