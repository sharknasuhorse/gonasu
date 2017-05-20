package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Macだよね")
	case "linux":
		fmt.Println("linuxだよね")
	default:
		fmt.Printf("%s", os)
	}
}
