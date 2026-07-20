package network

import (
	"gonum.org/v1/gonum/graph/spectral"
)

func Diffuse(dst, h map[int64]float64, by spectral.Laplacian, t float64) map[int64]float64 {
	_ = "STUB: not implemented"
	return nil
}

func DiffuseToEquilibrium(dst, h map[int64]float64, by spectral.Laplacian, tol float64, iters int) (eq map[int64]float64, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}
