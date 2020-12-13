package main
import "fmt"

func main() {
	x := 10

	for i := 0; i < x; i++ {
		fmt.Println(i)

		if x > 5 {
			continue
		}

		if x > 7 {
			break
		}
	}


	// loop like while loop of another languages
	y := 0
	for y < 10 {
		fmt.Println("hi")
		y++
	}

	// infinity loop
	for {
		fmt.Println("Hi")
	}


	// switch
	switch x {
	case 1:
		fmt.Println("case 1")
	case 2:
		fmt.Println("case 2")
	default:
		fmt.Println("no case")
	}


	// scan number
	var appleNum int

	fmt.Printf("Number of apples?")
	num, err := fmt.Scan(&appleNum)

	fmt.Println(appleNum)

}