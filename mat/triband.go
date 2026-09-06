package mat

import (
	"gonum.org/v1/gonum/blas/blas64"
)

var (
	triBand TriBanded
	_       Banded     = triBand
	_       Triangular = triBand

	triBandDense *TriBandDense
	_            Matrix           = triBandDense
	_            allMatrix        = triBandDense
	_            denseMatrix      = triBandDense
	_            Triangular       = triBandDense
	_            Banded           = triBandDense
	_            TriBanded        = triBandDense
	_            RawTriBander     = triBandDense
	_            MutableTriBanded = triBandDense
)

type TriBanded interface {
	Banded

	Triangle() (n int, kind TriKind)

	TTri() Triangular

	TriBand() (n, k int, kind TriKind)

	TTriBand() TriBanded
}

type RawTriBander interface {
	RawTriBand() blas64.TriangularBand
}

type MutableTriBanded interface {
	TriBanded
	SetTriBand(i, j int, v float64)
}

var (
	tTriBand TransposeTriBand
	_        Matrix               = tTriBand
	_        TriBanded            = tTriBand
	_        Untransposer         = tTriBand
	_        UntransposeTrier     = tTriBand
	_        UntransposeBander    = tTriBand
	_        UntransposeTriBander = tTriBand
)

type TransposeTriBand struct {
	TriBanded TriBanded
}

func (t TransposeTriBand) At(i, j int) float64 { _ = "STUB: not implemented"; return 0 }

func (t TransposeTriBand) Dims() (r, c int) { _ = "STUB: not implemented"; return 0, 0 }

func (t TransposeTriBand) T() Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

func (t TransposeTriBand) Triangle() (int, TriKind) {
	_ = "STUB: not implemented"
	return 0, *new(TriKind)
}

func (t TransposeTriBand) TTri() Triangular { _ = "STUB: not implemented"; return *new(Triangular) }

func (t TransposeTriBand) Bandwidth() (kl, ku int) { _ = "STUB: not implemented"; return 0, 0 }

func (t TransposeTriBand) TBand() Banded { _ = "STUB: not implemented"; return *new(Banded) }

func (t TransposeTriBand) TriBand() (n, k int, kind TriKind) {
	_ = "STUB: not implemented"
	return 0, 0, *new(TriKind)
}

func (t TransposeTriBand) TTriBand() TriBanded { _ = "STUB: not implemented"; return *new(TriBanded) }

func (t TransposeTriBand) Untranspose() Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

func (t TransposeTriBand) UntransposeTri() Triangular {
	_ = "STUB: not implemented"
	return *new(Triangular)
}

func (t TransposeTriBand) UntransposeBand() Banded { _ = "STUB: not implemented"; return *new(Banded) }

func (t TransposeTriBand) UntransposeTriBand() TriBanded {
	_ = "STUB: not implemented"
	return *new(TriBanded)
}

type TriBandDense struct {
	mat blas64.TriangularBand
}

func NewTriBandDense(n, k int, kind TriKind, data []float64) *TriBandDense {
	_ = "STUB: not implemented"
	return nil
}

func (t *TriBandDense) Dims() (r, c int) { _ = "STUB: not implemented"; return 0, 0 }

func (t *TriBandDense) T() Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

func (t *TriBandDense) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (t *TriBandDense) Reset() { _ = "STUB: not implemented"; return }

func (t *TriBandDense) ReuseAsTriBand(n, k int, kind TriKind) { _ = "STUB: not implemented"; return }

func (t *TriBandDense) reuseAsZeroed(n, k int, kind TriKind) { _ = "STUB: not implemented"; return }

//lint:ignore U1000 This will be used later.
func (t *TriBandDense) reuseAsNonZeroed(n, k int, kind TriKind) { _ = "STUB: not implemented"; return }

func (t *TriBandDense) DoNonZero(fn func(i, j int, v float64)) { _ = "STUB: not implemented"; return }

func (t *TriBandDense) DoRowNonZero(i int, fn func(i, j int, v float64)) {
	_ = "STUB: not implemented"
	return
}

func (t *TriBandDense) DoColNonZero(j int, fn func(i, j int, v float64)) {
	_ = "STUB: not implemented"
	return
}

func (t *TriBandDense) Zero() { _ = "STUB: not implemented"; return }

func (t *TriBandDense) isUpper() bool { _ = "STUB: not implemented"; return false }

func (t *TriBandDense) triKind() TriKind { _ = "STUB: not implemented"; return *new(TriKind) }

func (t *TriBandDense) Triangle() (n int, kind TriKind) {
	_ = "STUB: not implemented"
	return 0, *new(TriKind)
}

func (t *TriBandDense) TTri() Triangular { _ = "STUB: not implemented"; return *new(Triangular) }

func (t *TriBandDense) Bandwidth() (kl, ku int) { _ = "STUB: not implemented"; return 0, 0 }

func (t *TriBandDense) TBand() Banded { _ = "STUB: not implemented"; return *new(Banded) }

func (t *TriBandDense) TriBand() (n, k int, kind TriKind) {
	_ = "STUB: not implemented"
	return 0, 0, *new(TriKind)
}

func (t *TriBandDense) TTriBand() TriBanded { _ = "STUB: not implemented"; return *new(TriBanded) }

func (t *TriBandDense) RawTriBand() blas64.TriangularBand {
	_ = "STUB: not implemented"
	return *new(blas64.TriangularBand)
}

func (t *TriBandDense) SetRawTriBand(mat blas64.TriangularBand) { _ = "STUB: not implemented"; return }

func (t *TriBandDense) DiagView() Diagonal { _ = "STUB: not implemented"; return *new(Diagonal) }

func (t *TriBandDense) Norm(norm float64) float64 { _ = "STUB: not implemented"; return 0 }

func (t *TriBandDense) Trace() float64 { _ = "STUB: not implemented"; return 0 }

func (t *TriBandDense) SolveTo(dst *Dense, trans bool, b Matrix) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *TriBandDense) SolveVecTo(dst *VecDense, trans bool, b Vector) error {
	_ = "STUB: not implemented"
	return nil
}

func copySymBandIntoTriBand(dst *TriBandDense, s SymBanded) { _ = "STUB: not implemented"; return }
