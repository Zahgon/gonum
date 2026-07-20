package gonum

import (
	"gonum.org/v1/gonum/blas"
)

var _ blas.Complex64Level1 = Implementation{}

func (Implementation) Scasum(n int, x []complex64, incX int) float32 {
	_ = "STUB: not implemented"
	return 0
}

func (Implementation) Scnrm2(n int, x []complex64, incX int) float32 {
	_ = "STUB: not implemented"
	return 0
}

func (Implementation) Icamax(n int, x []complex64, incX int) int {
	_ = "STUB: not implemented"
	return 0
}

func (Implementation) Caxpy(n int, alpha complex64, x []complex64, incX int, y []complex64, incY int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Ccopy(n int, x []complex64, incX int, y []complex64, incY int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Cdotc(n int, x []complex64, incX int, y []complex64, incY int) complex64 {
	_ = "STUB: not implemented"
	return 0
}

func (Implementation) Cdotu(n int, x []complex64, incX int, y []complex64, incY int) complex64 {
	_ = "STUB: not implemented"
	return 0
}

func (Implementation) Csscal(n int, alpha float32, x []complex64, incX int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Cscal(n int, alpha complex64, x []complex64, incX int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Cswap(n int, x []complex64, incX int, y []complex64, incY int) {
	_ = "STUB: not implemented"
	return
}
