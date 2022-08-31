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

	// Normalmente o benchmark roda uma função que estamos testando em loop
	// b.N vezes.
	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}

/* SAÍDA ESPERADA PARA ESSE CÓDIGO:
$ go test -v
== RUN   TestIntMinBasic
--- PASS: TestIntMinBasic (0.00s)
=== RUN   TestIntMinTableDriven
=== RUN   TestIntMinTableDriven/0,1
=== RUN   TestIntMinTableDriven/1,0
=== RUN   TestIntMinTableDriven/2,-2
=== RUN   TestIntMinTableDriven/0,-1
=== RUN   TestIntMinTableDriven/-1,0
--- PASS: TestIntMinTableDriven (0.00s)
    --- PASS: TestIntMinTableDriven/0,1 (0.00s)
    --- PASS: TestIntMinTableDriven/1,0 (0.00s)
    --- PASS: TestIntMinTableDriven/2,-2 (0.00s)
    --- PASS: TestIntMinTableDriven/0,-1 (0.00s)
    --- PASS: TestIntMinTableDriven/-1,0 (0.00s)
PASS
ok      examples/testing-and-benchmarking    0.023s

$ go test -bench=.
goos: darwin
goarch: arm64
pkg: examples/testing
BenchmarkIntMin-8 1000000000 0.3136 ns/op
PASS
ok      examples/testing-and-benchmarking    0.351s

*/

// go test -v // roda todos os testes no projeto atual no modo verboso.

// go test -bench=. // Rota todos os benchmarks no projeto atual.
// Todos os tests são executados antes dos benchmarks. A flag bench
// Filtra as funções com nome benchmark com um regexp.

/* AVISO
A estrutura atual dos arquivos não permite executar estes testes.
*/
