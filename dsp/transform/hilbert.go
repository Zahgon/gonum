package transform

import (
	"gonum.org/v1/gonum/dsp/fourier"
)

type Hilbert struct {
	fft  *fourier.CmplxFFT
	work []complex128
}

func NewHilbert(n int) *Hilbert { _ = "STUB: not implemented"; return nil }

func (h *Hilbert) Len() int { _ = "STUB: not implemented"; return 0 }

func (h *Hilbert) AnalyticSignal(dst []complex128, signal []float64) []complex128 {
	_ = "STUB: not implemented"
	return nil
}
