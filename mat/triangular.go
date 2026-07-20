package mat

import (
	"gonum.org/v1/gonum/blas"
	"gonum.org/v1/gonum/blas/blas64"
)

var (
	triDense *TriDense
	_        Matrix            = triDense
	_        allMatrix         = triDense
	_        denseMatrix       = triDense
	_        Triangular        = triDense
	_        RawTriangular     = triDense
	_        MutableTriangular = triDense

	_ NonZeroDoer    = triDense
	_ RowNonZeroDoer = triDense
	_ ColNonZeroDoer = triDense
)

type TriDense struct {
	mat blas64.Triangular
	cap int
}

type Triangular interface {
	Matrix

	Triangle() (n int, kind TriKind)

	TTri() Triangular
}

type RawTriangular interface {
	RawTriangular() blas64.Triangular
}

type MutableTriangular interface {
	Triangular
	SetTri(i, j int, v float64)
}

var (
	_ Matrix           = TransposeTri{}
	_ Triangular       = TransposeTri{}
	_ UntransposeTrier = TransposeTri{}
)

type TransposeTri struct {
	Triangular Triangular
}

func (t TransposeTri) At(i, j int) float64 { _ = "STUB: not implemented"; return 0 }

func (t TransposeTri) Dims() (r, c int) { _ = "STUB: not implemented"; return 0, 0 }

func (t TransposeTri) T() Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

func (t TransposeTri) Triangle() (int, TriKind) { _ = "STUB: not implemented"; return 0, *new(TriKind) }

func (t TransposeTri) TTri() Triangular { _ = "STUB: not implemented"; return *new(Triangular) }

func (t TransposeTri) Untranspose() Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

func (t TransposeTri) UntransposeTri() Triangular {
	_ = "STUB: not implemented"
	return *new(Triangular)
}

func NewTriDense(n int, kind TriKind, data []float64) *TriDense {
	_ = "STUB: not implemented"
	return nil
}

func (t *TriDense) Dims() (r, c int) { _ = "STUB: not implemented"; return 0, 0 }

func (t *TriDense) Triangle() (n int, kind TriKind) {
	_ = "STUB: not implemented"
	return 0, *new(TriKind)
}

func (t *TriDense) isUpper() bool { _ = "STUB: not implemented"; return false }

func (t *TriDense) triKind() TriKind { _ = "STUB: not implemented"; return *new(TriKind) }

func isUpperUplo(u blas.Uplo) bool { _ = "STUB: not implemented"; return false }

func (t *TriDense) asSymBlas() blas64.Symmetric {
	_ = "STUB: not implemented"
	return *new(blas64.Symmetric)
}

func (t *TriDense) T() Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

func (t *TriDense) TTri() Triangular { _ = "STUB: not implemented"; return *new(Triangular) }

func (t *TriDense) RawTriangular() blas64.Triangular {
	_ = "STUB: not implemented"
	return *new(blas64.Triangular)
}

func (t *TriDense) SetRawTriangular(mat blas64.Triangular) { _ = "STUB: not implemented"; return }

func (t *TriDense) Reset() { _ = "STUB: not implemented"; return }

func (t *TriDense) Zero() { _ = "STUB: not implemented"; return }

func (t *TriDense) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func untransposeTri(a Triangular) (Triangular, bool) {
	_ = "STUB: not implemented"
	return *new(Triangular), false
}

func (t *TriDense) ReuseAsTri(n int, kind TriKind) { _ = "STUB: not implemented"; return }

func (t *TriDense) reuseAsNonZeroed(n int, kind TriKind) { _ = "STUB: not implemented"; return }

func (t *TriDense) reuseAsZeroed(n int, kind TriKind) { _ = "STUB: not implemented"; return }

func (t *TriDense) isolatedWorkspace(a Triangular) (w *TriDense, restore func()) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *TriDense) DiagView() Diagonal { _ = "STUB: not implemented"; return *new(Diagonal) }

func (t *TriDense) Copy(a Matrix) (r, c int) { _ = "STUB: not implemented"; return 0, 0 }

func (t *TriDense) InverseTri(a Triangular) error { _ = "STUB: not implemented"; return nil }

func (t *TriDense) MulTri(a, b Triangular) { _ = "STUB: not implemented"; return }

func (t *TriDense) ScaleTri(f float64, a Triangular) { _ = "STUB: not implemented"; return }

func (t *TriDense) SliceTri(i, k int) Triangular {
	_ = "STUB: not implemented"
	return *new(Triangular)
}

func (t *TriDense) sliceTri(i, k int) *TriDense { _ = "STUB: not implemented"; return nil }

func (t *TriDense) Norm(norm float64) float64 { _ = "STUB: not implemented"; return 0 }

func (t *TriDense) Trace() float64 { _ = "STUB: not implemented"; return 0 }

func copySymIntoTriangle(t *TriDense, s Symmetric) { _ = "STUB: not implemented"; return }

func (t *TriDense) DoNonZero(fn func(i, j int, v float64)) { _ = "STUB: not implemented"; return }

func (t *TriDense) DoRowNonZero(i int, fn func(i, j int, v float64)) {
	_ = "STUB: not implemented"
	return
}

func (t *TriDense) DoColNonZero(j int, fn func(i, j int, v float64)) {
	_ = "STUB: not implemented"
	return
}

func (t *TriDense) SolveTo(dst *Dense, trans bool, b Matrix) error {
	_ = "STUB: not implemented"
	return nil
}
