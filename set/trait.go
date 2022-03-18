package set

import "github.com/micln/collection/collection"

type Trait[E any] struct {
	collection.Trait[E]
}

func TraitWith[E any](instance collection.BaseCollection[E]) Trait[E] {
	return Trait[E]{collection.TraitWith(instance)}
}
