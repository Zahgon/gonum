package r2

type Vec struct {
	X, Y float64
}

func Add(p, q Vec) Vec { _ = "STUB: not implemented"; return *new(Vec) }

func Sub(p, q Vec) Vec { _ = "STUB: not implemented"; return *new(Vec) }

func Scale(f float64, p Vec) Vec { _ = "STUB: not implemented"; return *new(Vec) }

func Dot(p, q Vec) float64 { _ = "STUB: not implemented"; return 0 }

func Cross(p, q Vec) float64 { _ = "STUB: not implemented"; return 0 }

func Rotate(p Vec, alpha float64, q Vec) Vec { _ = "STUB: not implemented"; return *new(Vec) }

func Norm(p Vec) float64 { _ = "STUB: not implemented"; return 0 }

func Norm2(p Vec) float64 { _ = "STUB: not implemented"; return 0 }

func Unit(p Vec) Vec { _ = "STUB: not implemented"; return *new(Vec) }

func Cos(p, q Vec) float64 { _ = "STUB: not implemented"; return 0 }

type Rotation struct {
	sin, cos float64
	p        Vec
}

func NewRotation(alpha float64, p Vec) Rotation { _ = "STUB: not implemented"; return *new(Rotation) }

func (r Rotation) Rotate(p Vec) Vec { _ = "STUB: not implemented"; return *new(Vec) }

func (r Rotation) isIdentity() bool { _ = "STUB: not implemented"; return false }

func minElem(a, b Vec) Vec { _ = "STUB: not implemented"; return *new(Vec) }

func maxElem(a, b Vec) Vec { _ = "STUB: not implemented"; return *new(Vec) }

func absElem(a Vec) Vec { _ = "STUB: not implemented"; return *new(Vec) }

func mulElem(a, b Vec) Vec { _ = "STUB: not implemented"; return *new(Vec) }

func divElem(a, b Vec) Vec { _ = "STUB: not implemented"; return *new(Vec) }
