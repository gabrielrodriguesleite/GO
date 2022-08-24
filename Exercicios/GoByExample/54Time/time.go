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
