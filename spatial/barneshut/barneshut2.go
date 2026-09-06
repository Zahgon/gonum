package barneshut

import (
	"errors"

	"gonum.org/v1/gonum/spatial/r2"
)

type Particle2 interface {
	Coord2() r2.Vec
	Mass() float64
}

type Force2 func(p1, p2 Particle2, m1, m2 float64, v r2.Vec) r2.Vec

func Gravity2(_, _ Particle2, m1, m2 float64, v r2.Vec) r2.Vec {
	_ = "STUB: not implemented"
	return *new(r2.Vec)
}

type Plane struct {
	root tile

	Particles []Particle2
}

func NewPlane(p []Particle2) (*Plane, error) { _ = "STUB: not implemented"; return nil, nil }

func (q *Plane) Reset() (err error) { _ = "STUB: not implemented"; return nil }

var planeTooBig = errors.New("barneshut: plane too big")

func (q *Plane) ForceOn(p Particle2, theta float64, f Force2) (force r2.Vec) {
	_ = "STUB: not implemented"
	return *new(r2.Vec)
}

type tile struct {
	particle Particle2

	bounds r2.Box

	nodes [4]*tile

	center r2.Vec
	mass   float64
}

func (t *tile) insert(p Particle2) { _ = "STUB: not implemented"; return }

func (t *tile) passDown(p Particle2) { _ = "STUB: not implemented"; return }

const (
	ne = iota
	se
	sw
	nw
)

func quadrantOf(b r2.Box, p Particle2) int { _ = "STUB: not implemented"; return 0 }

func splitPlane(b r2.Box, dir int) r2.Box { _ = "STUB: not implemented"; return *new(r2.Box) }

func (t *tile) summarize() (center r2.Vec, mass float64) {
	_ = "STUB: not implemented"
	return *new(r2.Vec), 0
}

func (t *tile) forceOn(p Particle2, pt r2.Vec, m, theta float64, f Force2) (vector r2.Vec) {
	_ = "STUB: not implemented"
	return *new(r2.Vec)
}
