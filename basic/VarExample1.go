package main

import "fmt"

func main() {
	// var 선언
	var i1 int = 10
	var s1 string = "string"

	fmt.Println(i1)
	fmt.Println(s1)

	// 타입 생략 가능
	var i2 = 10
	var s2 = "string"

	fmt.Println(i2)
	fmt.Println(s2)

	// := 를 이용한 변수 선언
	i3 := 10
	s3 := "string"

	fmt.Println(i3)
	fmt.Println(s3)

	// 다수의 변수를 동시에 선언
	var i4, j4, k4 int = 10, 11, 12
	s4, s5, s6 := "string1", "string2", "string3"

	fmt.Println(i4, j4, k4)
	fmt.Println(s4, s5, s6)

	// var () 를 이용한 변수 선언
	var (
		i7            = 10
		j8            = 11
		k9            = 12
		s10, s11, s12 = "string1", "string2", "string3"
	)

	fmt.Println(i7, j8, k9)
	fmt.Println(s10, s11, s12)
}
