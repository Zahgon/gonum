package mat

import (
	"gonum.org/v1/gonum/lapack"
)

const (
	badSliceLength = "mat: improper slice length"
	badLU          = "mat: invalid LU factorization"
)

type LU struct {
	lu    *Dense
	swaps []int
	piv   []int
	cond  float64
	ok    bool
}

var _ Matrix = (*LU)(nil)

func (lu *LU) Dims() (r, c int) { _ = "STUB: not implemented"; return 0, 0 }

func (lu *LU) At(i, j int) float64 { _ = "STUB: not implemented"; return 0 }

func (lu *LU) T() Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

func (lu *LU) updateCond(anorm float64, norm lapack.MatrixNorm) { _ = "STUB: not implemented"; return }

func (lu *LU) Factorize(a Matrix) { _ = "STUB: not implemented"; return }

func (lu *LU) factorize(a Matrix, norm lapack.MatrixNorm) { _ = "STUB: not implemented"; return }

func (lu *LU) updatePivots(swaps []int) { _ = "STUB: not implemented"; return }

func (lu *LU) isValid() bool { _ = "STUB: not implemented"; return false }

func (lu *LU) Cond() float64 { _ = "STUB: not implemented"; return 0 }

func (lu *LU) Reset() { _ = "STUB: not implemented"; return }

func (lu *LU) isZero() bool { _ = "STUB: not implemented"; return false }

func (lu *LU) Det() float64 { _ = "STUB: not implemented"; return 0 }

func (lu *LU) LogDet() (det float64, sign float64) { _ = "STUB: not implemented"; return 0, 0 }

func (lu *LU) RowPivots(dst []int) []int { _ = "STUB: not implemented"; return nil }

func (lu *LU) Pivot(dst []int) []int { _ = "STUB: not implemented"; return nil }

func (lu *LU) RankOne(orig *LU, alpha float64, x, y Vector) { _ = "STUB: not implemented"; return }

func (lu *LU) LTo(dst *TriDense) *TriDense { _ = "STUB: not implemented"; return nil }

func (lu *LU) UTo(dst *TriDense) { _ = "STUB: not implemented"; return }

func (lu *LU) SolveTo(dst *Dense, trans bool, b Matrix) error {
	_ = "STUB: not implemented"
	return nil
}

func (lu *LU) SolveVecTo(dst *VecDense, trans bool, b Vector) error {
	_ = "STUB: not implemented"
	return nil
}
