package fd

type Point struct {
	Loc   float64
	Coeff float64
}

type Formula struct {
	Stencil    []Point
	Derivative int
	Step       float64
}

func (f Formula) isZero() bool { _ = "STUB: not implemented"; return false }

type Settings struct {
	Formula Formula

	Step float64

	OriginKnown bool
	OriginValue float64

	Concurrent bool
}

var Forward = Formula{
	Stencil:    []Point{{Loc: 0, Coeff: -1}, {Loc: 1, Coeff: 1}},
	Derivative: 1,
	Step:       2e-8,
}

var Forward2nd = Formula{
	Stencil:    []Point{{Loc: 0, Coeff: 1}, {Loc: 1, Coeff: -2}, {Loc: 2, Coeff: 1}},
	Derivative: 2,
	Step:       1e-4,
}

var Backward = Formula{
	Stencil:    []Point{{Loc: -1, Coeff: -1}, {Loc: 0, Coeff: 1}},
	Derivative: 1,
	Step:       2e-8,
}

var Backward2nd = Formula{
	Stencil:    []Point{{Loc: 0, Coeff: 1}, {Loc: -1, Coeff: -2}, {Loc: -2, Coeff: 1}},
	Derivative: 2,
	Step:       1e-4,
}

var Central = Formula{
	Stencil:    []Point{{Loc: -1, Coeff: -0.5}, {Loc: 1, Coeff: 0.5}},
	Derivative: 1,
	Step:       6e-6,
}

var Central2nd = Formula{
	Stencil:    []Point{{Loc: -1, Coeff: 1}, {Loc: 0, Coeff: -2}, {Loc: 1, Coeff: 1}},
	Derivative: 2,
	Step:       1e-4,
}

var negativeStep = "fd: negative step"

func checkFormula(formula Formula) { _ = "STUB: not implemented"; return }

func computeWorkers(concurrent bool, evals int) int { _ = "STUB: not implemented"; return 0 }

func usesOrigin(stencil []Point) bool { _ = "STUB: not implemented"; return false }

func getOrigin(originKnown bool, originValue float64, f func() float64, stencil []Point) float64 {
	_ = "STUB: not implemented"
	return 0
}

const (
	badDerivOrder = "fd: invalid derivative order"
)
