package unit

import (
	"fmt"
)

type Power float64

const Watt Power = 1

func (pw Power) Unit() *Unit { _ = "STUB: not implemented"; return nil }

func (pw Power) Power() Power { _ = "STUB: not implemented"; return *new(Power) }

func (pw *Power) From(u Uniter) error { _ = "STUB: not implemented"; return nil }

func (pw Power) Format(fs fmt.State, c rune) { _ = "STUB: not implemented"; return }
