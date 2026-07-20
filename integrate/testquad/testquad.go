package testquad

type Integral struct {
	Name  string
	A, B  float64
	F     func(float64) float64
	Value float64
}

func Constant(alpha float64) Integral { _ = "STUB: not implemented"; return *new(Integral) }

func Poly(degree int) Integral { _ = "STUB: not implemented"; return *new(Integral) }

func Sin() Integral { _ = "STUB: not implemented"; return *new(Integral) }

func XExpMinusX() Integral { _ = "STUB: not implemented"; return *new(Integral) }

func Sqrt() Integral { _ = "STUB: not implemented"; return *new(Integral) }

func ExpOverX2Plus1() Integral { _ = "STUB: not implemented"; return *new(Integral) }
