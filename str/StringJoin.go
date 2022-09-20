package main

import (
	"fmt"
	"strings"
)

func main() {
	strs := []string{"a", "b", "c"}
	fmt.Println(strings.Join(strs, ":"))
}
