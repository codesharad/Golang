package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a number: ")
	str1, _, _ := reader.ReadRune()

	// remove newline
	//str1 = strings.Replace(str1, "\n", "", -1)

	// convert string variable to int variable
	num, e := strconv.Atoi(string(str1))
	if e != nil {
		fmt.Println("conversion error:", str1)
	}

	if num >= 1 && num <= 10 {
		fmt.Println("correct")
	} else {
		fmt.Println("num not in range")
	}
}

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	reader := bufio.NewReader(os.Stdin)

// 	str, _ := reader.ReadString('\n')
// 	str = strings.Replace(str, "\n", "", -1)

// 	num, e := strconv.Atoi(str)

// 	if e != nil {
// 		fmt.Println("Conversion error: ", str)
// 	}

// 	if num >= 1 && num < 10 {
// 		fmt.Println("Number entered is between 1 and 10")
// 	} else {
// 		fmt.Println("Number entered is outside range 1 and 10")
// 	}

// }
