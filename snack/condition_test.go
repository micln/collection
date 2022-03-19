package snack

import (
	"fmt"
)

func ExampleIfThen() {
	fmt.Println(Max(8, 10))                 // 10
	fmt.Println(Max("8", "10"))             // 8
	fmt.Println(Min("8", "10"))             // 8
	fmt.Println(IfThen(true, false, true))  // false
	fmt.Println(IfThen(false, false, true)) // true
	// Output:
	// 10
	// 8
	// 10
	// false
	// true
}
