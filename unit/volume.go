package unit

import (
	"fmt"
)

type Volume float64

const Litre Volume = 1e-3

func (v Volume) Unit() *Unit { _ = "STUB: not implemented"; return nil }

func (v Volume) Volume() Volume { _ = "STUB: not implemented"; return *new(Volume) }

func (v *Volume) From(u Uniter) error { _ = "STUB: not implemented"; return nil }

func (v Volume) Format(fs fmt.State, c rune) { _ = "STUB: not implemented"; return }
