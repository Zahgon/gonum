//go:build safe
// +build safe

package mat

import "reflect"

var sizeOfFloat64 = int(reflect.TypeOf(float64(0)).Size())

func offset(a, b []float64) int { _ = "STUB: not implemented"; return 0 }

var sizeOfComplex128 = int(reflect.TypeOf(complex128(0)).Size())

func offsetComplex(a, b []complex128) int { _ = "STUB: not implemented"; return 0 }
