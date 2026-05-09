package recursion

import "fmt"

func Recursion() {
	// has 3 things in it :
	// 1. base condition
	// 2. function calling
	// 3. operation

	fmt.Println(Fact(5))

	// anonymous fucntions can also be recursive but this requires explicitly decalring a variable with var to store the function before it's defined

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(10))

}

func Fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * Fact(n-1)
}
