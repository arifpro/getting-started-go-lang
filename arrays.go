package main
import "fmt"

func main() {
	var x [5]int
	x[0] = 2
	fmt.Println(x[1])  // 0
	fmt.Println(x[3])  // 0
	fmt.Println(x[4])  // 0
	fmt.Println(x[0])  // 2

	
	y := [5]int{10, 20, 30, 40, 50}
	fmt.Println(y)  // [10 20 30 40 50]


	z := [...]int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	fmt.Println(z)  // [10 20 30 40 50 60 70 80 90]


	fmt.Println(len(z))  // 9


	for index, value := range y {
		fmt.Printf("index %d, value %d\n", index, value)
	}
	/*
		index 0, value 10
		index 1, value 20
		index 2, value 30
		index 3, value 40
		index 4, value 50
	*/
}