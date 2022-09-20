package concorrencia

type VerificadorWebsite func(string) bool
type resultado struct {
	string
	bool
}

func VerificadorWebsites(vw VerificadorWebsite, urls []string) map[string]bool {
	resultados := make(map[string]bool)
	canalResultado := make(chan resultado)

	for _, url := range urls {
		go func(u string) { // usando goroutine para executar em paralelo
			canalResultado <- resultado{u, vw(u)}
		}(url) // cada goroutine recebe o seu valor de url para trabalhar
	}

	for i := 0; i < len(urls); i++ { // i respostas serão aguardadas
		resultado := <-canalResultado // aguarda até receber a inésima resposta da goroutine
		resultados[resultado.string] = resultado.bool
	}
	return resultados
}

// é possível executar com sucesso ou obter uma condição de corrida
// detectado pelo detector de corrida nativo: quando mais de uma goroutine tenta
// escrever um recurso ao mesmo tempo:
// WARNING: DATA RACE

// Retorna um map de cada URL verificado com um valor boleano verdadeiro para
// uma resposta boa e falsa para uma resposta ruim.
// Recebe um VerificadorWebsite como parâmetro que leva uma url e retorna um boleano.
// Veja o teste →
