package main

import "github.com/poy/lightsoff"

func main() {
	someChannel := make(chan int, 10)
	count := 5
	lights := lightsoff.New(count, func() {
		close(someChannel)
	})

	for i := 0; i < count; i++ {
		go run(someChannel, lights)
	}

	for range someChannel {
		//NOP
	}
}

func run(someChannel chan<- int, lights *lightsoff.LightsOff) {
	defer lights.TurnOff()
	defer println("Done")

	for i := 0; i < 100; i++ {
		someChannel <- i
	}
}
