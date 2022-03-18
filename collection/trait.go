package collection

import (
	"encoding/json"

	"github.com/micln/collection/iter"
)

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

func (tr *Trait[E]) IsEmpty(e E) bool {
	return tr.instance.Size() == 0
}

func (tr *Trait[E]) Iterator() iter.Iterator[E] {
	return iter.SliceIteratorOf(tr.instance.ToSlice())
	// return iter.NewSliceListIterator(tr.instance.ToSlice())
}

func (tr *Trait[E]) Marshal() ([]byte, error) {
	return json.Marshal(tr.instance.ToSlice())
}

func (tr *Trait[E]) Unmarshal(data []byte) error {
	var list []E
	err := json.Unmarshal(data, &list)
	if err != nil {
		return err
	}

	tr.instance.Clear()
	for _, e := range list {
		tr.instance.Add(e)
	}

	return nil
}

func (tr *Trait[E]) String() string {
	s, _ := tr.Marshal()
	return string(s)
}

func (tr *Trait[E]) GoString() string {
	return tr.String()
}
