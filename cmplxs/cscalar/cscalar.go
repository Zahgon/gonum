package cscalar

func EqualWithinAbs(a, b complex128, tol float64) bool { _ = "STUB: not implemented"; return false }

const minNormalFloat64 = 0x1p-1022

func EqualWithinRel(a, b complex128, tol float64) bool { _ = "STUB: not implemented"; return false }

func EqualWithinAbsOrRel(a, b complex128, absTol, relTol float64) bool {
	_ = "STUB: not implemented"
	return false
}

func ParseWithNA(s, missing string) (value complex128, weight float64, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func Round(x complex128, prec int) complex128 { _ = "STUB: not implemented"; return 0 }

func RoundEven(x complex128, prec int) complex128 { _ = "STUB: not implemented"; return 0 }

func Same(a, b complex128) bool { _ = "STUB: not implemented"; return false }
