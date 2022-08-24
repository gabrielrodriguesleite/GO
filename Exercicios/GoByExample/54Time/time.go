package main

import (
	"fmt"
	"time"
)

// Go possui um amplo suporte para tempo e períodos de tempo

func main() {

	p := fmt.Println

	// Para obter a data e hora atual
	now := time.Now()
	p(now)

	// É possível construir uma estrutura de time por passar os dados necessários.
	// Horários sempre são associados com um Location ex.: hora local
	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	// É possível extrair os componentes da estrutura time
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// O dia da semana também está disponível
	p(then.Weekday())

	// Este métodos comparam dois time, testando se o primeiro ocorre antes, depois
	// ou ao mesmo segundo, respectivamente.
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// O método Sub retorna uma Duration representando o intervalo entre dois times
	diff := now.Sub(then)
	p(diff)

	// É possível computar o comprimento de uma Duration em várias unidades:
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// É possível utilizar Add para avançar um time por uma duration dada ou com uma
	// duration negativa retroceder um time.
	p(then.Add(diff))
	p(then.Add(-diff))

}

/* EXEMPLO DE SAÍDA DO CÓDIGO
2022-08-23 23:43:38.89026806 -0300 -03 m=+0.000036705
2009-11-17 20:34:58.651387237 +0000 UTC
2009
November
17
20
34
58
651387237
UTC
Tuesday
true
false
false
111894h8m40.238880823s
111894.14451080022
6.713648670648014e+06
4.028189202388808e+08
402818920238880823
2022-08-24 02:43:38.89026806 +0000 UTC
1997-02-11 14:26:18.412506414 +0000 UTC
*/
