package main

import (
	"fmt"
)

func main() {
	multiplier := 1

	for multiplier <= 10 {

		product := 5 * multiplier

		fmt.Printf("5 * %d = %d\n", multiplier, product)
		multiplier++
	}

}
