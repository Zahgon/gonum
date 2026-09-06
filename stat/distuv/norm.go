package distuv

import (
	"math/rand/v2"
)

var UnitNormal = Normal{Mu: 0, Sigma: 1}

type Normal struct {
	Mu    float64
	Sigma float64
	Src   rand.Source
}

func (n Normal) CDF(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (n *Normal) ConjugateUpdate(suffStat []float64, nSamples float64, priorStrength []float64) {
	_ = "STUB: not implemented"
	return
}

func (n Normal) Entropy() float64 { _ = "STUB: not implemented"; return 0 }

func (Normal) ExKurtosis() float64 { _ = "STUB: not implemented"; return 0 }

func (n *Normal) Fit(samples, weights []float64) { _ = "STUB: not implemented"; return }

func (n Normal) LogProb(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (n Normal) Mean() float64 { _ = "STUB: not implemented"; return 0 }

func (n Normal) Median() float64 { _ = "STUB: not implemented"; return 0 }

func (n Normal) Mode() float64 { _ = "STUB: not implemented"; return 0 }

func (Normal) NumParameters() int { _ = "STUB: not implemented"; return 0 }

func (Normal) NumSuffStat() int { _ = "STUB: not implemented"; return 0 }

func (n Normal) Prob(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (n Normal) Quantile(p float64) float64 { _ = "STUB: not implemented"; return 0 }

func (n Normal) Rand() float64 { _ = "STUB: not implemented"; return 0 }

func (n Normal) Score(deriv []float64, x float64) []float64 { _ = "STUB: not implemented"; return nil }

func (n Normal) ScoreInput(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (Normal) Skewness() float64 { _ = "STUB: not implemented"; return 0 }

func (n Normal) StdDev() float64 { _ = "STUB: not implemented"; return 0 }

func (Normal) SuffStat(suffStat, samples, weights []float64) (nSamples float64) {
	_ = "STUB: not implemented"
	return 0
}

func (n Normal) Survival(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (n *Normal) setParameters(p []Parameter) { _ = "STUB: not implemented"; return }

func (n Normal) Variance() float64 { _ = "STUB: not implemented"; return 0 }

func (n Normal) parameters(p []Parameter) []Parameter { _ = "STUB: not implemented"; return nil }
