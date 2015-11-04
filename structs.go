package main

import "fmt"

type Date struct {
	Year  int
	Day   int
	Month int
}

func (d Date) String() string {
	return fmt.Sprintf("%d-%d-%d", d.Year, d.Month, d.Day)
}

func main() {
	d := Date{
		Year:  2015,
		Day:   4,
		Month: 11,
	}

	fmt.Println(d)
}
