package main

import "time"

const PHILOSOPHERS = 5

func main() {

	go InitPhilChans()

	go giveQuery()

	//go AnswerQuery(0)

	time.Sleep(time.Millisecond * 5000)

	//go DefaultValue()

	//go AnswerQuery(1, 1)

}
