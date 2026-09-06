package combin

const (
	errNegInput             = "combin: negative input"
	badSetSize              = "combin: n < k"
	badInput                = "combin: wrong input slice length"
	errNonpositiveDimension = "combin: non-positive dimension"
)

func Binomial(n, k int) int { _ = "STUB: not implemented"; return 0 }

func GeneralizedBinomial(n, k float64) float64 { _ = "STUB: not implemented"; return 0 }

func LogGeneralizedBinomial(n, k float64) float64 { _ = "STUB: not implemented"; return 0 }

type CombinationGenerator struct {
	n         int
	k         int
	previous  []int
	remaining int
}

func NewCombinationGenerator(n, k int) *CombinationGenerator { _ = "STUB: not implemented"; return nil }

func (c *CombinationGenerator) Next() bool { _ = "STUB: not implemented"; return false }

func (c *CombinationGenerator) Combination(dst []int) []int { _ = "STUB: not implemented"; return nil }

func Combinations(n, k int) [][]int { _ = "STUB: not implemented"; return nil }

func nextCombination(s []int, n, k int) { _ = "STUB: not implemented"; return }

func CombinationIndex(comb []int, n, k int) int { _ = "STUB: not implemented"; return 0 }

func IndexToCombination(dst []int, idx, n, k int) []int { _ = "STUB: not implemented"; return nil }

func Cartesian(lens []int) [][]int { _ = "STUB: not implemented"; return nil }

func Card(dims []int) int { _ = "STUB: not implemented"; return 0 }

func NewCartesianGenerator(lens []int) *CartesianGenerator { _ = "STUB: not implemented"; return nil }

type CartesianGenerator struct {
	lens []int
	rows int
	idx  int
}

func (g *CartesianGenerator) Next() bool { _ = "STUB: not implemented"; return false }

func (g *CartesianGenerator) Product(dst []int) []int { _ = "STUB: not implemented"; return nil }

func IdxFor(sub, dims []int) int { _ = "STUB: not implemented"; return 0 }

func SubFor(sub []int, idx int, dims []int) []int { _ = "STUB: not implemented"; return nil }

func NumPermutations(n, k int) int { _ = "STUB: not implemented"; return 0 }

func Permutations(n, k int) [][]int { _ = "STUB: not implemented"; return nil }

type PermutationGenerator struct {
	n           int
	k           int
	nPerm       int
	idx         int
	permutation []int
}

func NewPermutationGenerator(n, k int) *PermutationGenerator { _ = "STUB: not implemented"; return nil }

func (p *PermutationGenerator) Next() bool { _ = "STUB: not implemented"; return false }

func (p *PermutationGenerator) Permutation(dst []int) []int { _ = "STUB: not implemented"; return nil }

func PermutationIndex(perm []int, n, k int) int { _ = "STUB: not implemented"; return 0 }

type sortInts struct {
	data []int
	idx  []int
}

func (s sortInts) Len() int { _ = "STUB: not implemented"; return 0 }

func (s sortInts) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (s sortInts) Swap(i, j int) { _ = "STUB: not implemented"; return }

func IndexToPermutation(dst []int, idx, n, k int) []int { _ = "STUB: not implemented"; return nil }

func equalPermutationIndex(perm []int) int { _ = "STUB: not implemented"; return 0 }

func indexToEqualPermutation(dst []int, idx int) { _ = "STUB: not implemented"; return }

func factorial(a int) int { _ = "STUB: not implemented"; return 0 }
