package scalar

func EqualWithinAbs(a, b, tol float64) bool { _ = "STUB: not implemented"; return false }

const minNormalFloat64 = 0x1p-1022

func EqualWithinRel(a, b, tol float64) bool { _ = "STUB: not implemented"; return false }

func EqualWithinAbsOrRel(a, b, absTol, relTol float64) bool {
	_ = "STUB: not implemented"
	return false
}

func EqualWithinULP(a, b float64, ulp uint) bool { _ = "STUB: not implemented"; return false }

func ulpDiff(a, b uint64) uint64 { _ = "STUB: not implemented"; return 0 }

const (
	nanBits = 0x7ff8000000000000
	nanMask = 0xfff8000000000000
)

func NaNWith(payload uint64) float64 { _ = "STUB: not implemented"; return 0 }

func NaNPayload(f float64) (payload uint64, ok bool) { _ = "STUB: not implemented"; return 0, false }

func ParseWithNA(s, missing string) (value, weight float64, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func Round(x float64, prec int) float64 { _ = "STUB: not implemented"; return 0 }

func RoundEven(x float64, prec int) float64 { _ = "STUB: not implemented"; return 0 }

func Same(a, b float64) bool { _ = "STUB: not implemented"; return false }
