package main

import "fmt"

func main() {
	fmt.Println("blah blah blah this is my test code")

	foo()

	fmt.Println("something else")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Println("I'm in foo")
}

//control flow
// 1) sequence
// 2) loop; iterative
// 3) conditional
