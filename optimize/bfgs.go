package optimize

import (
	"gonum.org/v1/gonum/mat"
)

var (
	_ Method          = (*BFGS)(nil)
	_ localMethod     = (*BFGS)(nil)
	_ NextDirectioner = (*BFGS)(nil)
)

type BFGS struct {
	Linesearcher Linesearcher

	GradStopThreshold float64

	ls *LinesearchMethod

	status Status
	err    error

	dim  int
	x    mat.VecDense
	grad mat.VecDense
	s    mat.VecDense
	y    mat.VecDense
	tmp  mat.VecDense

	invHess *mat.SymDense

	first bool
}

func (b *BFGS) Status() (Status, error) { _ = "STUB: not implemented"; return *new(Status), nil }

func (*BFGS) Uses(has Available) (uses Available, err error) {
	_ = "STUB: not implemented"
	return *new(Available), nil
}

func (b *BFGS) Init(dim, tasks int) int { _ = "STUB: not implemented"; return 0 }

func (b *BFGS) Run(operation chan<- Task, result <-chan Task, tasks []Task) {
	_ = "STUB: not implemented"
	return
}

func (b *BFGS) initLocal(loc *Location) (Operation, error) {
	_ = "STUB: not implemented"
	return *new(Operation), nil
}

func (b *BFGS) iterateLocal(loc *Location) (Operation, error) {
	_ = "STUB: not implemented"
	return *new(Operation), nil
}

func (b *BFGS) InitDirection(loc *Location, dir []float64) (stepSize float64) {
	_ = "STUB: not implemented"
	return 0
}

func (b *BFGS) NextDirection(loc *Location, dir []float64) (stepSize float64) {
	_ = "STUB: not implemented"
	return 0
}

func (*BFGS) needs() struct {
	Gradient bool
	Hessian  bool
} {
	_ = "STUB: not implemented"
	return nil
}
