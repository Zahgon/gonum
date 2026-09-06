package gonum

import (
	"gonum.org/v1/gonum/blas"
)

var _ blas.Float32Level1 = Implementation{}

func (Implementation) Snrm2(n int, x []float32, incX int) float32 {
	_ = "STUB: not implemented"
	return 0
}

func (Implementation) Sasum(n int, x []float32, incX int) float32 {
	_ = "STUB: not implemented"
	return 0
}

func (Implementation) Isamax(n int, x []float32, incX int) int { _ = "STUB: not implemented"; return 0 }

func (Implementation) Sswap(n int, x []float32, incX int, y []float32, incY int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Scopy(n int, x []float32, incX int, y []float32, incY int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Saxpy(n int, alpha float32, x []float32, incX int, y []float32, incY int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Srotg(a, b float32) (c, s, r, z float32) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0
}

func (Implementation) Srotmg(d1, d2, x1, y1 float32) (p blas.SrotmParams, rd1, rd2, rx1 float32) {
	_ = "STUB: not implemented"
	return *new(blas.SrotmParams), 0, 0, 0
}

func (Implementation) Srot(n int, x []float32, incX int, y []float32, incY int, c float32, s float32) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Srotm(n int, x []float32, incX int, y []float32, incY int, p blas.SrotmParams) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Sscal(n int, alpha float32, x []float32, incX int) {
	_ = "STUB: not implemented"
	return
}
