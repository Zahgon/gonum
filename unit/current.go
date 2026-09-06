package unit

import (
	"fmt"
)

type Current float64

const Ampere Current = 1

func (i Current) Unit() *Unit { _ = "STUB: not implemented"; return nil }

func (i Current) Current() Current { _ = "STUB: not implemented"; return *new(Current) }

func (i *Current) From(u Uniter) error { _ = "STUB: not implemented"; return nil }

func (i Current) Format(fs fmt.State, c rune) { _ = "STUB: not implemented"; return }
