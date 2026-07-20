package optimize

var (
	_ Method          = (*LBFGS)(nil)
	_ localMethod     = (*LBFGS)(nil)
	_ NextDirectioner = (*LBFGS)(nil)
)

type LBFGS struct {
	Linesearcher Linesearcher

	Store int

	GradStopThreshold float64

	status Status
	err    error

	ls *LinesearchMethod

	dim  int
	x    []float64
	grad []float64

	oldest int
	y      [][]float64
	s      [][]float64
	rho    []float64
	a      []float64
}

func (l *LBFGS) Status() (Status, error) { _ = "STUB: not implemented"; return *new(Status), nil }

func (*LBFGS) Uses(has Available) (uses Available, err error) {
	_ = "STUB: not implemented"
	return *new(Available), nil
}

func (l *LBFGS) Init(dim, tasks int) int { _ = "STUB: not implemented"; return 0 }

func (l *LBFGS) Run(operation chan<- Task, result <-chan Task, tasks []Task) {
	_ = "STUB: not implemented"
	return
}

func (l *LBFGS) initLocal(loc *Location) (Operation, error) {
	_ = "STUB: not implemented"
	return *new(Operation), nil
}

func (l *LBFGS) iterateLocal(loc *Location) (Operation, error) {
	_ = "STUB: not implemented"
	return *new(Operation), nil
}

func (l *LBFGS) InitDirection(loc *Location, dir []float64) (stepSize float64) {
	_ = "STUB: not implemented"
	return 0
}

func (l *LBFGS) initHistory(hist [][]float64) [][]float64 { _ = "STUB: not implemented"; return nil }

func (l *LBFGS) NextDirection(loc *Location, dir []float64) (stepSize float64) {
	_ = "STUB: not implemented"
	return 0
}

func (*LBFGS) needs() struct {
	Gradient bool
	Hessian  bool
} {
	_ = "STUB: not implemented"
	return nil
}
