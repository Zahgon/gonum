package functions

type MinimalSurface struct {
	bottom, top  []float64
	left, right  []float64
	origin, step [2]float64
}

func NewMinimalSurface(nx, ny int) *MinimalSurface { _ = "STUB: not implemented"; return nil }

func (ms *MinimalSurface) Func(x []float64) (area float64) { _ = "STUB: not implemented"; return 0 }

func (ms *MinimalSurface) Grad(grad, x []float64) []float64 { _ = "STUB: not implemented"; return nil }

func (ms *MinimalSurface) InitX() []float64 { _ = "STUB: not implemented"; return nil }

func (ms *MinimalSurface) ExactX() []float64 { _ = "STUB: not implemented"; return nil }

func (ms *MinimalSurface) ExactSolution(x, y float64) float64 { _ = "STUB: not implemented"; return 0 }

func (ms *MinimalSurface) Dims() (nx, ny int) { _ = "STUB: not implemented"; return 0, 0 }

func (ms *MinimalSurface) Steps() (hx, hy float64) { _ = "STUB: not implemented"; return 0, 0 }

func (ms *MinimalSurface) x(i int) float64 { _ = "STUB: not implemented"; return 0 }

func (ms *MinimalSurface) y(j int) float64 { _ = "STUB: not implemented"; return 0 }

func (ms *MinimalSurface) at(i, j int, x []float64) float64 { _ = "STUB: not implemented"; return 0 }

func (ms *MinimalSurface) index(i, j int) int { _ = "STUB: not implemented"; return 0 }

func (ms *MinimalSurface) initBoundary(b []float64, startX, startY, hx, hy float64) {
	_ = "STUB: not implemented"
	return
}
