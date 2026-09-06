package unit

import (
	"fmt"
)

type Pressure float64

const Pascal Pressure = 1

func (pr Pressure) Unit() *Unit { _ = "STUB: not implemented"; return nil }

func (pr Pressure) Pressure() Pressure { _ = "STUB: not implemented"; return *new(Pressure) }

func (pr *Pressure) From(u Uniter) error { _ = "STUB: not implemented"; return nil }

func (pr Pressure) Format(fs fmt.State, c rune) { _ = "STUB: not implemented"; return }
