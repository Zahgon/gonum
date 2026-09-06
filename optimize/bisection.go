package optimize

const defaultBisectionCurvature = 0.9

var _ Linesearcher = (*Bisection)(nil)

type Bisection struct {
	CurvatureFactor float64

	minStep  float64
	maxStep  float64
	currStep float64

	initF float64
	minF  float64
	maxF  float64
	lastF float64

	initGrad float64

	lastOp Operation
}

func (b *Bisection) Init(f, g float64, step float64) Operation {
	_ = "STUB: not implemented"
	return *new(Operation)
}

func (b *Bisection) Iterate(f, g float64) (Operation, float64, error) {
	_ = "STUB: not implemented"
	return *new(Operation), 0, nil
}

func (b *Bisection) nextStep(step float64) (Operation, float64, error) {
	_ = "STUB: not implemented"
	return *new(Operation), 0, nil
}
