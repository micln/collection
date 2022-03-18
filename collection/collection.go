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

	AddMany(...E)
	AddAll(Collection[E])

	ContainsMany(...E) bool
	ContainsAll(Collection[E]) bool

	RemoveMany(...E)
	RemoveAll(Collection[E])

	IsEmpty(E) bool
}
