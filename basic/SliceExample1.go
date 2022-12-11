package main

import "fmt"

func main() {

	// ------------------------------
	// 슬라이스 생성
	letters := []string{"a", "b", "c", "d"}
	fmt.Println(letters) // [a b c d]

	// ------------------------------
	// make 를 이용한 slice 생성
	s1 := make([]int, 0)
	fmt.Println(s1, len(s1), cap(s1)) // [] 0 0

	s2 := make([]int, 5)
	fmt.Println(s1, len(s1), cap(s2)) //[] 0 5

	s3 := make([]int, 3, 5)
	fmt.Println(s2, len(s2), cap(s3)) //[0 0 0 0 0] 5 5

	// append
	// append는 len으로 설정한 값 뒤에 추가 됨
	intArray := []int{100, 101, 102}
	s1 = append(s1, 100)
	s2 = append(s2, intArray...) // 슬라이스 끼리 더할 때는 ... 을 추가
	s3 = append(s3, 100, 101, 102)
	fmt.Println(s1, len(s1), cap(s1)) // [100] 1 1
	fmt.Println(s2, len(s2), cap(s2)) // [0 0 0 0 0 100 101 102] 8 10
	fmt.Println(s3, len(s3), cap(s3)) // [0 0 0 100 101 102] 6 10

	// ------------------------------
	// copy
	// copy를 위해서 생성한 슬라이스는 기존 슬라이스와 크기가 동일해야 함
	lettersCopy1 := make([]string, 0)
	copy(lettersCopy1, letters)
	fmt.Println(lettersCopy1) // []

	lettersCopy2 := make([]string, len(letters), len(letters))
	copy(lettersCopy2, letters)
	fmt.Println(lettersCopy2) // [a b c d]

	lettersCopy2[3] = "="
	fmt.Println(letters)      // [a b c d]
	fmt.Println(lettersCopy2) // [a b c =]

	// ------------------------------
	// 삭제, 추출
	integers := []int{1, 2, 3, 4, 5}
	sub1 := integers[1:4]
	sub2 := integers[2:4]
	fmt.Println(integers, len(integers), cap(integers)) // [1 2 3 4 5] 5 5
	fmt.Println(sub1, len(sub1), cap(sub1))             // [2 3 4] 3 4
	fmt.Println(sub2, len(sub2), cap(sub2))             // [3 4] 2 3

	sub1[2] = 100
	fmt.Println(integers, len(integers), cap(integers)) // [1 2 3 100 5] 5 5
	fmt.Println(sub1, len(sub1), cap(sub1))             // [2 3 100] 3 4
	fmt.Println(sub2, len(sub2), cap(sub2))             // [3 100] 2 3
}
