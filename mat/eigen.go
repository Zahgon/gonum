package mat

const (
	badFact   = "mat: use without successful factorization"
	noVectors = "mat: eigenvectors not computed"
)

type EigenSym struct {
	vectorsComputed bool

	values  []float64
	vectors *Dense
}

func (e *EigenSym) Dims() (r, c int) { _ = "STUB: not implemented"; return 0, 0 }

func (e *EigenSym) SymmetricDim() int { _ = "STUB: not implemented"; return 0 }

func (e *EigenSym) At(i, j int) float64 { _ = "STUB: not implemented"; return 0 }

func (e *EigenSym) T() Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

func (e *EigenSym) Factorize(a Symmetric, vectors bool) (ok bool) {
	_ = "STUB: not implemented"
	return false
}

func (e *EigenSym) succFact() bool { _ = "STUB: not implemented"; return false }

func (e *EigenSym) Values(dst []float64) []float64 { _ = "STUB: not implemented"; return nil }

func (e *EigenSym) RawValues() []float64 { _ = "STUB: not implemented"; return nil }

func (e *EigenSym) VectorsTo(dst *Dense) { _ = "STUB: not implemented"; return }

func (e *EigenSym) RawQ() Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

type EigenKind int

const (
	EigenNone EigenKind = 0

	EigenLeft EigenKind = 1 << iota

	EigenRight

	EigenBoth EigenKind = EigenLeft | EigenRight
)

type Eigen struct {
	n int

	kind EigenKind

	values   []complex128
	rVectors *CDense
	lVectors *CDense
}

func (e *Eigen) succFact() bool { _ = "STUB: not implemented"; return false }

func (e *Eigen) Factorize(a Matrix, kind EigenKind) (ok bool) {
	_ = "STUB: not implemented"
	return false
}

func (e *Eigen) Kind() EigenKind { _ = "STUB: not implemented"; return *new(EigenKind) }

func (e *Eigen) Values(dst []complex128) []complex128 { _ = "STUB: not implemented"; return nil }

func (e *Eigen) complexEigenTo(dst *CDense, d *Dense) { _ = "STUB: not implemented"; return }

func (e *Eigen) VectorsTo(dst *CDense) { _ = "STUB: not implemented"; return }

func (e *Eigen) LeftVectorsTo(dst *CDense) { _ = "STUB: not implemented"; return }
