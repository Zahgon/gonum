package gonum

import (
	"gonum.org/v1/gonum/lapack"
)

func (impl Implementation) Dlatdf(job lapack.MaximizeNormXJob, n int, z []float64, ldz int, rhs []float64, rdsum, rdscal float64, ipiv, jpiv []int) (scale, sum float64) {
	_ = "STUB: not implemented"
	return 0, 0
}
