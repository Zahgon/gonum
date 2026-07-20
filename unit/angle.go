package unit

import (
	"fmt"
)

type Angle float64

const Rad Angle = 1

func (a Angle) Unit() *Unit { _ = "STUB: not implemented"; return nil }

func (a Angle) Angle() Angle { _ = "STUB: not implemented"; return *new(Angle) }

func (a *Angle) From(u Uniter) error { _ = "STUB: not implemented"; return nil }

func (a Angle) Format(fs fmt.State, c rune) { _ = "STUB: not implemented"; return }
