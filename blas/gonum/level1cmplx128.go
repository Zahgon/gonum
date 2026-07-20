package gonum

import (
	"gonum.org/v1/gonum/blas"
)

var _ blas.Complex128Level1 = Implementation{}

func (Implementation) Dzasum(n int, x []complex128, incX int) float64 {
	_ = "STUB: not implemented"
	return 0
}

func (Implementation) Dznrm2(n int, x []complex128, incX int) float64 {
	_ = "STUB: not implemented"
	return 0
}

func (Implementation) Izamax(n int, x []complex128, incX int) int {
	_ = "STUB: not implemented"
	return 0
}

func (Implementation) Zaxpy(n int, alpha complex128, x []complex128, incX int, y []complex128, incY int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Zcopy(n int, x []complex128, incX int, y []complex128, incY int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Zdotc(n int, x []complex128, incX int, y []complex128, incY int) complex128 {
	_ = "STUB: not implemented"
	return 0
}

func (Implementation) Zdotu(n int, x []complex128, incX int, y []complex128, incY int) complex128 {
	_ = "STUB: not implemented"
	return 0
}

func (Implementation) Zdscal(n int, alpha float64, x []complex128, incX int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Zscal(n int, alpha complex128, x []complex128, incX int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Zswap(n int, x []complex128, incX int, y []complex128, incY int) {
	_ = "STUB: not implemented"
	return
}
