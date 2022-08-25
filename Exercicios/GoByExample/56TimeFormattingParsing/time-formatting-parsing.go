package main

import (
	"fmt"
	"time"
)

func main() {

	p := fmt.Println

	t := time.Now()
	p(t.Format(time.RFC3339))
}
