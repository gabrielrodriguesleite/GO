package main

import (
	"os"
)

type fita struct {
	arquivo *os.File
}

func (f *fita) Write(p []byte) (n int, err error) {
	f.arquivo.Truncate(0) // serve para esvaziar o aquivo
	f.arquivo.Seek(0, 0)  // move o "cursor" para a posição de escrita definida
	return f.arquivo.Write(p)
}
