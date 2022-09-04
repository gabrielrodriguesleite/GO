package main

const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "
const prefixoOlaFrances = "Bonjour, "

// Constantes melhoram a performance da aplicação por evitar de
// criar uma string "Olá, " toda vez que "OlaTu" é chamado.

// Definir o nome do retorno torna a intenção do código mais
// clara, já que a assinatura é exibida no "go doc"
// Além disso esta função tem o nome começado em minúsculo para
// a tornar privada.
func prefixoSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case "espanhol":
		prefixo = prefixoOlaEspanhol
	case "francês":
		prefixo = prefixoOlaFrances
	default:
		prefixo = prefixoOlaPortugues
	}
	return
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
