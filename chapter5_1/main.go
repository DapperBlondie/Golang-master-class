package main

import (
	"fmt"
	"time"
)

func main()  {

	ch := make(chan string)
	go  waitChannel(ch, "Hello, channel !")

	/*msg,_ := <- ch
	fmt.Println("The message from channel : ", msg)*/

	//time.Sleep(time.Second*1)
	fmt.Println("channel said : ", <-ch)
	time.Sleep(time.Second*1)
	_, ok := <-ch
	if ok {
		fmt.Println("channel ch is opened !")
	}else {
		fmt.Println("channel ch is closed !")
	}


	return
}

func waitChannel(ch chan string, value string)  {

	fmt.Println("The value you passed : ", value)
	ch <- value
	close(ch)
	fmt.Println("The value you passed : ", value)

	return
}
