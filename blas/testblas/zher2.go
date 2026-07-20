package testblas

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

var zher2TestCases = []struct {
	alpha      complex128
	incX, incY int
	x          []complex128
	y          []complex128
	a          []complex128

	want []complex128
}{
	{
		alpha: 1 + 2i,
		incX:  1,
		incY:  1,
	},
	{
		alpha: 1 + 2i,
		incX:  1,
		incY:  1,
		x: []complex128{
			-6 + 2i,
			-2 - 4i,
			0 + 0i,
			0 + 7i,
		},
		y: []complex128{
			2 - 5i,
			0 + 0i,
			-8 - 9i,
			6 + 6i,
		},
		a: []complex128{
			2 + 0i, -9 + 7i, 3 + 11i, 10 - 1i,
			-9 - 7i, 16 + 0i, -5 + 2i, -7 - 5i,
			3 - 11i, -5 - 2i, 14 + 0i, 2 - 1i,
			10 + 1i, -7 + 5i, 2 + 1i, 18 + 0i,
		},
		want: []complex128{
			62 + 0i, 43 - 7i, 173 + 1i, -173 + 55i,
			43 + 7i, 16 + 0i, 19 + 120i, -19 - 89i,
			173 - 1i, 19 - 120i, 14 + 0i, 51 + 181i,
			-173 - 55i, -19 + 89i, 51 - 181i, -66 + 0i,
		},
	},
	{
		alpha: 1 + 2i,
		incX:  1,
		incY:  1,
		x: []complex128{
			-6 + 2i,
			-2 - 4i,
			0 + 0i,
			0 + 7i,
		},
		y: []complex128{
			2 - 5i,
			-8 - 9i,
			0 + 0i,
			6 + 6i,
		},
		a: []complex128{
			2 + 0i, -9 + 7i, 3 + 11i, 10 - 1i,
			-9 - 7i, 16 + 0i, -5 + 2i, -7 - 5i,
			3 - 11i, -5 - 2i, 14 + 0i, 2 - 1i,
			10 + 1i, -7 + 5i, 2 + 1i, 18 + 0i,
		},
		want: []complex128{
			62 + 0i, 213 - 17i, 3 + 11i, -173 + 55i,
			213 + 17i, 64 + 0i, -5 + 2i, 30 + 93i,
			3 - 11i, -5 - 2i, 14 + 0i, 2 - 1i,
			-173 - 55i, 30 - 93i, 2 + 1i, -66 + 0i,
		},
	},
	{
		alpha: 1 + 2i,
		incX:  2,
		incY:  4,
		x: []complex128{
			-6 + 2i,
			-2 - 4i,
			0 + 0i,
			0 + 7i,
		},
		y: []complex128{
			2 - 5i,
			-8 - 9i,
			0 + 0i,
			6 + 6i,
		},
		a: []complex128{
			2 + 0i, -9 + 7i, 3 + 11i, 10 - 1i,
			-9 - 7i, 16 + 0i, -5 + 2i, -7 - 5i,
			3 - 11i, -5 - 2i, 14 + 0i, 2 - 1i,
			10 + 1i, -7 + 5i, 2 + 1i, 18 + 0i,
		},
		want: []complex128{
			62 + 0i, 213 - 17i, 3 + 11i, -173 + 55i,
			213 + 17i, 64 + 0i, -5 + 2i, 30 + 93i,
			3 - 11i, -5 - 2i, 14 + 0i, 2 - 1i,
			-173 - 55i, 30 - 93i, 2 + 1i, -66 + 0i,
		},
	},
	{
		alpha: 1 + 2i,
		incX:  3,
		incY:  7,
		x: []complex128{
			-6 + 2i,
			-2 - 4i,
			0 + 0i,
			0 + 7i,
		},
		y: []complex128{
			2 - 5i,
			0 + 0i,
			-8 - 9i,
			6 + 6i,
		},
		a: []complex128{
			2 + 0i, -9 + 7i, 3 + 11i, 10 - 1i,
			-9 - 7i, 16 + 0i, -5 + 2i, -7 - 5i,
			3 - 11i, -5 - 2i, 14 + 0i, 2 - 1i,
			10 + 1i, -7 + 5i, 2 + 1i, 18 + 0i,
		},
		want: []complex128{
			62 + 0i, 43 - 7i, 173 + 1i, -173 + 55i,
			43 + 7i, 16 + 0i, 19 + 120i, -19 - 89i,
			173 - 1i, 19 - 120i, 14 + 0i, 51 + 181i,
			-173 - 55i, -19 + 89i, 51 - 181i, -66 + 0i,
		},
	},
	{
		alpha: 1 + 2i,
		incX:  -3,
		incY:  7,
		x: []complex128{
			0 + 7i,
			0 + 0i,
			-2 - 4i,
			-6 + 2i,
		},
		y: []complex128{
			2 - 5i,
			0 + 0i,
			-8 - 9i,
			6 + 6i,
		},
		a: []complex128{
			2 + 0i, -9 + 7i, 3 + 11i, 10 - 1i,
			-9 - 7i, 16 + 0i, -5 + 2i, -7 - 5i,
			3 - 11i, -5 - 2i, 14 + 0i, 2 - 1i,
			10 + 1i, -7 + 5i, 2 + 1i, 18 + 0i,
		},
		want: []complex128{
			62 + 0i, 43 - 7i, 173 + 1i, -173 + 55i,
			43 + 7i, 16 + 0i, 19 + 120i, -19 - 89i,
			173 - 1i, 19 - 120i, 14 + 0i, 51 + 181i,
			-173 - 55i, -19 + 89i, 51 - 181i, -66 + 0i,
		},
	},
	{
		alpha: 1 + 2i,
		incX:  3,
		incY:  -7,
		x: []complex128{
			-6 + 2i,
			-2 - 4i,
			0 + 0i,
			0 + 7i,
		},
		y: []complex128{
			6 + 6i,
			-8 - 9i,
			0 + 0i,
			2 - 5i,
		},
		a: []complex128{
			2 + 0i, -9 + 7i, 3 + 11i, 10 - 1i,
			-9 - 7i, 16 + 0i, -5 + 2i, -7 - 5i,
			3 - 11i, -5 - 2i, 14 + 0i, 2 - 1i,
			10 + 1i, -7 + 5i, 2 + 1i, 18 + 0i,
		},
		want: []complex128{
			62 + 0i, 43 - 7i, 173 + 1i, -173 + 55i,
			43 + 7i, 16 + 0i, 19 + 120i, -19 - 89i,
			173 - 1i, 19 - 120i, 14 + 0i, 51 + 181i,
			-173 - 55i, -19 + 89i, 51 - 181i, -66 + 0i,
		},
	},
	{
		alpha: 1 + 2i,
		incX:  -3,
		incY:  -7,
		x: []complex128{
			0 + 7i,
			0 + 0i,
			-2 - 4i,
			-6 + 2i,
		},
		y: []complex128{
			6 + 6i,
			-8 - 9i,
			0 + 0i,
			2 - 5i,
		},
		a: []complex128{
			2 + 0i, -9 + 7i, 3 + 11i, 10 - 1i,
			-9 - 7i, 16 + 0i, -5 + 2i, -7 - 5i,
			3 - 11i, -5 - 2i, 14 + 0i, 2 - 1i,
			10 + 1i, -7 + 5i, 2 + 1i, 18 + 0i,
		},
		want: []complex128{
			62 + 0i, 43 - 7i, 173 + 1i, -173 + 55i,
			43 + 7i, 16 + 0i, 19 + 120i, -19 - 89i,
			173 - 1i, 19 - 120i, 14 + 0i, 51 + 181i,
			-173 - 55i, -19 + 89i, 51 - 181i, -66 + 0i,
		},
	},
	{
		alpha: 0,
		incX:  1,
		incY:  1,
		x: []complex128{
			-6 + 2i,
			-2 - 4i,
			0 + 0i,
			0 + 7i,
		},
		y: []complex128{
			2 - 5i,
			0 + 0i,
			-8 - 9i,
			6 + 6i,
		},
		a: []complex128{
			2 + 0i, -9 + 7i, 3 + 11i, 10 - 1i,
			-9 - 7i, 16 + 0i, -5 + 2i, -7 - 5i,
			3 - 11i, -5 - 2i, 14 + 0i, 2 - 1i,
			10 + 1i, -7 + 5i, 2 + 1i, 18 + 0i,
		},
		want: []complex128{
			2 + 0i, -9 + 7i, 3 + 11i, 10 - 1i,
			-9 - 7i, 16 + 0i, -5 + 2i, -7 - 5i,
			3 - 11i, -5 - 2i, 14 + 0i, 2 - 1i,
			10 + 1i, -7 + 5i, 2 + 1i, 18 + 0i,
		},
	},
}

type Zher2er interface {
	Zher2(uplo blas.Uplo, n int, alpha complex128, x []complex128, incX int, y []complex128, incY int, a []complex128, lda int)
}

func Zher2Test(t *testing.T, impl Zher2er) { _ = "STUB: not implemented"; return }
