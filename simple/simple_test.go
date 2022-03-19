package simple

import (
	"fmt"
)

func ExampleSimple() {
	fmt.Println(Max(8, 10))                 // 10
	fmt.Println(Max("8", "10"))             // 8
	fmt.Println(IfThen(true, false, true))  // false
	fmt.Println(IfThen(false, false, true)) // true
	// Output:
	// 10
	// 8
	// false
	// true
}
