package fd

func Gradient(dst []float64, f func([]float64) float64, x []float64, settings *Settings) []float64 {
	_ = "STUB: not implemented"
	return nil
}

type fdrun struct {
	idx    int
	pt     Point
	result float64
}
