package main

import (
	"fmt"
)

func main() {

	// 1. Create slice of integers with len and capacity 5
	// 2. Append the value '1' to slice
	// 3. Print slice contents
	// 4. Set value of 3rd element to value '5'
	// 5. Print slice contents
	// 6. Remove 3rd element of slice
	// 7. Print slice contents

	// 1.
	sliceInt := make([]int, 5)
	fmt.Println(len(sliceInt), cap(sliceInt))

	// 2.
	sliceInt = append(sliceInt, 1)
	fmt.Println(sliceInt)

	// 4.
	sliceInt[2] = 5
	fmt.Println(sliceInt)

	// 6.
	sliceInt = append(sliceInt[:2], sliceInt[3:]...)
	// 7.
	fmt.Println(sliceInt)

}
