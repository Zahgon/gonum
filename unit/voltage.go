package unit

import (
	"fmt"
)

type Voltage float64

const Volt Voltage = 1

func (v Voltage) Unit() *Unit { _ = "STUB: not implemented"; return nil }

func (v Voltage) Voltage() Voltage { _ = "STUB: not implemented"; return *new(Voltage) }

func (v *Voltage) From(u Uniter) error { _ = "STUB: not implemented"; return nil }

func (v Voltage) Format(fs fmt.State, c rune) { _ = "STUB: not implemented"; return }
