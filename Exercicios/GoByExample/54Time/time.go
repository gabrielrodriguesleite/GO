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

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	p(then.Weekday())

	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

}
