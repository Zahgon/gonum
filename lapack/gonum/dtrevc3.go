package gonum

import (
	"gonum.org/v1/gonum/lapack"
)

func (impl Implementation) Dtrevc3(side lapack.EVSide, howmny lapack.EVHowMany, selected []bool, n int, t []float64, ldt int, vl []float64, ldvl int, vr []float64, ldvr int, mm int, work []float64, lwork int) (m int) {
	_ = "STUB: not implemented"
	return 0
}
