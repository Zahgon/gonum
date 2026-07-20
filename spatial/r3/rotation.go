package r3

import (
	"gonum.org/v1/gonum/num/quat"
)

type Rotation quat.Number

func NewRotation(alpha float64, axis Vec) Rotation {
	_ = "STUB: not implemented"
	return *new(Rotation)
}

func (r Rotation) Rotate(p Vec) Vec { _ = "STUB: not implemented"; return *new(Vec) }

func (r Rotation) isIdentity() bool { _ = "STUB: not implemented"; return false }

func raise(p Vec) quat.Number { _ = "STUB: not implemented"; return *new(quat.Number) }
