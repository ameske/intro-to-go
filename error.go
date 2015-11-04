package main

import (
	"log"
	"os"
)

func main() {
	fd, err := os.Open("i/cant/open/this/here.txt")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("WHAT? IMPOSSIBLE!?!")
}
