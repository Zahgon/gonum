package mat

import (
	"gonum.org/v1/gonum/blas/blas64"
)

const badRcond = "mat: invalid rcond value"

type SVD struct {
	kind SVDKind

	s  []float64
	u  blas64.General
	vt blas64.General
}

type SVDKind int

const (
	SVDNone SVDKind = 0

	SVDThinU SVDKind = 1 << (iota - 1)

	SVDFullU

	SVDThinV

	SVDFullV

	SVDThin SVDKind = SVDThinU | SVDThinV

	SVDFull SVDKind = SVDFullU | SVDFullV
)

func (svd *SVD) succFact() bool { _ = "STUB: not implemented"; return false }

func (svd *SVD) Factorize(a Matrix, kind SVDKind) (ok bool) {
	_ = "STUB: not implemented"
	return false
}

func (svd *SVD) Kind() SVDKind { _ = "STUB: not implemented"; return *new(SVDKind) }

func (svd *SVD) Rank(rcond float64) int { _ = "STUB: not implemented"; return 0 }

func (svd *SVD) Cond() float64 { _ = "STUB: not implemented"; return 0 }

func (svd *SVD) Values(s []float64) []float64 { _ = "STUB: not implemented"; return nil }

func (svd *SVD) UTo(dst *Dense) { _ = "STUB: not implemented"; return }

func (svd *SVD) VTo(dst *Dense) { _ = "STUB: not implemented"; return }

func (svd *SVD) SolveTo(dst *Dense, b Matrix, rank int) []float64 {
	_ = "STUB: not implemented"
	return nil
}

type repVector struct {
	vec  []float64
	cols int
}

func (m repVector) Dims() (r, c int)    { _ = "STUB: not implemented"; return 0, 0 }
func (m repVector) At(i, j int) float64 { _ = "STUB: not implemented"; return 0 }

func (m repVector) T() Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

func (svd *SVD) SolveVecTo(dst *VecDense, b Vector, rank int) float64 {
	_ = "STUB: not implemented"
	return 0
}
