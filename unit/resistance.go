package unit

import (
	"fmt"
)

type Resistance float64

const Ohm Resistance = 1

func (r Resistance) Unit() *Unit { _ = "STUB: not implemented"; return nil }

func (r Resistance) Resistance() Resistance { _ = "STUB: not implemented"; return *new(Resistance) }

func (r *Resistance) From(u Uniter) error { _ = "STUB: not implemented"; return nil }

func (r Resistance) Format(fs fmt.State, c rune) { _ = "STUB: not implemented"; return }
