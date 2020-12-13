package main
import "fmt"

func main() {
	const x = 1.3
	const (
		y = 4
		z = "Hi"
	)


	// Iota Example
	type Grades int
	const (
		A Grades = iota
		B
		C
		D
		F
	)
}
