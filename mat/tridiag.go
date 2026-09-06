package mat

import (
	"gonum.org/v1/gonum/lapack/lapack64"
)

var (
	tridiagDense *Tridiag
	_            Matrix           = tridiagDense
	_            allMatrix        = tridiagDense
	_            denseMatrix      = tridiagDense
	_            Banded           = tridiagDense
	_            MutableBanded    = tridiagDense
	_            RawTridiagonaler = tridiagDense
)

type RawTridiagonaler interface {
	RawTridiagonal() lapack64.Tridiagonal
}

type Tridiag struct {
	mat lapack64.Tridiagonal
}

func NewTridiag(n int, dl, d, du []float64) *Tridiag { _ = "STUB: not implemented"; return nil }

func (a *Tridiag) Dims() (r, c int) { _ = "STUB: not implemented"; return 0, 0 }

func (a *Tridiag) Bandwidth() (kl, ku int) { _ = "STUB: not implemented"; return 0, 0 }

func (a *Tridiag) T() Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

func (a *Tridiag) TBand() Banded { _ = "STUB: not implemented"; return *new(Banded) }

func (a *Tridiag) RawTridiagonal() lapack64.Tridiagonal {
	_ = "STUB: not implemented"
	return *new(lapack64.Tridiagonal)
}

func (a *Tridiag) SetRawTridiagonal(mat lapack64.Tridiagonal) { _ = "STUB: not implemented"; return }

func (a *Tridiag) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (a *Tridiag) Reset() { _ = "STUB: not implemented"; return }

func (a *Tridiag) CloneFromTridiag(from *Tridiag) { _ = "STUB: not implemented"; return }

func (a *Tridiag) DiagView() Diagonal { _ = "STUB: not implemented"; return *new(Diagonal) }

func (a *Tridiag) Zero() { _ = "STUB: not implemented"; return }

func (a *Tridiag) Trace() float64 { _ = "STUB: not implemented"; return 0 }

func (a *Tridiag) Norm(norm float64) float64 { _ = "STUB: not implemented"; return 0 }

func (a *Tridiag) MulVecTo(dst *VecDense, trans bool, x Vector) { _ = "STUB: not implemented"; return }

func (a *Tridiag) SolveTo(dst *Dense, trans bool, b Matrix) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *Tridiag) SolveVecTo(dst *VecDense, trans bool, b Vector) error {
	_ = "STUB: not implemented"
	return nil
}

func (a *Tridiag) DoNonZero(fn func(i, j int, v float64)) { _ = "STUB: not implemented"; return }

func (a *Tridiag) DoRowNonZero(i int, fn func(i, j int, v float64)) {
	_ = "STUB: not implemented"
	return
}

func (a *Tridiag) DoColNonZero(j int, fn func(i, j int, v float64)) {
	_ = "STUB: not implemented"
	return
}
