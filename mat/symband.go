package mat

import (
	"gonum.org/v1/gonum/blas/blas64"
)

var (
	symBandDense *SymBandDense
	_            Matrix           = symBandDense
	_            allMatrix        = symBandDense
	_            denseMatrix      = symBandDense
	_            Symmetric        = symBandDense
	_            Banded           = symBandDense
	_            SymBanded        = symBandDense
	_            RawSymBander     = symBandDense
	_            MutableSymBanded = symBandDense

	_ NonZeroDoer    = symBandDense
	_ RowNonZeroDoer = symBandDense
	_ ColNonZeroDoer = symBandDense
)

type SymBandDense struct {
	mat blas64.SymmetricBand
}

type SymBanded interface {
	Banded

	SymmetricDim() int

	SymBand() (n, k int)
}

type MutableSymBanded interface {
	SymBanded
	SetSymBand(i, j int, v float64)
}

type RawSymBander interface {
	RawSymBand() blas64.SymmetricBand
}

func NewSymBandDense(n, k int, data []float64) *SymBandDense { _ = "STUB: not implemented"; return nil }

func (s *SymBandDense) Dims() (r, c int) { _ = "STUB: not implemented"; return 0, 0 }

func (s *SymBandDense) SymmetricDim() int { _ = "STUB: not implemented"; return 0 }

func (s *SymBandDense) Bandwidth() (kl, ku int) { _ = "STUB: not implemented"; return 0, 0 }

func (s *SymBandDense) SymBand() (n, k int) { _ = "STUB: not implemented"; return 0, 0 }

func (s *SymBandDense) T() Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

func (s *SymBandDense) TBand() Banded { _ = "STUB: not implemented"; return *new(Banded) }

func (s *SymBandDense) RawSymBand() blas64.SymmetricBand {
	_ = "STUB: not implemented"
	return *new(blas64.SymmetricBand)
}

func (s *SymBandDense) SetRawSymBand(mat blas64.SymmetricBand) { _ = "STUB: not implemented"; return }

func (s *SymBandDense) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (s *SymBandDense) Reset() { _ = "STUB: not implemented"; return }

func (s *SymBandDense) Zero() { _ = "STUB: not implemented"; return }

func (s *SymBandDense) DiagView() Diagonal { _ = "STUB: not implemented"; return *new(Diagonal) }

func (s *SymBandDense) DoNonZero(fn func(i, j int, v float64)) { _ = "STUB: not implemented"; return }

func (s *SymBandDense) DoRowNonZero(i int, fn func(i, j int, v float64)) {
	_ = "STUB: not implemented"
	return
}

func (s *SymBandDense) DoColNonZero(j int, fn func(i, j int, v float64)) {
	_ = "STUB: not implemented"
	return
}

func (s *SymBandDense) Norm(norm float64) float64 { _ = "STUB: not implemented"; return 0 }

func (s *SymBandDense) Trace() float64 { _ = "STUB: not implemented"; return 0 }

func (s *SymBandDense) MulVecTo(dst *VecDense, _ bool, x Vector) { _ = "STUB: not implemented"; return }
