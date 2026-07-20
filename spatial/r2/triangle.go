package r2

type Triangle [3]Vec

func (t Triangle) Centroid() Vec { _ = "STUB: not implemented"; return *new(Vec) }

func (t Triangle) Area() float64 { _ = "STUB: not implemented"; return 0 }

func (t Triangle) orderedLengths() (a, b, c float64) { _ = "STUB: not implemented"; return 0, 0, 0 }

func (t Triangle) sides() (Vec, Vec, Vec) {
	_ = "STUB: not implemented"
	return *new(Vec), *new(Vec), *new(Vec)
}

func (t Triangle) IsDegenerate(tol float64) bool { _ = "STUB: not implemented"; return false }

type line [2]Vec

func (l line) vecOnLine(t float64) Vec { _ = "STUB: not implemented"; return *new(Vec) }

func (l line) distance(p Vec) float64 { _ = "STUB: not implemented"; return 0 }
