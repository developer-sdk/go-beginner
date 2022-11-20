package main

import "fmt"

func main() {

	strArray := []string{"A", "B", "C"}
	for index, str := range strArray {
		fmt.Printf("index: %d, str: %s\n", index, str)
	}

	dictionary := map[string]string{
		"key_A": "value_A",
		"key_B": "value_B",
	}

	for key, value := range dictionary {
		fmt.Printf("key: %s, value: %s\n", key, value)
	}

}
