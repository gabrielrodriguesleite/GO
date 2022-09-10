package main

type Dicionario map[string]string

func (d Dicionario) Busca(palavra string) (retorno string, erro error) {
	return d[palavra], nil
}
