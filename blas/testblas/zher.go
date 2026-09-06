package testblas

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

var zherTestCases = []struct {
	alpha float64
	x     []complex128
	a     []complex128

	want    []complex128
	wantRev []complex128
}{
	{
		alpha: 1,
	},
	{
		alpha: 3,
		x: []complex128{
			0 - 3i,
			6 + 10i,
			-2 - 7i,
		},
		a: []complex128{
			-2 + 3i, -3 - 11i, 0 + 4i,
			-3 + 11i, -6 + 3i, 7 + 2i,
			0 - 4i, 7 - 2i, 18 + 3i,
		},
		want: []complex128{
			25 + 0i, -93 - 65i, 63 + 22i,
			-93 + 65i, 402 + 0i, -239 + 68i,
			63 - 22i, -239 - 68i, 177 + 0i},
		wantRev: []complex128{
			157 + 0i, -249 - 77i, 63 - 14i,
			-249 + 77i, 402 + 0i, -83 + 56i,
			63 + 14i, -83 - 56i, 45 + 0i,
		},
	},
	{
		alpha: 3,
		x: []complex128{
			-6 + 2i,
			-2 - 4i,
			0 + 0i,
			0 + 7i,
		},
		a: []complex128{
			2 + 3i, -9 + 7i, 3 + 11i, 10 - 1i,
			-9 - 7i, 16 + 3i, -5 + 2i, -7 - 5i,
			3 - 11i, -5 - 2i, 14 + 3i, 2 - 1i,
			10 + 1i, -7 + 5i, 2 + 1i, 18 + 3i,
		},
		want: []complex128{
			122 + 0i, 3 - 77i, 3 + 11i, 52 + 125i,
			3 + 77i, 76 + 0i, -5 + 2i, -91 + 37i,
			3 - 11i, -5 - 2i, 14 + 0i, 2 - 1i,
			52 - 125i, -91 - 37i, 2 + 1i, 165 + 0i,
		},
		wantRev: []complex128{
			149 + 0i, -9 + 7i, -81 - 31i, 52 - 127i,
			-9 - 7i, 16 + 0i, -5 + 2i, -7 - 5i,
			-81 + 31i, -5 - 2i, 74 + 0i, 14 + 83i,
			52 + 127i, -7 + 5i, 14 - 83i, 138 + 0i,
		},
	},
	{
		alpha: 0,
		x: []complex128{
			-6 + 2i,
			-2 - 4i,
			0 + 0i,
			0 + 7i,
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
		wantRev: []complex128{
			2 + 0i, -9 + 7i, 3 + 11i, 10 - 1i,
			-9 - 7i, 16 + 0i, -5 + 2i, -7 - 5i,
			3 - 11i, -5 - 2i, 14 + 0i, 2 - 1i,
			10 + 1i, -7 + 5i, 2 + 1i, 18 + 0i,
		},
	},
}

type Zherer interface {
	Zher(uplo blas.Uplo, n int, alpha float64, x []complex128, incX int, a []complex128, lda int)
}

func ZherTest(t *testing.T, impl Zherer) { _ = "STUB: not implemented"; return }
