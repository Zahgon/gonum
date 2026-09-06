package unit

import (
	"fmt"
)

type Velocity float64

func (v Velocity) Unit() *Unit { _ = "STUB: not implemented"; return nil }

func (v Velocity) Velocity() Velocity { _ = "STUB: not implemented"; return *new(Velocity) }

func (v *Velocity) From(u Uniter) error { _ = "STUB: not implemented"; return nil }

func (v Velocity) Format(fs fmt.State, c rune) { _ = "STUB: not implemented"; return }
