package unit

import (
	"fmt"
)

type Area float64

func (a Area) Unit() *Unit { _ = "STUB: not implemented"; return nil }

func (a Area) Area() Area { _ = "STUB: not implemented"; return *new(Area) }

func (a *Area) From(u Uniter) error { _ = "STUB: not implemented"; return nil }

func (a Area) Format(fs fmt.State, c rune) { _ = "STUB: not implemented"; return }
