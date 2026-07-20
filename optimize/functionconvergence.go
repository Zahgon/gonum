package optimize

type Converger interface {
	Init(dim int)
	Converged(loc *Location) Status
}

var (
	_ Converger = NeverTerminate{}
	_ Converger = (*FunctionConverge)(nil)
)

type NeverTerminate struct{}

func (NeverTerminate) Init(dim int) { _ = "STUB: not implemented"; return }

func (NeverTerminate) Converged(loc *Location) Status {
	_ = "STUB: not implemented"
	return *new(Status)
}

type FunctionConverge struct {
	Absolute   float64
	Relative   float64
	Iterations int

	first bool
	best  float64
	iter  int
}

func (fc *FunctionConverge) Init(dim int) { _ = "STUB: not implemented"; return }

func (fc *FunctionConverge) Converged(l *Location) Status {
	_ = "STUB: not implemented"
	return *new(Status)
}
