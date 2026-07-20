package distuv

import (
	"math/rand/v2"
)

type Weibull struct {
	K float64

	Lambda float64

	Src rand.Source
}

func (w Weibull) CDF(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (w Weibull) Entropy() float64 { _ = "STUB: not implemented"; return 0 }

func (w Weibull) ExKurtosis() float64 { _ = "STUB: not implemented"; return 0 }

func (w Weibull) gammaIPow(i, pow float64) float64 { _ = "STUB: not implemented"; return 0 }

func (w Weibull) LogProb(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (w Weibull) LogSurvival(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (w Weibull) Mean() float64 { _ = "STUB: not implemented"; return 0 }

func (w Weibull) Median() float64 { _ = "STUB: not implemented"; return 0 }

func (w Weibull) Mode() float64 { _ = "STUB: not implemented"; return 0 }

func (Weibull) NumParameters() int { _ = "STUB: not implemented"; return 0 }

func (w Weibull) Prob(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (w Weibull) Quantile(p float64) float64 { _ = "STUB: not implemented"; return 0 }

func (w Weibull) Rand() float64 { _ = "STUB: not implemented"; return 0 }

func (w Weibull) Score(deriv []float64, x float64) []float64 { _ = "STUB: not implemented"; return nil }

func (w Weibull) ScoreInput(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (w Weibull) Skewness() float64 { _ = "STUB: not implemented"; return 0 }

func (w Weibull) StdDev() float64 { _ = "STUB: not implemented"; return 0 }

func (w Weibull) Survival(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (w *Weibull) setParameters(p []Parameter) { _ = "STUB: not implemented"; return }

func (w Weibull) Variance() float64 { _ = "STUB: not implemented"; return 0 }

func (w Weibull) parameters(p []Parameter) []Parameter { _ = "STUB: not implemented"; return nil }
