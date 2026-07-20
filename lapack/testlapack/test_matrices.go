package testlapack

import (
	"math/rand/v2"

	"gonum.org/v1/gonum/blas/blas64"
)

type A123 struct{}

func (A123) Matrix() blas64.General { _ = "STUB: not implemented"; return *new(blas64.General) }

func (A123) Eigenvalues() []complex128 { _ = "STUB: not implemented"; return nil }

func (A123) LeftEV() blas64.General { _ = "STUB: not implemented"; return *new(blas64.General) }

func (A123) RightEV() blas64.General { _ = "STUB: not implemented"; return *new(blas64.General) }

type AntisymRandom struct {
	mat blas64.General
}

func NewAntisymRandom(n int, rnd *rand.Rand) AntisymRandom {
	_ = "STUB: not implemented"
	return *new(AntisymRandom)
}

func (a AntisymRandom) Matrix() blas64.General {
	_ = "STUB: not implemented"
	return *new(blas64.General)
}

func (AntisymRandom) Eigenvalues() []complex128 { _ = "STUB: not implemented"; return nil }

type Circulant int

func (c Circulant) Matrix() blas64.General { _ = "STUB: not implemented"; return *new(blas64.General) }

func (c Circulant) Eigenvalues() []complex128 { _ = "STUB: not implemented"; return nil }

type Clement int

func (c Clement) Matrix() blas64.General { _ = "STUB: not implemented"; return *new(blas64.General) }

func (c Clement) Eigenvalues() []complex128 { _ = "STUB: not implemented"; return nil }

type Creation int

func (c Creation) Matrix() blas64.General { _ = "STUB: not implemented"; return *new(blas64.General) }

func (c Creation) Eigenvalues() []complex128 { _ = "STUB: not implemented"; return nil }

type Diagonal int

func (d Diagonal) Matrix() blas64.General { _ = "STUB: not implemented"; return *new(blas64.General) }

func (d Diagonal) Eigenvalues() []complex128 { _ = "STUB: not implemented"; return nil }

type Downshift int

func (d Downshift) Matrix() blas64.General { _ = "STUB: not implemented"; return *new(blas64.General) }

func (d Downshift) Eigenvalues() []complex128 { _ = "STUB: not implemented"; return nil }

type Fibonacci int

func (f Fibonacci) Matrix() blas64.General { _ = "STUB: not implemented"; return *new(blas64.General) }

func (f Fibonacci) Eigenvalues() []complex128 { _ = "STUB: not implemented"; return nil }

type Gear int

func (g Gear) Matrix() blas64.General { _ = "STUB: not implemented"; return *new(blas64.General) }

func (g Gear) Eigenvalues() []complex128 { _ = "STUB: not implemented"; return nil }

type Grcar struct {
	N int
	K int
}

func (g Grcar) Matrix() blas64.General { _ = "STUB: not implemented"; return *new(blas64.General) }

func (Grcar) Eigenvalues() []complex128 { _ = "STUB: not implemented"; return nil }

type Hanowa struct {
	N     int
	Alpha float64
}

func (h Hanowa) Matrix() blas64.General { _ = "STUB: not implemented"; return *new(blas64.General) }

func (h Hanowa) Eigenvalues() []complex128 { _ = "STUB: not implemented"; return nil }

type Lesp int

func (l Lesp) Matrix() blas64.General { _ = "STUB: not implemented"; return *new(blas64.General) }

func (Lesp) Eigenvalues() []complex128 { _ = "STUB: not implemented"; return nil }

type Rutis struct{}

func (Rutis) Matrix() blas64.General { _ = "STUB: not implemented"; return *new(blas64.General) }

func (Rutis) Eigenvalues() []complex128 { _ = "STUB: not implemented"; return nil }

type Tris struct {
	N       int
	X, Y, Z float64
}

func (t Tris) Matrix() blas64.General { _ = "STUB: not implemented"; return *new(blas64.General) }

func (t Tris) Eigenvalues() []complex128 { _ = "STUB: not implemented"; return nil }

type Wilk4 struct{}

func (Wilk4) Matrix() blas64.General { _ = "STUB: not implemented"; return *new(blas64.General) }

func (Wilk4) Eigenvalues() []complex128 { _ = "STUB: not implemented"; return nil }

type Wilk12 struct{}

func (Wilk12) Matrix() blas64.General { _ = "STUB: not implemented"; return *new(blas64.General) }

func (Wilk12) Eigenvalues() []complex128 { _ = "STUB: not implemented"; return nil }

type Wilk20 float64

func (w Wilk20) Matrix() blas64.General { _ = "STUB: not implemented"; return *new(blas64.General) }

func (w Wilk20) Eigenvalues() []complex128 { _ = "STUB: not implemented"; return nil }

type Zero int

func (z Zero) Matrix() blas64.General { _ = "STUB: not implemented"; return *new(blas64.General) }

func (z Zero) Eigenvalues() []complex128 { _ = "STUB: not implemented"; return nil }
