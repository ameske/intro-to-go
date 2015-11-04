package main

import "fmt"

type Date struct {
	year  int
	day   int
	month int
}

func (d Date) String() string {
	return fmt.Sprintf("%d-%d-%d", year, month, day)
}
