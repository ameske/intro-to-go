package main

import "log"

func main() {
	x := MakeNumberFunc(42)
	x()

	y := MakeNumberFunc(24)
	y()
}

func MakeNumberFunc(x int) func() {
	return func() {
		log.Println(x)
	}
}
