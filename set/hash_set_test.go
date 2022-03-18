package set

import (
	"fmt"
)

func ExampleHashSet() {
	set1 := NewHashSet("e", "c", "b", "b", "c")
	fmt.Println(toSortedSlice[string](set1)) // [b c e]

	set1.AddMany("t", "a")
	fmt.Println(toSortedSlice[string](set1)) // [a b c e t]

	fmt.Println(set1.Contains("b")) // true
	set1.Remove("b")
	fmt.Println(toSortedSlice[string](set1)) // [a c e t]
	fmt.Println(set1.Contains("b"))          // false

	fmt.Println(set1.Size()) // 4
	set1.Clear()
	fmt.Println(set1.Size()) // 0

	// Output:
	// [b c e]
	// [a b c e t]
	// true
	// [a c e t]
	// false
	// 4
	// 0
}
