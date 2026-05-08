package closures

import "fmt"

func Closures() {
	// Go supports anonymous functions, which can form closures.
	// useful when we want to define a functions inline without having to name it.

	// here we will call intSeq, assigning its results to another function nextInt, this func will capture its own i value,
	// which will be updated each time we call nextInt

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// the state is unique to one particular function, to confirm we will create and test new one.
	newInt := intSeq()
	fmt.Println(newInt())
	fmt.Println(newInt())

}

// This function returns another function, which we define anonymously in its body.
// the returned fucntion closes over the variable i to form a closure.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
