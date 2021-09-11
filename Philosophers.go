package main

import "fmt"

//var eatCount  map key philID value = count
// array of philosopher_Input channels
var P_OUT [PHILOSOPHERS]chan int

// array of philosopher_output channels
//var P_OUT [PHILOSOPHERS]chan int

// each philosopher must include two channels
//(one for input and one for output, both usable from the outside)
func InitPhilChans() {
	for i := 0; i < PHILOSOPHERS; i++ {
		P_OUT[i] = make(chan int)

	}

	// test  if input is correctly inserted into channels
	a := <-P_OUT[0]
	b := <-P_OUT[1]
	c := <-P_OUT[2]
	d := <-P_OUT[3]
	e := <-P_OUT[4]

	fmt.Println(a, b, c, d, e)
}

func AnswerQuery(PhilID int) {
	status := <-P_OUT[PhilID]
	fmt.Println("status", status)

}

// default value, 1000 = thinking initially
func DefaultValue() {
	for i := 0; i < PHILOSOPHERS; i++ {
		P_OUT[i] <- 1

	}

}
