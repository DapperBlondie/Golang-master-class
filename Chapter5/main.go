package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

func mainRoutines()  {

	/*go firstRoutine()

	go func() {

		for i:=10;i<20;i++ {

			fmt.Print(i, " ")
		}

		return
	}()*/

	n := flag.Int("n", 5, "numberOfRoutines")
	flag.Parse()

	count := *n

	for i:=0;i<count;i++ {

		go func(x int) {

			fmt.Printf("The routine id : %d\n", x)
		}(i)
	}

	time.Sleep(time.Second*2)

	return
}

func firstRoutine()  {

	for i:=0;i<10;i++ {

		fmt.Print(i)
	}

	fmt.Println()
	return
}

func main()  {

	n := flag.Int("n", 5, "numberOfRoutines")
	flag.Parse()

	count := *n
	var waitGroup sync.WaitGroup

	for i:=0;i<count;i++ {
		waitGroup.Add(2)
		go func(x int) {

			defer waitGroup.Done()
			fmt.Printf("The routine id : %d\n", x)
			return
		}(i)

		go func(x int) {

			defer waitGroup.Done()
			fmt.Printf("The routine id : %d\n", 2*x+1)
			return
		}(i)
	}

	//time.Sleep(time.Second*3)
	waitGroup.Wait()
	fmt.Println("Go routines finished their jobs !")
	return
}
