package kdtree

var (
	_ Interface  = Points(nil)
	_ Comparable = Point(nil)
)

type Point []float64

func (p Point) Compare(c Comparable, d Dim) float64 { _ = "STUB: not implemented"; return 0 }

func (p Point) Dims() int { _ = "STUB: not implemented"; return 0 }

func (p Point) Distance(c Comparable) float64 { _ = "STUB: not implemented"; return 0 }

func (p Point) Extend(b *Bounding) *Bounding { _ = "STUB: not implemented"; return nil }

type Points []Point

func (p Points) Bounds() *Bounding { _ = "STUB: not implemented"; return nil }

func (p Points) Index(i int) Comparable         { _ = "STUB: not implemented"; return *new(Comparable) }
func (p Points) Len() int                       { _ = "STUB: not implemented"; return 0 }
func (p Points) Pivot(d Dim) int                { _ = "STUB: not implemented"; return 0 }
func (p Points) Slice(start, end int) Interface { _ = "STUB: not implemented"; return *new(Interface) }

type Plane struct {
	Dim
	Points
}

const randoms = 100

func (p Plane) Less(i, j int) bool              { _ = "STUB: not implemented"; return false }
func (p Plane) Pivot() int                      { _ = "STUB: not implemented"; return 0 }
func (p Plane) Slice(start, end int) SortSlicer { _ = "STUB: not implemented"; return *new(SortSlicer) }
func (p Plane) Swap(i, j int)                   { _ = "STUB: not implemented"; return }
