package unit

import (
	"fmt"
)

type Time float64

const (
	Second Time = 1

	Minute = 60 * Second
	Hour   = 60 * Minute
)

func (t Time) Unit() *Unit { _ = "STUB: not implemented"; return nil }

func (t Time) Time() Time { _ = "STUB: not implemented"; return *new(Time) }

func (t *Time) From(u Uniter) error { _ = "STUB: not implemented"; return nil }

func (t Time) Format(fs fmt.State, c rune) { _ = "STUB: not implemented"; return }
