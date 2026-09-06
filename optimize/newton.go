package optimize

import (
	"gonum.org/v1/gonum/mat"
)

const maxNewtonModifications = 20

var (
	_ Method          = (*Newton)(nil)
	_ localMethod     = (*Newton)(nil)
	_ NextDirectioner = (*Newton)(nil)
)

type Newton struct {
	Linesearcher Linesearcher

	Increase float64

	GradStopThreshold float64

	status Status
	err    error

	ls *LinesearchMethod

	hess *mat.SymDense
	chol mat.Cholesky
	tau  float64
}

func (n *Newton) Status() (Status, error) { _ = "STUB: not implemented"; return *new(Status), nil }

func (*Newton) Uses(has Available) (uses Available, err error) {
	_ = "STUB: not implemented"
	return *new(Available), nil
}

func (n *Newton) Init(dim, tasks int) int { _ = "STUB: not implemented"; return 0 }

func (n *Newton) Run(operation chan<- Task, result <-chan Task, tasks []Task) {
	_ = "STUB: not implemented"
	return
}

func (n *Newton) initLocal(loc *Location) (Operation, error) {
	_ = "STUB: not implemented"
	return *new(Operation), nil
}

func (n *Newton) iterateLocal(loc *Location) (Operation, error) {
	_ = "STUB: not implemented"
	return *new(Operation), nil
}

func (n *Newton) InitDirection(loc *Location, dir []float64) (stepSize float64) {
	_ = "STUB: not implemented"
	return 0
}

func (n *Newton) NextDirection(loc *Location, dir []float64) (stepSize float64) {
	_ = "STUB: not implemented"
	return 0
}

func (n *Newton) needs() struct {
	Gradient bool
	Hessian  bool
} {
	_ = "STUB: not implemented"
	return nil
}
