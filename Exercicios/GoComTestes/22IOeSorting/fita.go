package main

import "io"

type fita struct {
	arquivo io.ReadWriteSeeker
}

func (f *fita) Write(p []byte) (n int, err error) {
	f.arquivo.Seek(0, 0)
	return f.arquivo.Write(p)
}
