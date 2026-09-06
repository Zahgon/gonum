package distmv

type Quantiler interface {
	Quantile(x, p []float64) []float64
}

type LogProber interface {
	LogProb(x []float64) float64
}

type Rander interface {
	Rand(x []float64) []float64
}

type RandLogProber interface {
	Rander
	LogProber
}
