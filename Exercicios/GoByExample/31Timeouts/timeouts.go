package main

import "time"

func main() {

	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
	}()
}
