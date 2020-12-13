package main
import "fmt"

func main() {
	var x int = 1
	var y int
	var ip *int  // ip is pointer to int

	ip = &x  // ip now points to x
	y = *ip  // y is now 1

	fmt.Println("ip = ", ip)  // 0xc0000b4008
	fmt.Println("y = ", y)  // 1
}