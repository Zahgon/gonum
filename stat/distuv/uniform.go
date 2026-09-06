package distuv

import (
	"math/rand/v2"
)

var UnitUniform = Uniform{Min: 0, Max: 1}

type Uniform struct {
	Min float64
	Max float64
	Src rand.Source
}

func (u Uniform) CDF(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (u Uniform) Entropy() float64 { _ = "STUB: not implemented"; return 0 }

func (Uniform) ExKurtosis() float64 { _ = "STUB: not implemented"; return 0 }

func (u Uniform) LogProb(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (u Uniform) parameters(p []Parameter) []Parameter { _ = "STUB: not implemented"; return nil }

func (u Uniform) Mean() float64 { _ = "STUB: not implemented"; return 0 }

func (u Uniform) Median() float64 { _ = "STUB: not implemented"; return 0 }

func (Uniform) NumParameters() int { _ = "STUB: not implemented"; return 0 }

func (u Uniform) Prob(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (u Uniform) Quantile(p float64) float64 { _ = "STUB: not implemented"; return 0 }

func (u Uniform) Rand() float64 { _ = "STUB: not implemented"; return 0 }

func (u Uniform) Score(deriv []float64, x float64) []float64 { _ = "STUB: not implemented"; return nil }

func (u Uniform) ScoreInput(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (Uniform) Skewness() float64 { _ = "STUB: not implemented"; return 0 }

func (u Uniform) StdDev() float64 { _ = "STUB: not implemented"; return 0 }

func (u Uniform) Survival(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (u *Uniform) setParameters(p []Parameter) { _ = "STUB: not implemented"; return }

func (u Uniform) Variance() float64 { _ = "STUB: not implemented"; return 0 }
