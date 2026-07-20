package optimize

import (
	"gonum.org/v1/gonum/stat/distmv"
)

var _ Method = (*GuessAndCheck)(nil)

type GuessAndCheck struct {
	Rander distmv.Rander

	bestF float64
	bestX []float64
}

func (*GuessAndCheck) Uses(has Available) (uses Available, err error) {
	_ = "STUB: not implemented"
	return *new(Available), nil
}

func (g *GuessAndCheck) Init(dim, tasks int) int { _ = "STUB: not implemented"; return 0 }

func (g *GuessAndCheck) sendNewLoc(operation chan<- Task, task Task) {
	_ = "STUB: not implemented"
	return
}

func (g *GuessAndCheck) updateMajor(operation chan<- Task, task Task) {
	_ = "STUB: not implemented"
	return
}

func (g *GuessAndCheck) Run(operation chan<- Task, result <-chan Task, tasks []Task) {
	_ = "STUB: not implemented"
	return
}
