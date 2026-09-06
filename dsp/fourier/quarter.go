package fourier

type QuarterWaveFFT struct {
	work []float64
	ifac [15]int
}

func NewQuarterWaveFFT(n int) *QuarterWaveFFT { _ = "STUB: not implemented"; return nil }

func (t *QuarterWaveFFT) Len() int { _ = "STUB: not implemented"; return 0 }

func (t *QuarterWaveFFT) Reset(n int) { _ = "STUB: not implemented"; return }

func (t *QuarterWaveFFT) CosCoefficients(dst, seq []float64) []float64 {
	_ = "STUB: not implemented"
	return nil
}

func (t *QuarterWaveFFT) CosSequence(dst, coeff []float64) []float64 {
	_ = "STUB: not implemented"
	return nil
}

func (t *QuarterWaveFFT) SinCoefficients(dst, seq []float64) []float64 {
	_ = "STUB: not implemented"
	return nil
}

func (t *QuarterWaveFFT) SinSequence(dst, coeff []float64) []float64 {
	_ = "STUB: not implemented"
	return nil
}
