package concorrencia

type VerificadorWebsite func(string) bool

func VerificadorWebsites(vw VerificadorWebsite, urls []string) map[string]bool {
	resultados := make(map[string]bool)

	for _, url := range urls {
		go func() {
			resultados[url] = vw(url)
		}()
	}

	return resultados
}

// Retorna um map de cada URL verificado com um valor boleano verdadeiro para
// uma resposta boa e falsa para uma resposta ruim.
// Recebe um VerificadorWebsite como parâmetro que leva uma url e retorna um boleano.
// Veja o teste →
