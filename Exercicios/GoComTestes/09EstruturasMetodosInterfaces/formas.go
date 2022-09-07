package formas

type Retangulo struct {
	Largura float64
	Altura  float64
}

func Perimetro(ret Retangulo) (per float64) {
	return 2 * (ret.Largura + ret.Altura)
}

func Area(ret Retangulo) (area float64) {
	return ret.Largura * ret.Altura
}
