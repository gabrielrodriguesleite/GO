package concorrencia

import (
	"reflect"
	"testing"
	"time"
)

// Usando a injeção de dependência é possível testar uma função sem fazer chamadas
// HTTP de verdade, tornando o teste seguro e rápido.

func mockVerificadorWebsite(url string) bool {
	return url != "waat://furhurterwe.geds"
}

func TestVerificaWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	esperado := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	resultado := VerificadorWebsites(mockVerificadorWebsite, websites)

	if !reflect.DeepEqual(esperado, resultado) {
		t.Fatalf("esperado %v, resultado %v", esperado, resultado)
	}
}

// Como melhorar a velocidade do teste para poder verificar centenas de websites em menos tempo?

func slowStubVerificadorWebsite(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

// go test -bench="."
// sem concorrência 2017804909 ns/op B800 @1.5Ghz
// com paralelismo  	20608179 ns/op B800 @1.5Ghz
// umas 100x mais rápido.
func BenchmarkVerificaWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "uma url"
	}

	for i := 0; i < b.N; i++ {
		VerificadorWebsites(slowStubVerificadorWebsite, urls)
	}

}
