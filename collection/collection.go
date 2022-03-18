package collection

import "github.com/micln/collection/iter"

type BaseCollection[E any] interface {
	Add(E)
	Contains(E) bool
	Remove(E)
	Clear()
	Size() int
	ToSlice() []E
}

type Collection[E any] interface {
	BaseCollection[E]

	iter.Iterable[E]

	// Add(E)
	AddMany(...E)
	AddAll(Collection[E])

	// Contains(E) bool
	ContainsMany(...E) bool
	ContainsAll(Collection[E]) bool

	// Remove(E)
	RemoveMany(...E)
	RemoveAll(Collection[E])

	IsEmpty(E) bool
	Size() int
	Clear()

	// ToSlice() []E
}
