package main

import (
	"reflect"
	"testing"
)

// func TestPercorre(t *testing.T) {

// 	esperado := "Leite"

// 	var resultado []string

// 	x := struct {
// 		Nome string
// 	}{esperado}

// 	percorre(x, func(entrada string) {
// 		resultado = append(resultado, entrada)
// 	})

// 	if len(resultado) != 1 {
// 		t.Errorf("número incorreto de chamadas da função: resultado: %d, esperado %d", len(resultado), 1)
// 	}

// 	if resultado[0] != esperado {
// 		t.Errorf("resultado '%s', esperado '%s'", resultado[0], esperado)
// 	}
// }

// Este teste tem o objetivo de garantir que:

// + É possível armazenar um slice de strings(resultado) que armazena quais strings foram passadas
// dentro de fn pelo percorre. Algumas vezes, nos exemplos anteriores foram criados tipos dedicados
// para isso para espionar chamadas de função/método, mas nesse casso será apenas passádo em uma
// função anônima para fn que acaba em resultado.

// + É usado uma struct anônima com um campo Nome do tipo string para partir para o caminho "feliz"
// e mais simples.

// + Então é chamada percorre com x e o espião e por enquanto só é verificado o tamanho do resultado.
// As verificações serão mais precisas quando se possui algo simples funcionando.

func TestPercorre(t *testing.T) {
	casos := []struct {
		Nome              string
		Entrada           interface{}
		ChamadasEsperadas []string
	}{
		{
			"Struct com um campo string",
			struct{ Nome string }{"Leite"},
			[]string{"Leite"},
		},
		{
			"Struct com dois campos tipo string",
			struct {
				Nome   string
				Cidade string
			}{"Leite", "Nova Hartz"},
			[]string{"Leite", "Nova Hartz"},
		},
		{
			"Struct sem campo tipo string",
			struct {
				Nome  string
				Idade int
			}{"Leite", 32},
			[]string{"Leite"},
		},
		{
			"Campos aninhados",
			Pessoa{"Leite", Perfil{32, "Nova Hartz"}},
			[]string{"Leite", "Nova Hartz"},
		},
		{
			"Ponteiros para coisas",
			&Pessoa{
				"Leite",
				Perfil{32, "Nova Hartz"},
			},
			[]string{"Leite", "Nova Hartz"},
		},
		{
			"Slices",
			[]Perfil{
				{32, "Nova Hartz"},
				{35, "Primavera"},
			},
			[]string{"Nova Hartz", "Primavera"},
		},
		{
			"Arrays",
			[2]Perfil{
				{32, "Nova Hartz"},
				{35, "Primavera"},
			},
			[]string{"Nova Hartz", "Primavera"},
		},
		{
			"Maps",
			map[string]string{
				"Foo": "Bar",
				"Baz": "Boz",
			},
			[]string{"Bar", "Boz"},
		},
	}

	for _, teste := range casos {
		t.Run(teste.Nome, func(t *testing.T) {
			var resultado []string
			percorre(teste.Entrada, func(entrada string) {
				resultado = append(resultado, entrada)
			})

			if !reflect.DeepEqual(resultado, teste.ChamadasEsperadas) {
				t.Errorf("resultado %v, esperado %v", resultado, teste.ChamadasEsperadas)
			}
		})
	}
}

type Pessoa struct {
	Nome   string
	Perfil Perfil
}

type Perfil struct {
	Idade  int
	Cidade string
}
