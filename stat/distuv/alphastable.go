package distuv

import (
	"math/rand/v2"
)

type AlphaStable struct {
	Alpha float64

	Beta float64

	C float64

	Mu  float64
	Src rand.Source
}

func (a AlphaStable) ExKurtosis() float64 { _ = "STUB: not implemented"; return 0 }

func (a AlphaStable) Mean() float64 { _ = "STUB: not implemented"; return 0 }

func (a AlphaStable) Median() float64 { _ = "STUB: not implemented"; return 0 }

func (a AlphaStable) Mode() float64 { _ = "STUB: not implemented"; return 0 }

func (a AlphaStable) NumParameters() int { _ = "STUB: not implemented"; return 0 }

func (a AlphaStable) Rand() float64 { _ = "STUB: not implemented"; return 0 }

func (a AlphaStable) Skewness() float64 { _ = "STUB: not implemented"; return 0 }

func (a AlphaStable) StdDev() float64 { _ = "STUB: not implemented"; return 0 }

func (a AlphaStable) Variance() float64 { _ = "STUB: not implemented"; return 0 }
