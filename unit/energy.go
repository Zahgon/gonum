package unit

import (
	"fmt"
)

type Energy float64

const Joule Energy = 1

func (e Energy) Unit() *Unit { _ = "STUB: not implemented"; return nil }

func (e Energy) Energy() Energy { _ = "STUB: not implemented"; return *new(Energy) }

func (e *Energy) From(u Uniter) error { _ = "STUB: not implemented"; return nil }

func (e Energy) Format(fs fmt.State, c rune) { _ = "STUB: not implemented"; return }
