package main

import (
	"encoding/json"
	"fmt"
	"io"
)

type Liga []Jogador

func NovaLiga(rdr io.Reader) ([]Jogador, error) {
	var liga []Jogador
	err := json.NewDecoder(rdr).Decode(&liga)
	if err != nil {
		err = fmt.Errorf("problema parseando a liga, %v", err)
	}
	return liga, err
}
