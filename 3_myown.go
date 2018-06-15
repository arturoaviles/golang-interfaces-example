package main

import "fmt"

type Cheetah struct {
	Distance int
	Time     int
}

func (c Cheetah) Velocity() int {
	return c.Distance / c.Time
}

type Lion struct {
	Distance int
	Time     int
}

func (l Lion) Velocity() int {
	return l.Distance / l.Time
}

func IsItFaster(a, b Animal) bool {
	return a.Velocity() > b.Velocity()
}

type Animal interface {
	Velocity() int
}

func main() {
	c := Cheetah{100, 1}
	l := Lion{60, 1}

	fmt.Println(IsItFaster(c, l))
}
