package distuv

import (
	"math/rand/v2"
)

type Exponential struct {
	Rate float64
	Src  rand.Source
}

func (e Exponential) CDF(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (e *Exponential) ConjugateUpdate(suffStat []float64, nSamples float64, priorStrength []float64) {
	_ = "STUB: not implemented"
	return
}

func (e Exponential) Entropy() float64 { _ = "STUB: not implemented"; return 0 }

func (Exponential) ExKurtosis() float64 { _ = "STUB: not implemented"; return 0 }

func (e *Exponential) Fit(samples, weights []float64) { _ = "STUB: not implemented"; return }

func (e Exponential) LogProb(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (e Exponential) Mean() float64 { _ = "STUB: not implemented"; return 0 }

func (e Exponential) Median() float64 { _ = "STUB: not implemented"; return 0 }

func (Exponential) Mode() float64 { _ = "STUB: not implemented"; return 0 }

func (Exponential) NumParameters() int { _ = "STUB: not implemented"; return 0 }

func (Exponential) NumSuffStat() int { _ = "STUB: not implemented"; return 0 }

func (e Exponential) Prob(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (e Exponential) Quantile(p float64) float64 { _ = "STUB: not implemented"; return 0 }

func (e Exponential) Rand() float64 { _ = "STUB: not implemented"; return 0 }

func (e Exponential) Score(deriv []float64, x float64) []float64 {
	_ = "STUB: not implemented"
	return nil
}

func (e Exponential) ScoreInput(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (Exponential) Skewness() float64 { _ = "STUB: not implemented"; return 0 }

func (e Exponential) StdDev() float64 { _ = "STUB: not implemented"; return 0 }

func (Exponential) SuffStat(suffStat, samples, weights []float64) (nSamples float64) {
	_ = "STUB: not implemented"
	return 0
}

func (e Exponential) Survival(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (e *Exponential) setParameters(p []Parameter) { _ = "STUB: not implemented"; return }

func (e Exponential) Variance() float64 { _ = "STUB: not implemented"; return 0 }

func (e Exponential) parameters(p []Parameter) []Parameter { _ = "STUB: not implemented"; return nil }
