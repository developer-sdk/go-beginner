package main

import "fmt"

type Obj struct {
	Name string
	Age  int
}

func PrintObject(list []Obj) {
	for index, object := range list {
		fmt.Printf("index: %d, object: %+v\n", index, object)
	}
}

func main() {
	list := []Obj{
		{"Beckham", 11},
		{"Zidane", 7},
		{"Ronaldo", 9},
	}

	for _, object := range list {
		object.Age = object.Age * 2
	}

	// 출력, 11, 7, 9
	PrintObject(list)

	for index := range list {
		object := &list[index]
		object.Age = object.Age * 2
	}

	// 출력, 22, 14, 18
	PrintObject(list)
}
