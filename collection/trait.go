package collection

type Trait[E any] struct {
	instance BaseCollection[E]
}

func TraitWith[E any](instance BaseCollection[E]) Trait[E] {
	return Trait[E]{instance: instance}
}

func (tr *Trait[E]) AddMany(elems ...E) {
	for _, e := range elems {
		tr.instance.Add(e)
	}
}

func (tr *Trait[E]) AddAll(c Collection[E]) {
	tr.AddMany(c.ToSlice()...)
}

func (tr *Trait[E]) ContainsMany(elems ...E) bool {
	for _, e := range elems {
		if !tr.instance.Contains(e) {
			return false
		}
	}
	return true
}

func (tr *Trait[E]) ContainsAll(c Collection[E]) bool {
	return tr.ContainsMany(c.ToSlice()...)
}

func (tr *Trait[E]) RemoveMany(elems ...E) {
	for _, e := range elems {
		tr.instance.Remove(e)
	}
}

func (tr *Trait[E]) RemoveAll(c Collection[E]) {
	tr.RemoveMany(c.ToSlice()...)
}

func (tr *Trait[E]) IsEmpty() bool {
	return tr.instance.Size() == 0
}
