package fd

func CrossLaplacian(f func(x, y []float64) float64, x, y []float64, settings *Settings) float64 {
	_ = "STUB: not implemented"
	return 0
}

func crossLaplacianSerial(f func(x, y []float64) float64, x, y []float64, stencil []Point, step float64, originKnown bool, originValue float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func crossLaplacianConcurrent(nWorkers, evals int, f func(x, y []float64) float64, x, y []float64, stencil []Point, step float64, originKnown bool, originValue float64) float64 {
	_ = "STUB: not implemented"
	return 0
}
