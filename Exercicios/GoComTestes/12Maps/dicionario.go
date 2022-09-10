package main

type Dicionario map[string]string

func Busca(dicionario Dicionario, palavra string) (retorno string) {
	return dicionario[palavra]
}
