package main

import (
	"bufio"
	"fmt"
	"os"
)

const prompt = " and press enter when ready."

func main() {
	reader := bufio.NewReader(os.Stdin)

	//3 types of variable declaration
	var firstNumer int
	firstNumer = 2

	var secondNumber = 5

	thirdNumber := 7
	var ans int

	fmt.Println("Guess the number")

	fmt.Println("Think of a number b/w 1 and 10.", prompt)
	reader.ReadString('\n')

	fmt.Println("Multiply your number by ", firstNumer, prompt)
	reader.ReadString('\n')

	fmt.Println("Now, multiply the result by ", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Devide the number you initially thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now substract", thirdNumber, prompt)
	reader.ReadString('\n')

	ans = firstNumer*secondNumber-thirdNumber
	fmt.Println(ans)
}
