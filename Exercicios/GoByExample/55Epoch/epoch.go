package main

import (
	"fmt"
	"time"
)

// Uma necessidade comum ao programar é conseguir o número em segundos,
// milisegundos ou nanosegundos desde uma Unix epoch.
// Mais aqui: https://en.wikipedia.org/wiki/Unix_time

func main() {

	p := fmt.Println

	now := time.Now()
	p(now)

	// Usando time.now com Unix, UnixMilli ou UnixNano para computar o tempo
	// decorrido desde Unix epoch em segundos, milisegundos ou nanosegundos:
	p(now.Unix())
	p(now.UnixMilli())
	p(now.UnixNano())

	// É possível converter inteiros segundos ou nanosegundos desde epoch
	// no correspondente time formato.
	p(time.Unix(now.Unix(), 0))
	p(time.Unix(0, now.UnixNano()))
}
