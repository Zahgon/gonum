package distuv

import (
	"math/rand/v2"
)

type Triangle struct {
	a, b, c float64
	src     rand.Source
}

func NewTriangle(a, b, c float64, src rand.Source) Triangle {
	_ = "STUB: not implemented"
	return *new(Triangle)
}

func checkTriangleParameters(a, b, c float64) { _ = "STUB: not implemented"; return }

func (t Triangle) CDF(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (t Triangle) Entropy() float64 { _ = "STUB: not implemented"; return 0 }

func (Triangle) ExKurtosis() float64 { _ = "STUB: not implemented"; return 0 }

func (t Triangle) LogProb(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (t Triangle) Mean() float64 { _ = "STUB: not implemented"; return 0 }

func (t Triangle) Median() float64 { _ = "STUB: not implemented"; return 0 }

func (t Triangle) Mode() float64 { _ = "STUB: not implemented"; return 0 }

func (Triangle) NumParameters() int { _ = "STUB: not implemented"; return 0 }

func (t Triangle) Prob(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (t Triangle) Quantile(p float64) float64 { _ = "STUB: not implemented"; return 0 }

func (t Triangle) Rand() float64 { _ = "STUB: not implemented"; return 0 }

func (t Triangle) Score(deriv []float64, x float64) []float64 {
	_ = "STUB: not implemented"
	return nil
}

func (t Triangle) ScoreInput(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (t Triangle) Skewness() float64 { _ = "STUB: not implemented"; return 0 }

func (t Triangle) StdDev() float64 { _ = "STUB: not implemented"; return 0 }

func (t Triangle) Survival(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (t Triangle) parameters(p []Parameter) []Parameter { _ = "STUB: not implemented"; return nil }

func (t *Triangle) setParameters(p []Parameter) { _ = "STUB: not implemented"; return }

func (t Triangle) Variance() float64 { _ = "STUB: not implemented"; return 0 }
