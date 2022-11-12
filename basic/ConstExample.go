package main

import "fmt"

const i = 1
const s string = "STRING"

// iota 가 0이 할당되고, 순서대로 증가
// A: 0, B: 1, C: 2, D: 3
const (
	A = iota
	B
	C
	D
)

func main() {
	fmt.Println(i)
	fmt.Println(s)

	fmt.Println(A, B, C, D)
}
