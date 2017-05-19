package main

import (
	"log"
	"time"
)

func main() {
	log.Print("started")
	x := 1

	log.Print(x, "のスリープ初め")
	time.Sleep(1 * time.Second)
	log.Print(x, "のスリープ終わり")
	x = x + 1

	log.Print(x, "のスリープ初め")
	time.Sleep(2 * time.Second)
	log.Print(x, "のスリープ終わり")
	x = x + 1

	log.Print(x, "のスリープ初め"
	time.Sleep(3 * time.Second)
	log.Print(x, "のスリープ終わり")
	x = x + 1

}
