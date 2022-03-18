package set

import (
	"golang.org/x/exp/constraints"

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
