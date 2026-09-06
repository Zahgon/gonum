package stat

import (
	"gonum.org/v1/gonum/mat"
)

func CovarianceMatrix(dst *mat.SymDense, x mat.Matrix, weights []float64) {
	_ = "STUB: not implemented"
	return
}

func CorrelationMatrix(dst *mat.SymDense, x mat.Matrix, weights []float64) {
	_ = "STUB: not implemented"
	return
}

func covToCorr(c *mat.SymDense) { _ = "STUB: not implemented"; return }

func corrToCov(c *mat.SymDense, sigma []float64) { _ = "STUB: not implemented"; return }

func Mahalanobis(x, y mat.Vector, chol *mat.Cholesky) float64 { _ = "STUB: not implemented"; return 0 }
