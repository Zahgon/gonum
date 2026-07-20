package distuv

import (
	"math/rand/v2"
)

type Gamma struct {
	Alpha float64

	Beta float64

	Src rand.Source
}

func (g Gamma) CDF(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (g Gamma) ExKurtosis() float64 { _ = "STUB: not implemented"; return 0 }

func (g Gamma) LogProb(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (g Gamma) Mean() float64 { _ = "STUB: not implemented"; return 0 }

func (g Gamma) Mode() float64 { _ = "STUB: not implemented"; return 0 }

func (Gamma) NumParameters() int { _ = "STUB: not implemented"; return 0 }

func (g Gamma) Prob(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (g Gamma) Quantile(p float64) float64 { _ = "STUB: not implemented"; return 0 }

func (g Gamma) Rand() float64 { _ = "STUB: not implemented"; return 0 }

func (g Gamma) Survival(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (g Gamma) StdDev() float64 { _ = "STUB: not implemented"; return 0 }

func (g Gamma) Variance() float64 { _ = "STUB: not implemented"; return 0 }
