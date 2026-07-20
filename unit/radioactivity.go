package unit

import (
	"fmt"
)

type Radioactivity float64

const Becquerel Radioactivity = 1

func (r Radioactivity) Unit() *Unit { _ = "STUB: not implemented"; return nil }

func (r Radioactivity) Radioactivity() Radioactivity {
	_ = "STUB: not implemented"
	return *new(Radioactivity)
}

func (r *Radioactivity) From(u Uniter) error { _ = "STUB: not implemented"; return nil }

func (r Radioactivity) Format(fs fmt.State, c rune) { _ = "STUB: not implemented"; return }
