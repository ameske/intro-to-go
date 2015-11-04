package main

import (
	"log"
)

func main() {
	input := make(chan int)
	output := make(chan int)
	wdone := make(chan bool)
	pdone := make(chan bool)

	go Worker(input, output, wdone)
	go Printer(output, pdone)

	for i := 0; i < 10; i++ {
		input <- i
	}

	close(input)
	<-wdone
	close(output)
	<-pdone
}

func Worker(in chan int, out chan int, done chan bool) {
	for x := range in {
		out <- x + 2
	}

	log.Println("Worker done!")

	done <- true
}

func Printer(values chan int, done chan bool) {
	for x := range values {
		log.Println("Received", x)
	}

	log.Println("Printer done!")

	done <- true
}
