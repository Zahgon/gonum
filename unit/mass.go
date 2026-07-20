package unit

import (
	"fmt"
)

type Mass float64

const (
	Gram Mass = 1e-3

	Kilogram = Kilo * Gram
)

func (m Mass) Unit() *Unit { _ = "STUB: not implemented"; return nil }

func (m Mass) Mass() Mass { _ = "STUB: not implemented"; return *new(Mass) }

func (m *Mass) From(u Uniter) error { _ = "STUB: not implemented"; return nil }

func (m Mass) Format(fs fmt.State, c rune) { _ = "STUB: not implemented"; return }
