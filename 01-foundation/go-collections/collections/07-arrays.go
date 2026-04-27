package collections

import "fmt"

func Arrays() {

	// by default arrays are zero-valued
	var a [5]int
	fmt.Println("emp:", a) // 0 0 0 0 0

	// to set value
	// array[index] = value
	a[4] = 3

	// to get value
	// array[index]
	fmt.Println("value at index 4 : ", a[4])

	// to return lenght of array : use len(array)
	fmt.Println("lenth of array 'a' : ", len(a))

	// to declare and initialze array in one line
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("elements of array b: ", b)

	// You can also have the compiler count the number of elements for you with ...
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// If you specify the index with :, the elements in between will be zeroed.
	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)
}
