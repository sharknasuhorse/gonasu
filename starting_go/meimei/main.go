package main

import (
	"fmt"
)

const (
	朝の挨拶 = "おはよう"
	昼の挨拶 = "こんにちは"
	夜の挨拶 = "こんばんわ"
)

func あいさつ(m string) {
	fmt.Println(m)
}

func main() {
	あいさつ(夜の挨拶)
}
