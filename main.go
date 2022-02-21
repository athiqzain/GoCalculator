package main

import (
	"fmt"
)

func main() {

	number1 := 0
	number2 := 0
	number := 0

	fmt.Print("Enter first number: ")
	fmt.Scan(&number1)

	fmt.Print("Enter second number: ")
	fmt.Scan(&number2)

	fmt.Print("Select the operation 1-Addition 2-Subtraction 3-Multiplication 4:Division")
	fmt.Scan(&number)

	sum := number1 + number2
	diff := number1 - number2
	mul := number1 * number2
	div := number1 / number2

	switch number {

	case 1:
		fmt.Println("Sum is", sum)
	case 2:
		fmt.Println("Diff is", diff)
	case 3:
		fmt.Println("prod is", mul)
	case 4:
		fmt.Println("Div is", div)
	}
}
