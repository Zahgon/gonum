package testlapack

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dbdsqrer interface {
	Dbdsqr(uplo blas.Uplo, n, ncvt, nru, ncc int, d, e, vt []float64, ldvt int, u []float64, ldu int, c []float64, ldc int, work []float64) (ok bool)
}

func DbdsqrTest(t *testing.T, impl Dbdsqrer) { _ = "STUB: not implemented"; return }
