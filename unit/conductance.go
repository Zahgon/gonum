package unit

import (
	"fmt"
)

type Conductance float64

const Siemens Conductance = 1

func (co Conductance) Unit() *Unit { _ = "STUB: not implemented"; return nil }

func (co Conductance) Conductance() Conductance {
	_ = "STUB: not implemented"
	return *new(Conductance)
}

func (co *Conductance) From(u Uniter) error { _ = "STUB: not implemented"; return nil }

func (co Conductance) Format(fs fmt.State, c rune) { _ = "STUB: not implemented"; return }
