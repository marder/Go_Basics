package main

import "fmt"

func main() {
	for i := 1; i <= 8; i++ {
		for j := 1; j <= 8; j++ {

			if j == 5 {
				continue
			}

			fmt.Println("i=", i, "j=", j)

		}
	}
}
