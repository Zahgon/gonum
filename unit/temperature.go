package unit

import (
	"fmt"
)

type Temperature float64

const Kelvin Temperature = 1

func (t Temperature) Unit() *Unit { _ = "STUB: not implemented"; return nil }

func (t Temperature) Temperature() Temperature { _ = "STUB: not implemented"; return *new(Temperature) }

func (t *Temperature) From(u Uniter) error { _ = "STUB: not implemented"; return nil }

func (t Temperature) Format(fs fmt.State, c rune) { _ = "STUB: not implemented"; return }
