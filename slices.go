/*
	* Pointer - indicates the start of the slice.
	* Length - is the number of elts in the slice.
	* Capacity - is maximum number of elts
*/

/*
	* len() - function returns the length
	* cap() - function returns the capacity
*/

package main
import "fmt"

func main() {
	arr := [...]string{"a", "b", "c", "d", "e", "f", "g"}

	s1 := arr[1:3] // index 1, 2
	s2 := arr[2:5]  // index 2, 3, 4

	fmt.Println(s1)  // [b c]
	fmt.Println(s2)  // [c d e]

	fmt.Println(len(s1), cap(s1))  // 2 3
	fmt.Println(len(s2), cap(s2))  // 3 5


	// slice
	slice := []int{1, 2, 3, 4}
	fmt.Println(slice)  //  [1 2 3 4]
}