package gen

import (
	"math/rand/v2"
)

func NavigableSmallWorld(dst GraphBuilder, dims []int, p, q int, r float64, src rand.Source) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func iterateOver(dims []int, fn func(state []int)) { _ = "STUB: not implemented"; return }

func iterator(d int, dims, state []int, fn func(state []int)) { _ = "STUB: not implemented"; return }

func manhattanBetween(a, b []int) int { _ = "STUB: not implemented"; return 0 }

func manhattanDelta(a, delta, dims []int, translate int) int { _ = "STUB: not implemented"; return 0 }

func idxFrom(n, dims []int) int { _ = "STUB: not implemented"; return 0 }

func idxFromDelta(base, delta, dims []int, translate int) int { _ = "STUB: not implemented"; return 0 }
