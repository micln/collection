package collection

type X[E any] struct {
	Collection[E]
}

var _ CollectionX[int] = (*X[int])(nil)

func Xof[E any](base Collection[E]) X[E] {
	return X[E]{Collection: base}
}

func (tr *X[E]) AddMany(elems ...E) {
	for _, e := range elems {
		tr.Add(e)
	}
}

func (tr *X[E]) AddAll(c Collection[E]) {
	tr.AddMany(c.ToSlice()...)
}

func (tr *X[E]) ContainsMany(elems ...E) bool {
	for _, e := range elems {
		if !tr.Contains(e) {
			return false
		}
	}
	return true
}

func (tr *X[E]) ContainsAll(c Collection[E]) bool {
	return tr.ContainsMany(c.ToSlice()...)
}

func (tr *X[E]) RemoveMany(elems ...E) {
	for _, e := range elems {
		tr.Remove(e)
	}
}

func (tr *X[E]) RemoveAll(c Collection[E]) {
	tr.RemoveMany(c.ToSlice()...)
}

func (tr *X[E]) IsEmpty() bool {
	return tr.Size() == 0
}
