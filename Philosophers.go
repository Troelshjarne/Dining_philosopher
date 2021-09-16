package main

import (
	"fmt"
	"math/rand"
	"time"
)

var timesEaten [PHILOSOPHERS]int
var status [PHILOSOPHERS]int

var P_OUT [PHILOSOPHERS]chan int
var P_IN [PHILOSOPHERS]chan int

// each philosopher must include two channels
//(one for input and one for output, both usable from the outside)
func InitPhilChans() {
	for i := 0; i < PHILOSOPHERS; i++ {

		P_OUT[i] = make(chan int, 10)
		P_IN[i] = make(chan int, 10)
		timesEaten[i] = 0
		status[i] = 1000
		P_OUT[i] <- timesEaten[i]
		P_OUT[i] <- status[i]
		//fmt.Println(<-P_OUT[i])
		//fmt.Println(<-P_OUT[i])
	}

}

func giveQuery() {

	fmt.Println("Enter query: philosopher ID and 1: for times eaten. 2: current status ")

	var PhilID int

	fmt.Scanln(&PhilID)

	var query int
	fmt.Scanln(&query)

	P_IN[PhilID] <- query
	go AnswerQuery(PhilID)
	time.Sleep(time.Millisecond * 500)
}

// take a phil_id and answers time eaten and status. 1 = timeEaten and 2 = status
func AnswerQuery(PhilID int) {
	query := <-P_IN[PhilID]
	switch query {
	case 1:
		fmt.Print("Times eaten ")
		fmt.Println(<-P_OUT[PhilID])
		<-P_OUT[PhilID]
		P_OUT[PhilID] <- timesEaten[PhilID]
		P_OUT[PhilID] <- status[PhilID]
	case 2:
		fmt.Print("status")
		<-P_OUT[PhilID]
		fmt.Println(<-P_OUT[PhilID])
		P_OUT[PhilID] <- timesEaten[PhilID]
		P_OUT[PhilID] <- status[PhilID]

	}
}

func randPhil() int {

	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 6
	return rand.Intn(max-min) + min

}
