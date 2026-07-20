//go:build !safe
// +build !safe

package r3

import (
	"gonum.org/v1/gonum/blas/blas64"
	"gonum.org/v1/gonum/mat"
)

type array [3][3]float64

func (m *Mat) At(i, j int) float64 { _ = "STUB: not implemented"; return 0 }

func (m *Mat) Set(i, j int, v float64) { _ = "STUB: not implemented"; return }

func Eye() *Mat { _ = "STUB: not implemented"; return nil }

func Skew(v Vec) (M *Mat) { _ = "STUB: not implemented"; return nil }

func (m *Mat) Mul(a, b mat.Matrix) { _ = "STUB: not implemented"; return }

func (m *Mat) RawMatrix() blas64.General { _ = "STUB: not implemented"; return *new(blas64.General) }

func (r Rotation) Mat() *Mat { _ = "STUB: not implemented"; return nil }

func arrayFrom(vals []float64) *array { _ = "STUB: not implemented"; return nil }

func (m *Mat) slice() []float64 { _ = "STUB: not implemented"; return nil }
