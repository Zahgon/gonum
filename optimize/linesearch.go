package optimize

type LinesearchMethod struct {
	NextDirectioner NextDirectioner

	Linesearcher Linesearcher

	x   []float64
	dir []float64

	first     bool
	nextMajor bool
	eval      Operation

	lastStep float64
	lastOp   Operation
}

func (ls *LinesearchMethod) Init(loc *Location) (Operation, error) {
	_ = "STUB: not implemented"
	return *new(Operation), nil
}

func (ls *LinesearchMethod) Iterate(loc *Location) (Operation, error) {
	_ = "STUB: not implemented"
	return *new(Operation), nil
}

func (ls *LinesearchMethod) error(err error) (Operation, error) {
	_ = "STUB: not implemented"
	return *new(Operation), nil
}

func (ls *LinesearchMethod) initNextLinesearch(loc *Location) (Operation, error) {
	_ = "STUB: not implemented"
	return *new(Operation), nil
}

func ArmijoConditionMet(currObj, initObj, initGrad, step, decrease float64) bool {
	_ = "STUB: not implemented"
	return false
}

func StrongWolfeConditionsMet(currObj, currGrad, initObj, initGrad, step, decrease, curvature float64) bool {
	_ = "STUB: not implemented"
	return false
}

func WeakWolfeConditionsMet(currObj, currGrad, initObj, initGrad, step, decrease, curvature float64) bool {
	_ = "STUB: not implemented"
	return false
}
