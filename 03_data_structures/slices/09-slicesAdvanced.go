package collections

import "fmt"

func SlicesAdvanced() {

	// Internal Representation of Slices:

	// - a slice id represented by a struct that contains a pointer to the underlying array, the length of slice and its capacity

	// type slice struct {
	// 	ptr *ElementType
	// 	len int
	// 	cap int
	// }

	// slices have the pointer to the array and passing slices to a function feels like passing by reference as the values are not copied but the slice struct itself is passed by value
	nums_array := [...]int{1, 2, 3}
	nums_slice := []int{1, 2, 3}
	modify(nums_array, nums_slice)
	fmt.Println(nums_array) // 1 2 3     -- becuse arrays are passed by reference
	fmt.Println(nums_slice) // 99 2 3    --  because slices are passed by value

	// Copy a Slice
	// copying a slice creates a new slice with the same elements. can be done using built-in 'copy' function
	s1 := []int{1, 2, 3}
	s2 := make([]int, len(s1))
	copy(s2, s1) // copy(copyInto, copyFrom)
	fmt.Println(s2)

	// NOTES :

	// - the capacity of destination slice is not automatically adjusted. if the destination slice gas a smaller capacity than the source slice, it will only copy up to the capacity of the destination slice

	// - if the source slice is nil, the copy function will not panic, but the destination slice will remain unchanged

	// - if the source and destination slices overlap, the behavior is undefined. To avoid this, make sure to copy to a separate slice.

	// Multi-Dimensional Slices
	matrix := [][]int{
		{7, 8, 9},
		{4, 5, 6},
		{1, 2, 3},
	}

	fmt.Println(matrix)

	rows := 3
	cols := 4
	matrix2 := make([][]int, rows)
	for i := range matrix2 {
		matrix2[i] = make([]int, cols)
	}

	fmt.Println(matrix2)

	//matrix are useful when we need flexible, dynamic grids of data.
}

func modify(s1 [3]int, s2 []int) {
	s1[0] = 99
	s2[0] = 99
}
