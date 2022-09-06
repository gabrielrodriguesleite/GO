package iteracao

import "testing"

func TestRepetir(t *testing.T) {
	repeticoes := Repetir("a")
	esperado := "aaaaa"

	if repeticoes != esperado {
		t.Errorf("Esperado: '%s' mas obteve: '%s'", esperado, repeticoes)
	}
}

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
