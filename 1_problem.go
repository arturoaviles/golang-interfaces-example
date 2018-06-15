package main

import "fmt"

// Rock resembles a typical rock with mass and volume
type Rock struct {
	Mass   int
	Volume int
}

// Density returns the rock density (mass/volume)
func (r Rock) Density() int {
	return r.Mass / r.Volume
}

// IsItDenser returns true if a density is bigger than b density
func IsItDenser(a, b Rock) bool {
	return a.Density() > b.Density()
}

type Water struct {
	Mass   int
	Volume int
}

func (w Water) Density() int {
	return w.Mass / w.Volume
}

func main() {
	r1 := Rock{10, 1}
	r2 := Rock{12, 2}

	fmt.Println(IsItDenser(r1, r2))

	w1 := Water{44, 3}

	fmt.Println(IsItDenser(w1, r1))
}
