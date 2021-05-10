package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var readValue = make(chan int)
var writeValue = make(chan int)

func set(newValue int) {

	writeValue <- newValue
	return
}

func get() int {

	return <-readValue
}

func monitor() {

	var value int = 200

	for {

		select {

		case newValue := <-writeValue:

			value = newValue
			fmt.Printf("value : %d\n", value)
		case readValue <- value:
		}
	}
}

func main() {

	num := 5
	var waitGroup sync.WaitGroup

	go monitor()
	for j := 0; j < num; j++ {

		waitGroup.Add(1)
		go func(jj int) {

			defer waitGroup.Done()
			set(rand.Intn(27*jj + 12))
			//fmt.Printf("The current value j : %d, %d\n", get(), jj)
		}(j)
	}

	waitGroup.Wait()
	fmt.Printf("The last value : %d\n", get())

	return
}
