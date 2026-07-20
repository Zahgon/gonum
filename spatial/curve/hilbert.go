package curve

import (
	"errors"
)

var ErrUnderflow = errors.New("order is less than 1")

var ErrOverflow = errors.New("overflow int")

const intSize = 32 << (^uint(0) >> 63)

type Hilbert2D struct{ order int }

func NewHilbert2D(order int) (Hilbert2D, error) {
	_ = "STUB: not implemented"
	return *new(Hilbert2D), nil
}

func (h Hilbert2D) Dims() []int { _ = "STUB: not implemented"; return nil }

func (h Hilbert2D) Len() int { _ = "STUB: not implemented"; return 0 }

func (h Hilbert2D) rot(n int, v []int, d int) { _ = "STUB: not implemented"; return }

func (h Hilbert2D) Pos(v []int) int { _ = "STUB: not implemented"; return 0 }

func (h Hilbert2D) Coord(dst []int, pos int) []int { _ = "STUB: not implemented"; return nil }

type Hilbert3D struct{ order int }

func NewHilbert3D(order int) (Hilbert3D, error) {
	_ = "STUB: not implemented"
	return *new(Hilbert3D), nil
}

func (h Hilbert3D) Dims() []int { _ = "STUB: not implemented"; return nil }

func (h Hilbert3D) Len() int { _ = "STUB: not implemented"; return 0 }

func (h Hilbert3D) rot(reverse bool, n int, v []int, d int) { _ = "STUB: not implemented"; return }

func (h Hilbert3D) Pos(v []int) int { _ = "STUB: not implemented"; return 0 }

func (h Hilbert3D) Coord(dst []int, pos int) []int { _ = "STUB: not implemented"; return nil }

type Hilbert4D struct{ order int }

func NewHilbert4D(order int) (Hilbert4D, error) {
	_ = "STUB: not implemented"
	return *new(Hilbert4D), nil
}

func (h Hilbert4D) Dims() []int { _ = "STUB: not implemented"; return nil }

func (h Hilbert4D) Len() int { _ = "STUB: not implemented"; return 0 }

func (h Hilbert4D) rot(reverse bool, n int, v []int, d int) { _ = "STUB: not implemented"; return }

func (h Hilbert4D) Pos(v []int) int { _ = "STUB: not implemented"; return 0 }

func (h Hilbert4D) Coord(dst []int, pos int) []int { _ = "STUB: not implemented"; return nil }

type op interface{ do(int, []int) }

type invert struct{ i, j int }

func (c invert) do(n int, v []int) { _ = "STUB: not implemented"; return }

type swap struct{ i, j int }

func (c swap) do(n int, v []int) { _ = "STUB: not implemented"; return }

type flip struct{ i, j int }

func (c flip) do(n int, v []int) { _ = "STUB: not implemented"; return }

func do2[A, B op](reverse bool, n int, v []int, a A, b B) { _ = "STUB: not implemented"; return }
