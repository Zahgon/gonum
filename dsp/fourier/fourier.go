package fourier

type FFT struct {
	work []float64
	ifac [15]int

	real []float64
}

func NewFFT(n int) *FFT { _ = "STUB: not implemented"; return nil }

func (t *FFT) Len() int { _ = "STUB: not implemented"; return 0 }

func (t *FFT) Reset(n int) { _ = "STUB: not implemented"; return }

func (t *FFT) Coefficients(dst []complex128, seq []float64) []complex128 {
	_ = "STUB: not implemented"
	return nil
}

func (t *FFT) Sequence(dst []float64, coeff []complex128) []float64 {
	_ = "STUB: not implemented"
	return nil
}

func (t *FFT) Freq(i int) float64 { _ = "STUB: not implemented"; return 0 }

type CmplxFFT struct {
	work []float64
	ifac [15]int

	real []float64
}

func NewCmplxFFT(n int) *CmplxFFT { _ = "STUB: not implemented"; return nil }

func (t *CmplxFFT) Len() int { _ = "STUB: not implemented"; return 0 }

func (t *CmplxFFT) Reset(n int) { _ = "STUB: not implemented"; return }

func (t *CmplxFFT) Coefficients(dst, seq []complex128) []complex128 {
	_ = "STUB: not implemented"
	return nil
}

func (t *CmplxFFT) Sequence(dst, coeff []complex128) []complex128 {
	_ = "STUB: not implemented"
	return nil
}

func (t *CmplxFFT) Freq(i int) float64 { _ = "STUB: not implemented"; return 0 }

func (t *CmplxFFT) ShiftIdx(i int) int { _ = "STUB: not implemented"; return 0 }

func (t *CmplxFFT) UnshiftIdx(i int) int { _ = "STUB: not implemented"; return 0 }
