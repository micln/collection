package kvk

import "golang.org/x/exp/constraints"

func Large[E constraints.Ordered](a, b E) bool {
	return a > b
}

func Less[E constraints.Ordered](a, b E) bool {
	return a < b
}

func Self[E any](e E) E {
	return e
}
