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

	p(now.Unix())
	p(now.UnixMilli())
	p(now.UnixNano())
}
