package main

import "fmt"
import "os"

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	fmt.Printf("%%v %v\n", p)
	fmt.Printf("%%+v %+v\n", p)
	fmt.Printf("%%#v %#v\n", p)
	fmt.Printf("%%T %T\n", p)
	fmt.Printf("%%t %t\n", true)
	fmt.Printf("%%d %d\n", 123)
	fmt.Printf("%%b %b\n", 14)
	fmt.Printf("%%c %c\n", 33)
	fmt.Printf("%%x %x\n", 456)
	fmt.Printf("%%f %f\n", 78.9)
	fmt.Printf("%%e %e\n", 123400000.0)
	fmt.Printf("%%E %E\n", 123400000.0)
	fmt.Printf("%%s %s\n", "\"string\"")
	fmt.Printf("%%q %q\n", "\"string\"")
	fmt.Printf("%%x %x\n", "hex this")
	fmt.Printf("%%p %p\n", &p)
	fmt.Printf("|%6d|%6d|\n", 12, 345)
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)
	fmt.Printf("|%6s|%6s|\n", "foo", "b")
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)
	fprintf, err := fmt.Fprintf(os.Stderr, "an %s\n", "error")
	if err != nil {
		return
	}
	fmt.Printf("%d written.", fprintf)
}
