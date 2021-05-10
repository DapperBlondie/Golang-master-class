package main

import (
	"fmt"
	"runtime"
)

var CLOSE = false

func main()  {

	inCh := make(chan int)
	outCh := make(chan int)

	go input(inCh)
	go output(outCh, inCh)

	fmt.Println("The result of Summation : ", <-outCh)

	fmt.Println("The operation finished !")
	fmt.Println("The GOMAXPROCS : ", runtime.GOMAXPROCS(0)) // returns the number of CPUs for executing threads

	return
}

func input(inCH chan int)  {

	for j:=1;j<10;j++ {

		inCH <- j*j
	}

	CLOSE = true

	close(inCH)
	return
}

func output(outCh chan int, inCh chan int)  {

	var Summation int = 0

	for j:= range inCh{

		Summation = Summation + j

		if CLOSE {
			break
		}
	}

	outCh <- Summation
	close(outCh)

	return
}
