package main

import (
	"log"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup
var totalPrinted = 0
var limit = 100

//The concept is that the first for loop starts the printRandom1 goroutine and waits for it to finish before executing the next one by adding the goroutine to the wait group
//Meanwhile the printRandom1 goroutine sends a message which contains the total of printed numbers after every for loop finishes running
//When the total of printed numbers is 100, the first loop breaks and problem1 finishes running
func problem1() {

	printedCountChannel := make(chan int)
	log.Printf("problem1: started --------------------------------------------")

	for inx := 0; inx < 10; inx++ {

		wg.Add(1)
		go printRandom1(inx, printedCountChannel)

		//Receive printedCount and stop the loop if it's equal to the limit
		printedCount := <-printedCountChannel
		if printedCount == limit {
			log.Printf("Finished printing exactly %3d random numbers from problem1..", limit)
			break
		}
	}
	wg.Wait()
	log.Printf("problem1: finished --------------------------------------------")

}

func printRandom1(slot int, printedCountChannel chan int) {
	defer wg.Done()
	for inx := 0; inx < 25; inx++ {
		log.Printf("problem1: slot=%03d count=%05d rand=%f", slot, inx, rand.Float32())
		totalPrinted++

		if totalPrinted == limit {
			break
		}
	}
	//Send totalPrinted through the channel after the end of every for loop from this goroutine
	printedCountChannel <- totalPrinted
}
