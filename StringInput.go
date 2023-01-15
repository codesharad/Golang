package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("ENter your name")
	str, _ := reader.ReadString('\n')
	str = strings.Replace(str, "\n", "", -1)

	//name, _ := strconv.Atoi(str)

	fmt.Println("Your name is ", str)

}
