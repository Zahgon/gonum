package quad

type Hermite struct{}

func (h Hermite) FixedLocations(x, weight []float64, min, max float64) {
	_ = "STUB: not implemented"
	return
}

func (h Hermite) locations(x, weights []float64) { _ = "STUB: not implemented"; return }

func (h Hermite) locationsAsy(x, w []float64) { _ = "STUB: not implemented"; return }

func (h Hermite) locationsAsy0(i, n int) (x, w float64) { _ = "STUB: not implemented"; return 0, 0 }

func (h Hermite) hermpolyAsyAiry(i, n int, t float64) (valVec, dvalVec float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (h Hermite) hermiteInitialGuess(i, n int) float64 { _ = "STUB: not implemented"; return 0 }

var airyRtsExact = []float64{
	-2.338107410459762,
	-4.087949444130970,
	-5.520559828095555,
	-6.786708090071765,
	-7.944133587120863,
	-9.022650853340979,
	-10.040174341558084,
	-11.008524303733260,
	-11.936015563236262,
	-12.828776752865757,
}
