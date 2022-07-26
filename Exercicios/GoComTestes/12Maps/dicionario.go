package main

// Organizar os erros:
// Mais em: https://dave.cheney.net/2016/04/07/constant-errors
type ErrDicionario string

func (e ErrDicionario) Error() string {
	return string(e)
}

const (
	ErrNaoEncontrado      = ErrDicionario("não foi possível encontrar a palavra que você buscou")
	ErrPalavraExistente   = ErrDicionario("não foi possível adicionar a palavra pois ela já existe")
	ErrPalavraInexistente = ErrDicionario("não foi possível atualizar a palavra pois ela não existe")
)

// Maps é um tipo de referência por isso podemos modificá-lo sem passar como ponteiro.
// Maps usa uma estrutura de dados chamada tabela de dispersão ou mapa de hash.
// Mais aqui: https://pt.wikipedia.org/wiki/Tabela_de_dispers%C3%A3o
type Dicionario map[string]string

// IMPORTANTE: nunca inicie um map vazio como:
// var m map[string]string
// ao invés disso usamos make ou dessa forma:
// dicionario = map[string]string{} // chaves indica que está vazio (nil)
// dicionario = make(map[string]string)

func (d Dicionario) Busca(palavra string) (retorno string, erro error) {
	definicao, existe := d[palavra]
	if !existe {
		return "", ErrNaoEncontrado
	}
	return definicao, nil
}

func (d Dicionario) Adiciona(chave, valor string) (erro error) {
	_, err := d.Busca(chave)
	switch err {
	case ErrNaoEncontrado:
		d[chave] = valor
	case nil:
		return ErrPalavraExistente
	default:
		return err
	}
	return
}

func (d Dicionario) Atualiza(chave, valor string) (erro error) {
	_, erro = d.Busca(chave)
	switch erro {
	case ErrNaoEncontrado:
		return ErrPalavraInexistente
	case nil:
		d[chave] = valor
	default:
		return erro
	}
	return
}

func (d Dicionario) Deleta(chave string) {
	delete(d, chave)
}
