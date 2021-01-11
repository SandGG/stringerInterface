package main

import (
	"fmt"
)

var bestAvg int = 0

type people interface {
	average()
}

type student struct {
	name                  string
	cal1, cal2, cal3, avg int
}

func (s *student) average() {
	s.avg = (s.cal1 + s.cal2 + s.cal3) / 3
}

func (s student) String() string {
	s.average()
	if bestAvg < s.avg {
		bestAvg = s.avg
	}
	return fmt.Sprintf("%v's average: %v", s.name, s.avg)
}

func main() {
	fmt.Println(&student{name: "Marco Diaz", cal1: 7, cal2: 8, cal3: 9})
	fmt.Println(&student{name: "Lucia Mendez", cal1: 8, cal2: 7, cal3: 10})
	fmt.Println(&student{name: "Juan Ortiz", cal1: 6, cal2: 10, cal3: 7})
	fmt.Println(&student{name: "Ana Gomez", cal1: 9, cal2: 10, cal3: 10})

	fmt.Println("The best average was:", bestAvg)
}
