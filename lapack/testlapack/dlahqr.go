package testlapack

import (
	"testing"

	"gonum.org/v1/gonum/blas/blas64"
)

type Dlahqrer interface {
	Dlahqr(wantt, wantz bool, n, ilo, ihi int, h []float64, ldh int, wr, wi []float64, iloz, ihiz int, z []float64, ldz int) int
}

type dlahqrTest struct {
	h            blas64.General
	ilo, ihi     int
	iloz, ihiz   int
	wantt, wantz bool

	evWant []complex128
}

func DlahqrTest(t *testing.T, impl Dlahqrer) { _ = "STUB: not implemented"; return }

func testDlahqr(t *testing.T, impl Dlahqrer, test dlahqrTest) { _ = "STUB: not implemented"; return }
