package variadicfunctions

import "fmt"

func VariadicFunctions() {
	// those functions which can be called with any number of trailing arguments. for eg.: fmt.Println

	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)

}

// variadic function
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
