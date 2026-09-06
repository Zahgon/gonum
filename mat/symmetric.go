package mat

import (
	"gonum.org/v1/gonum/blas/blas64"
)

var (
	symDense *SymDense

	_ Matrix           = symDense
	_ allMatrix        = symDense
	_ denseMatrix      = symDense
	_ Symmetric        = symDense
	_ RawSymmetricer   = symDense
	_ MutableSymmetric = symDense
)

const badSymTriangle = "mat: blas64.Symmetric not upper"

type SymDense struct {
	mat blas64.Symmetric
	cap int
}

type Symmetric interface {
	Matrix

	SymmetricDim() int
}

type RawSymmetricer interface {
	RawSymmetric() blas64.Symmetric
}

type MutableSymmetric interface {
	Symmetric
	SetSym(i, j int, v float64)
}

func NewSymDense(n int, data []float64) *SymDense { _ = "STUB: not implemented"; return nil }

func (s *SymDense) Dims() (r, c int) { _ = "STUB: not implemented"; return 0, 0 }

func (s *SymDense) Caps() (r, c int) { _ = "STUB: not implemented"; return 0, 0 }

func (s *SymDense) T() Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

func (s *SymDense) SymmetricDim() int { _ = "STUB: not implemented"; return 0 }

func (s *SymDense) RawSymmetric() blas64.Symmetric {
	_ = "STUB: not implemented"
	return *new(blas64.Symmetric)
}

func (s *SymDense) SetRawSymmetric(mat blas64.Symmetric) { _ = "STUB: not implemented"; return }

func (s *SymDense) Reset() { _ = "STUB: not implemented"; return }

func (s *SymDense) ReuseAsSym(n int) { _ = "STUB: not implemented"; return }

func (s *SymDense) Zero() { _ = "STUB: not implemented"; return }

func (s *SymDense) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (s *SymDense) reuseAsNonZeroed(n int) { _ = "STUB: not implemented"; return }

func (s *SymDense) reuseAsZeroed(n int) { _ = "STUB: not implemented"; return }

func (s *SymDense) isolatedWorkspace(a Symmetric) (w *SymDense, restore func()) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SymDense) DiagView() Diagonal { _ = "STUB: not implemented"; return *new(Diagonal) }

func (s *SymDense) AddSym(a, b Symmetric) { _ = "STUB: not implemented"; return }

func (s *SymDense) CopySym(a Symmetric) int { _ = "STUB: not implemented"; return 0 }

func (s *SymDense) SymRankOne(a Symmetric, alpha float64, x Vector) {
	_ = "STUB: not implemented"
	return
}

func (s *SymDense) SymRankK(a Symmetric, alpha float64, x Matrix) {
	_ = "STUB: not implemented"
	return
}

func (s *SymDense) SymOuterK(alpha float64, x Matrix) { _ = "STUB: not implemented"; return }

func (s *SymDense) RankTwo(a Symmetric, alpha float64, x, y Vector) {
	_ = "STUB: not implemented"
	return
}

func (s *SymDense) ScaleSym(f float64, a Symmetric) { _ = "STUB: not implemented"; return }

func (s *SymDense) SubsetSym(a Symmetric, set []int) { _ = "STUB: not implemented"; return }

func (s *SymDense) SliceSym(i, k int) Symmetric { _ = "STUB: not implemented"; return *new(Symmetric) }

func (s *SymDense) sliceSym(i, k int) *SymDense { _ = "STUB: not implemented"; return nil }

func (s *SymDense) Norm(norm float64) float64 { _ = "STUB: not implemented"; return 0 }

func (s *SymDense) Trace() float64 { _ = "STUB: not implemented"; return 0 }

func (s *SymDense) GrowSym(n int) Symmetric { _ = "STUB: not implemented"; return *new(Symmetric) }

func (s *SymDense) PowPSD(a Symmetric, pow float64) error { _ = "STUB: not implemented"; return nil }
