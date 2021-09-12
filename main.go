package main

import "time"

const PHILOSOPHERS = 5

func main() {

	go InitPhilChans()
	go DefaultValue()

	go AnswerQuery(0)
	time.Sleep(time.Millisecond * 1500)

}
