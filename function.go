package main

import (
	"fmt"
)

func main() {
	x := myFunction()

	fmt.Printf("The answer to the meaning of life is %d\n", x)
}

func myFunction() int {
	return 42
}
