package main

type Dicionario map[string]string

func (d Dicionario) Busca(palavra string) (retorno string) {
	return d[palavra]
}
