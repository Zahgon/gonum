package mat

import (
	"gonum.org/v1/gonum/blas/blas64"
)

var (
	diagDense *DiagDense
	_         Matrix          = diagDense
	_         allMatrix       = diagDense
	_         denseMatrix     = diagDense
	_         Diagonal        = diagDense
	_         MutableDiagonal = diagDense
	_         Triangular      = diagDense
	_         TriBanded       = diagDense
	_         Symmetric       = diagDense
	_         SymBanded       = diagDense
	_         Banded          = diagDense
	_         RawBander       = diagDense
	_         RawSymBander    = diagDense

	diag Diagonal
	_    Matrix     = diag
	_    Diagonal   = diag
	_    Triangular = diag
	_    TriBanded  = diag
	_    Symmetric  = diag
	_    SymBanded  = diag
	_    Banded     = diag
)

type Diagonal interface {
	Matrix

	Diag() int

	Banded
	SymBanded
	Symmetric
	Triangular
	TriBanded
}

type MutableDiagonal interface {
	Diagonal
	SetDiag(i int, v float64)
}

type DiagDense struct {
	mat blas64.Vector
}

func NewDiagDense(n int, data []float64) *DiagDense { _ = "STUB: not implemented"; return nil }

func (d *DiagDense) Diag() int { _ = "STUB: not implemented"; return 0 }

func (d *DiagDense) Dims() (r, c int) { _ = "STUB: not implemented"; return 0, 0 }

func (d *DiagDense) T() Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

func (d *DiagDense) TTri() Triangular { _ = "STUB: not implemented"; return *new(Triangular) }

func (d *DiagDense) TBand() Banded { _ = "STUB: not implemented"; return *new(Banded) }

func (d *DiagDense) TTriBand() TriBanded { _ = "STUB: not implemented"; return *new(TriBanded) }

func (d *DiagDense) Bandwidth() (kl, ku int) { _ = "STUB: not implemented"; return 0, 0 }

func (d *DiagDense) SymmetricDim() int { _ = "STUB: not implemented"; return 0 }

func (d *DiagDense) SymBand() (n, k int) { _ = "STUB: not implemented"; return 0, 0 }

func (d *DiagDense) Triangle() (int, TriKind) { _ = "STUB: not implemented"; return 0, *new(TriKind) }

func (d *DiagDense) TriBand() (n, k int, kind TriKind) {
	_ = "STUB: not implemented"
	return 0, 0, *new(TriKind)
}

func (d *DiagDense) Reset() { _ = "STUB: not implemented"; return }

func (d *DiagDense) Zero() { _ = "STUB: not implemented"; return }

func (d *DiagDense) DiagView() Diagonal { _ = "STUB: not implemented"; return *new(Diagonal) }

func (d *DiagDense) DiagFrom(m Matrix) { _ = "STUB: not implemented"; return }

func (d *DiagDense) RawBand() blas64.Band { _ = "STUB: not implemented"; return *new(blas64.Band) }

func (d *DiagDense) RawSymBand() blas64.SymmetricBand {
	_ = "STUB: not implemented"
	return *new(blas64.SymmetricBand)
}

func (d *DiagDense) reuseAsNonZeroed(r int) { _ = "STUB: not implemented"; return }

func (d *DiagDense) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (d *DiagDense) Trace() float64 { _ = "STUB: not implemented"; return 0 }

func (d *DiagDense) Norm(norm float64) float64 { _ = "STUB: not implemented"; return 0 }
