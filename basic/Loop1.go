package main

import "fmt"

func main() {

	for index := 1; index <= 10; index++ {
		fmt.Printf("index: %d\n", index)
	}

	for index := 1; index <= 10; {
		fmt.Printf("index: %d\n", index)
		index++
	}

	index := 1
	for index <= 10 {
		fmt.Printf("index: %d\n", index)
		index++
	}

	index = 1
	for true {
		fmt.Printf("index: %d\n", index)
		index++

		if index > 10 {
			break
		}
	}
}
