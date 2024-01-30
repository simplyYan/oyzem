# oyzem
memoizing has never been easier

```go
package main

import (
	"fmt"
	"github.com/simplyYan/oyzem"
)

// Function to calculate the factorial of a number
func calculateFactorial(n int) int {
	fmt.Printf("Calculating factorial of %d...\n", n)
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	// Create a new Memoizer instance
	memoizer := oyzem.New()

	// Memoize the calculateFactorial function
	memoizedFactorial, _ := memoizer.Memoize(calculateFactorial)

	// Calculate factorial of 5 (first invocation)
	result1, _ := memoizer.Run(memoizedFactorial, 5)
	fmt.Println("Factorial Result:", result1)

	// Calculate factorial of 5 again (result retrieved from cache)
	result2, _ := memoizer.Run(memoizedFactorial, 5)
	fmt.Println("Factorial Result (from cache):", result2)

	// Calculate factorial of 7 (new calculation)
	result3, _ := memoizer.Run(memoizedFactorial, 7)
	fmt.Println("Factorial Result:", result3)
}
```
