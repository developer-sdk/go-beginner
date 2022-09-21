package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	strInt := "100"
	fmt.Println(string(strInt))

	i, err := strconv.Atoi(strInt)
	fmt.Println(i, err, reflect.TypeOf(i))

	i8, err := strconv.ParseInt(strInt, 0, 8)
	i16, err := strconv.ParseInt(strInt, 0, 16)
	i32, err := strconv.ParseInt(strInt, 0, 32)
	i64, err := strconv.ParseInt(strInt, 0, 64)

	fmt.Println(i8, err, reflect.TypeOf(i8))
	fmt.Println(i16, err, reflect.TypeOf(i16))
	fmt.Println(i32, err, reflect.TypeOf(i32))
	fmt.Println(i64, err, reflect.TypeOf(i64))
}
