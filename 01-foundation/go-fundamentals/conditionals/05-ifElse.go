package conditionals

import "fmt"

func IfElse() {

	// no need of paranthese around conditions, but need curly braces

	// same as other languages, so no need of explanation and examples

	// we can initialize a variable in if statement and it will be available in all branches of if-else
	// after intialization, we can also use that variable in condition

	// ................................................................................
	// NOTE : GO REQUIRES IDENTATION OF CODE BLOCKS, OTHERWISE IT WILL THROW ERROR
	// ................................................................................

	if num := 10; num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// There is NO TERNARY OPERATOR in Go, so we have to use if-else for all conditional logic
}
