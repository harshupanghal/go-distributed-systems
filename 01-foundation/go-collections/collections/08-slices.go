package collections

import "fmt"

func ArraySlices() {

	// arrays are fixed size so Go introduced flexible, dynamic sequences built on top of arrays.
	// think of them as view into arrays, like views in sql.

	// A slice keeps three things :
	// Pointer : a reference to the underlying array
	// Length : the number of elements in the slice
	// Capacity : the max number of elements the slice can hold (always greater than or equal to the length)

	// arrays are mainly useful when you need a fixed size known at compile time, for almost everything else, slices are the go to choice

	// declaring a slice :
	var s []int         // slice of integers
	fmt.Println(s)      // prints slice
	fmt.Println(len(s)) // length of slice
	fmt.Println(cap(s)) // capacity

	// until this point we are only declaring a slice, which means we have introduced a variable s of type "slice of int" but we haven't given it any backing array, it is nil and doesn't point to any actual storage. That's why its length and capacity are both zero, until we allocate or append to it.

	// we can declare a slice using var s[] int{}, which intializes the slice with zero elements,
	// but we can't create an empty array using this syntax : var s[...] int{}. this latter is invalid in Go,
	// we can't use [...] with var and an empty intializer.

	// Allocate (make)

	s2 := make([]int, 3) // length 3, capacity 3
	fmt.Println(s2)

	s3 := make([]int, 3, 5)
	fmt.Println(s3)
	fmt.Println(len(s3))
	fmt.Println(cap(s3))

	// for slices, make does three things :
	// 1. allocates an array of the given size (either the length you specify or the capacity if you provide both)
	// 2. creates a slice header (pointer, length, capacity) that points to that array
	// 3. returns the slice header , ready to use.

	// APPEND

	// slices can grow dynamically
	// built-in 'append' function that takes an existing slice and one or more new elements and returns a new slice with those elements added.

	fmt.Println("...append ......")
	s4 := make([]int, 3, 5)
	fmt.Println(s4)
	fmt.Println(len(s4))
	fmt.Println(cap(s4))

	s4 = append(s4, 1)
	fmt.Println(s4)
	fmt.Println(len(s4))
	fmt.Println(cap(s4))

	s4 = append(s4, 2, 3)
	fmt.Println(s4)
	fmt.Println(len(s4))
	fmt.Println(cap(s4))

	// if there is enough capacity, append just write into the existng array. If not Go automatically allocates a new larger array, copies the old elements over and add new values.
	// this can lead to performance overhead due to need for memory allocation and copying.
	// it's good practice to preallocate slices with an appropriate capacity when you know the size in advance.

	// SLICING The SLICES

	// we can create a new slice by slicing an existing one.
	// syntax : slice[low:high]
	// low -  inclusive ; includes index low, excludes index high
	// high - exclusive ; result contains elements from low to high-1
	// if low is omitted -  defaults to 0
	// if high is omitted - defaults to length of slice

	s5 := []int{1, 2, 3, 4, 5}
	s6 := s5[1:4] // 2 3 4
	s7 := s5[:3]  // 1 2 3
	s8 := s5[2:]  // 3 4 5
	fmt.Println("......slices........")
	fmt.Println(s5)
	fmt.Println(s6)
	fmt.Println(s7)
	fmt.Println(s8)

	// if 2 slices share the same underlying array, changes to elements of one slice will be reflected in other. This is because both slices point to the same data in memory.

	s6[1] = 6
	fmt.Println(s6)
	fmt.Println(s7)

}
