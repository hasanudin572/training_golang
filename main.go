package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {

		if i%2 != 0 {
			fmt.Println(i, " ganjil\n")
		} else {
			fmt.Println(i, " Genap\n")
		}

	}
}
