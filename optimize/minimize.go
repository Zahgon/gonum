package optimize

import (
	"time"

	"gonum.org/v1/gonum/mat"
)

const (
	nonpositiveDimension string = "optimize: non-positive input dimension"
	negativeTasks        string = "optimize: negative input number of tasks"
)

type Task struct {
	ID int
	Op Operation
	*Location
}

type Location struct {
	X []float64

	F float64

	Gradient []float64

	Hessian *mat.SymDense
}

type Method interface {
	Init(dim, tasks int) (concurrent int)

	Run(operation chan<- Task, result <-chan Task, tasks []Task)

	Uses(has Available) (uses Available, err error)
}

func Minimize(p Problem, initX []float64, settings *Settings, method Method) (*Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getDefaultMethod(p *Problem) Method { _ = "STUB: not implemented"; return *new(Method) }

func minimize(prob *Problem, method Method, settings *Settings, converger Converger, stats *Stats, initOp Operation, initLoc, optLoc *Location, startTime time.Time) (Status, error) {
	_ = "STUB: not implemented"
	return *new(Status), nil
}

func defaultFunctionConverge() *FunctionConverge { _ = "STUB: not implemented"; return nil }

func newLocation(dim int) *Location { _ = "STUB: not implemented"; return nil }

func getInitLocation(dim int, initX []float64, initValues *Location) (Operation, *Location) {
	_ = "STUB: not implemented"
	return *new(Operation), nil
}

func checkOptimization(p Problem, dim int, recorder Recorder) error {
	_ = "STUB: not implemented"
	return nil
}

func evaluate(p *Problem, loc *Location, op Operation, x []float64) {
	_ = "STUB: not implemented"
	return
}

func updateEvaluationStats(stats *Stats, op Operation) { _ = "STUB: not implemented"; return }

func checkLocationConvergence(loc *Location, settings *Settings, converger Converger) Status {
	_ = "STUB: not implemented"
	return *new(Status)
}

func checkEvaluationLimits(p *Problem, stats *Stats, settings *Settings) (Status, error) {
	_ = "STUB: not implemented"
	return *new(Status), nil
}

func checkIterationLimits(loc *Location, stats *Stats, settings *Settings) Status {
	_ = "STUB: not implemented"
	return *new(Status)
}

func performMajorIteration(optLoc, loc *Location, stats *Stats, converger Converger, startTime time.Time, settings *Settings) Status {
	_ = "STUB: not implemented"
	return *new(Status)
}
