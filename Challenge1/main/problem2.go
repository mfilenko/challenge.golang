package main

import (
	"log"
	"time"
	"math/rand"
)

func problem2() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	log.Printf("problem2: started --------------------------------------------")

	for inx := 0; inx < 10; inx++ {

		go printRandom2(inx, ticker)

	}

	//
	// Todo:
	//
	// Remove this quick and dirty sleep
	// against a synchronized wait until all
	// go routines are finished.
	//
	// Same as problem1...
	//

	time.Sleep(5 * time.Second)

	log.Printf("problem2: finished -------------------------------------------")
}

func printRandom2(slot int, ticker *time.Ticker) {

	for inx := 0; inx < 10; inx++ {
		// The following line blocks further execution until a value will be
		// delivered to the channel by a ticker.
		<-ticker.C
		log.Printf("problem2: slot=%03d count=%05d rand=%f", slot, inx, rand.Float32())

	}
}
