package unit

import (
	"fmt"
)

type Torque float64

const Newtonmetre Torque = 1

func (t Torque) Unit() *Unit { _ = "STUB: not implemented"; return nil }

func (t Torque) Torque() Torque { _ = "STUB: not implemented"; return *new(Torque) }

func (t *Torque) From(u Uniter) error { _ = "STUB: not implemented"; return nil }

func (t Torque) Format(fs fmt.State, c rune) { _ = "STUB: not implemented"; return }
