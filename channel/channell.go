package main

import "fmt"

func main() {

	c := make(chan string)
	go introduce("AB", c)
	go introduce("BC", c)

	go introduce("CD", c)

	msg1 := <-c
	fmt.Println(msg1)
	msg2 := <-c
	fmt.Println(msg2)
	msg3 := <-c
	fmt.Println(msg3)
	close(c)
}

func introduce(student string, c chan string) {
	result := fmt.Sprintf("hai, my name", student)
	c <- result
}
