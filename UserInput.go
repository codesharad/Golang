package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
  
  fmt.Println("ENter your name")
	str, _ := reader.ReadString('\n')
	str = strings.Replace(str, "\n", "", -1)

	num, e := strconv.Atoi(str)

	if e != nil {
		fmt.Println("Conversion error: ", str)
	}

	if num >= 1 && num < 10 {
		fmt.Println("Number entered is between 1 and 10")
	} else {
		fmt.Println("Number entered is outside range 1 and 10")
	}

}
