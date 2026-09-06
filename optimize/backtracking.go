package optimize

const (
	defaultBacktrackingContraction = 0.5
	defaultBacktrackingDecrease    = 1e-4
	minimumBacktrackingStepSize    = 1e-20
)

var _ Linesearcher = (*Backtracking)(nil)

type Backtracking struct {
	DecreaseFactor    float64
	ContractionFactor float64
	MinimumStepSize   float64

	stepSize float64
	initF    float64
	initG    float64

	lastOp Operation
}

func (b *Backtracking) Init(f, g float64, step float64) Operation {
	_ = "STUB: not implemented"
	return *new(Operation)
}

func (b *Backtracking) Iterate(f, _ float64) (Operation, float64, error) {
	_ = "STUB: not implemented"
	return *new(Operation), 0, nil
}
