package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "a.b.c"
	r := strings.Replace(str, ".", "_", -1)
	fmt.Println(r)
}
