package transform

import (
	"errors"
)

type DegenerateInputError float64

func (e DegenerateInputError) Error() string { _ = "STUB: not implemented"; return "" }

var ErrFactorizationFailed = errors.New("transform: factorization failed")
