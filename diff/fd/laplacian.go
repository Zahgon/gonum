package fd

func Laplacian(f func(x []float64) float64, x []float64, settings *Settings) float64 {
	_ = "STUB: not implemented"
	return 0
}

func laplacianSerial(f func(x []float64) float64, x []float64, stencil []Point, step float64, originKnown bool, originValue float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func laplacianConcurrent(nWorkers, evals int, f func(x []float64) float64, x []float64, stencil []Point, step float64, originKnown bool, originValue float64) float64 {
	_ = "STUB: not implemented"
	return 0
}
