package main

import "time"

const PHILOSOPHERS = 5

func main() {

	go InitPhilChans()
	go Default_value()

	time.Sleep(time.Millisecond * 1500)

}
