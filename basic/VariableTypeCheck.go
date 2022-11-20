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

	v := 100

	switch conv := interface{}(v).(type) {
	case int:
		fmt.Println(conv, "int")
	case bool:
		fmt.Println(conv, "bool")
	case string:
		fmt.Println(conv, "string")
	default:
		fmt.Println(conv, "unknown")
	}
}
