package main

import (
	"fmt"
)

func main() {
	// 1. Create an empty map of strings to integer
	// 2. Implement and call method which assigns key-value pairs to map
	// 3. Print contents of map
	// 4. Expected output: `{"one":1,"two":2,"three":3}`

	// 1.
	menu := map[string]int{} // an empty map literal

	// 2.
	assign(menu, "one", 1)
	assign(menu, "two", 2)

	fmt.Printf("menu: %v", menu)

}

// 2.
func assign(m map[string]int, key string, value int) {
	m[key] = value
}
