package mat

import (
	"gonum.org/v1/gonum/blas/cblas128"
)

var (
	cDense *CDense

	_ CMatrix   = cDense
	_ allMatrix = cDense
)

type CDense struct {
	mat cblas128.General

	capRows, capCols int
}

func (m *CDense) Dims() (r, c int) { _ = "STUB: not implemented"; return 0, 0 }

func (m *CDense) Caps() (r, c int) { _ = "STUB: not implemented"; return 0, 0 }

func (m *CDense) H() CMatrix { _ = "STUB: not implemented"; return *new(CMatrix) }

func (m *CDense) T() CMatrix { _ = "STUB: not implemented"; return *new(CMatrix) }

func (m *CDense) Conj(a CMatrix) { _ = "STUB: not implemented"; return }

func (m *CDense) Slice(i, k, j, l int) CMatrix { _ = "STUB: not implemented"; return *new(CMatrix) }

func (m *CDense) slice(i, k, j, l int) *CDense { _ = "STUB: not implemented"; return nil }

func NewCDense(r, c int, data []complex128) *CDense { _ = "STUB: not implemented"; return nil }

func (m *CDense) ReuseAs(r, c int) { _ = "STUB: not implemented"; return }

func (m *CDense) reuseAsNonZeroed(r, c int) { _ = "STUB: not implemented"; return }

func (m *CDense) reuseAsZeroed(r, c int) { _ = "STUB: not implemented"; return }

func (m *CDense) isolatedWorkspace(a CMatrix) (w *CDense, restore func()) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *CDense) Reset() { _ = "STUB: not implemented"; return }

func (m *CDense) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (m *CDense) Zero() { _ = "STUB: not implemented"; return }

func (m *CDense) Copy(a CMatrix) (r, c int) { _ = "STUB: not implemented"; return 0, 0 }

func (m *CDense) SetRawCMatrix(b cblas128.General) { _ = "STUB: not implemented"; return }

func (m *CDense) RawCMatrix() cblas128.General {
	_ = "STUB: not implemented"
	return *new(cblas128.General)
}

func (m *CDense) Grow(r, c int) CMatrix { _ = "STUB: not implemented"; return *new(CMatrix) }
