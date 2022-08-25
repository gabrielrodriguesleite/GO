package main

import (
	"fmt"
	"time"
)

// Go suporta formatação e análise de formatos de tempo
// via layouts baseados em padrão.

func main() {

	p := fmt.Println

	// Este é um exemplo básico de formatação de tempo de acordo com
	// RFC3339, usando a constante do perfil correspondente.
	t := time.Now()
	p(t.Format(time.RFC3339))

	// Análise de tempo utiliza o mesmo valor de perfil como formato.
	t1, _ := time.Parse(time.RFC3339, "2021-11-01T22:08:41+00:00")
	p(t1)

	// Forma e Parse ussando perfil baseado em exemplo
	// Normalmente usamos uma constante do pacote time para definir
	// o perfil, mas também é possível personalizar um perfil.
	p(t.Format("3:04PM"))
	// Perfis devem usar a referência "Mon Jan _2 15:04:05 2006"
	// para mostrar o padrão ao qual o format/parse deve formatar um time/string
	// O exemplo deve ter exatamente estes valores com mostrado.
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, _ := time.Parse(form, "8 41 PM")
	p(t2)

	// Para representações puramente numericas é possível usar formatação padrão
	// de strings com os valores extraidos dos componentes de time.
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	// Parse vai retornar um erro explicando o problema quando lhe é passado uma
	// entrada malformada.
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e := time.Parse(ansic, "8:41PM")
	p(e)
}

/* SAÍDA DE REFERÊNCIA PARA ESSE CÓDIGO
2022-08-24T21:53:48-03:00
2021-11-01 22:08:41 +0000 +0000
9:53PM
Wed Aug 24 21:53:48 2022
2022-08-24T21:53:48.323578-03:00
0000-01-01 20:41:00 +0000 UTC
2022-08-24T21:53:48-00:00
parsing time "8:41PM" as "Mon Jan _2 15:04:05 2006": cannot parse "8:41PM" as "Mon"
*/
