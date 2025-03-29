/*
Hands-on exercise #56 - anonymous struct
Create and use an anonymous struct with these fields:
	● first string
	● friends map[string]int
	● favDrinks []string
Print things
*/

package main

import "fmt"

func main() {

	p1 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first:     "James",
		friends:   map[string]int{"Jenny": 27, "Q": 21, "M": 47},
		favDrinks: []string{"Gin and Tonic", "Vodka", "Red Wine"},
	}

	fmt.Println(p1)

	for k, friend := range p1.friends {
		fmt.Println(p1.first, " - friend - ", k, friend)
	}

	for _, drink := range p1.favDrinks {
		fmt.Println(p1.first, " - drink - ", drink)
	}

}
