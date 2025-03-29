/*
Hands-on exercise #55 - embed struct
	● Create a type engine struct, and include this field
		○ electric bool
	● Create a type vehicle struct, and include these fields
		○ engine
		○ make
		○ model
		○ doors
		○ color
	● Create two VALUES of TYPE vehicle
		○ use a composite literal
	● Print out each of these values.
	● Print out a single field from each of these values.
*/

package main

import "fmt"

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	make   string
	model  string
	doors  int
	colour string
}

func main() {

	v1 := vehicle{
		engine: engine{electric: true},
		make:   "Ford",
		model:  "Puma",
		doors:  5,
		colour: "Navy Blue",
	}
	fmt.Printf("%#v\n", v1)
	fmt.Printf("Colour: %v\n", v1.colour)

	v2 := vehicle{
		engine: engine{electric: false},
		make:   "Vauxhall",
		model:  "Corsa",
		doors:  2,
		colour: "Silver",
	}

	fmt.Printf("%#v\n", v2)
	fmt.Printf("Electric: %v\n", v2.electric)
}
