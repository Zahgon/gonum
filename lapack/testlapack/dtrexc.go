package testlapack

import (
	"math/rand/v2"
	"testing"

	"gonum.org/v1/gonum/lapack"
)

type Dtrexcer interface {
	Dtrexc(compq lapack.UpdateSchurComp, n int, t []float64, ldt int, q []float64, ldq int, ifst, ilst int, work []float64) (ifstOut, ilstOut int, ok bool)
}

func DtrexcTest(t *testing.T, impl Dtrexcer) { _ = "STUB: not implemented"; return }

func dtrexcTest(t *testing.T, impl Dtrexcer, rnd *rand.Rand, n, ifst, ilst, extra int) {
	_ = "STUB: not implemented"
	return
}
