package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix()) // rand seed 는 매번 다른 값을 주지 않으면 결과가 동일함
	i := rand.Intn(15)

	fmt.Println(i)

	switch i {
	case 0, 1:
		fmt.Println("A")
	case 2, 3, 4:
		fmt.Println("B")
	case 5:
		fmt.Println("C")
	default:
		fmt.Println("D")
	}
}
