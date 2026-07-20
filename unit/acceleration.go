package unit

import (
	"fmt"
)

type Acceleration float64

func (a Acceleration) Unit() *Unit { _ = "STUB: not implemented"; return nil }

func (a Acceleration) Acceleration() Acceleration {
	_ = "STUB: not implemented"
	return *new(Acceleration)
}

func (a *Acceleration) From(u Uniter) error { _ = "STUB: not implemented"; return nil }

func (a Acceleration) Format(fs fmt.State, c rune) { _ = "STUB: not implemented"; return }
