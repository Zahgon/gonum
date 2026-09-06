package testblas

import (
	"testing"
)

type Dgerer interface {
	Dger(m, n int, alpha float64, x []float64, incX int, y []float64, incY int, a []float64, lda int)
}

func DgerTest(t *testing.T, blasser Dgerer) { _ = "STUB: not implemented"; return }

func dgercomp(t *testing.T, x, xCopy, y, yCopy []float64, ans [][]float64, trueAns [][]float64, name string) {
	_ = "STUB: not implemented"
	return
}
