package unit

import (
	"fmt"
)

type Frequency float64

const Hertz Frequency = 1

func (f Frequency) Unit() *Unit { _ = "STUB: not implemented"; return nil }

func (f Frequency) Frequency() Frequency { _ = "STUB: not implemented"; return *new(Frequency) }

func (f *Frequency) From(u Uniter) error { _ = "STUB: not implemented"; return nil }

func (f Frequency) Format(fs fmt.State, c rune) { _ = "STUB: not implemented"; return }
