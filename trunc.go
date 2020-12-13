/*
1. Write a program which prompts the user to enter a floating point number
2. and prints the integer which is a truncated version of the floating point number that was entered.

Note: Truncation is the process of removing the digits to the right of the decimal place.
*/

package main
import "fmt"

func main() {
	var floatNum float64

	fmt.Printf("Enter a floating point number: ")
	fmt.Scan(&floatNum)

	fmt.Println(int(floatNum))
}