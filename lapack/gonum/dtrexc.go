package gonum

import "gonum.org/v1/gonum/lapack"

func (impl Implementation) Dtrexc(compq lapack.UpdateSchurComp, n int, t []float64, ldt int, q []float64, ldq int, ifst, ilst int, work []float64) (ifstOut, ilstOut int, ok bool) {
	_ = "STUB: not implemented"
	return 0, 0, false
}
