package collection_test

import (
	"fmt"

	"github.com/micln/collection/collection"
	"github.com/micln/collection/set"
)

func ExampleCollectionSet() {
	set1 := set.NewHashSet("e", "c", "b", "b", "c")
	fmt.Println(set.ToSortedSlice[string](set1)) // [b c e]

	x := collection.Xof[string](set1)
	x.AddMany("a", "t")
	fmt.Println(set.ToSortedSlice[string](x)) // [a b c e t]
	// Output:
	// [b c e]
	// [a b c e t]
}
