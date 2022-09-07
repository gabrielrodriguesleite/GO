package formas

import "math"

type Forma interface {
	Area() float64
}

type Retangulo struct {
	Largura float64
	Altura  float64
}

type Triangulo struct {
	Largura float64
	Altura  float64
}

type Circulo struct {
	Raio float64
}

func Perimetro(ret Retangulo) (per float64) {
	return 2 * (ret.Largura + ret.Altura)
}

func (r Retangulo) Area() (area float64) {
	return r.Largura * r.Altura
}

func (c Circulo) Area() (area float64) {
	return math.Pi * c.Raio * c.Raio
}

func (t Triangulo) Area() (area float64) {
	return t.Largura * t.Altura / 2
}
