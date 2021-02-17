package main

import (
	"fmt"
)

type customErr struct {
	err string
}

func (c customErr) Error() string {
	return c.err
}

func main() {
	err := customErr{"CustomErr"}
	foo(err)
}

func foo(e error) {
	fmt.Println(e)
}
