/*
Hands-on exercise #54 - map struct
Take the code from the previous exercise, then store the VALUES of type person in a map with
the KEY of last name. Access each value in the map. Print out the values, ranging over the slice.
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
	//fmt.Println(jb)

	//for _, flavour := range jb.favoriteIceCream {
	//	fmt.Printf("Favourite is %v\n", flavour)
	//}

	jm := person{
		firstName: "Jenny",
		lastName:  "Moneypenny",
	}
	//fmt.Println(jm)

	jm.favoriteIceCream = append(jm.favoriteIceCream, "Mint", "Chocolate")

	//for _, flavour := range jm.favoriteIceCream {
	//	fmt.Printf("Favourite is %v\n", flavour)
	//}

	var myMap = make(map[string]person)
	myMap[jb.lastName] = jb
	myMap[jm.lastName] = jm
	//fmt.Println(myMap)

	for k, v := range myMap {
		fmt.Printf("key: %v\t First: %v\t Last: %v\nFavourites: ", k, v.firstName, v.lastName)
		for _, favourite := range v.favoriteIceCream {
			fmt.Printf("%v\t", favourite)
		}
		fmt.Println()
	}

}
