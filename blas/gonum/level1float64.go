package gonum

import (
	"gonum.org/v1/gonum/blas"
)

var _ blas.Float64Level1 = Implementation{}

func (Implementation) Dnrm2(n int, x []float64, incX int) float64 {
	_ = "STUB: not implemented"
	return 0
}

func (Implementation) Dasum(n int, x []float64, incX int) float64 {
	_ = "STUB: not implemented"
	return 0
}

func (Implementation) Idamax(n int, x []float64, incX int) int { _ = "STUB: not implemented"; return 0 }

func (Implementation) Dswap(n int, x []float64, incX int, y []float64, incY int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Dcopy(n int, x []float64, incX int, y []float64, incY int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Daxpy(n int, alpha float64, x []float64, incX int, y []float64, incY int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Drotg(a, b float64) (c, s, r, z float64) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0
}

func (Implementation) Drotmg(d1, d2, x1, y1 float64) (p blas.DrotmParams, rd1, rd2, rx1 float64) {
	_ = "STUB: not implemented"
	return *new(blas.DrotmParams), 0, 0, 0
}

func (Implementation) Drot(n int, x []float64, incX int, y []float64, incY int, c float64, s float64) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Drotm(n int, x []float64, incX int, y []float64, incY int, p blas.DrotmParams) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Dscal(n int, alpha float64, x []float64, incX int) {
	_ = "STUB: not implemented"
	return
}
