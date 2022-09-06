package iteracao

func Repetir(caractere string, quantidade int) string {
	var repeticoes string
	for i := 0; i < quantidade; i++ {
		repeticoes += caractere
	}
	return repeticoes
}
