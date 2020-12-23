package main

import "fmt"

func main() {
	jb := []string{`Shaken, not stirred`, `Martinis`, `Women`}
	mp := []string{`James Bond`, `Literature`, `Computer Science`}
	no := []string{`Being evil`, `Ice cream`, `Sunsets`}

	x := map[string][]string{"bond_james": jb, "moneypenny_miss": mp, "no_dr": no}

	x["bk"] = []string{"Time", "money", "wife"}

	for i, v := range x {
		fmt.Println("For ", i)
		for ii, vv := range v {
			fmt.Println("Index: ", ii, ", Value: ", vv)
		}
	}

	delete(x, "bk")
	fmt.Println("deleting.......")

	for i, v := range x {
		fmt.Println("For ", i)
		for ii, vv := range v {
			fmt.Println("Index: ", ii, ", Value: ", vv)
		}
	}

}
