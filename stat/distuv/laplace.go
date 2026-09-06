package distuv

import (
	"math/rand/v2"
)

type Laplace struct {
	Mu    float64
	Scale float64
	Src   rand.Source
}

func (l Laplace) CDF(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (l Laplace) Entropy() float64 { _ = "STUB: not implemented"; return 0 }

func (l Laplace) ExKurtosis() float64 { _ = "STUB: not implemented"; return 0 }

func (l *Laplace) Fit(samples, weights []float64) { _ = "STUB: not implemented"; return }

func (l Laplace) LogProb(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (l Laplace) parameters(p []Parameter) []Parameter { _ = "STUB: not implemented"; return nil }

func (l Laplace) Mean() float64 { _ = "STUB: not implemented"; return 0 }

func (l Laplace) Median() float64 { _ = "STUB: not implemented"; return 0 }

func (l Laplace) Mode() float64 { _ = "STUB: not implemented"; return 0 }

func (l Laplace) NumParameters() int { _ = "STUB: not implemented"; return 0 }

func (l Laplace) Quantile(p float64) float64 { _ = "STUB: not implemented"; return 0 }

func (l Laplace) Prob(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (l Laplace) Rand() float64 { _ = "STUB: not implemented"; return 0 }

func (l Laplace) Score(deriv []float64, x float64) []float64 { _ = "STUB: not implemented"; return nil }

func (l Laplace) ScoreInput(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (Laplace) Skewness() float64 { _ = "STUB: not implemented"; return 0 }

func (l Laplace) StdDev() float64 { _ = "STUB: not implemented"; return 0 }

func (l Laplace) Survival(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (l *Laplace) setParameters(p []Parameter) { _ = "STUB: not implemented"; return }

func (l Laplace) Variance() float64 { _ = "STUB: not implemented"; return 0 }
