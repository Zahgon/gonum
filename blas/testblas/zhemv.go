package testblas

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

var zhemvTestCases = []struct {
	uplo  blas.Uplo
	alpha complex128
	a     []complex128
	x     []complex128
	beta  complex128
	y     []complex128

	want      []complex128
	wantXNeg  []complex128
	wantYNeg  []complex128
	wantXYNeg []complex128
}{
	{
		uplo:  blas.Upper,
		alpha: 6 + 2i,
		beta:  -6 - 7i,
	},
	{
		uplo:  blas.Lower,
		alpha: 6 + 2i,
		beta:  -6 - 7i,
	},
	{
		uplo:  blas.Upper,
		alpha: 6 + 2i,
		a: []complex128{
			7, 8 + 4i, -9 - 6i, -9 + 3i,
			znan, -3, -10 - 6i, 0 + 3i,
			znan, znan, 6, 2 + 8i,
			znan, znan, znan, -4,
		},
		x: []complex128{
			-4 + 0i,
			-2 - 5i,
			8 + 0i,
			6 - 1i,
		},
		beta: -6 - 7i,
		y: []complex128{
			1 - 5i,
			-2 - 5i,
			0 - 4i,
			7 + 7i,
		},
		want: []complex128{
			-785 - 685i,
			-643 - 156i,
			776 + 692i,
			169 - 317i,
		},
		wantXNeg: []complex128{
			599 + 703i,
			1 + 172i,
			-978 - 86i,
			-449 - 423i,
		},
		wantYNeg: []complex128{
			121 - 203i,
			781 + 712i,
			-648 - 176i,
			-737 - 799i,
		},
		wantXYNeg: []complex128{
			-497 - 309i,
			-973 - 66i,
			-4 + 152i,
			647 + 589i,
		},
	},
	{
		uplo:  blas.Lower,
		alpha: 6 + 2i,
		a: []complex128{
			7, znan, znan, znan,
			8 - 4i, -3, znan, znan,
			-9 + 6i, -10 + 6i, 6, znan,
			-9 - 3i, 0 - 3i, 2 - 8i, -4,
		},
		x: []complex128{
			-4 + 0i,
			-2 - 5i,
			8 + 0i,
			6 - 1i,
		},
		beta: -6 - 7i,
		y: []complex128{
			1 - 5i,
			-2 - 5i,
			0 - 4i,
			7 + 7i,
		},
		want: []complex128{
			-785 - 685i,
			-643 - 156i,
			776 + 692i,
			169 - 317i,
		},
		wantXNeg: []complex128{
			599 + 703i,
			1 + 172i,
			-978 - 86i,
			-449 - 423i,
		},
		wantYNeg: []complex128{
			121 - 203i,
			781 + 712i,
			-648 - 176i,
			-737 - 799i,
		},
		wantXYNeg: []complex128{
			-497 - 309i,
			-973 - 66i,
			-4 + 152i,
			647 + 589i,
		},
	},
	{
		uplo:  blas.Upper,
		alpha: 0,
		a: []complex128{
			7, 8 + 4i, -9 - 6i, -9 + 3i,
			znan, -3, -10 - 6i, 0 + 3i,
			znan, znan, 6, 2 + 8i,
			znan, znan, znan, -4,
		},
		x: []complex128{
			-4 + 0i,
			-2 - 5i,
			8 + 0i,
			6 - 1i,
		},
		beta: -6 - 7i,
		y: []complex128{
			1 - 5i,
			-2 - 5i,
			0 - 4i,
			7 + 7i,
		},
		want: []complex128{
			-41 + 23i,
			-23 + 44i,
			-28 + 24i,
			7 - 91i,
		},
		wantXNeg: []complex128{
			-41 + 23i,
			-23 + 44i,
			-28 + 24i,
			7 - 91i,
		},
		wantYNeg: []complex128{
			-41 + 23i,
			-23 + 44i,
			-28 + 24i,
			7 - 91i,
		},
		wantXYNeg: []complex128{
			-41 + 23i,
			-23 + 44i,
			-28 + 24i,
			7 - 91i,
		},
	},
	{
		uplo:  blas.Upper,
		alpha: 6 + 2i,
		a: []complex128{
			7, 8 + 4i, -9 - 6i, -9 + 3i,
			znan, -3, -10 - 6i, 0 + 3i,
			znan, znan, 6, 2 + 8i,
			znan, znan, znan, -4,
		},
		x: []complex128{
			-4 + 0i,
			-2 - 5i,
			8 + 0i,
			6 - 1i,
		},
		beta: 0,
		y: []complex128{
			1 - 5i,
			-2 - 5i,
			0 - 4i,
			7 + 7i,
		},
		want: []complex128{
			-744 - 708i,
			-620 - 200i,
			804 + 668i,
			162 - 226i,
		},
		wantXNeg: []complex128{
			640 + 680i,
			24 + 128i,
			-950 - 110i,
			-456 - 332i,
		},
		wantYNeg: []complex128{
			162 - 226i,
			804 + 668i,
			-620 - 200i,
			-744 - 708i,
		},
		wantXYNeg: []complex128{
			-456 - 332i,
			-950 - 110i,
			24 + 128i,
			640 + 680i,
		},
	},
}

type Zhemver interface {
	Zhemv(uplo blas.Uplo, n int, alpha complex128, a []complex128, lda int, x []complex128, incX int, beta complex128, y []complex128, incY int)
}

func ZhemvTest(t *testing.T, impl Zhemver) { _ = "STUB: not implemented"; return }
