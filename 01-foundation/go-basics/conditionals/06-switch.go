package conditionals

import (
	"fmt"
	"time"
)

func Switch() {

	// basic switch

	i := 2

	fmt.Print("Write ", i, " as ")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("out of scope")
	}

	// you can use commas to seperate multiple expressions in the same case statement.
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("its the weekend")
	default:
		fmt.Println("its the weekday")
	}

	// switch without an expression (switch {}) behaves like an if/else-if chain.
	// Use switch when:
	// - checking multiple cases of the same value (cleaner than long if/else-if)
	// - handling many conditional branches for better readability
	// Use if/else for:
	// - simple or few conditions
	// - cases where logic is not structured around distinct cases
	//
	// Choice is based on readability and structure, not constant vs non-constant expressions.

	// type switch compares type instead of values.
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("i am a boolean")
		case int:
			fmt.Println("i am an integer")
		default:
			fmt.Printf("Don't know who am i %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("goLang")
}
