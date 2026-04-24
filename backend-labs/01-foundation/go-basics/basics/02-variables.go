package basics

import "fmt"

func Variable() {

	// var can declare more than one variable, the Go compiler will automatically infer the tyype of initialized variables
	var a, b int = 1, 2
	fmt.Println(a + b)

	// variables declared without an initialization are 0 valued.
	var c int
	fmt.Println(c)

	// + operator is used to contactenate 2 strings
	var d = "go"
	var e = "lang"
	fmt.Println(d + e)

	// := is a shorthand to declare and initalize a variable, but this syntax is only available inside functions
	f := "golang"
	fmt.Println(f)
}
