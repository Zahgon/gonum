package optimize

import (
	"gonum.org/v1/gonum/mat"
)

var _ Method = (*ListSearch)(nil)

type ListSearch struct {
	Locs mat.Matrix

	eval    int
	rows    int
	bestF   float64
	bestIdx int
}

func (*ListSearch) Uses(has Available) (uses Available, err error) {
	_ = "STUB: not implemented"
	return *new(Available), nil
}

func (l *ListSearch) Init(dim, tasks int) int { _ = "STUB: not implemented"; return 0 }

func (l *ListSearch) sendNewLoc(operation chan<- Task, task Task) {
	_ = "STUB: not implemented"
	return
}

func (l *ListSearch) updateMajor(operation chan<- Task, task Task) {
	_ = "STUB: not implemented"
	return
}

func (l *ListSearch) Status() (Status, error) { _ = "STUB: not implemented"; return *new(Status), nil }

func (l *ListSearch) Run(operation chan<- Task, result <-chan Task, tasks []Task) {
	_ = "STUB: not implemented"
	return
}
