package mat

import (
	"gonum.org/v1/gonum/blas"
	"gonum.org/v1/gonum/blas/blas64"
)

var (
	dense *Dense

	_ Matrix      = dense
	_ allMatrix   = dense
	_ denseMatrix = dense
	_ Mutable     = dense

	_ ClonerFrom   = dense
	_ RowViewer    = dense
	_ ColViewer    = dense
	_ RawRowViewer = dense
	_ Grower       = dense

	_ RawMatrixSetter = dense
	_ RawMatrixer     = dense

	_ Reseter = dense
)

type Dense struct {
	mat blas64.General

	capRows, capCols int
}

func NewDense(r, c int, data []float64) *Dense { _ = "STUB: not implemented"; return nil }

func (m *Dense) ReuseAs(r, c int) { _ = "STUB: not implemented"; return }

func (m *Dense) reuseAsNonZeroed(r, c int) { _ = "STUB: not implemented"; return }

func (m *Dense) reuseAsZeroed(r, c int) { _ = "STUB: not implemented"; return }

func (m *Dense) Zero() { _ = "STUB: not implemented"; return }

func (m *Dense) isolatedWorkspace(a Matrix) (w *Dense, restore func()) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *Dense) Reset() { _ = "STUB: not implemented"; return }

func (m *Dense) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (m *Dense) asTriDense(n int, diag blas.Diag, uplo blas.Uplo) *TriDense {
	_ = "STUB: not implemented"
	return nil
}

func DenseCopyOf(a Matrix) *Dense { _ = "STUB: not implemented"; return nil }

func (m *Dense) SetRawMatrix(b blas64.General) { _ = "STUB: not implemented"; return }

func (m *Dense) RawMatrix() blas64.General { _ = "STUB: not implemented"; return *new(blas64.General) }

func (m *Dense) Dims() (r, c int) { _ = "STUB: not implemented"; return 0, 0 }

func (m *Dense) Caps() (r, c int) { _ = "STUB: not implemented"; return 0, 0 }

func (m *Dense) T() Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

func (m *Dense) ColView(j int) Vector { _ = "STUB: not implemented"; return *new(Vector) }

func (m *Dense) SetCol(j int, src []float64) { _ = "STUB: not implemented"; return }

func (m *Dense) SetRow(i int, src []float64) { _ = "STUB: not implemented"; return }

func (m *Dense) RowView(i int) Vector { _ = "STUB: not implemented"; return *new(Vector) }

func (m *Dense) RawRowView(i int) []float64 { _ = "STUB: not implemented"; return nil }

func (m *Dense) rawRowView(i int) []float64 { _ = "STUB: not implemented"; return nil }

func (m *Dense) DiagView() Diagonal { _ = "STUB: not implemented"; return *new(Diagonal) }

func (m *Dense) Slice(i, k, j, l int) Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

func (m *Dense) slice(i, k, j, l int) *Dense { _ = "STUB: not implemented"; return nil }

func (m *Dense) Grow(r, c int) Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

func (m *Dense) CloneFrom(a Matrix) { _ = "STUB: not implemented"; return }

func (m *Dense) Copy(a Matrix) (r, c int) { _ = "STUB: not implemented"; return 0, 0 }

func (m *Dense) Stack(a, b Matrix) { _ = "STUB: not implemented"; return }

func (m *Dense) Augment(a, b Matrix) { _ = "STUB: not implemented"; return }

func (m *Dense) Trace() float64 { _ = "STUB: not implemented"; return 0 }

func (m *Dense) Norm(norm float64) float64 { _ = "STUB: not implemented"; return 0 }

func (m *Dense) Permutation(n int, p []int) { _ = "STUB: not implemented"; return }

func (m *Dense) PermuteRows(p []int, inverse bool) { _ = "STUB: not implemented"; return }

func (m *Dense) PermuteCols(p []int, inverse bool) { _ = "STUB: not implemented"; return }
