/*
Hands-on exercise #53 - struct with slice
Create your own type “person” which will have an underlying type of “struct” so that it can store
the following data:
	● first name
	● last name
	● favorite ice cream flavors
Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice
which stores the favorite flavors.
*/

package main

import "fmt"

type person struct {
	firstName        string
	lastName         string
	favoriteIceCream []string
}

func main() {

	jb := person{
		firstName:        "James",
		lastName:         "Bond",
		favoriteIceCream: []string{"Vanilla", "Strawberry"},
	}
	fmt.Println(jb)

	for _, flavour := range jb.favoriteIceCream {
		fmt.Printf("Favourite is %v\n", flavour)
	}

	jm := person{
		firstName: "Jenny",
		lastName:  "Moneypenny",
	}
	fmt.Println(jm)

	jm.favoriteIceCream = append(jm.favoriteIceCream, "Mint", "Chocolate")

	for _, flavour := range jm.favoriteIceCream {
		fmt.Printf("Favourite is %v\n", flavour)
	}

}
