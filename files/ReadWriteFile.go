package main

import (
	"fmt"
	"io/ioutil"
)

func Read(fileName string) {
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}

func Write(fileName string, message string) {
	bytes := []byte(message)

	err := ioutil.WriteFile(fileName, bytes, 0644)

	if err != nil {
		panic(err)
	}
}

func main() {
	Write("./test.txt", "abcdefg")
	Read("./test.txt")
}
