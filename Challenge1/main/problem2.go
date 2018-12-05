package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

var wg2 sync.WaitGroup

//The concept is that the first for loop starts the printRandom2 goroutine and waits for it to finish before executing the next one by adding the goroutine to the wait group
//Meanwhile the printRandom2 goroutine waits for one second after printing each number, then sends a message (with a struct) to the first goroutine to signal the end of the second
//The first goroutine moves to the next loop only upon receiving the message from the second goroutine
func problem2() {

	//Channel of type struct{} because no information is included in it
	secondPassedEventChannel := make(chan struct{})
	log.Printf("problem2: started --------------------------------------------")

	for inx := 0; inx < 10; inx++ {

		wg.Add(1)
		go printRandom2(inx, secondPassedEventChannel)
		//Receive printedCount
		<-secondPassedEventChannel
	}
	wg.Wait()
	log.Printf("problem2: finished --------------------------------------------")

}

func printRandom2(slot int, secondPassedEventChannel chan struct{}) {
	defer wg.Done()
	for inx := 0; inx < 25; inx++ {
		log.Printf("problem2: slot=%03d count=%05d rand=%f", slot, inx, rand.Float32())
		time.Sleep(1 * time.Second)
	}
	//Send a struct through the channel after the end of every for loop from this goroutine
	secondPassedEventChannel <- struct{}{}
}
