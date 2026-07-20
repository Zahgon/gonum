package fd

import (
	"gonum.org/v1/gonum/mat"
)

type JacobianSettings struct {
	Formula     Formula
	OriginValue []float64
	Step        float64
	Concurrent  bool
}

func Jacobian(dst *mat.Dense, f func(y, x []float64), x []float64, settings *JacobianSettings) {
	_ = "STUB: not implemented"
	return
}

func jacobianSerial(dst *mat.Dense, f func([]float64, []float64), x, origin []float64, formula Formula, step float64) {
	_ = "STUB: not implemented"
	return
}

func jacobianConcurrent(dst *mat.Dense, f func([]float64, []float64), x, origin []float64, formula Formula, step float64, nWorkers int) {
	_ = "STUB: not implemented"
	return
}

type jacJob struct {
	j  int
	pt Point
}
