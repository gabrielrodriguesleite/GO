package formas

import "math"

type Retangulo struct {
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
