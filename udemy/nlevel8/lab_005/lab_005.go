package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type byage []user

func (a byage) Len() int           { return len(a) }
func (a byage) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byage) Less(i, j int) bool { return a[i].Age < a[j].Age }

type byname []user

func (a byname) Len() int           { return len(a) }
func (a byname) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byname) Less(i, j int) bool { return a[i].Last < a[j].Last }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	sort.Sort(byage(users))
	for i, v := range users {
		fmt.Println("No.", i+1)
		fmt.Println(v.Age, v.First, v.Last)
		sort.Strings(v.Sayings)
		fmt.Println(v.Sayings)
	}

	fmt.Println("-----------------")

	sort.Sort(byname(users))
	for i, v := range users {
		fmt.Println("No.", i+1)
		fmt.Println(v.Age, v.First, v.Last)
		sort.Strings(v.Sayings)
		fmt.Println(v.Sayings)
	}
}
