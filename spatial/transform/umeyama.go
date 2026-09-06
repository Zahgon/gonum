package transform

import (
	"gonum.org/v1/gonum/mat"
)

type Umeyama struct {
	x, y     *mat.Dense
	n, m     int
	muX, muY *mat.VecDense
	varX     float64
}

func NewUmeyama(x, y *mat.Dense) *Umeyama { _ = "STUB: not implemented"; return nil }

func (u *Umeyama) Var() float64 { _ = "STUB: not implemented"; return 0 }

func (u *Umeyama) Transform() (c float64, r *mat.Dense, t *mat.VecDense, ok bool) {
	_ = "STUB: not implemented"
	return 0, nil, nil, false
}

func UmeyamaTransform(x, y *mat.Dense) (c float64, r *mat.Dense, t *mat.VecDense, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil, nil
}
