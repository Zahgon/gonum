package quat

import (
	"fmt"
)

var zero Number

type Number struct {
	Real, Imag, Jmag, Kmag float64
}

func (q Number) Format(fs fmt.State, c rune) { _ = "STUB: not implemented"; return }

func fmtString(fs fmt.State, c rune, prec, width int, wantPlus bool) string {
	_ = "STUB: not implemented"
	return ""
}

func Add(x, y Number) Number { _ = "STUB: not implemented"; return *new(Number) }

func Sub(x, y Number) Number { _ = "STUB: not implemented"; return *new(Number) }

func Mul(x, y Number) Number { _ = "STUB: not implemented"; return *new(Number) }

func Scale(f float64, q Number) Number { _ = "STUB: not implemented"; return *new(Number) }

func Parse(s string) (Number, error) { _ = "STUB: not implemented"; return *new(Number), nil }

func floatPart(s string) (beg, end int, part uint, err error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, nil
}

func isSign(r rune) bool { _ = "STUB: not implemented"; return false }

func isDigit(r rune) bool { _ = "STUB: not implemented"; return false }

func isExponent(r rune) bool { _ = "STUB: not implemented"; return false }

func isDot(r rune) bool { _ = "STUB: not implemented"; return false }

type parseError struct {
	string string
	state  int
	rune   rune
}

func (e parseError) Error() string { _ = "STUB: not implemented"; return "" }
