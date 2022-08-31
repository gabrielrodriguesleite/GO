package main

// Teste unitários são uma importante parte dos princípios de programação em Go.
// O pacote de testes fornece as ferramentas necessárias para escrever testes
// unitários. O comando "go test" roda os testes.

// Neste exemplo o cógido está no package main, mas poderia estar em quaquer
// pacote. Geralmente o código de teste fica no mesmo pacote que ele testa.

// Esta função simples de seleção do mínimo inteiro será usada nos testes.
// Geralmente o código que precisamos testar fica em um arquivo fonte com o
// nome do estilo "intutils.go", e o arquivo de teste para ele terá o nome
// "intutils_test.go".
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
