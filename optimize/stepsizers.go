package optimize

const (
	initialStepFactor = 1

	quadraticMinimumStepSize = 1e-3
	quadraticMaximumStepSize = 1
	quadraticThreshold       = 1e-12

	firstOrderMinimumStepSize = quadraticMinimumStepSize
	firstOrderMaximumStepSize = quadraticMaximumStepSize
)

var (
	_ StepSizer = ConstantStepSize{}
	_ StepSizer = (*QuadraticStepSize)(nil)
	_ StepSizer = (*FirstOrderStepSize)(nil)
)

type ConstantStepSize struct {
	Size float64
}

func (c ConstantStepSize) Init(_ *Location, _ []float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func (c ConstantStepSize) StepSize(_ *Location, _ []float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

type QuadraticStepSize struct {
	Threshold float64

	InitialStepFactor float64

	MinStepSize float64

	MaxStepSize float64

	fPrev        float64
	dirPrevNorm  float64
	projGradPrev float64
	xPrev        []float64
}

func (q *QuadraticStepSize) Init(loc *Location, dir []float64) (stepSize float64) {
	_ = "STUB: not implemented"
	return 0
}

func (q *QuadraticStepSize) StepSize(loc *Location, dir []float64) (stepSize float64) {
	_ = "STUB: not implemented"
	return 0
}

type FirstOrderStepSize struct {
	InitialStepFactor float64

	MinStepSize float64

	MaxStepSize float64

	dirPrevNorm  float64
	projGradPrev float64
	xPrev        []float64
}

func (fo *FirstOrderStepSize) Init(loc *Location, dir []float64) (stepSize float64) {
	_ = "STUB: not implemented"
	return 0
}

func (fo *FirstOrderStepSize) StepSize(loc *Location, dir []float64) (stepSize float64) {
	_ = "STUB: not implemented"
	return 0
}
