package main

import (
	"fmt"
	"testing"
)

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

// Um teste é criado por escrever uma função com prefixo "Test" no nome
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		// t.Error* vai reportar um teste com falha mas continuará executando os
		// próximos testes. t.Fatal* vai reportar o teste com falha e parar o
		// teste imediatamente.
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

// Escrever testes pode ser repetitivo, então é interessante escrever testes
// no estilo dirigido a tabela, onde os valores do input e os resultados
// experados são litados numa tabela e um loop único percorre cada linha
// e executa a lógica do teste.
func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	// t.Run permite rodar "subtestes" um para cada linha da tabela.
	// Então são visualizados separadamente quando se executa "go test -v"
	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// Testes de Benchmark são definidos geralmente nos arquivos "_test.go" e
// possuem "Benchmark" como prefixo dos seus nomes. Go roda os testes de
// benchmark várias vezes, aumentando b.N a cada vez até que seja capturada
// uma medida precisa.
func BenchmarkIntMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}
