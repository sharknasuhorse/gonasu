package main

import (
	"log"
	"time"
)

func main() {
	log.Print("started")

	sleep1_finished := make(chan bool)
	sleep2_finished := make(chan bool)
	sleep3_finished := make(chan bool)

	go func() {
		log.Print("1秒ストップ初め")
		time.Sleep(1 * time.Second)
		log.Print("1秒ストップ終わり")
		sleep1_finished <- true
	}()

	go func() {
		log.Print("2秒ストップ初め")
		time.Sleep(2 * time.Second)
		log.Print("2秒ストップ終わり")
		sleep2_finished <- true
	}()

	go func() {
		log.Print("3秒ストップ初め")
		time.Sleep(3 * time.Second)
		log.Print("3秒ストップ終わり")
		sleep3_finished <- true
	}()

	<-sleep1_finished
	<-sleep2_finished
	<-sleep3_finished
	log.Print("全部終わり")
}
