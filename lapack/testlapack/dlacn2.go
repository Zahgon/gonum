package testlapack

import (
	"testing"
)

type Dlacn2er interface {
	Dlacn2(n int, v, x []float64, isgn []int, est float64, kase int, isave *[3]int) (float64, int)
}

func Dlacn2Test(t *testing.T, impl Dlacn2er) { _ = "STUB: not implemented"; return }
