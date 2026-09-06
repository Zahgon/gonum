package optimize

import (
	"time"

	"gonum.org/v1/gonum/mat"
)

const defaultGradientAbsTol = 1e-12

type Operation uint64

const (
	NoOperation Operation = 0

	InitIteration Operation = 1 << (iota - 1)

	PostIteration

	MajorIteration

	MethodDone

	FuncEvaluation

	GradEvaluation

	HessEvaluation

	signalDone

	evalMask = FuncEvaluation | GradEvaluation | HessEvaluation
)

func (op Operation) isEvaluation() bool { _ = "STUB: not implemented"; return false }

func (op Operation) String() string { _ = "STUB: not implemented"; return "" }

var operationNames = map[Operation]string{
	NoOperation:    "NoOperation",
	InitIteration:  "InitIteration",
	MajorIteration: "MajorIteration",
	PostIteration:  "PostIteration",
	MethodDone:     "MethodDone",
	signalDone:     "signalDone",
}

type Result struct {
	Location
	Stats
	Status Status
}

type Stats struct {
	MajorIterations int
	FuncEvaluations int
	GradEvaluations int
	HessEvaluations int
	Runtime         time.Duration
}

func complementEval(loc *Location, eval Operation) (complEval Operation) {
	_ = "STUB: not implemented"
	return *new(Operation)
}

type Problem struct {
	Func func(x []float64) float64

	Grad func(grad, x []float64)

	Hess func(hess *mat.SymDense, x []float64)

	Status func() (Status, error)
}

type Available struct {
	Grad bool
	Hess bool
}

func availFromProblem(prob Problem) Available { _ = "STUB: not implemented"; return *new(Available) }

func (has Available) function() (uses Available, err error) {
	_ = "STUB: not implemented"
	return *new(Available), nil
}

func (has Available) gradient() (uses Available, err error) {
	_ = "STUB: not implemented"
	return *new(Available), nil
}

func (has Available) hessian() (uses Available, err error) {
	_ = "STUB: not implemented"
	return *new(Available), nil
}

type Settings struct {
	InitValues *Location

	GradientThreshold float64

	Converger Converger

	MajorIterations int

	Runtime time.Duration

	FuncEvaluations int

	GradEvaluations int

	HessEvaluations int

	Recorder Recorder

	Concurrent int
}

func resize(x []float64, dim int) []float64 { _ = "STUB: not implemented"; return nil }

func resizeSymDense(m *mat.SymDense, dim int) *mat.SymDense { _ = "STUB: not implemented"; return nil }
