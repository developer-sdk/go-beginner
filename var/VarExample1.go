package main

func main() {
	var i1 int = 10
	var s1 string = "string"

	println(i1)
	println(s1)

	// 타입 없이 변수 선언
	var i2 = 10
	var s2 = "string"

	println(i2)
	println(s2)

	// := 를 이용한 변수 선언
	i3 := 10
	s3 := "string"

	println(i3)
	println(s3)

	// 다수의 변수를 동시에 선언
	var i4, j4, k4 int = 10, 11, 12
	s4, s5, s6 := "string1", "string2", "string3"

	println(i4, j4, k4)
	println(s4, s5, s6)

	// var () 를 이용한 변수 선언
	var (
		i7            = 10
		j8            = 11
		k9            = 12
		s10, s11, s12 = "string1", "string2", "string3"
	)

	println(i7, j8, k9)
	println(s10, s11, s12)
}
