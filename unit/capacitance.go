package unit

import (
	"fmt"
)

type Capacitance float64

const Farad Capacitance = 1

func (cp Capacitance) Unit() *Unit { _ = "STUB: not implemented"; return nil }

func (cp Capacitance) Capacitance() Capacitance {
	_ = "STUB: not implemented"
	return *new(Capacitance)
}

func (cp *Capacitance) From(u Uniter) error { _ = "STUB: not implemented"; return nil }

func (cp Capacitance) Format(fs fmt.State, c rune) { _ = "STUB: not implemented"; return }
