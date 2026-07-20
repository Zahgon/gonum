package dualcmplx

import (
	"fmt"
)

type Number struct {
	Real, Dual complex128
}

func (d Number) Format(fs fmt.State, c rune) { _ = "STUB: not implemented"; return }

func fmtString(fs fmt.State, c rune, prec, width int, wantPlus bool) string {
	_ = "STUB: not implemented"
	return ""
}

func Add(x, y Number) Number { _ = "STUB: not implemented"; return *new(Number) }

func Sub(x, y Number) Number { _ = "STUB: not implemented"; return *new(Number) }

func Mul(x, y Number) Number { _ = "STUB: not implemented"; return *new(Number) }

func Inv(d Number) Number { _ = "STUB: not implemented"; return *new(Number) }

func Conj(d Number) Number { _ = "STUB: not implemented"; return *new(Number) }

func Scale(f float64, d Number) Number { _ = "STUB: not implemented"; return *new(Number) }

func Abs(d Number) float64 { _ = "STUB: not implemented"; return 0 }

func PowReal(d Number, p float64) Number { _ = "STUB: not implemented"; return *new(Number) }

func Pow(d, p Number) Number { _ = "STUB: not implemented"; return *new(Number) }

func Sqrt(d Number) Number { _ = "STUB: not implemented"; return *new(Number) }

func Exp(d Number) Number { _ = "STUB: not implemented"; return *new(Number) }

func Log(d Number) Number { _ = "STUB: not implemented"; return *new(Number) }
