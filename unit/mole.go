package unit

import (
	"fmt"
)

type Mole float64

const Mol Mole = 1

func (n Mole) Unit() *Unit { _ = "STUB: not implemented"; return nil }

func (n Mole) Mole() Mole { _ = "STUB: not implemented"; return *new(Mole) }

func (n *Mole) From(u Uniter) error { _ = "STUB: not implemented"; return nil }

func (n Mole) Format(fs fmt.State, c rune) { _ = "STUB: not implemented"; return }
