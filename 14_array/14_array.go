package main

import "fmt"

func main() {
	family := [3]string{"mama", "tata", "syn"}

	fmt.Println(family[0])

	fmt.Println(family[2])

	length := len(family)

	fmt.Println(length)

	fmt.Println("________")

	for i := 0; i < len(family); i++ {
		fmt.Println(family[i])
	}

}
