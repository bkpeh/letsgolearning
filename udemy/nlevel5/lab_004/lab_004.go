package main

import "fmt"

func main() {
	x := struct {
		name string
		ID   int
	}{
		name: "Pete",
		ID:   123,
	}

	fmt.Println(x)
	fmt.Println(x.name)
	fmt.Println(x.ID)

}
