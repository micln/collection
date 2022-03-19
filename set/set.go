package set

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"

	"github.com/micln/collection/collection"
)

type (
	Set[E any] interface {
		collection.Collection[E]
	}

	ComparableSet[E comparable] interface {
		Set[E]
	}

	OrderedSet[E constraints.Ordered] interface {
		Set[E]
	}
)

type X[E any] struct {
	collection.X[E]
}

func Xof[E any](base collection.Collection[E]) X[E] {
	return X[E]{collection.Xof(base)}
}

func ToSortedSlice[E constraints.Ordered](s OrderedSet[E]) (list []E) {
	list = s.ToSlice()
	slices.Sort(list)
	return
}
