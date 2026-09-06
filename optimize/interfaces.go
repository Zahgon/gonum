package optimize

type localMethod interface {
	initLocal(loc *Location) (Operation, error)

	iterateLocal(loc *Location) (Operation, error)

	needser
}

type needser interface {
	needs() struct {
		Gradient bool
		Hessian  bool
	}
}

type Statuser interface {
	Status() (Status, error)
}

type Linesearcher interface {
	Init(value, derivative float64, step float64) Operation

	Iterate(value, derivative float64) (op Operation, step float64, err error)
}

type NextDirectioner interface {
	InitDirection(loc *Location, dir []float64) (step float64)

	NextDirection(loc *Location, dir []float64) (step float64)
}

type StepSizer interface {
	Init(loc *Location, dir []float64) float64
	StepSize(loc *Location, dir []float64) float64
}

type Recorder interface {
	Init() error
	Record(*Location, Operation, *Stats) error
}
