package barneshut

import (
	"errors"

	"gonum.org/v1/gonum/spatial/r3"
)

type Particle3 interface {
	Coord3() r3.Vec
	Mass() float64
}

type Force3 func(p1, p2 Particle3, m1, m2 float64, v r3.Vec) r3.Vec

func Gravity3(_, _ Particle3, m1, m2 float64, v r3.Vec) r3.Vec {
	_ = "STUB: not implemented"
	return *new(r3.Vec)
}

type Volume struct {
	root bucket

	Particles []Particle3
}

func NewVolume(p []Particle3) (*Volume, error) { _ = "STUB: not implemented"; return nil, nil }

func (q *Volume) Reset() (err error) { _ = "STUB: not implemented"; return nil }

var volumeTooBig = errors.New("barneshut: volume too big")

func (q *Volume) ForceOn(p Particle3, theta float64, f Force3) (force r3.Vec) {
	_ = "STUB: not implemented"
	return *new(r3.Vec)
}

type bucket struct {
	particle Particle3

	bounds r3.Box

	nodes [8]*bucket

	center r3.Vec
	mass   float64
}

func (b *bucket) insert(p Particle3) { _ = "STUB: not implemented"; return }

func (b *bucket) passDown(p Particle3) { _ = "STUB: not implemented"; return }

const (
	lne = iota
	lse
	lsw
	lnw
	une
	use
	usw
	unw
)

func octantOf(b r3.Box, p Particle3) int { _ = "STUB: not implemented"; return 0 }

func splitVolume(b r3.Box, dir int) r3.Box { _ = "STUB: not implemented"; return *new(r3.Box) }

func (b *bucket) summarize() (center r3.Vec, mass float64) {
	_ = "STUB: not implemented"
	return *new(r3.Vec), 0
}

func (b *bucket) forceOn(p Particle3, pt r3.Vec, m, theta float64, f Force3) (vector r3.Vec) {
	_ = "STUB: not implemented"
	return *new(r3.Vec)
}
