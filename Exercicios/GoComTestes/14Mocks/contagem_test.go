package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

//spy sleeper
//mock
type SleeperSpy struct {
	Chamadas int
}

func (s *SleeperSpy) Sleep() {
	s.Chamadas++
}

type TempoSpy struct {
	ducacaoPausa time.Duration
}

func (t *TempoSpy) Pausa(duracao time.Duration) {
	t.ducacaoPausa = duracao
}

// Spies são um tipo de mock no qual podemos gravar informações como por
// exemplo: quantas vezes a função foi chamada. Também podem gravar os
// argumentos definidos, etc.

const pausa = "pausa"
const escrita = "escrita"

type SpyContagemOperacoes struct {
	Chamadas []string
}

func (s *SpyContagemOperacoes) Pausa() {
	s.Chamadas = append(s.Chamadas, pausa)
}

func (s *SpyContagemOperacoes) Write(p []byte) (n int, err error) {
	s.Chamadas = append(s.Chamadas, escrita)
	return
}

func TestContagem(t *testing.T) {
	t.Run("teste simples", func(t *testing.T) {

		buffer := &bytes.Buffer{}
		Contagem(buffer, &SpyContagemOperacoes{})

		resultado := buffer.String()
		esperado := `3
2
1
Vai!`

		if resultado != esperado {
			t.Errorf("resultado: '%s', esperado '%s'", resultado, esperado)
		}
	})

	t.Run("pausa antes de cada impressão", func(t *testing.T) {
		spyImpressoraSleep := &SpyContagemOperacoes{}
		Contagem(spyImpressoraSleep, spyImpressoraSleep)
		esperado := []string{
			pausa,
			escrita,
			pausa,
			escrita,
			pausa,
			escrita,
			pausa,
			escrita,
		}

		if !reflect.DeepEqual(esperado, spyImpressoraSleep.Chamadas) {
			t.Errorf("esperado '%v' chamadas, resultado '%v'", esperado, spyImpressoraSleep.Chamadas)
		}
	})
}

func TestSleeprerConfiguravel(t *testing.T) {
	tempoPausa := 5 * time.Second
	tempoSpy := TempoSpy{}
	sleeper := SleeperConfiguravel{tempoPausa, tempoSpy.Pausa}
	sleeper.Pausa()

	if tempoSpy.ducacaoPausa != tempoPausa {
		t.Errorf("deveria ter pausado %v, mas pausou por %v", tempoPausa, tempoSpy.ducacaoPausa)
	}

}
