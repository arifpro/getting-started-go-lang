/*
1. Write a program which prompts the user to enter a string.
2. The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’.
3. The program should print “Found!” if the entered string starts with the character ‘i’, ends with the character ‘n’, and contains the character ‘a’.
4. The program should print “Not Found!” otherwise.
5. The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.

Examples: The program should print “Found!” for the following example entered strings, “ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”. The program should print “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian”.
*/
package main

import ( 
	"fmt"
	"strings"
	"bufio"
	"io"
  	"os"
) 

func main() {
	var stdin *bufio.Reader
	var line []rune
	var c rune
	var err error

	stdin = bufio.NewReader(os.Stdin)
	
	fmt.Printf("Enter a string: ")

	for {
		c, _, err = stdin.ReadRune()
		if err == io.EOF || c == '\n' { break }
		if err != nil {
			fmt.Fprintf(os.Stderr,"Error reading standard input\n")
			os.Exit(1)
		}
		line = append(line,c)
	}

	inputStr := string(line[:len(line)])
	str := strings.ToLower(inputStr)
	count := 0

	// starts with 'i'
	if strings.HasPrefix(str, "i") {
		count++
	}

	// end with 'n'
	if strings.HasSuffix(str, "n") {
		count++
	}

	// check if the string contains character 'a'
	for i := 1; i < len(str)-1; i++ {
		if 'a' == str[i] {
			count++
		}
	}

	if count == 3 {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}