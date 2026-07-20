package unit

import (
	"fmt"
)

type Charge float64

const Coulomb Charge = 1

func (ch Charge) Unit() *Unit { _ = "STUB: not implemented"; return nil }

func (ch Charge) Charge() Charge { _ = "STUB: not implemented"; return *new(Charge) }

func (ch *Charge) From(u Uniter) error { _ = "STUB: not implemented"; return nil }

func (ch Charge) Format(fs fmt.State, c rune) { _ = "STUB: not implemented"; return }
