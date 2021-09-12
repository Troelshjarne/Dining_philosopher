package main

import "fmt"

var P_OUT [PHILOSOPHERS]chan int

// array of philosopher_output channels
var P_IN [PHILOSOPHERS]chan int

// each philosopher must include two channels
//(one for input and one for output, both usable from the outside)
func InitPhilChans() {
	for i := 0; i < PHILOSOPHERS; i++ {
		P_OUT[i] = make(chan int)

	}

	// test  if input is correctly inserted into channels

}

// take a phil_id and answers time eaten and status
func AnswerQuery(PhilID int) {
	status := <-P_OUT[PhilID]

	if status == 1 {
		fmt.Println("thinking", status)
	}

}

// default value, 1000 = thinking initially
func DefaultValue() {
	for i := 0; i < PHILOSOPHERS; i++ {
		P_OUT[i] <- 1

	}

}
