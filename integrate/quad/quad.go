package quad

type FixedLocationer interface {
	FixedLocations(x, weight []float64, min, max float64)
}

type FixedLocationSingler interface {
	FixedLocationSingle(n, k int, min, max float64) (x, weight float64)
}

func Fixed(f func(float64) float64, min, max float64, n int, rule FixedLocationer, concurrent int) float64 {
	_ = "STUB: not implemented"
	return 0
}
