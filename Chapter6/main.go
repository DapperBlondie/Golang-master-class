package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func mainHTTP()  {

	ch := make(chan string)

	go func() {

		res, err := http.Get("https://amazon.com")
		if err != nil {
			ch <- err.Error()
			return
		}

		ch <- res.Status
		return
	}()

	go func() {

		res, err := http.Get("https://github.com")
		if err != nil {
			ch <- err.Error()
			return
		}

		ch <- res.Status
		return
	}()

	for  {

		select {
		case res := <-ch:

			fmt.Println(res)
		case <- time.After(time.Millisecond*5000):

			fmt.Println("Timed Out !")
			return
		}
	}
}

func main()  {

	var max int = 10
	chNumber := make(chan int)
	end := make(chan string)

	go generate(0, max, chNumber, end)

	for i := 0;i<max;i++ {
		fmt.Println("Number ", i,"th : ", <-chNumber)
	}

	time.Sleep(time.Second*4)

	end <- "Ending ..."
	fmt.Println("It just a BLOCK !")
	time.Sleep(time.Millisecond*2)

	return
}

func generate(min , max int, chNumber chan int, end chan string)  {

	for  {

		select {
		case chNumber <- rand.Intn(max-min) + min:
		case <-end:
			fmt.Println("end Channel Received a MSG !")
			close(end)
			return
		case <-time.After(time.Second*3):
			fmt.Println("time.After() Called !")
		}
	}
}
