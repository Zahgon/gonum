package stat

import (
	"gonum.org/v1/gonum/mat"
)

type PC struct {
	n, d    int
	weights []float64
	svd     *mat.SVD
	ok      bool
}

func (c *PC) PrincipalComponents(a mat.Matrix, weights []float64) (ok bool) {
	_ = "STUB: not implemented"
	return false
}

func (c *PC) VectorsTo(dst *mat.Dense) { _ = "STUB: not implemented"; return }

func (c *PC) VarsTo(dst []float64) []float64 { _ = "STUB: not implemented"; return nil }

type CC struct {
	n int

	xd, yd int

	x, y, c *mat.SVD
	ok      bool
}

func (c *CC) CanonicalCorrelations(x, y mat.Matrix, weights []float64) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *CC) CorrsTo(dst []float64) []float64 { _ = "STUB: not implemented"; return nil }

func (c *CC) LeftTo(dst *mat.Dense, spheredSpace bool) { _ = "STUB: not implemented"; return }

func (c *CC) RightTo(dst *mat.Dense, spheredSpace bool) { _ = "STUB: not implemented"; return }

func svdFactorizeCentered(work *mat.SVD, m mat.Matrix, weights []float64) (svd *mat.SVD, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func scaleColsReciSqrt(cols *mat.Dense, vals []float64) { _ = "STUB: not implemented"; return }
