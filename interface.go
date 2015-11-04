package main

import "log"

type Speaker interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "woof"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "meow"
}

func main() {
	d := Dog{}
	c := Cat{}

	CommandToSpeak(d)
	CommandToSpeak(c)
}

func CommandToSpeak(s Speaker) {
	log.Println(s.Speak())
}
