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
			case channel <- j:
				//default:
				//fmt.Println("We can not accept another value, ", j)
			}
		}

		return
	}()

	go func() {

		for j := 0; j < 10; j++ {

			select {
			case channel <- j:
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
