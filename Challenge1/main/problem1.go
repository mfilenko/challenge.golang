package main

import (
	"log"
	"time"
	"math/rand"
)

func problem1() {
	const limit = 100
	c := make(chan int, limit)

	log.Printf("problem1: started --------------------------------------------")

	for i := 0; i < limit; i++ {
		c <- i
	}

	for inx := 0; inx < 10; inx++ {

		go printRandom1(inx, c)

	}

	//
	// Todo:
	//
	// Remove this quick and dirty sleep
	// against a synchronized wait until all
	// go routines are finished.
	//

	time.Sleep(5 * time.Second)

	log.Printf("problem1: finised --------------------------------------------")
}

func printRandom1(slot int, c <-chan int) {
	for inx := 0; inx < 25; inx++ {
		select {
		case <-c:
			log.Printf("problem1: slot=%03d count=%05d rand=%f", slot, inx, rand.Float32())
		default:
			return
		}
	}
}
