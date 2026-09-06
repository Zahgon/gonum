package mds

import (
	"gonum.org/v1/gonum/blas/blas64"
	"gonum.org/v1/gonum/mat"
)

func TorgersonScaling(dst *mat.Dense, eigdst []float64, dis mat.Symmetric) (k int, eig []float64) {
	_ = "STUB: not implemented"
	return 0, nil
}

func reverse(values []float64, vectors blas64.General) { _ = "STUB: not implemented"; return }
