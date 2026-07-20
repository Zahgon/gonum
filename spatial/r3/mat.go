package r3

import "gonum.org/v1/gonum/mat"

type Mat struct {
	data *array
}

var _ mat.Matrix = (*Mat)(nil)

func NewMat(val []float64) *Mat { _ = "STUB: not implemented"; return nil }

func (m *Mat) Dims() (r, c int) { _ = "STUB: not implemented"; return 0, 0 }

func (m *Mat) T() mat.Matrix { _ = "STUB: not implemented"; return *new(mat.Matrix) }

func (m *Mat) Scale(f float64, a mat.Matrix) { _ = "STUB: not implemented"; return }

func (m *Mat) MulVec(v Vec) Vec { _ = "STUB: not implemented"; return *new(Vec) }

func (m *Mat) MulVecTrans(v Vec) Vec { _ = "STUB: not implemented"; return *new(Vec) }

func (m *Mat) CloneFrom(a mat.Matrix) { _ = "STUB: not implemented"; return }

func (m *Mat) Sub(a, b mat.Matrix) { _ = "STUB: not implemented"; return }

func (m *Mat) Add(a, b mat.Matrix) { _ = "STUB: not implemented"; return }

func (m *Mat) VecRow(i int) Vec { _ = "STUB: not implemented"; return *new(Vec) }

func (m *Mat) VecCol(j int) Vec { _ = "STUB: not implemented"; return *new(Vec) }

func (m *Mat) Outer(alpha float64, x, y Vec) { _ = "STUB: not implemented"; return }

func (m *Mat) Det() float64 { _ = "STUB: not implemented"; return 0 }

func (m *Mat) Skew(v Vec) { _ = "STUB: not implemented"; return }

func (m *Mat) Hessian(p, step Vec, field func(Vec) float64) { _ = "STUB: not implemented"; return }

func (m *Mat) Jacobian(p, step Vec, field func(Vec) Vec) { _ = "STUB: not implemented"; return }
