package main

import (
	"fmt"
	"sync"
	"time"
)

//Creating the forks as structs
type Fork struct {
	sync.Mutex
	id        int
	timesUsed int
	chanIn    chan int
	chanOut   chan int
}

//Creating philosophers as structs
type Philosopher struct {
	id           int
	timesEaten   int
	eatingStatus int
	leftFork     *Fork
	rightFork    *Fork
	chanIn       chan int
	chanOut      chan int
}

/* func communication(p Philosopher, status int) {
	p.chanOut <- status
} */

func (p Philosopher) eat() {

	//Comment for debugging
	fmt.Printf("Philosopher number '%d' is still thinking....\n", p.id)

	//Synthetic wait time
	time.Sleep(2 * time.Second)

	//Locking the forks for the philosopher to eat
	p.leftFork.Lock()
	p.rightFork.Lock()

	//Updating eating status
	p.eatingStatus = 9999

	//Writing status to console. REPLACE THIS WITH CHANNEL COMMUNICATION
	if p.eatingStatus == 9999 {
		fmt.Printf("Philosopher number '%d' is eating now...\n", p.id)
	} else {
		fmt.Println("Oops something is ffffddd up")
	}

	//Synthetic wait time
	time.Sleep(4 * time.Second)

	//Unlocking the forks so the next philosopher can take them
	p.rightFork.Unlock()
	p.leftFork.Unlock()

	//Updating use count on forks
	p.leftFork.timesUsed++
	p.rightFork.timesUsed++

	//Updating eat-count for the philosopher
	p.timesEaten++

	//Printing fork status
	fmt.Printf("Philosopher number '%d' is done eating!\nLeft fork with id='%d' has been used '%d' times.\nRight fork with id='%d' has been used '%d' times.\n", p.id, p.leftFork.id, p.leftFork.timesUsed, p.rightFork.id, p.rightFork.timesUsed)

	//p.chanOut <- p.eatingStatus
	//communication(p, p.eatingStatus)

	//p.chanOut <- p.timesEaten

	//Printing times eaten
	fmt.Printf("Times philosopher number '%d' has eaten today = %d\n\n", p.id, p.timesEaten)

	//Updating eating status
	p.eatingStatus = 10000

	//Printing eating status. REPLACE WITH CHANNEL COMMUNICATION
	if p.eatingStatus == 10000 {
		fmt.Printf("Philosopher number '%d' starts thinking again...\n", p.id)
	} else {
		fmt.Println("Oops something is ffffddd up")
	}

	//Repeat the execution, for eating continuesly
	p.eat()

}

func main() {

	fmt.Println("Welcome to the Philosophers Dining party")

	//Edit for other party sizes
	var partySize int = 5

	//Forks are created as a slice pointing to the "Fork" type defined earlier
	forks := make([]*Fork, partySize)
	//Forks are created according to party size
	for i := 0; i < partySize; i++ {
		//Initialization of each fork
		forks[i] = &Fork{
			id:        i,
			timesUsed: 0,
			chanIn:    make(chan int),
			chanOut:   make(chan int),
		}
	}

	//Philosophers are created as a slice pointing to the "Philosopher" type defined earlier
	philosophers := make([]*Philosopher, partySize)
	//Philosophers are created according to party size
	for i := 0; i < partySize; i++ {
		//Initialization of each philosopher
		philosophers[i] = &Philosopher{
			id:           i,
			timesEaten:   0,
			eatingStatus: 1000,
			leftFork:     forks[i],
			rightFork:    forks[(i+1)%partySize],
			chanIn:       make(chan int),
			chanOut:      make(chan int),
		}

		//Start goroutines for all philosophers
		go philosophers[i].eat()
	}

	//Hack to make the dining continue endlessly
	endless := make(chan int)
	<-endless

	fmt.Println("All done eating")

}
