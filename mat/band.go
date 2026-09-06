package mat

import (
	"gonum.org/v1/gonum/blas/blas64"
)

var (
	bandDense *BandDense
	_         Matrix      = bandDense
	_         allMatrix   = bandDense
	_         denseMatrix = bandDense
	_         Banded      = bandDense
	_         RawBander   = bandDense

	_ NonZeroDoer    = bandDense
	_ RowNonZeroDoer = bandDense
	_ ColNonZeroDoer = bandDense
)

type BandDense struct {
	mat blas64.Band
}

type Banded interface {
	Matrix

	Bandwidth() (kl, ku int)

	TBand() Banded
}

type RawBander interface {
	RawBand() blas64.Band
}

type MutableBanded interface {
	Banded

	SetBand(i, j int, v float64)
}

var (
	_ Matrix            = TransposeBand{}
	_ Banded            = TransposeBand{}
	_ UntransposeBander = TransposeBand{}
)

type TransposeBand struct {
	Banded Banded
}

func (t TransposeBand) At(i, j int) float64 { _ = "STUB: not implemented"; return 0 }

func (t TransposeBand) Dims() (r, c int) { _ = "STUB: not implemented"; return 0, 0 }

func (t TransposeBand) T() Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

func (t TransposeBand) Bandwidth() (kl, ku int) { _ = "STUB: not implemented"; return 0, 0 }

func (t TransposeBand) TBand() Banded { _ = "STUB: not implemented"; return *new(Banded) }

func (t TransposeBand) Untranspose() Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

func (t TransposeBand) UntransposeBand() Banded { _ = "STUB: not implemented"; return *new(Banded) }

func NewBandDense(r, c, kl, ku int, data []float64) *BandDense {
	_ = "STUB: not implemented"
	return nil
}

func NewDiagonalRect(r, c int, data []float64) *BandDense { _ = "STUB: not implemented"; return nil }

func (b *BandDense) Dims() (r, c int) { _ = "STUB: not implemented"; return 0, 0 }

func (b *BandDense) Bandwidth() (kl, ku int) { _ = "STUB: not implemented"; return 0, 0 }

func (b *BandDense) T() Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

func (b *BandDense) TBand() Banded { _ = "STUB: not implemented"; return *new(Banded) }

func (b *BandDense) RawBand() blas64.Band { _ = "STUB: not implemented"; return *new(blas64.Band) }

func (b *BandDense) SetRawBand(mat blas64.Band) { _ = "STUB: not implemented"; return }

func (b *BandDense) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (b *BandDense) Reset() { _ = "STUB: not implemented"; return }

func (b *BandDense) DiagView() Diagonal { _ = "STUB: not implemented"; return *new(Diagonal) }

func (b *BandDense) DoNonZero(fn func(i, j int, v float64)) { _ = "STUB: not implemented"; return }

func (b *BandDense) DoRowNonZero(i int, fn func(i, j int, v float64)) {
	_ = "STUB: not implemented"
	return
}

func (b *BandDense) DoColNonZero(j int, fn func(i, j int, v float64)) {
	_ = "STUB: not implemented"
	return
}

func (b *BandDense) Zero() { _ = "STUB: not implemented"; return }

func (b *BandDense) Norm(norm float64) float64 { _ = "STUB: not implemented"; return 0 }

func (b *BandDense) Trace() float64 { _ = "STUB: not implemented"; return 0 }

func (b *BandDense) MulVecTo(dst *VecDense, trans bool, x Vector) {
	_ = "STUB: not implemented"
	return
}
