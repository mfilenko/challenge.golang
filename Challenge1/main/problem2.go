package main

import (
	"log"
	"time"
	"sync"
	"math/rand"
)

func problem2() {
	var wg sync.WaitGroup
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	log.Printf("problem2: started --------------------------------------------")

	for inx := 0; inx < 10; inx++ {

		wg.Add(1)
		go printRandom2(inx, ticker, &wg)

	}

	wg.Wait()

	log.Printf("problem2: finished -------------------------------------------")
}

func printRandom2(slot int, ticker *time.Ticker, wg *sync.WaitGroup) {

	for inx := 0; inx < 10; inx++ {
		// The following line blocks further execution until a value will be
		// delivered to the channel by a ticker.
		<-ticker.C
		log.Printf("problem2: slot=%03d count=%05d rand=%f", slot, inx, rand.Float32())

	}

	wg.Done()
}
