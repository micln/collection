package collection

type Collection[E any] interface {
	Add(E)
	Contains(E) bool
	Remove(E)
	Clear()
	Size() int
	ToSlice() []E
}

type CollectionX[E any] interface {
	Collection[E]

	AddMany(...E)
	AddAll(Collection[E])

	ContainsMany(...E) bool
	ContainsAll(Collection[E]) bool

	RemoveMany(...E)
	RemoveAll(Collection[E])

	IsEmpty() bool
}
