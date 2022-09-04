package main

const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "

// Constantes melhoram a performance da aplicação por evitar de
// criar uma string "Olá, " toda vez que "OlaTu" é chamado.

func Ola(nome, idioma string) string {
	prefixo := prefixoOlaPortugues
	if nome == "" {
		nome = "mundo"
	}
	switch idioma {
	case "espanhol":
		prefixo = prefixoOlaEspanhol
	}
	return prefixo + nome
}

// SOBRE O FLUXO TDD:

/*
- Escrever o teste
- Compilar o código sem erros
- Rodar o teste, ver o teste falhar e certificar que a mensagem de erro faz sentido
- Escrever a quantidade mínima de código para o teste passar
- Refatorar
*/
