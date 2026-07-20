package mat

import (
	"fmt"
)

func Formatted(m Matrix, options ...FormatOption) fmt.Formatter {
	_ = "STUB: not implemented"
	return *new(fmt.Formatter)
}

type formatter struct {
	matrix  Matrix
	prefix  string
	margin  int
	dot     byte
	squeeze bool

	format func(m Matrix, prefix string, margin int, dot byte, squeeze bool, fs fmt.State, c rune)
}

type FormatOption func(*formatter)

func Prefix(p string) FormatOption { _ = "STUB: not implemented"; return *new(FormatOption) }

func Excerpt(m int) FormatOption { _ = "STUB: not implemented"; return *new(FormatOption) }

func DotByte(b byte) FormatOption { _ = "STUB: not implemented"; return *new(FormatOption) }

func Squeeze() FormatOption { _ = "STUB: not implemented"; return *new(FormatOption) }

func FormatMATLAB() FormatOption { _ = "STUB: not implemented"; return *new(FormatOption) }

func FormatPython() FormatOption { _ = "STUB: not implemented"; return *new(FormatOption) }

func (f formatter) Format(fs fmt.State, c rune) { _ = "STUB: not implemented"; return }

func format(m Matrix, prefix string, margin int, dot byte, squeeze bool, fs fmt.State, c rune) {
	_ = "STUB: not implemented"
	return
}

func formatMATLAB(m Matrix, prefix string, _ int, _ byte, squeeze bool, fs fmt.State, c rune) {
	_ = "STUB: not implemented"
	return
}

func formatPython(m Matrix, prefix string, _ int, _ byte, squeeze bool, fs fmt.State, c rune) {
	_ = "STUB: not implemented"
	return
}

func fmtString(fs fmt.State, c rune, prec, width int) string { _ = "STUB: not implemented"; return "" }

func maxCellWidth(m Matrix, c rune, printed, prec int, w widther) ([]byte, int) {
	_ = "STUB: not implemented"
	return nil, 0
}

type widther interface {
	width(i int) int
	setWidth(i, w int)
}

type uniformWidth int

func (u *uniformWidth) width(_ int) int   { _ = "STUB: not implemented"; return 0 }
func (u *uniformWidth) setWidth(_, w int) { _ = "STUB: not implemented"; return }

type columnWidth []int

func (c columnWidth) width(i int) int   { _ = "STUB: not implemented"; return 0 }
func (c columnWidth) setWidth(i, w int) { _ = "STUB: not implemented"; return }
