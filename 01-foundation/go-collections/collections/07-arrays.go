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

	// to return length of array : use len(array)
	fmt.Println("length of array 'a' : ", len(a))

	// to declare and initialze array in one line
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("elements of array b: ", b)

	// if we omit the size when intializing an array, Go will infer it from the number of elements you provide.
	// use [...] when not declaring size
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// if we specify the size explicitly, the compile will enforce that size.
	// if a size is specified but elements entered is lower than that, then all remaining elements will be Zero.
	nums := [3]int{1, 2}
	fmt.Println(nums)

	// If you specify the index with :, the elements in between will be zeroed.
	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	// unlike in other languages like C, if we try to get address of an array, &nums, it doesn't give same result as &nums[0]

	// fmt package will recognizw it as a pointer to an array and shows the array's contents (&[1 2 3]) rather than a raw memory address.

}
