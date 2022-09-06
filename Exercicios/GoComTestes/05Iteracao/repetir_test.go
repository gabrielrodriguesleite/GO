package iteracao

import "testing"

func TestRepetir(t *testing.T) {
	repeticoes := Repetir("a", 9)
	esperado := "aaaaa"

	if repeticoes != esperado {
		t.Errorf("Esperado: '%s' mas obteve: '%s'", esperado, repeticoes)
	}
}

// SAÍDA ESPERADA PARA O TESTE
/* go test -v 05Iteracao/*
=== RUN   TestRepetir
--- PASS: TestRepetir (0.00s)
PASS
ok      command-line-arguments  0.004s
*/

func BenchmarkRepetir(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Repetir("a")
	}
}

// SAÍDA ESPERADA PARA O BENCHMARK
/*  go test -bench=. 05Iteracao/*
goos: linux
goarch: amd64
cpu: Intel(R) Celeron(R) CPU B800 @ 1.50GHz
BenchmarkRepetir-2       3026282               396.3 ns/op
PASS
ok      command-line-arguments  1.613s
*/

// BenchmarkRepetir executada 3026282 levou uma média de 396.3 ns/op
