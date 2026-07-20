package distuv

import (
	"math/rand/v2"
)

const logPi = 1.1447298858494001741

type StudentsT struct {
	Mu float64

	Sigma float64

	Nu float64

	Src rand.Source
}

func (s StudentsT) CDF(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (s StudentsT) LogProb(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (s StudentsT) Mean() float64 { _ = "STUB: not implemented"; return 0 }

func (s StudentsT) Mode() float64 { _ = "STUB: not implemented"; return 0 }

func (StudentsT) NumParameters() int { _ = "STUB: not implemented"; return 0 }

func (s StudentsT) Prob(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (s StudentsT) Quantile(p float64) float64 { _ = "STUB: not implemented"; return 0 }

func (s StudentsT) Rand() float64 { _ = "STUB: not implemented"; return 0 }

func (s StudentsT) StdDev() float64 { _ = "STUB: not implemented"; return 0 }

func (s StudentsT) Survival(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (s StudentsT) Variance() float64 { _ = "STUB: not implemented"; return 0 }
