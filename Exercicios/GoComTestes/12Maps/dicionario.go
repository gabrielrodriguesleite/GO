package main

import "errors"

var ErrNaoEncontrado = errors.New("não foi possível encontrar a palavra que você buscou")

type Dicionario map[string]string

func (d Dicionario) Busca(palavra string) (retorno string, erro error) {
	definicao, existe := d[palavra]
	if !existe {
		return "", ErrNaoEncontrado
	}
	return definicao, nil
}
