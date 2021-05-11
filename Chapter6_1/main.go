package main

import (
	"fmt"
)

func main() {

	channel := make(chan int, 5)

	for j := 0; j < 10; j++ {

		select {
		case channel <- j:

		default:
			fmt.Println("We can not accept another value, ", j)
		}
	}

	for j := 0; j < 10; j++ {

		select {
		case channel <- j:

		default:
			fmt.Println("We can not accept another value, ", j)
		}
	}

	for {
		select {
		case num := <-channel:

			fmt.Println(num)
		default:

			fmt.Println("End of the Streaming !")
			return
		}
	}
}
