package loops

import "fmt"

func Loops() {

	// for loops in Go doesn't need paranthese around conditions and post statements, but they do need curly braces

	// basic loop
	i := 1
	for i <= 3 {
		fmt.Print(i, " ")
		i++
	}

	fmt.Println()
	// for loop with init, condition and post statement
	for j := 0; j < 5; j++ {
		fmt.Print(j, " ")
	}

	fmt.Println()
	// for loop with range
	for i := range 3 {
		fmt.Println("range", i, " ")
	}

	fmt.Println()
	// for loop with string range
	for i, c := range "hello" {
		fmt.Printf("index: %d, char: %c\n", i, c)
	}
}
