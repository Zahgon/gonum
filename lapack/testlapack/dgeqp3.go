package testlapack

import (
	"math/rand/v2"
	"testing"
)

type Dgeqp3er interface {
	Dlapmter
	Dgeqp3(m, n int, a []float64, lda int, jpvt []int, tau, work []float64, lwork int)
}

func Dgeqp3Test(t *testing.T, impl Dgeqp3er) { _ = "STUB: not implemented"; return }

func dgeqp3Test(t *testing.T, impl Dgeqp3er, rnd *rand.Rand, m, n, lda int) {
	_ = "STUB: not implemented"
	return
}
