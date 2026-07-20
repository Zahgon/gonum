package optimize

import (
	"errors"
)

var (
	ErrZeroDimensional = errors.New("optimize: zero dimensional input")

	ErrLinesearcherFailure = errors.New("linesearch: failed to converge")

	ErrNonDescentDirection = errors.New("linesearch: non-descent search direction")

	ErrNoProgress = errors.New("linesearch: no change in location after Linesearcher step")

	ErrLinesearcherBound = errors.New("linesearch: step out of bounds")

	ErrMissingGrad = errors.New("optimize: problem does not provide needed Grad function")

	ErrMissingHess = errors.New("optimize: problem does not provide needed Hess function")
)

type ErrFunc float64

func (err ErrFunc) Error() string { _ = "STUB: not implemented"; return "" }

type ErrGrad struct {
	Grad  float64
	Index int
}

func (err ErrGrad) Error() string { _ = "STUB: not implemented"; return "" }

const badProblem = "optimize: objective function is undefined"
