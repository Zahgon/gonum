package distuv

import (
	"math/rand/v2"
)

type InverseGamma struct {
	Alpha float64

	Beta float64

	Src rand.Source
}

func (g InverseGamma) CDF(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (g InverseGamma) ExKurtosis() float64 { _ = "STUB: not implemented"; return 0 }

func (g InverseGamma) LogProb(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (g InverseGamma) Mean() float64 { _ = "STUB: not implemented"; return 0 }

func (g InverseGamma) Mode() float64 { _ = "STUB: not implemented"; return 0 }

func (InverseGamma) NumParameters() int { _ = "STUB: not implemented"; return 0 }

func (g InverseGamma) Prob(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (g InverseGamma) Quantile(p float64) float64 { _ = "STUB: not implemented"; return 0 }

func (g InverseGamma) Rand() float64 { _ = "STUB: not implemented"; return 0 }

func (g InverseGamma) Survival(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (g InverseGamma) StdDev() float64 { _ = "STUB: not implemented"; return 0 }

func (g InverseGamma) Variance() float64 { _ = "STUB: not implemented"; return 0 }
