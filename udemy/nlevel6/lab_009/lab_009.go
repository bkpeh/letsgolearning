package main

import "fmt"

func main() {
	ff("Hello", "World", ss)
}

func ss(s1 string, s2 string) {
	fmt.Println(s1, s2)
}

func ff(s1 string, s2 string, f func(s3 string, s4 string)) {
	f(s1, s2)
}
