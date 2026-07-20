package kdtree

import (
	"sort"
)

func Partition(list sort.Interface, pivot int) int { _ = "STUB: not implemented"; return 0 }

type SortSlicer interface {
	sort.Interface
	Slice(start, end int) SortSlicer
}

func Select(list SortSlicer, k int) int { _ = "STUB: not implemented"; return 0 }

func MedianOfMedians(list SortSlicer) int { _ = "STUB: not implemented"; return 0 }

func MedianOfRandoms(list SortSlicer, n int) int { _ = "STUB: not implemented"; return 0 }
