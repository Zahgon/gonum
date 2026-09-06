package testlapack

import (
	"math/rand/v2"
	"testing"

	"gonum.org/v1/gonum/lapack"
)

type Dgghrder interface {
	Dgghrd(compq, compz lapack.OrthoComp, n, ilo, ihi int, a []float64, lda int, b []float64, ldb int, q []float64, ldq int, z []float64, ldz int)
}

func DgghrdTest(t *testing.T, impl Dgghrder) { _ = "STUB: not implemented"; return }

func testDgghrd(t *testing.T, impl Dgghrder, rnd *rand.Rand, compq, compz lapack.OrthoComp, n, ilo, ihi, lda, ldb, ldq, ldz int) {
	_ = "STUB: not implemented"
	return
}
