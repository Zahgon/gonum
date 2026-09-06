package testblas

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dgemmer interface {
	Dgemm(tA, tB blas.Transpose, m, n, k int, alpha float64, a []float64, lda int, b []float64, ldb int, beta float64, c []float64, ldc int)
}

type DgemmCase struct {
	m, n, k     int
	alpha, beta float64
	a           [][]float64
	b           [][]float64
	c           [][]float64
	ans         [][]float64
}

var DgemmCases = []DgemmCase{

	{
		m:     4,
		n:     3,
		k:     2,
		alpha: 2,
		beta:  0.5,
		a: [][]float64{
			{1, 2},
			{4, 5},
			{7, 8},
			{10, 11},
		},
		b: [][]float64{
			{1, 5, 6},
			{5, -8, 8},
		},
		c: [][]float64{
			{4, 8, -9},
			{12, 16, -8},
			{1, 5, 15},
			{-3, -4, 7},
		},
		ans: [][]float64{
			{24, -18, 39.5},
			{64, -32, 124},
			{94.5, -55.5, 219.5},
			{128.5, -78, 299.5},
		},
	},
	{
		m:     4,
		n:     2,
		k:     3,
		alpha: 2,
		beta:  0.5,
		a: [][]float64{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
			{10, 11, 12},
		},
		b: [][]float64{
			{1, 5},
			{5, -8},
			{6, 2},
		},
		c: [][]float64{
			{4, 8},
			{12, 16},
			{1, 5},
			{-3, -4},
		},
		ans: [][]float64{
			{60, -6},
			{136, -8},
			{202.5, -19.5},
			{272.5, -30},
		},
	},
	{
		m:     3,
		n:     2,
		k:     4,
		alpha: 2,
		beta:  0.5,
		a: [][]float64{
			{1, 2, 3, 4},
			{4, 5, 6, 7},
			{8, 9, 10, 11},
		},
		b: [][]float64{
			{1, 5},
			{5, -8},
			{6, 2},
			{8, 10},
		},
		c: [][]float64{
			{4, 8},
			{12, 16},
			{9, -10},
		},
		ans: [][]float64{
			{124, 74},
			{248, 132},
			{406.5, 191},
		},
	},
	{
		m:     3,
		n:     4,
		k:     2,
		alpha: 2,
		beta:  0.5,
		a: [][]float64{
			{1, 2},
			{4, 5},
			{8, 9},
		},
		b: [][]float64{
			{1, 5, 2, 1},
			{5, -8, 2, 1},
		},
		c: [][]float64{
			{4, 8, 2, 2},
			{12, 16, 8, 9},
			{9, -10, 10, 10},
		},
		ans: [][]float64{
			{24, -18, 13, 7},
			{64, -32, 40, 22.5},
			{110.5, -69, 73, 39},
		},
	},
	{
		m:     2,
		n:     4,
		k:     3,
		alpha: 2,
		beta:  0.5,
		a: [][]float64{
			{1, 2, 3},
			{4, 5, 6},
		},
		b: [][]float64{
			{1, 5, 8, 8},
			{5, -8, 9, 10},
			{6, 2, -3, 2},
		},
		c: [][]float64{
			{4, 8, 7, 8},
			{12, 16, -2, 6},
		},
		ans: [][]float64{
			{60, -6, 37.5, 72},
			{136, -8, 117, 191},
		},
	},
	{
		m:     2,
		n:     3,
		k:     4,
		alpha: 2,
		beta:  0.5,
		a: [][]float64{
			{1, 2, 3, 4},
			{4, 5, 6, 7},
		},
		b: [][]float64{
			{1, 5, 8},
			{5, -8, 9},
			{6, 2, -3},
			{8, 10, 2},
		},
		c: [][]float64{
			{4, 8, 1},
			{12, 16, 6},
		},
		ans: [][]float64{
			{124, 74, 50.5},
			{248, 132, 149},
		},
	},
	{
		m:     2,
		n:     3,
		k:     4,
		alpha: 2,
		beta:  0,
		a: [][]float64{
			{1, 2, 3, 4},
			{4, 5, 6, 7},
		},
		b: [][]float64{
			{1, 5, 8},
			{5, -8, 9},
			{6, 2, -3},
			{8, 10, 2},
		},
		c: [][]float64{
			{4, 8, 1},
			{12, 16, 6},
		},
		ans: [][]float64{
			{122, 70, 50},
			{242, 124, 146},
		},
	},
}

func transpose(a [][]float64) [][]float64 { _ = "STUB: not implemented"; return nil }

func TestDgemm(t *testing.T, blasser Dgemmer) { _ = "STUB: not implemented"; return }

func dgemmcomp(i int, name string, t *testing.T, blasser Dgemmer, tA, tB blas.Transpose, m, n, k int,
	alpha, beta float64, a [][]float64, b [][]float64, c [][]float64, ans [][]float64) {
	_ = "STUB: not implemented"
	return
}
