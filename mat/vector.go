package mat

import (
	"gonum.org/v1/gonum/blas/blas64"
)

var (
	vector *VecDense

	_ Matrix        = vector
	_ allMatrix     = vector
	_ Vector        = vector
	_ Reseter       = vector
	_ MutableVector = vector
)

type Vector interface {
	Matrix
	AtVec(int) float64
	Len() int
}

type MutableVector interface {
	Vector
	SetVec(i int, v float64)
}

type TransposeVec struct {
	Vector Vector
}

func (t TransposeVec) At(i, j int) float64 { _ = "STUB: not implemented"; return 0 }

func (t TransposeVec) AtVec(i int) float64 { _ = "STUB: not implemented"; return 0 }

func (t TransposeVec) Dims() (r, c int) { _ = "STUB: not implemented"; return 0, 0 }

func (t TransposeVec) T() Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

func (t TransposeVec) Len() int { _ = "STUB: not implemented"; return 0 }

func (t TransposeVec) TVec() Vector { _ = "STUB: not implemented"; return *new(Vector) }

func (t TransposeVec) Untranspose() Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

func (t TransposeVec) UntransposeVec() Vector { _ = "STUB: not implemented"; return *new(Vector) }

type VecDense struct {
	mat blas64.Vector
}

func NewVecDense(n int, data []float64) *VecDense { _ = "STUB: not implemented"; return nil }

func (v *VecDense) SliceVec(i, k int) Vector { _ = "STUB: not implemented"; return *new(Vector) }

func (v *VecDense) sliceVec(i, k int) *VecDense { _ = "STUB: not implemented"; return nil }

func (v *VecDense) Dims() (r, c int) { _ = "STUB: not implemented"; return 0, 0 }

func (v *VecDense) Caps() (r, c int) { _ = "STUB: not implemented"; return 0, 0 }

func (v *VecDense) Len() int { _ = "STUB: not implemented"; return 0 }

func (v *VecDense) Cap() int { _ = "STUB: not implemented"; return 0 }

func (v *VecDense) T() Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

func (v *VecDense) TVec() Vector { _ = "STUB: not implemented"; return *new(Vector) }

func (v *VecDense) Reset() { _ = "STUB: not implemented"; return }

func (v *VecDense) Zero() { _ = "STUB: not implemented"; return }

func (v *VecDense) CloneFromVec(a Vector) { _ = "STUB: not implemented"; return }

func VecDenseCopyOf(a Vector) *VecDense { _ = "STUB: not implemented"; return nil }

func (v *VecDense) RawVector() blas64.Vector { _ = "STUB: not implemented"; return *new(blas64.Vector) }

func (v *VecDense) SetRawVector(a blas64.Vector) { _ = "STUB: not implemented"; return }

func (v *VecDense) CopyVec(a Vector) int { _ = "STUB: not implemented"; return 0 }

func (v *VecDense) Norm(norm float64) float64 { _ = "STUB: not implemented"; return 0 }

func (v *VecDense) ScaleVec(alpha float64, a Vector) { _ = "STUB: not implemented"; return }

func (v *VecDense) AddScaledVec(a Vector, alpha float64, b Vector) {
	_ = "STUB: not implemented"
	return
}

func (v *VecDense) AddVec(a, b Vector) { _ = "STUB: not implemented"; return }

func (v *VecDense) SubVec(a, b Vector) { _ = "STUB: not implemented"; return }

func (v *VecDense) MulElemVec(a, b Vector) { _ = "STUB: not implemented"; return }

func (v *VecDense) DivElemVec(a, b Vector) { _ = "STUB: not implemented"; return }

func (v *VecDense) MulVec(a Matrix, b Vector) { _ = "STUB: not implemented"; return }

func (v *VecDense) ReuseAsVec(n int) { _ = "STUB: not implemented"; return }

func (v *VecDense) reuseAsNonZeroed(r int) { _ = "STUB: not implemented"; return }

func (v *VecDense) reuseAsZeroed(r int) { _ = "STUB: not implemented"; return }

func (v *VecDense) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (v *VecDense) isolatedWorkspace(a Vector) (n *VecDense, restore func()) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (v *VecDense) asDense() *Dense { _ = "STUB: not implemented"; return nil }

func (v *VecDense) asGeneral() blas64.General {
	_ = "STUB: not implemented"
	return *new(blas64.General)
}

func (v *VecDense) ColViewOf(m RawMatrixer, j int) { _ = "STUB: not implemented"; return }

func (v *VecDense) RowViewOf(m RawMatrixer, i int) { _ = "STUB: not implemented"; return }

func (v *VecDense) Permute(p []int, inverse bool) { _ = "STUB: not implemented"; return }
