package set

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

func toSortedSlice[E constraints.Ordered](s OrderedSet[E]) (list []E) {
	list = s.ToSlice()
	slices.Sort(list)
	return
}
