package r2

type Box struct {
	Min, Max Vec
}

func NewBox(x0, y0, x1, y1 float64) Box { _ = "STUB: not implemented"; return *new(Box) }

func (a Box) Size() Vec { _ = "STUB: not implemented"; return *new(Vec) }

func (a Box) Center() Vec { _ = "STUB: not implemented"; return *new(Vec) }

func (a Box) Empty() bool { _ = "STUB: not implemented"; return false }

func (a Box) Vertices() []Vec { _ = "STUB: not implemented"; return nil }

func (a Box) Union(b Box) Box { _ = "STUB: not implemented"; return *new(Box) }

func (a Box) Add(v Vec) Box { _ = "STUB: not implemented"; return *new(Box) }

func (a Box) Scale(v Vec) Box { _ = "STUB: not implemented"; return *new(Box) }

func centeredBox(center, size Vec) Box { _ = "STUB: not implemented"; return *new(Box) }

func (a Box) Contains(v Vec) bool { _ = "STUB: not implemented"; return false }

func (b Box) Canon() Box { _ = "STUB: not implemented"; return *new(Box) }
