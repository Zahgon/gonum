package distuv

type LogProber interface {
	LogProb(x float64) float64
}

type Rander interface {
	Rand() float64
}

type RandLogProber interface {
	Rander
	LogProber
}

type Quantiler interface {
	Quantile(p float64) float64
}
