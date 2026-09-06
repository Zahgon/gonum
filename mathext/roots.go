package mathext

type objectiveFunc func(float64, []float64) float64

type fSolveResult uint8

const (
	fSolveExact fSolveResult = iota + 1

	fSolveConverged

	fSolveMaxIterations
)

const (
	machEp = 1.0 / (1 << 53)
)

func falsePosition(x1, x2, f1, f2, absErr, relErr, bisectTil float64, f objectiveFunc, fExtra []float64) (fSolveResult, float64, float64, float64) {
	_ = "STUB: not implemented"
	return *new(fSolveResult), 0, 0, 0
}
