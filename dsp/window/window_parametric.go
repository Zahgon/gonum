package window

type Gaussian struct {
	Sigma float64
}

func (g Gaussian) Transform(seq []float64) []float64 { _ = "STUB: not implemented"; return nil }

func (g Gaussian) TransformComplex(seq []complex128) []complex128 {
	_ = "STUB: not implemented"
	return nil
}

type Tukey struct {
	Alpha float64
}

func (t Tukey) Transform(seq []float64) []float64 { _ = "STUB: not implemented"; return nil }

func (t Tukey) TransformComplex(seq []complex128) []complex128 {
	_ = "STUB: not implemented"
	return nil
}

type Values []float64

func NewValues(window func([]float64) []float64, n int) Values {
	_ = "STUB: not implemented"
	return *new(Values)
}

func (v Values) Transform(seq []float64) []float64 { _ = "STUB: not implemented"; return nil }

func (v Values) TransformTo(dst, src []float64) { _ = "STUB: not implemented"; return }

func (v Values) TransformComplex(seq []complex128) []complex128 {
	_ = "STUB: not implemented"
	return nil
}

func (v Values) TransformComplexTo(dst, src []complex128) { _ = "STUB: not implemented"; return }
