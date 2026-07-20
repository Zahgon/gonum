package unit

import (
	"fmt"
	"sync"
)

type Uniter interface {
	Unit() *Unit
}

type Dimension int

func NewDimension(symbol string) Dimension { _ = "STUB: not implemented"; return *new(Dimension) }

func (d Dimension) String() string { _ = "STUB: not implemented"; return "" }

func SymbolExists(symbol string) bool { _ = "STUB: not implemented"; return false }

const (
	reserved Dimension = iota
	CurrentDim
	LengthDim
	LuminousIntensityDim
	MassDim
	MoleDim
	TemperatureDim
	TimeDim

	AngleDim
)

var (
	mu      sync.RWMutex
	symbols = []string{
		CurrentDim:           "A",
		LengthDim:            "m",
		LuminousIntensityDim: "cd",
		MassDim:              "kg",
		MoleDim:              "mol",
		TemperatureDim:       "K",
		TimeDim:              "s",
		AngleDim:             "rad",
	}

	dimensions = map[string]Dimension{
		"A":   CurrentDim,
		"m":   LengthDim,
		"cd":  LuminousIntensityDim,
		"kg":  MassDim,
		"mol": MoleDim,
		"K":   TemperatureDim,
		"s":   TimeDim,
		"rad": AngleDim,

		"Y":  reserved,
		"Z":  reserved,
		"E":  reserved,
		"P":  reserved,
		"T":  reserved,
		"G":  reserved,
		"M":  reserved,
		"k":  reserved,
		"h":  reserved,
		"da": reserved,
		"d":  reserved,
		"c":  reserved,
		"μ":  reserved,
		"n":  reserved,
		"p":  reserved,
		"f":  reserved,
		"a":  reserved,
		"z":  reserved,
		"y":  reserved,

		"sr":  reserved,
		"F":   reserved,
		"C":   reserved,
		"S":   reserved,
		"H":   reserved,
		"V":   reserved,
		"Ω":   reserved,
		"J":   reserved,
		"N":   reserved,
		"Hz":  reserved,
		"lx":  reserved,
		"lm":  reserved,
		"Wb":  reserved,
		"W":   reserved,
		"Pa":  reserved,
		"Bq":  reserved,
		"Gy":  reserved,
		"Sv":  reserved,
		"kat": reserved,

		"ha": reserved,
		"L":  reserved,
		"l":  reserved,

		"bar": reserved,
		"b":   reserved,
		"Ci":  reserved,
		"R":   reserved,
		"rd":  reserved,
		"rem": reserved,
	}
)

type Dimensions map[Dimension]int

func (d Dimensions) clone() Dimensions { _ = "STUB: not implemented"; return *new(Dimensions) }

func (d Dimensions) matches(o Dimensions) bool { _ = "STUB: not implemented"; return false }

func (d Dimensions) String() string { _ = "STUB: not implemented"; return "" }

type atom struct {
	Dimension
	pow int
}

type Unit struct {
	dimensions Dimensions
	value      float64
}

func New(value float64, d Dimensions) *Unit { _ = "STUB: not implemented"; return nil }

func DimensionsMatch(a, b Uniter) bool { _ = "STUB: not implemented"; return false }

func (u *Unit) Dimensions() Dimensions { _ = "STUB: not implemented"; return *new(Dimensions) }

func (u *Unit) Add(uniter Uniter) *Unit { _ = "STUB: not implemented"; return nil }

func (u *Unit) Unit() *Unit { _ = "STUB: not implemented"; return nil }

func (u *Unit) Copy() *Unit { _ = "STUB: not implemented"; return nil }

func (u *Unit) Mul(uniter Uniter) *Unit { _ = "STUB: not implemented"; return nil }

func (u *Unit) Div(uniter Uniter) *Unit { _ = "STUB: not implemented"; return nil }

func (u *Unit) Value() float64 { _ = "STUB: not implemented"; return 0 }

func (u *Unit) SetValue(v float64) { _ = "STUB: not implemented"; return }

func (u *Unit) Format(fs fmt.State, c rune) { _ = "STUB: not implemented"; return }

func pos(a int) int { _ = "STUB: not implemented"; return 0 }
