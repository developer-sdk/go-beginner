package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fileName := "test.txt"
	fileInfo, err := os.Stat(fileName)

	// 파일 정보 확인
	if fileInfo != nil {
		fmt.Printf("%+v\n", fileInfo)
	} else {
		fmt.Println(err)
	}

	if os.IsNotExist(err) {
		// 생성
		file, err := os.Create(fileName)
		if err != nil {
			panic(err)
		}

		defer func(file *os.File) {
			if err := file.Close(); err != nil {
				panic(err)
			}
		}(file)
	} else {
		currentTime := time.Now().Local()

		// 시간 변경
		err = os.Chtimes(fileName, currentTime, currentTime)
		if err != nil {
			panic(err)
		}
	}
}
