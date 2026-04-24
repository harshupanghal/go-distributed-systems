package basics

import (
	"fmt"
	"math"
)

const s string = "constant"

func Constants() {

	fmt.Println(s)

	// const value can be decalred inside and outside functions
	const n = 50000

	const d = 3e20 / n
	fmt.Println(d)

	// numeric const has no type until its given one, by expplicit conversion.
	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

}
