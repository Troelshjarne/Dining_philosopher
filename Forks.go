package main

import "sync"

var mu sync.Mutex

//hello
// out input query inUse
// array of fork_output channels
var F_IN [PHILOSOPHERS]chan int

// array of fork_input channels
var F_OUT [PHILOSOPHERS]chan int

func tryEat() {

	//fork.
	mu.Lock()

	// mutex
	//lock

	//status

	// eaten

}
