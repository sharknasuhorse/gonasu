package main

import (
	"fmt"
)

func main() {
	fprint()
	intprint()
	hensu()
}

func fprint() {
	fmt.Printf("10進数=%d 2進数=%b¥n", 17, 17)
	fmt.Println()
}

func intprint() {
	var x int
	x = 10
	x -= 20
	fmt.Println(x)

	y := 50
	y = 20
	println(y)
}

func hensu() {
	o := 20
	p := 90
	println(o, p)
}
