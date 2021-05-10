package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func main()  {

	channel := make(chan string)

	fmt.Println("the number of cpu core in mt local machine : ", returnNumberOfCpu())

	go func() {
		time.Sleep(time.Second*2)
		fmt.Println("hello from another routine.")
		UID, PPID:= os.Getuid(), os.Getppid()
		fmt.Println("The UID & PPID : ", UID, PPID)
		channel <- "Pushed"
	}()
	fmt.Println("The trigger : ", <-channel)
	printMachineInfo()
}

func returnNumberOfCpu() int {
	return runtime.NumCPU()
}

func printMachineInfo()  {
	fmt.Println("The Go lang architecture we use : ", runtime.GOARCH)
	fmt.Println("The number of go routine in our program : ", runtime.NumGoroutine())
	fmt.Println("The os utilities : ", os.Getgid(), "\nPID : ", os.Getpid())
}