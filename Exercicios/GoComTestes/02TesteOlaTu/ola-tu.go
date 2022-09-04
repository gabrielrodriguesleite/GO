package main

const prefixoOlaPortugues = "Olá, "

// Constantes melhoram a performance da aplicação por evitar de
// criar uma string "Olá, " toda vez que "OlaTu" é chamado.

func OlaTu(nome string) string {
	return prefixoOlaPortugues + nome
}
