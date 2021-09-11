package main

import "time"

const PHILOSOPHERS = 5

func main() {

	go InitPhilChans()
	go DefaultValue()

	time.Sleep(time.Millisecond * 1500)

}
