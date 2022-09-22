package main

import "testing"

func TestPercorre(t *testing.T) {

	esperado := "Leite"

	var resultado []string

	x := struct {
		Nome string
	}{esperado}

	percorre(x, func(entrada string) {
		resultado = append(resultado, entrada)
	})

	if len(resultado) != 1 {
		t.Errorf("número incorreto de chamadas da função: resultado: %d, esperado %d", len(resultado), 1)
	}
}

// Este teste tem o objetivo de garantir que:

// + É possível armazenar um slice de strings(resultado) que armazena quais strings foram passadas
// dentro de fn pelo percorre. Algumas vezes, nos exemplos anteriores foram criados tipos dedicados
// para isso para espionar chamadas de função/método, mas nesse casso será apenas passádo em uma
// função anônima para fn que acaba em resultado.

// + É usado uma struct anônima com um campo Nome do tipo string para partir para o caminho "feliz"
// e mais simples.

// + Então é chamada percorre com x e o espião e por enquanto só é verificado o tamanho do resultado.
// As verificações serão mais precisas quando se possui algo simples funcionando.
