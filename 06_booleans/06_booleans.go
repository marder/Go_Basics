// Program to illustrate the working of Relational Operators

package main

import "fmt"

func main() {

	number1 := 13
	number2 := 56
	var result bool

	result = (number1 == number2)

	fmt.Printf("%d == %d returns %t \n", number1, number2, result)

	result = (number1 != number2)

	fmt.Printf("%d != %d returns %t \n", number1, number2, result)

	result = (number1 > number2)

	fmt.Printf("%d > %d returns %t \n", number1, number2, result)

	result = (number1 < number2)

	fmt.Printf("%d < %d returns %t \n", number1, number2, result)

}
