package generic_types

type Pair[T comparable] struct {
	A, B T
}

type DistinctPair[T, K comparable] struct {
	A T
	B K
}
