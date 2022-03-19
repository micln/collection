package group6

import (
	"fmt"

	"github.com/micln/collection/kvk"
)

func ExampleGroup() {
	g := map[string][]int{
		"alex": {84, 82, 83},
		"bob":  {94, 93, 95},
		"tom":  {70, 80, 93, 92},
	}

	fmt.Println(Max(g))                       // map[alex:84 bob:95 tom:93]
	fmt.Println(Min(g))                       // map[alex:82 bob:93 tom:70]
	fmt.Println(Avg(g))                       // map[alex:83 bob:94 tom:83.75]
	fmt.Println(TopBy(g, kvk.Less[int]))      // map[alex:82 bob:93 tom:70]
	fmt.Println(TopNBy(g, 2, kvk.Large[int])) // map[alex:[84 83] bob:[95 94] tom:[93 92]]

	// Output:
	// map[alex:84 bob:95 tom:93]
	// map[alex:82 bob:93 tom:70]
	// map[alex:83 bob:94 tom:83.75]
	// map[alex:82 bob:93 tom:70]
	// map[alex:[84 83] bob:[95 94] tom:[93 92]]
}
