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

/*
		+ Escrever um teste que falha e ver ele falhar para saber que é um teste relevante para
	o requisito, além de produzir uma descrição clara da falha.

		+ Escrever a solução que satisfaz o teste passar.

		+ Em seguida, refatorar o código tendo a certeza que o teste vai garantir que o código
	funciona, podendo se preocupar apenas em deixá-lo bem feito e fácil de dar manutenção.
*/
