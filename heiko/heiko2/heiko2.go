package main

import (
	"log"
	"time"
)

func main() {
	log.Print("started")

	finished := make(chan bool)

	go func() {
		log.Print("1秒ストップ初め")
		time.Sleep(1 * time.Second)
		log.Print("1秒ストップ終わり")
		finished <- true
	}()

	go func() {
		log.Print("2秒ストップ初め")
		time.Sleep(2 * time.Second)
		log.Print("2秒ストップ終わり")
		finished <- true
	}()

	go func() {
		log.Print("3秒ストップ初め")
		time.Sleep(3 * time.Second)
		log.Print("3秒ストップ終わり")
		finished <- true
	}()

	for i := 0; i < 3; i++ {
		<-finished

	}

	log.Print("全部終わり")
}
