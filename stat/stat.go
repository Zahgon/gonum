package stat

import (
	"gonum.org/v1/gonum/mat"
)

type CumulantKind int

const (
	Empirical CumulantKind = 1

	LinInterp CumulantKind = 4

	Linear CumulantKind = 7
)

func WassersteinDistance(p, q, pWeights, qWeights []float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func WassersteinDistanceND(p, q mat.Matrix, pWeights, qWeights []float64, tol float64) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func normalizeWeights(weights []float64) { _ = "STUB: not implemented"; return }

func solveOptimalTransportLP(costMatrix *mat.Dense, supply, demand []float64, tol float64) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func bhattacharyyaCoeff(p, q []float64) float64 { _ = "STUB: not implemented"; return 0 }

func Bhattacharyya(p, q []float64) float64 { _ = "STUB: not implemented"; return 0 }

func CDF(q float64, c CumulantKind, x, weights []float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func ChiSquare(obs, exp []float64) float64 { _ = "STUB: not implemented"; return 0 }

func CircularMean(x, weights []float64) float64 { _ = "STUB: not implemented"; return 0 }

func Correlation(x, y, weights []float64) float64 { _ = "STUB: not implemented"; return 0 }

func Kendall(x, y, weights []float64) float64 { _ = "STUB: not implemented"; return 0 }

func Covariance(x, y, weights []float64) float64 { _ = "STUB: not implemented"; return 0 }

func covarianceMeans(x, y, weights []float64, xu, yu float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func CrossEntropy(p, q []float64) float64 { _ = "STUB: not implemented"; return 0 }

func Entropy(p []float64) float64 { _ = "STUB: not implemented"; return 0 }

func ExKurtosis(x, weights []float64) float64 { _ = "STUB: not implemented"; return 0 }

func kurtosisCorrection(n float64) (mul, offset float64) { _ = "STUB: not implemented"; return 0, 0 }

func GeometricMean(x, weights []float64) float64 { _ = "STUB: not implemented"; return 0 }

func HarmonicMean(x, weights []float64) float64 { _ = "STUB: not implemented"; return 0 }

func Hellinger(p, q []float64) float64 { _ = "STUB: not implemented"; return 0 }

func Histogram(count, dividers, x, weights []float64) []float64 {
	_ = "STUB: not implemented"
	return nil
}

func JensenShannon(p, q []float64) float64 { _ = "STUB: not implemented"; return 0 }

func KolmogorovSmirnov(x, xWeights, y, yWeights []float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func updateKS(idx int, cdf, sum float64, values, weights []float64, isNil bool) (val, newCdf float64, newIdx int) {
	_ = "STUB: not implemented"
	return 0, 0, 0
}

func KullbackLeibler(p, q []float64) float64 { _ = "STUB: not implemented"; return 0 }

func LinearRegression(x, y, weights []float64, origin bool) (alpha, beta float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func RSquared(x, y, weights []float64, alpha, beta float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func RSquaredFrom(estimates, values, weights []float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func RNoughtSquared(x, y, weights []float64, beta float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func Mean(x, weights []float64) float64 { _ = "STUB: not implemented"; return 0 }

func Mode(x, weights []float64) (val float64, count float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func BivariateMoment(r, s float64, x, y, weights []float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func Moment(moment float64, x, weights []float64) float64 { _ = "STUB: not implemented"; return 0 }

func MomentAbout(moment float64, x []float64, mean float64, weights []float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func Quantile(p float64, c CumulantKind, x, weights []float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func empiricalQuantile(p float64, x, weights []float64, sumWeights float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func linInterpQuantile(p float64, x, weights []float64, sumWeights float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func type7Quantile(p float64, x, weights []float64, sumWeights float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func Skew(x, weights []float64) float64 { _ = "STUB: not implemented"; return 0 }

func skewCorrection(n float64) float64 { _ = "STUB: not implemented"; return 0 }

func SortWeighted(x, weights []float64) { _ = "STUB: not implemented"; return }

type weightSorter struct {
	x []float64
	w []float64
}

func (w weightSorter) Len() int           { _ = "STUB: not implemented"; return 0 }
func (w weightSorter) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
func (w weightSorter) Swap(i, j int)      { _ = "STUB: not implemented"; return }

func SortWeightedLabeled(x []float64, labels []bool, weights []float64) {
	_ = "STUB: not implemented"
	return
}

type labelSorter struct {
	x []float64
	l []bool
}

func (a labelSorter) Len() int           { _ = "STUB: not implemented"; return 0 }
func (a labelSorter) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
func (a labelSorter) Swap(i, j int)      { _ = "STUB: not implemented"; return }

type weightLabelSorter struct {
	x []float64
	l []bool
	w []float64
}

func (a weightLabelSorter) Len() int           { _ = "STUB: not implemented"; return 0 }
func (a weightLabelSorter) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
func (a weightLabelSorter) Swap(i, j int)      { _ = "STUB: not implemented"; return }

func StdDev(x, weights []float64) float64 { _ = "STUB: not implemented"; return 0 }

func MeanStdDev(x, weights []float64) (mean, std float64) { _ = "STUB: not implemented"; return 0, 0 }

func StdErr(std, sampleSize float64) float64 { _ = "STUB: not implemented"; return 0 }

func StdScore(x, mean, std float64) float64 { _ = "STUB: not implemented"; return 0 }

func Variance(x, weights []float64) float64 { _ = "STUB: not implemented"; return 0 }

func MeanVariance(x, weights []float64) (mean, variance float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func PopMeanVariance(x, weights []float64) (mean, variance float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func PopMeanStdDev(x, weights []float64) (mean, std float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func PopStdDev(x, weights []float64) float64 { _ = "STUB: not implemented"; return 0 }

func PopVariance(x, weights []float64) float64 { _ = "STUB: not implemented"; return 0 }

func meanUnnormalisedVarianceSumWeights(x, weights []float64) (mean, unnormalisedVariance, sumWeights float64) {
	_ = "STUB: not implemented"
	return 0, 0, 0
}
