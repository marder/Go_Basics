package main

import "fmt"

func main() {

	numbers := [5]int{13, 67, 69, 2, 100}

	for index, item := range numbers {
		fmt.Printf("numbers[%d] = %d \n", index, item)
	}

}
