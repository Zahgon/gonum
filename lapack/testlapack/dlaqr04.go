package testlapack

import (
	"testing"

	"gonum.org/v1/gonum/blas/blas64"
)

type Dlaqr04er interface {
	Dlaqr04(wantt, wantz bool, n, ilo, ihi int, h []float64, ldh int, wr, wi []float64, iloz, ihiz int, z []float64, ldz int, work []float64, lwork int, recur int) int

	Dlahqrer
}

type dlaqr04Test struct {
	h            blas64.General
	ilo, ihi     int
	iloz, ihiz   int
	wantt, wantz bool

	evWant []complex128
}

func Dlaqr04Test(t *testing.T, impl Dlaqr04er) { _ = "STUB: not implemented"; return }

func testDlaqr04(t *testing.T, impl Dlaqr04er, test dlaqr04Test, optwork bool, recur int) {
	_ = "STUB: not implemented"
	return
}
