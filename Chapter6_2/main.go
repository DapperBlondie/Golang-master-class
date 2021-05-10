package main

import (
	"fmt"
	"math/rand"
	"time"
)

func mainMain()  {

	channel := make(chan int)
	go send(channel)
	go add(channel)

	time.Sleep(time.Second*4)

	return
}

func send(ch chan int)  {

	for ch != nil {

		//fmt.Println("sign of the rand")
		ch <- rand.Intn(10)
		//fmt.Println("Hello")
	}

	fmt.Println("The Sender Function Broke !")
	return
}

func add(ch chan int) {

	sum := 0
	t := time.NewTimer(time.Second*3)
	for {

		select {

		case input := <-ch:

			sum = sum + input
		case <-t.C:

			ch = nil
			fmt.Println(sum)
		}
	}
}

func main()  {

	channel := make(chan int, 1)

	go func() {

		for  {

			num := rand.Intn(20)
			if num == 15 {

				channel = nil
				break
			}else {

				channel <- num
			}
		}

		return
	}()

	for  {

		if channel==nil {

			fmt.Println("ForLoop for receiving data was broken !")
			break
		}else {

			fmt.Println(<-channel)
		}
	}

	return
}
