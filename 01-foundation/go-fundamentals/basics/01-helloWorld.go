//package _name_
// package declares/describes
//1. namespace - scope boundary

//2. Compilation unit
//Go compiles packages, not individual files
//All files in a package are built together

//3. code organization and reusability

// package main Marks an executable program, Must contain:
// func main() {}

package basics

// importing necessary packages, modules, libraries, code files etc.
import "fmt"

// the entry package (package main) requires a main() function.
// every executable program must have a main method
// in other files, libraries,and packages, no main() needed

func SayHello() {
	fmt.Println("Hello World, i ran successfully")
}

// ..... COMMANDS ......
// to run a Go file :
// go run    --  go run hello-world.go

// to build the Go file
// go build   -- go build hello-world.go

// go build will generate a binary executable file that we can run later as executable.
// ./hello-world
