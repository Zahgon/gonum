package optimize

var (
	_ Method          = (*GradientDescent)(nil)
	_ localMethod     = (*GradientDescent)(nil)
	_ NextDirectioner = (*GradientDescent)(nil)
)

type GradientDescent struct {
	Linesearcher Linesearcher

	StepSizer StepSizer

	GradStopThreshold float64

	ls *LinesearchMethod

	status Status
	err    error
}

func (g *GradientDescent) Status() (Status, error) {
	_ = "STUB: not implemented"
	return *new(Status), nil
}

func (*GradientDescent) Uses(has Available) (uses Available, err error) {
	_ = "STUB: not implemented"
	return *new(Available), nil
}

func (g *GradientDescent) Init(dim, tasks int) int { _ = "STUB: not implemented"; return 0 }

func (g *GradientDescent) Run(operation chan<- Task, result <-chan Task, tasks []Task) {
	_ = "STUB: not implemented"
	return
}

func (g *GradientDescent) initLocal(loc *Location) (Operation, error) {
	_ = "STUB: not implemented"
	return *new(Operation), nil
}

func (g *GradientDescent) iterateLocal(loc *Location) (Operation, error) {
	_ = "STUB: not implemented"
	return *new(Operation), nil
}

func (g *GradientDescent) InitDirection(loc *Location, dir []float64) (stepSize float64) {
	_ = "STUB: not implemented"
	return 0
}

func (g *GradientDescent) NextDirection(loc *Location, dir []float64) (stepSize float64) {
	_ = "STUB: not implemented"
	return 0
}

func (*GradientDescent) needs() struct {
	Gradient bool
	Hessian  bool
} {
	_ = "STUB: not implemented"
	return nil
}
