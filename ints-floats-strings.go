package main
import "fmt"

func main() {
	// Type Conversion
	var x int32 = 1
	var y int16 = 2
	x = int32(y)


	// Floating Point
	/*
		float32 ~6 digits of precision
		float64 ~15 digits of precision
	*/


	// Complex numbers represented as two floats: real and imaginary
	var z complex128 = complex(2, 3)  // 2+3i


	// ASCII and Unicode
	/*
		'A' = 0x41
	*/

	/*
		* Unicode is a 32-bit character code
		* UTF-8 is variable length (8-bit UTF codes are same as ASCII)
		* Code points - Unicode characters
		* Rune - a code point in Go
	*/


	// Assign a string
	str := "Hi Khan"

}