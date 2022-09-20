package concorrencia

type VerificadorWebsite func(string) bool

func VerificadorWebsites(vw VerificadorWebsite, urls []string) map[string]bool {
	resultados := make(map[string]bool)

	for _, url := range urls {
		resultados[url] = vw(url)
	}

	return resultados
}
