package main

func NovoArmazenamentoJogadorEmMemoria() *ArmazenamentoJogadorEmMemoria {
	return &ArmazenamentoJogadorEmMemoria{map[string]int{}}
}

type ArmazenamentoJogadorEmMemoria struct {
	armazenamento map[string]int
}

func (a *ArmazenamentoJogadorEmMemoria) ObterPontuacaoDoJogador(nome string) int {
	return a.armazenamento[nome]
}

func (a *ArmazenamentoJogadorEmMemoria) GravarVitoria(nome string) {
	a.armazenamento[nome]++
}

func (a *ArmazenamentoJogadorEmMemoria) ObterLiga() []Jogador {
	var liga []Jogador
	for nome, vitorias := range a.armazenamento {
		liga = append(liga, Jogador{nome, vitorias})
	}
	return liga
}
