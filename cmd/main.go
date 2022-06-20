package main

import (
	"log"
	"time"
)

func main() {
	tickerChan := make(chan int)
	go ticker(tickerChan)

	for {
		select {
		case tick := <-tickerChan:
			log.Printf("Ticker value: %d", tick)
		}
	}
}

func ticker(c chan int) {
	counter := 0
	for {
		time.Sleep(1 * time.Second)
		c <- counter
		counter++
	}
}
