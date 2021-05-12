package main

import (
	"fmt"
	"time"
)

func main() {

	channel := make(chan int, 5)

	go func() {

		for j := 0; j < 10; j++ {

			select {
			case channel <- j: // this is a blocking call
				//default:
				//fmt.Println("We can not accept another value, ", j)
			}
		}

		return
	}()

	go func() {

		for j := 0; j < 10; j++ {

			select {
			case channel <- j: // this is a blocking call, because it is possible our buffered channel be full
				// and we need to exhaust it !
				//default:
				//fmt.Println("We can not accept another value, ", j)
			}
		}

		return
	}()

	for {
		select {
		case num := <-channel:

			fmt.Println(num)
		case <-time.After(time.Second * 3):

			fmt.Println("End of Streaming ...")
			return
		}
	}
}
