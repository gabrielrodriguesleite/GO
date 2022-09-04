package main

const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "
const prefixoOlaFrances = "Bonjour, "

// Constantes melhoram a performance da aplicação por evitar de
// criar uma string "Olá, " toda vez que "OlaTu" é chamado.

func prefixoSaudacao(idioma string) string {
	switch idioma {
	case "espanhol":
		return prefixoOlaEspanhol
	case "francês":
		return prefixoOlaFrances
	default:
		return prefixoOlaPortugues
	}
}

func Ola(nome, idioma string) string {
	if nome == "" {
		nome = "mundo"
	}

	return prefixoSaudacao(idioma) + nome
}

// SOBRE O FLUXO TDD:

/*
- Escrever o teste
- Compilar o código sem erros
- Rodar o teste, ver o teste falhar e certificar que a mensagem de erro faz sentido
- Escrever a quantidade mínima de código para o teste passar
- Refatorar
*/
