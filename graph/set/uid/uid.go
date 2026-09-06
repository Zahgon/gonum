package uid

import (
	"math"

	"gonum.org/v1/gonum/graph/internal/set"
)

const Max = math.MaxInt64

type Set struct {
	maxID      int64
	used, free set.Ints[int64]
}

func NewSet() *Set { _ = "STUB: not implemented"; return nil }

func (s *Set) NewID() int64 { _ = "STUB: not implemented"; return 0 }

func (s *Set) Use(id int64) { _ = "STUB: not implemented"; return }

func (s *Set) Release(id int64) { _ = "STUB: not implemented"; return }
