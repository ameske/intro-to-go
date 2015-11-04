package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fd, err := os.Open("i/cant/open/this/here.txt")
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadAll(fd)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(b)
}
