package unit

import (
	"fmt"
)

type Length float64

const Metre Length = 1

func (l Length) Unit() *Unit { _ = "STUB: not implemented"; return nil }

func (l Length) Length() Length { _ = "STUB: not implemented"; return *new(Length) }

func (l *Length) From(u Uniter) error { _ = "STUB: not implemented"; return nil }

func (l Length) Format(fs fmt.State, c rune) { _ = "STUB: not implemented"; return }
