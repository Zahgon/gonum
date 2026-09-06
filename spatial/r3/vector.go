package r3

type Vec struct {
	X, Y, Z float64
}

func Add(p, q Vec) Vec { _ = "STUB: not implemented"; return *new(Vec) }

func Sub(p, q Vec) Vec { _ = "STUB: not implemented"; return *new(Vec) }

func Scale(f float64, p Vec) Vec { _ = "STUB: not implemented"; return *new(Vec) }

func Dot(p, q Vec) float64 { _ = "STUB: not implemented"; return 0 }

func Cross(p, q Vec) Vec { _ = "STUB: not implemented"; return *new(Vec) }

func Rotate(p Vec, alpha float64, axis Vec) Vec { _ = "STUB: not implemented"; return *new(Vec) }

func Norm(p Vec) float64 { _ = "STUB: not implemented"; return 0 }

func Norm2(p Vec) float64 { _ = "STUB: not implemented"; return 0 }

func Unit(p Vec) Vec { _ = "STUB: not implemented"; return *new(Vec) }

func Cos(p, q Vec) float64 { _ = "STUB: not implemented"; return 0 }

func Divergence(p, step Vec, field func(Vec) Vec) float64 { _ = "STUB: not implemented"; return 0 }

func Gradient(p, step Vec, field func(Vec) float64) Vec {
	_ = "STUB: not implemented"
	return *new(Vec)
}

func minElem(a, b Vec) Vec { _ = "STUB: not implemented"; return *new(Vec) }

func maxElem(a, b Vec) Vec { _ = "STUB: not implemented"; return *new(Vec) }

func absElem(a Vec) Vec { _ = "STUB: not implemented"; return *new(Vec) }

func mulElem(a, b Vec) Vec { _ = "STUB: not implemented"; return *new(Vec) }

func divElem(a, b Vec) Vec { _ = "STUB: not implemented"; return *new(Vec) }
