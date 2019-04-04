package main

import (
	"log"
	"sync"
	"math/rand"
)

func problem1() {
	var wg sync.WaitGroup
	const limit = 100
	c := make(chan int, limit)

	log.Printf("problem1: started --------------------------------------------")

	for i := 0; i < limit; i++ {
		c <- i
	}

	for inx := 0; inx < 10; inx++ {

		wg.Add(1)
		go printRandom1(inx, c, &wg)

	}

	wg.Wait()

	log.Printf("problem1: finised --------------------------------------------")
}

func printRandom1(slot int, c <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for inx := 0; inx < 25; inx++ {
		select {
		case <-c:
			log.Printf("problem1: slot=%03d count=%05d rand=%f", slot, inx, rand.Float32())
		default:
			return
		}
	}
}
