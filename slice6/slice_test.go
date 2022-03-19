package slice6

import (
	"fmt"

	"github.com/micln/collection/simple"
)

func ExampleFilterBy() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(FilterBy(s, func(e int) bool { return e&1 == 1 }))
	fmt.Println(FilterNotBy(s, func(e int) bool { return e&1 == 1 }))
	// Output:
	// [1 3 5]
	// [2 4]
}

func ExampleGroupBy() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(GroupBy(s, func(e int) string { return simple.IfThen(e&1 == 1, "odd", "even") }))
	// Output:
	// map[even:[2 4] odd:[1 3 5]]
}

func ExampleKeyBy() {
	s := []int{1, 2, 3, 5, 5}
	fmt.Println(KeyBy(s, func(e int) int { return e*10 + e }))
	// Output:
	// map[11:1 22:2 33:3 55:5]
}

func ExampleDistinct() {
	s := []int{1, 2, 3, 3, 5}
	fmt.Println(Distinct(s))
	// Output:
	// [1 2 3 5]
}
