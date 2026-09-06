package optimize

import "errors"

type Status int

const (
	NotTerminated Status = iota
	Success
	FunctionThreshold
	FunctionConvergence
	GradientThreshold
	StepConvergence
	FunctionNegativeInfinity
	MethodConverge
	Failure
	IterationLimit
	RuntimeLimit
	FunctionEvaluationLimit
	GradientEvaluationLimit
	HessianEvaluationLimit
)

func (s Status) String() string { _ = "STUB: not implemented"; return "" }

func (s Status) Early() bool { _ = "STUB: not implemented"; return false }

func (s Status) Err() error { _ = "STUB: not implemented"; return nil }

var statuses = []struct {
	name  string
	early bool
	err   error
}{
	{
		name: "NotTerminated",
	},
	{
		name: "Success",
	},
	{
		name: "FunctionThreshold",
	},
	{
		name: "FunctionConvergence",
	},
	{
		name: "GradientThreshold",
	},
	{
		name: "StepConvergence",
	},
	{
		name: "FunctionNegativeInfinity",
	},
	{
		name: "MethodConverge",
	},
	{
		name:  "Failure",
		early: true,
		err:   errors.New("optimize: termination ended in failure"),
	},
	{
		name:  "IterationLimit",
		early: true,
		err:   errors.New("optimize: maximum number of major iterations reached"),
	},
	{
		name:  "RuntimeLimit",
		early: true,
		err:   errors.New("optimize: maximum runtime reached"),
	},
	{
		name:  "FunctionEvaluationLimit",
		early: true,
		err:   errors.New("optimize: maximum number of function evaluations reached"),
	},
	{
		name:  "GradientEvaluationLimit",
		early: true,
		err:   errors.New("optimize: maximum number of gradient evaluations reached"),
	},
	{
		name:  "HessianEvaluationLimit",
		early: true,
		err:   errors.New("optimize: maximum number of Hessian evaluations reached"),
	},
}

func NewStatus(name string, early bool, err error) Status {
	_ = "STUB: not implemented"
	return *new(Status)
}
