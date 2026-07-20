package unit

import (
	"fmt"
)

type Inductance float64

const Henry Inductance = 1

func (i Inductance) Unit() *Unit { _ = "STUB: not implemented"; return nil }

func (i Inductance) Inductance() Inductance { _ = "STUB: not implemented"; return *new(Inductance) }

func (i *Inductance) From(u Uniter) error { _ = "STUB: not implemented"; return nil }

func (i Inductance) Format(fs fmt.State, c rune) { _ = "STUB: not implemented"; return }
