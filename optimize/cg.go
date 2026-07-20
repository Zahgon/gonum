package optimize

const (
	iterationRestartFactor = 6
	angleRestartThreshold  = -0.9
)

var (
	_ Method          = (*CG)(nil)
	_ localMethod     = (*CG)(nil)
	_ NextDirectioner = (*CG)(nil)
)

type CGVariant interface {
	Init(loc *Location)

	Beta(grad, gradPrev, dirPrev []float64) float64
}

var (
	_ CGVariant = (*FletcherReeves)(nil)
	_ CGVariant = (*PolakRibierePolyak)(nil)
	_ CGVariant = (*HestenesStiefel)(nil)
	_ CGVariant = (*DaiYuan)(nil)
	_ CGVariant = (*HagerZhang)(nil)
)

type CG struct {
	Linesearcher Linesearcher

	Variant CGVariant

	InitialStep StepSizer

	IterationRestartFactor float64

	AngleRestartThreshold float64

	GradStopThreshold float64

	ls *LinesearchMethod

	status Status
	err    error

	restartAfter    int
	iterFromRestart int

	dirPrev      []float64
	gradPrev     []float64
	gradPrevNorm float64
}

func (cg *CG) Status() (Status, error) { _ = "STUB: not implemented"; return *new(Status), nil }

func (*CG) Uses(has Available) (uses Available, err error) {
	_ = "STUB: not implemented"
	return *new(Available), nil
}

func (cg *CG) Init(dim, tasks int) int { _ = "STUB: not implemented"; return 0 }

func (cg *CG) Run(operation chan<- Task, result <-chan Task, tasks []Task) {
	_ = "STUB: not implemented"
	return
}

func (cg *CG) initLocal(loc *Location) (Operation, error) {
	_ = "STUB: not implemented"
	return *new(Operation), nil
}

func (cg *CG) iterateLocal(loc *Location) (Operation, error) {
	_ = "STUB: not implemented"
	return *new(Operation), nil
}

func (cg *CG) InitDirection(loc *Location, dir []float64) (stepSize float64) {
	_ = "STUB: not implemented"
	return 0
}

func (cg *CG) NextDirection(loc *Location, dir []float64) (stepSize float64) {
	_ = "STUB: not implemented"
	return 0
}

func (*CG) needs() struct {
	Gradient bool
	Hessian  bool
} {
	_ = "STUB: not implemented"
	return nil
}

type FletcherReeves struct {
	prevNorm float64
}

func (fr *FletcherReeves) Init(loc *Location) { _ = "STUB: not implemented"; return }

func (fr *FletcherReeves) Beta(grad, _, _ []float64) (beta float64) {
	_ = "STUB: not implemented"
	return 0
}

type PolakRibierePolyak struct {
	prevNorm float64
}

func (pr *PolakRibierePolyak) Init(loc *Location) { _ = "STUB: not implemented"; return }

func (pr *PolakRibierePolyak) Beta(grad, gradPrev, _ []float64) (beta float64) {
	_ = "STUB: not implemented"
	return 0
}

type HestenesStiefel struct {
	y []float64
}

func (hs *HestenesStiefel) Init(loc *Location) { _ = "STUB: not implemented"; return }

func (hs *HestenesStiefel) Beta(grad, gradPrev, dirPrev []float64) (beta float64) {
	_ = "STUB: not implemented"
	return 0
}

type DaiYuan struct {
	y []float64
}

func (dy *DaiYuan) Init(loc *Location) { _ = "STUB: not implemented"; return }

func (dy *DaiYuan) Beta(grad, gradPrev, dirPrev []float64) (beta float64) {
	_ = "STUB: not implemented"
	return 0
}

type HagerZhang struct {
	y []float64
}

func (hz *HagerZhang) Init(loc *Location) { _ = "STUB: not implemented"; return }

func (hz *HagerZhang) Beta(grad, gradPrev, dirPrev []float64) (beta float64) {
	_ = "STUB: not implemented"
	return 0
}
