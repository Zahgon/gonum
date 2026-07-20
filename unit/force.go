package unit

import (
	"fmt"
)

type Force float64

const Newton Force = 1

func (f Force) Unit() *Unit { _ = "STUB: not implemented"; return nil }

func (f Force) Force() Force { _ = "STUB: not implemented"; return *new(Force) }

func (f *Force) From(u Uniter) error { _ = "STUB: not implemented"; return nil }

func (f Force) Format(fs fmt.State, c rune) { _ = "STUB: not implemented"; return }
