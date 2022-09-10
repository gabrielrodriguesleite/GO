package main

import "errors"

var ErrNaoEncontrado = errors.New("não foi possível encontrar a palavra que você buscou")

// Maps é um tipo de referência por isso podemos modificá-lo sem passar como ponteiro.
// Maps usa uma estrutura de dados chamada tabela de dispersão ou mapa de hash.
// Mais aqui: https://pt.wikipedia.org/wiki/Tabela_de_dispers%C3%A3o
type Dicionario map[string]string

func (d Dicionario) Busca(palavra string) (retorno string, erro error) {
	definicao, existe := d[palavra]
	if !existe {
		return "", ErrNaoEncontrado
	}
	return definicao, nil
}

func (d Dicionario) Adiciona(chave, valor string) {
	d[chave] = valor
}
