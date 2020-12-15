package main
import "fmt"

func main() {
	/*
	// Example: Person Struct
		* Name, Address, Phone
		* Option1: Have 3 separate variables, programmer remembers that they are related
		* Option2: Make a single struct which contains all 3 vars
	*/

	type struct Person {
		name string
		address string
		phone string
	}

	var p1 = Person

	p1.name = "Arif"
	x := p1.address
	fmt.Println(p1.name)


	p2 := new(Person)

	p3 := Person(
		name = "Arif",
		address = "....."
		phone = "123..."
	)
}