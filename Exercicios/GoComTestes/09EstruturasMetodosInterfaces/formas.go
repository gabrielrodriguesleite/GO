package formas

type Retangulo struct {
	Largura float64
	Altura  float64
}

func Perimetro(lar float64, alt float64) (per float64) {
	return 2 * (lar + alt)
}

func Area(lar float64, alt float64) (area float64) {
	return lar * alt
}
