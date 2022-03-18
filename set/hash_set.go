package set

type (
	some = struct{}

	HashSet[E comparable] struct {
		Trait[E]

		store map[E]some
	}
)

var _ ComparableSet[int] = (*HashSet[int])(nil)

func NewHashSet[E comparable](elements ...E) (hashSet *HashSet[E]) {
	hashSet = &HashSet[E]{store: make(map[E]some, len(elements))}
	hashSet.Trait = TraitWith[E](hashSet)

	for _, e := range elements {
		hashSet.store[e] = some{}
	}
	return
}

func (hashSet *HashSet[E]) Add(e E) {
	hashSet.store[e] = some{}
}

func (hashSet *HashSet[E]) Clear() {
	hashSet.store = make(map[E]some)
}

func (hashSet *HashSet[E]) Contains(e E) bool {
	_, ok := hashSet.store[e]
	return ok
}

func (hashSet *HashSet[E]) Remove(e E) {
	delete(hashSet.store, e)
}

func (hashSet *HashSet[E]) Size() int {
	return len(hashSet.store)
}

func (hashSet *HashSet[E]) ToSlice() []E {
	list := make([]E, 0, len(hashSet.store))
	for e := range hashSet.store {
		list = append(list, e)
	}

	return list
}
