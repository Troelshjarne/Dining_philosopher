package main

//var eatCount  map key philID value = count
// array of philosopher_Input channels
var P_IN [PHILOSOPHERS]chan int

// array of philosopher_output channels
var P_OUT [PHILOSOPHERS]chan int

// each philosopher must include two channels
//(one for input and one for output, both usable from the outside)
func InitPhilChans(philosophers int) {
	for i := 0; i < PHILOSOPHERS; i++ {
		P_IN[i] = make(chan int)
		P_OUT[i] = make(chan int)
	}
}
