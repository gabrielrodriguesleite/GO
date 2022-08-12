package main

import (
	"fmt"
	"time"
)

// Timeouts são importantes que se conectam com recursos externos
// ou que estejam vinculados de outra forma ao tempo de execução
// Implementar timeouts em  go é facil e elegante graças a channels e select
func main() {

	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	}
}
