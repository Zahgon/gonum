package cscalar

func parse(s string) (complex128, error) { _ = "STUB: not implemented"; return 0, nil }

func floatPart(s string) (beg, end int, part uint, err error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, nil
}

func isSign(r rune) bool { _ = "STUB: not implemented"; return false }

func isDigit(r rune) bool { _ = "STUB: not implemented"; return false }

func isExponent(r rune) bool { _ = "STUB: not implemented"; return false }

func isDot(r rune) bool { _ = "STUB: not implemented"; return false }

type parseError struct {
	string string
	state  int
	rune   rune
}

func (e parseError) Error() string { _ = "STUB: not implemented"; return "" }
