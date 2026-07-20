//go:generate ./generate_64bit.sh

package card

import (
	"reflect"
	"sync"
)

const (
	w32 = 32
	w64 = 64
)

func alpha(m uint64) float64 { _ = "STUB: not implemented"; return 0 }

var alphaValues = [...]float64{
	16: 0.673,
	32: 0.697,
	64: 0.709,
}

func linearCounting(m, v float64) float64 { _ = "STUB: not implemented"; return 0 }

func typeNameOf(v interface{}) string { _ = "STUB: not implemented"; return "" }

var hashes sync.Map

type userType struct {
	fn  reflect.Value
	typ reflect.Type
}

func RegisterHash(fn interface{}) { _ = "STUB: not implemented"; return }
