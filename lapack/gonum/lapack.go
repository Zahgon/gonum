package gonum

import "gonum.org/v1/gonum/lapack"

type Implementation struct{}

var _ lapack.Float64 = Implementation{}

func abs(a int) int { _ = "STUB: not implemented"; return 0 }

const (
	dlamchE = 0x1p-53

	dlamchB = 2

	dlamchP = dlamchB * dlamchE

	dlamchS = 0x1p-1022

	dtsml = 0x1p-511
	dtbig = 0x1p486
	dssml = 0x1p537
	dsbig = 0x1p-538
)
