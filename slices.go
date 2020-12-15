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


	/*
	// make()
		* create a slice (and array) using make()
		* make(type, length/capacity)
		* make(type, length, capacity)
	*/
	sli1 := make([]int, 10)
	sli2 := make([]int, 10, 15)

	fmt.Println(sli1)  // [0 0 0 0 0 0 0 0 0 0]
	fmt.Println(sli2)  // [0 0 0 0 0 0 0 0 0 0]

	fmt.Println(len(sli1), cap(sli1))  // 10 10
	fmt.Println(len(sli2), cap(sli2))  // 10 15


	/*
	// append()
		* size of a slice can be increased by append()
		* adds elements to the end of a slice
		* inserts into underlying array
	*/

	// increases size of array if necessary
	sli := make([]int, 0, 3)
	fmt.Println(sli)  // []
	fmt.Println(len(sli))  // 0

	// length of sli is 0
	sli = append(sli, 100)
	fmt.Println(sli)  // [100]
	fmt.Println(len(sli))  // 1
}