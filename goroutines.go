package main

import (
	"log"
)

func main() {
	input := make(chan int)
	output := make(chan int)

	go Worker(input, output)
	go Printer(output)

	for i := 0; i < 10; i++ {
		input <- i
	}
}

func Worker(in chan int, out chan int) {
	for x := range in {
		out <- x + 2
	}

	log.Println("Worker done!")
}

func Printer(values chan int) {
	for x := range values {
		log.Println("Received", x)
	}

	log.Println("Printer done!")
}
