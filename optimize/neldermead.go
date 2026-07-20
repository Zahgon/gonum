package optimize

type nmIterType int

const (
	nmReflected = iota
	nmExpanded
	nmContractedInside
	nmContractedOutside
	nmInitialize
	nmShrink
	nmMajor
)

type nmVertexSorter struct {
	vertices [][]float64
	values   []float64
}

func (n nmVertexSorter) Len() int { _ = "STUB: not implemented"; return 0 }

func (n nmVertexSorter) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (n nmVertexSorter) Swap(i, j int) { _ = "STUB: not implemented"; return }

var _ Method = (*NelderMead)(nil)

type NelderMead struct {
	InitialVertices [][]float64
	InitialValues   []float64
	Reflection      float64
	Expansion       float64
	Contraction     float64
	Shrink          float64
	SimplexSize     float64

	status Status
	err    error

	reflection  float64
	expansion   float64
	contraction float64
	shrink      float64

	vertices [][]float64
	values   []float64
	centroid []float64

	fillIdx        int
	lastIter       nmIterType
	reflectedPoint []float64
	reflectedValue float64
}

func (n *NelderMead) Status() (Status, error) { _ = "STUB: not implemented"; return *new(Status), nil }

func (*NelderMead) Uses(has Available) (uses Available, err error) {
	_ = "STUB: not implemented"
	return *new(Available), nil
}

func (n *NelderMead) Init(dim, tasks int) int { _ = "STUB: not implemented"; return 0 }

func (n *NelderMead) Run(operation chan<- Task, result <-chan Task, tasks []Task) {
	_ = "STUB: not implemented"
	return
}

func (n *NelderMead) initLocal(loc *Location) (Operation, error) {
	_ = "STUB: not implemented"
	return *new(Operation), nil
}

func computeCentroid(vertices [][]float64, centroid []float64) { _ = "STUB: not implemented"; return }

func (n *NelderMead) iterateLocal(loc *Location) (Operation, error) {
	_ = "STUB: not implemented"
	return *new(Operation), nil
}

func (n *NelderMead) returnNext(iter nmIterType, loc *Location) (Operation, error) {
	_ = "STUB: not implemented"
	return *new(Operation), nil
}

func (n *NelderMead) replaceWorst(x []float64, f float64) { _ = "STUB: not implemented"; return }

func (*NelderMead) needs() struct {
	Gradient bool
	Hessian  bool
} {
	_ = "STUB: not implemented"
	return nil
}
