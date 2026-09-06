package fd

import (
	"gonum.org/v1/gonum/mat"
)

func Hessian(dst *mat.SymDense, f func(x []float64) float64, x []float64, settings *Settings) {
	_ = "STUB: not implemented"
	return
}

func hessianSerial(dst *mat.SymDense, f func(x []float64) float64, x []float64, stencil []Point, step float64, originKnown bool, originValue float64) {
	_ = "STUB: not implemented"
	return
}

func hessianConcurrent(dst *mat.SymDense, nWorkers, evals int, f func(x []float64) float64, x []float64, stencil []Point, step float64, originKnown bool, originValue float64) {
	_ = "STUB: not implemented"
	return
}
