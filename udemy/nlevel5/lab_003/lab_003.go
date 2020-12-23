package main

import "fmt"

type vehicle struct {
	door  string
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	rav := truck{
		vehicle:   vehicle{"2 Door", "Red"},
		fourWheel: true,
	}

	camry := sedan{
		vehicle: vehicle{"4 Door", "Black"},
		luxury:  true,
	}

	fmt.Println(rav)
	fmt.Println(camry)
	fmt.Println(rav.door)
	fmt.Println(camry.color)
}
