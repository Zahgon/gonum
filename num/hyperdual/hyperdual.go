package hyperdual

import (
	"fmt"
	"math"
)

type Number struct {
	Real, E1mag, E2mag, E1E2mag float64
}

var negZero = math.Float64frombits(1 << 63)

func (d Number) Format(fs fmt.State, c rune) { _ = "STUB: not implemented"; return }

func fmtString(fs fmt.State, c rune, prec, width int, wantPlus bool) string {
	_ = "STUB: not implemented"
	return ""
}

func Add(x, y Number) Number { _ = "STUB: not implemented"; return *new(Number) }

func Sub(x, y Number) Number { _ = "STUB: not implemented"; return *new(Number) }

func Mul(x, y Number) Number { _ = "STUB: not implemented"; return *new(Number) }

func Inv(d Number) Number { _ = "STUB: not implemented"; return *new(Number) }

func Scale(f float64, d Number) Number { _ = "STUB: not implemented"; return *new(Number) }

func Abs(d Number) Number { _ = "STUB: not implemented"; return *new(Number) }
