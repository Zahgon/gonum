package testlapack

import (
	"testing"
)

type Dlapller interface {
	Dgesvder
	Dlapll(n int, x []float64, incX int, y []float64, incY int) float64
}

func DlapllTest(t *testing.T, impl Dlapller) { _ = "STUB: not implemented"; return }
