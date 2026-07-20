package optimize

var _ Linesearcher = (*MoreThuente)(nil)

type MoreThuente struct {
	DecreaseFactor float64

	CurvatureFactor float64

	StepTolerance float64

	MinimumStep float64

	MaximumStep float64

	bracketed bool
	fInit     float64
	gInit     float64

	stage int

	step         float64
	lower, upper float64
	x            float64
	fx, gx       float64
	y            float64
	fy, gy       float64
	width        [2]float64
}

const (
	mtMinGrowthFactor float64 = 1.1
	mtMaxGrowthFactor float64 = 4
)

func (mt *MoreThuente) Init(f, g float64, step float64) Operation {
	_ = "STUB: not implemented"
	return *new(Operation)
}

func (mt *MoreThuente) Iterate(f, g float64) (Operation, float64, error) {
	_ = "STUB: not implemented"
	return *new(Operation), 0, nil
}

func (mt *MoreThuente) nextStep(fx, gx, fy, gy, f, g float64) { _ = "STUB: not implemented"; return }
