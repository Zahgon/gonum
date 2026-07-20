package unit

import (
	"fmt"
)

type Dimless float64

func (d Dimless) Unit() *Unit { _ = "STUB: not implemented"; return nil }

func (d Dimless) Dimless() Dimless { _ = "STUB: not implemented"; return *new(Dimless) }

func (d *Dimless) From(u Uniter) error { _ = "STUB: not implemented"; return nil }

func (d Dimless) Format(fs fmt.State, c rune) { _ = "STUB: not implemented"; return }
