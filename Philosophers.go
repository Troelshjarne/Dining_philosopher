package main

import "fmt"

//var eatCount  map key philID value = count
// array of philosopher_Input channels
var P_IN [PHILOSOPHERS]chan int

// array of philosopher_output channels
//var P_OUT [PHILOSOPHERS]chan int

// each philosopher must include two channels
//(one for input and one for output, both usable from the outside)
func InitPhilChans() {
	for i := 0; i < PHILOSOPHERS; i++ {
		P_IN[i] = make(chan int)
		//P_OUT[i] = make(chan int)

	}

	// test  if input is correctly inserted into channels
	a := <-P_IN[0]
	b := <-P_IN[1]
	c := <-P_IN[2]
	d := <-P_IN[3]
	e := <-P_IN[4]

	fmt.Println(a, b, c, d, e)
}

// default value, 1000 = thinking initially
func Default_value() {
	for i := 0; i < PHILOSOPHERS; i++ {
		P_IN[i] <- 1000
	}

}
