//go:generate ./single_precision.bash

package gonum

type Implementation struct{}

const (
	blockSize   = 64
	minParBlock = 4
)

func blocks(dim, bsize int) int { _ = "STUB: not implemented"; return 0 }

func dcabs1(z complex128) float64 { _ = "STUB: not implemented"; return 0 }

func scabs1(z complex64) float32 { _ = "STUB: not implemented"; return 0 }
