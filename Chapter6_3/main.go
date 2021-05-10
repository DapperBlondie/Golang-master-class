package main

import (
	"fmt"
	"time"
)

func main()  {

	x := make(chan struct{})
	y := make(chan struct{})
	z := make(chan struct{})

	go C(z)
	go A(x, y)
	go C(z)
	go B(y, z)
	go C(z)

	close(x)
	time.Sleep(3 * time.Second)

	return
}

func A(a, b chan struct{})  {

	<-a
	fmt.Println("We are in A()")
	time.Sleep(time.Second*1)
	close(b)
	return
}

func B(a, b chan struct{})  {

	<-a
	fmt.Println("We are in B()")
	close(b)
	return
}

func C(b chan struct{})  {

	<-b
	fmt.Println("We are in C()")
	return
}